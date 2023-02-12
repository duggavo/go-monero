# go-monero

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/duggavo/go-monero)


A multi-platform Go library for interacting with Monero servers
either on clearnet or not, supporting daemon and wallet RPC,
p2p commands and ZeroMQ.

## Quick start

### Library

To consume `go-monero` as a library for your Go project:

```console
$ go get -u -v github.com/duggavo/go-monero
```

`go-monero` exposes two high-level packages: `levin` and `rpc`.

The first (`levin`) is used for interacting with the p2p network via plain TCP
(optionally, Tor and I2P can also be used via socks5 proxy - see options). 

For instance, to reach out to a node (of a particular address `addr`) and grab
its list of connected peers (information that comes out of the initial
handshake):

```golang
import (
	"context"
	"fmt"

	"github.com/duggavo/go-monero/levin"
)

func ListNodePeers(ctx context.Context, addr string) error {
	// start a client - this will actually establish a TCP `connect()`ion
	// with the other node.
	//
	client, err := levin.NewClient(ctx, addr)
	if err != nil {
		return fmt.Errorf("new client '%s': %w", addr, err)
	}

	// close the connection when done
	//
	defer client.Close()

	// perform the handshake
	//
	pl, err := client.Handshake(ctx)
	if err != nil {
		return fmt.Errorf("handshake: %w", err)
	}

	// list the peers reported back (250 max per monero's implementation)
	//
	for addr := range pl.Peers {
		fmt.Println(addr)
	}

	return nil
}
```

The second (`rpc`), is used to communicate with `monerod` via its HTTP
endpoints. Note that not all endpoints/fields are exposed on a given port - if
it's being served in a restricted manner, you'll have access to less endpoints
than you see in the documentation
(https://www.getmonero.org/resources/developer-guides/daemon-rpc.html)

`rpc` itself is subdivided in two other packages: `wallet` and `daemon`, exposing `monero-wallet-rpc` and `monerod` RPCs accordingly.

For instance, to get the the height of the main chain:

```go
package daemon_test

import (
	"context"
	"fmt"

	"github.com/duggavo/go-monero/rpc"
	"github.com/duggavo/go-monero/rpc/daemon"
)

func ExampleGetHeight() {
	ctx := context.Background()
	addr := "http://localhost:18081"

	// instantiate a generic RPC client
	//
	client, err := rpc.NewClient(addr)
	if err != nil {
		panic(fmt.Errorf("new client for '%s': %w", addr, err))
	}

	// instantiate a daemon-specific client and call the `get_height`
	// remote procedure.
	//
	height, err := daemon.NewClient(client).GetHeight(ctx)
	if err != nil {
		panic(fmt.Errorf("get height: %w", err))
	}

	fmt.Printf("height=%d hash=%s\n", height.Height, height.Hash)
}
```

## License

See [LICENSE](./LICENSE).


## Thanks

Huge thanks to [Ciro Costa](https://github.com/cirocosta/go-monero) for writing the original implementation!
