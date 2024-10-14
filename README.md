
# サーバーの起動

```shell
$ go run cmd/api-server/main.go 
2024/10/14 15:36:57 Starting web server on :8080
```

# Request Sample

* `GET /api/invoices`

```shell
$ curl -i -X GET localhost:8080/api/invoices
HTTP/1.1 200 OK
Date: Mon, 14 Oct 2024 06:37:06 GMT
Content-Length: 58
Content-Type: text/plain; charset=utf-8

[{"id":1,"amount":100,"due_date":"2024-10-14T00:00:00Z"}]
```

* `POST /api/invoices`

```shell
$ curl -i -X POST -H "Content-Type: application/json" -d '{"id": 1, "Amount":100, "due_date": "2024-10-14T00:00:00Z"}' localhost:8080/api/invoices

HTTP/1.1 201 Created
Date: Mon, 14 Oct 2024 06:35:20 GMT
Content-Length: 0
```