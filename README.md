# go-monero

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/duggavo/go-monero)


A multi-platform Go library for interacting with Monero servers
either on clearnet or not, supporting daemon and wallet RPC,
p2p commands and ZeroMQ.

## Quick start

### Library

To consume `go-monero` as a library for your Go project:

```bash
go get -u -v github.com/duggavo/go-monero
```

`go-monero` exposes an high-level package: `rpc`.

The package `rpc`, is used to communicate with `monerod` and `monero-wallet-rpc` via its HTTP
endpoints. Note that not all endpoints/fields are exposed on a given port - if
it's being served in a restricted manner, you'll have access to less endpoints
than you see in the documentation
([daemon RPC](https://www.getmonero.org/resources/developer-guides/daemon-rpc.html), )

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
	client, err := rpc.NewClient(addr)
	if err != nil {
		panic(fmt.Errorf("new client for '%s': %w", addr, err))
	}

	// instantiate a daemon-specific client and call the `get_height`
	// remote procedure.
	height, err := daemon.NewClient(client).GetHeight(ctx)
	if err != nil {
		panic(fmt.Errorf("get height: %w", err))
	}

	fmt.Printf("height=%d hash=%s\n", height.Height, height.Hash)
}
```

And to get the height from `monero-wallet-rpc`:
```go
package wallet_test

import (
	"context"
	"fmt"

	"github.com/duggavo/go-monero/rpc"
	"github.com/duggavo/go-monero/rpc/wallet"
)

func ExampleGetHeight() {
	ctx := context.Background()
	addr := "http://localhost:18086"

	// instantiate a generic RPC client
	client, err := rpc.NewClient(addr)
	if err != nil {
		panic(fmt.Errorf("new client for '%s': %w", addr, err))
	}

	// instantiate a wallet-specific client and call the `get_height`
	// remote procedure.
	height, err := wallet.NewClient(client).GetHeight(ctx)
	if err != nil {
		panic(fmt.Errorf("get height: %w", err))
	}

	fmt.Printf("height=%d\n", height.Height)
}
```

## License

See [LICENSE](./LICENSE).


## Thanks

Huge thanks to [Ciro Costa](https://github.com/cirocosta/go-monero) for writing the original implementation!
