# tcp-proxy
Simple TCP Proxy built with Golang.

It is specialized to forward the request to the target address and return the response.

## Usage
Run the proxy with the following command:
```bash
go run main.go -listen {ip}:{port} -target {ip}:{port}

# Example
go run main.go -listen localhost:8080 -target httpbin.org:80
```

## Example
Open a new terminal and send a request to the proxy:

```bash
curl -X GET http://localhost:8080/get

# In case of multiple requests.
ab -n 5 -c 2 "http://127.0.0.1:8080/get"
```




