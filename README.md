# CLI-tool for making HTTP-requests

### supports:
1) method selection
2) headers
3) body of request
4) showing response
5) possibly saving response to a file

### Project structure


### Using
make requests with flags:
1) -X - request method
2) -d - body of request
3) -H - header of request(may be several)
4) -o - file for saving name to

#### Example of request:
minicurl -X POST -d '{"name":"Sarmat"}' -H "Content-Type: application/json" https://primerapi.com/api

