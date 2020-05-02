# min-rpc
a minimal calculator service for learning jsonrpc in go. 

### directory:
- rpc.go
- client
  - main.go
- server
  - main.go

### So what is RPC 
RPC means Remote Procedure Call. It allows you(client) to call functions(service) that is on another computer(server) and receive result, just like calling any other functions, by sending parameters(json), without knowing the lower level communication details(Socket, Http). 

### How to play with this
run server on port:1234,
run client, dial at port:1234
call service like it is a function
```
// call div service
err = client.Call("DemoService.Div", rpcdemo.Args{10, 0}, &result)
// call add service
err = client.Call("DemoService.Add", rpcdemo.Args{10, 0}, &result)
```
