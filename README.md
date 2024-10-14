
# サーバーの起動

```shell
$ go run cmd/api-server/main.go 
2024/10/14 15:36:57 Starting web server on :8080
```

# Request Sample

* `GET /api/invoices`

```shell
$ curl -s -X GET localhost:8080/api/invoices | jq .
[
  {
    "ID": 1,
    "Amount": 100,
    "DueDate": "2024-10-14T00:00:00+09:00",
    "IssueDate": "2024-10-14T00:00:00+09:00",
    "TotalAmount": 104.4,
    "FeeRate": 0.04,
    "TaxRate": 0.1,
    "Status": "未処理"
  },
  {
    "ID": 2,
    "Amount": 100,
    "DueDate": "2024-10-14T00:00:00+09:00",
    "IssueDate": "2024-10-14T00:00:00+09:00",
    "TotalAmount": 104.4,
    "FeeRate": 0.04,
    "TaxRate": 0.1,
    "Status": "未処理"
  }
]
```

* `POST /api/invoices`

```shell
$ curl -i -X POST -H "Content-Type: application/json" -d '{"Amount":100, "due_date": "2024-10-14T00:00:00Z"}' localhost:8080/api/invoices

HTTP/1.1 201 Created
Date: Mon, 14 Oct 2024 06:35:20 GMT
Content-Length: 0
```