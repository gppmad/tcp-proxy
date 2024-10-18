# tcp-proxy
Simple TCP Proxy built with Golang.

It is specialized to forward the request to httbin.org and return the response.

## Usage
Run the proxy with the following command:
```bash
go run main.go
```

## Example
Open a new terminal and send a request to the proxy:

```bash
curl -X GET http://localhost:8080/get
```



