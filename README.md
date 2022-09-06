# Winning Fish Backend

## 環境変数

```
FRONTEND_URL
```


```
# migration version up
GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=./webapp.db goose up

# migration version down
GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=./webapp.db goose down
```