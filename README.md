# api-server

CURL Query:
```
GET -> curl -X GET <url> -H "Authorization:Bearer <token>"

POST -> curl -X POST -H "Authorization:Bearer <token>"  -d '{"id":"2","name":"add-data","type":"add-type","window":{"type":"Door","height":"20","width":"10"}}' <url>

PUT -> curl -X PUT -H "Authorization:Bearer <token>"  -d '{"id":"100","name":"update-data","type":"update-type","window":{"type":"Door","height":"20","width":"10"}}' <url>

DELETE -> curl -X DELETE <url> -H "Authorization:Bearer <token>"
```
..............................

Example:
```
Token Flag: -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRvaG9yaXplZCI6dHJ1ZSwiZXhwIjoxNjc3NTgxNzk4LCJ1c2VycyI6Im9iYXlkdWxsYWgifQ.DmJa4Q_iT6IDWyHOLLmxR1rvVqaVL_6AIxNbKh8UQ8s"

GET -> curl -X GET http://localhost:9091/sales 

POST -> curl -X POST http://localhost:9091/sales -d '{"id":"2","name":"add-data","type":"add-type","window":{"type":"Door","height":"20","width":"10"}}' 

PUT -> curl -X PUT -d '{"id":"100","name":"update-data","type":"update-type","window":{"type":"Door","height":"20","width":"10"}}' http://localhost:9091/sales/2 

DELETE -> curl -X DELETE http://localhost:9091/sales/2
```
