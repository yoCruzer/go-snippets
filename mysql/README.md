# mysql

## Dependencies

```bash
$ go get github.com/go-sql-driver/mysql
```
[Wiki Link](https://github.com/go-sql-driver/mysql/wiki/Examples)

## Getting Started

```go
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

db, err := sql.Open("mysql", "user:password@/dbname")
```