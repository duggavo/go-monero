package levin_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
	"github.com/stretchr/testify/assert"

	"github.com/duggavo/go-monero/pkg/levin"
)

func TestPortableStorage(t *testing.T) {
	spec.Run(t, "NewPortableStorageFromBytes", func(t *testing.T, when spec.G, it spec.S) {
		it("fails w/ wrong sigA", func() {
			bytes := []byte{
				0xaa, 0xaa, 0xaa, 0xaa,
			}

			_, err := levin.NewPortableStorageFromBytes(bytes)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "sig-a doesn't match")
		})

		it("fails w/ wrong sigB", func() {
			bytes := []byte{
				0x01, 0x11, 0x01, 0x01,
				0xaa, 0xaa, 0xaa, 0xaa,
			}

			_, err := levin.NewPortableStorageFromBytes(bytes)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "sig-b doesn't match")
		})

		it("fails w/ wrong format ver", func() {
			bytes := []byte{
				0x01, 0x11, 0x01, 0x01,
				0x01, 0x01, 0x02, 0x01,
				0xaa,
			}

			_, err := levin.NewPortableStorageFromBytes(bytes)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "version doesn't match")
		})

		it("reads the contents", func() {
			bytes := []byte{
				0x01, 0x11, 0x01, 0x01, // sig a
				0x01, 0x01, 0x02, 0x01, // sig b
				0x01, // format ver

				0x08, // var_in(len(entries))

				// node_data
				0x09,                                                 // len("node_data")
				0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, // "node_data"
				0x0c, // boost_serialized_obj
				0x04, // var_in(node_data.entries)

				// for i in range node_data
				0x03,             // len("foo")
				0x66, 0x6f, 0x6f, // "foo"
				0x0a,             // boost_serialized_string
				0xc,              // var_in(len("bar"))
				0x62, 0x61, 0x72, // "bar"

				// payload_data
				0x0c,                                                                   // len("payload_data")
				0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, // "payload_data"
				0x0c, // boost_serialized_obj
				0x04, // var_in(payload_data.entries)

				// for i in range payload_data.entries
				0x06,                               // len("number")
				0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, // "number"
				0x06,                   // boost_serialized_uint32
				0x01, 0x00, 0x00, 0x00, // uint32(1)
			}

			ps, err := levin.NewPortableStorageFromBytes(bytes)
			assert.NoError(t, err)

			assert.Len(t, ps.Entries, 2)
			assert.Equal(t, ps.Entries[0].Name, "node_data")
			assert.EqualValues(t, ps.Entries[0].Value, levin.Entries{
				{
					Name:  "foo",
					Value: "bar",
				},
			})

			assert.Equal(t, ps.Entries[1].Name, "payload_data")
			assert.EqualValues(t, ps.Entries[1].Value, levin.Entries{
				{
					Name:  "number",
					Value: uint32(1),
				},
			})
		})
	}, spec.Report(report.Log{}), spec.Parallel(), spec.Random())

	spec.Run(t, "ReadVarIn", func(t *testing.T, when spec.G, it spec.S) {
		it("i <= 63", func() {
			b := []byte{0x08}
			n, v := levin.ReadVarInt(b)

			assert.Equal(t, n, 1)
			assert.Equal(t, v, 2)
		})

		it("64 <= i <= 16383", func() {
			b := []byte{0x01, 0x02}
			n, v := levin.ReadVarInt(b)
			assert.Equal(t, n, 2)
			assert.Equal(t, v, 128)
		})

		it("16384 <= i <= 1073741823", func() {
			b := []byte{0x02, 0x00, 0x01, 0x00}
			n, v := levin.ReadVarInt(b)
			assert.Equal(t, n, 4)
			assert.Equal(t, v, 16384)
		})
	}, spec.Report(report.Log{}), spec.Parallel(), spec.Random())

	spec.Run(t, "VarrIn", func(t *testing.T, when spec.G, it spec.S) {
		it("i <= 63", func() {
			i := 2 // 0b00000010

			b, err := levin.VarIn(i)
			assert.NoError(t, err)
			assert.Equal(t, b, []byte{
				0x08, // 0b00001000	(shift left twice, union 0)
			})
		})

		it("64 <= i <= 16383", func() {
			i := 128 // 0b010000000

			b, err := levin.VarIn(i)
			assert.NoError(t, err)
			assert.Equal(t, b, []byte{
				0x01, 0x02, // 0b1000000001 ((128 * 2 * 2) | 1) == 513
				// '    '
				// 1   2 * 256
			})
		})

		it("16384 <= i <= 1073741823", func() {
			i := 16384 // 1 << 14

			b, err := levin.VarIn(i)
			assert.NoError(t, err)
			assert.Equal(t, b, []byte{
				0x02, 0x00, 0x01, 0x00, // (1 << 16) | 2
			})
		})
	}, spec.Report(report.Log{}), spec.Parallel(), spec.Random())

	spec.Run(t, "PortableStorage", func(t *testing.T, when spec.G, it spec.S) {
		it("bytes", func() {
			ps := &levin.PortableStorage{
				Entries: []levin.Entry{
					{
						Name: "node_data",
						Serializable: &levin.Section{
							Entries: []levin.Entry{
								{
									Name:         "foo",
									Serializable: levin.BoostString("bar"),
								},
							},
						},
					},
					{
						Name: "payload_data",
						Serializable: &levin.Section{
							Entries: []levin.Entry{
								{
									Name:         "number",
									Serializable: levin.BoostUint32(1),
								},
							},
						},
					},
				},
			}

			assert.Equal(t, []byte{
				0x01, 0x11, 0x01, 0x01, // sig a
				0x01, 0x01, 0x02, 0x01, // sig b
				0x01, // format ver
				0x08, // var_in(len(entries))

				// node_data
				0x09,                                                 // len("node_data")
				0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, // "node_data"
				0x0c, // boost_serialized_obj
				0x04, // var_in(node_data.entries)

				// for i in range node_data
				0x03,             // len("foo")
				0x66, 0x6f, 0x6f, // "foo"
				0x0a,             // boost_serialized_string
				0xc,              // var_in(len("bar"))
				0x62, 0x61, 0x72, // "bar"

				// payload_data
				0x0c,                                                                   // len("payload_data")
				0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, // "payload_data"
				0x0c, // boost_serialized_obj
				0x04, // var_in(payload_data.entries)

				// for i in range payload_data.entries
				0x06,                               // len("number")
				0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, // "number"
				0x06,                   // boost_serialized_uint32
				0x01, 0x00, 0x00, 0x00, // uint32(1)

			}, ps.Bytes())
		})
	}, spec.Report(report.Log{}), spec.Parallel(), spec.Random())
}
