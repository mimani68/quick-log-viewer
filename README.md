# Log system

# Usage

## Submit a log

```bash
curl -XPOST \
    --header "Content-Type: application/json" \
    --data '{"data":{"id":1,"message":"salam"}, "project":"app.io", "environment":"stage","service":"ai"}' \
        http://localhost:3000/log/submit
```

## Query specific value

```bash
curl -XPOST \
    --header "Content-Type: application/json" \
    --data '{"query":"salam","project":"app.io", "environment":"stage", "service":"ai"}' \
        http://localhost:3000/log/query
```
