# Winning Fish Backend

## 環境変数

```
FRONTEND_URL=localhost:1234
GOOSE_DRIVER=sqlite3
GOOSE_DBSTRING=./webapp.db
```

```
# migration version up
GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=./webapp.db goose up

# migration version down
GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=./webapp.db goose down
```


```
# 対話型shell
gore

# 型出力
fmt.Println(reflect.TypeOf(hoge))
```