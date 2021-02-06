# autocomplete-proxy
A proxy written in Go to augment the API of an existing service

### Instructions
- Run the user API with `docker run -ti --rm -p 8080:8080 cvstom/interview:latest autocomplete`
- Run the proxy API using the binary with `proxy`
- Make a request to the proxy API's user endpoint `curl localhost:8081/users?search=dan`
- Have a great time
