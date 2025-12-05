## Bookstore App

### API

- gRPC
- REST

#### How to call endpoints

```shell
# gRPC
grpcurl -protoset <(buf build -o -) -plaintext -d '{"shelf": "1"}' localhost:8080 bookstore.v1.BookstoreService/ListBooks

# REST
curl \
    -H 'Content-Type: application/json' \
    -d '{"shelf": "1"}' \
    http://localhost:8080/bookstore.v1.BookstoreService/ListBooks
```
