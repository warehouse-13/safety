## safety

A "real" server to use when doing acceptance level testing/sanity checking
of things which interact with [`flintlock`](https://github.com/weaveworks-liquidmetal/flintlock).

Use this when you don't necessarily care what `flintlock` should do, you
just care that it is being asked to do _something_.

### As a local server

This will start a server on `localhost:9090` or whichever port you set with `--port`.

```bash
git clone https://github.com/warehouse-13/safety
cd safety
make build
./safety-on
```

### As a test server

This will start a server on a real network with a randomly assigned port.

```go
server := fakeserver.New()
address := server.Start("")
// use the address for something
defer server.Stop()
```

### As a programmatic test server

This will start a server on a fake network (buffer).

```go
server := fakeserver.New()
dialer := server.StartBuf("")
defer server.Stop()

conn, _ := grpc.DialContext(context.TODO, "bufnet", grpc.WithContextDialer(dialer), grpc.WithInsecure())
defer conn.Close()
```
