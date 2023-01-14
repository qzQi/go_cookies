go目前就先作为自己的取代Java的语言，go web先算是自己的gui吧。

学完go操作数据库再学一下前端！！！

go access database 里面的命令。

```bash
go run filename.go
go run .
# the difference ?
```

### Accessing databases

#### Tutorial: Accessing a relational database
讲了最基本的，连接数据库。涉及了数据的查改。
```go
import(
    "database/sql"
    "github.com/go-sql-driver/mysql"
    // sql-driver
)
```
都需要database/sql，这个module定义了很多，包括一些接口。我们使用不同的数据库需要import不同的sql-driver。

**使用步骤，以mysql为例**      
```go
import(
    "database/sql"
    "github.com/go-sql-driver/mysql"
)

var db *sql.DB
//more about db handler: a connection pool
//https://go.dev/doc/database/open-handle

func main() {
	conf := mysql.Config{
		User:   "qzyDB",
		Passwd: "helloQzy",
		Net:    "tcp",
		Addr:   "120.24.178.74:3306",
		DBName: "recordings",
	}
    db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
    // ping一下看看连接没有
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
    // 接下来就是增查改删
}
```



```go
// 查询数据库，query 
```





#### Accessing relational databases

https://go.dev/doc/database/，这个tutorial算是overall吧。

* supported DBMS
* function for query && exec
* transaction
* Query cancellation
* managed connection-pool



**单条sql语句的执行**

select使用query，update/insert/delete使用exec。

* exec

[再看这个tutorial吧](https://go.dev/doc/database/change-data)





When you perform database actions that don’t return data, use an `Exec` or `ExecContext` method from the `database/sql` package. SQL statements you’d execute this way include `INSERT`, `DELETE`, and `UPDATE`.



有点忘记了，难道每个row都有一个id？id, err := result.LastInsertId()

`DB.Exec` returns values: an [`sql.Result`](https://pkg.go.dev/database/sql#Result) and an error. When the error is `nil`, you can use the `Result` to get the ID of the last inserted item (as in the example) or to retrieve the number of rows affected by the operation.



If your code will be executing the same SQL statement repeatedly, consider using an `sql.Stmt` to create a reusable prepared statement from the SQL statement. For more, see [Using prepared statements](https://go.dev/doc/database/prepared-statements).

对于需要重复执行的sql语句可以使用sql.Stmt，be cautious，不要使用fmt.Sprintf获得string，sql injection。





* [query](https://go.dev/doc/database/querying)









#### database handler

https://go.dev/doc/database/open-handle



执行一条sql语句，以及执行一个transaction。都不用自己操心，执行一条不过是在pool的选出一条不要管上下文；执行一个transaction需要上下文？



```go
// three step to open a handler
1. locate a driver
import(
    _ "github.com/go-sql-driver/mysql"
    //如果不使用mysql.func 使用空白import
    // 对数据库的操作，我们的database/sql里面都有，为了与
    //具体的dbms解耦合，使用sql里的func
)

2、open a database handler
var dbHandler *sql.DB

//格式字符串的format根据不同的dbms也不同，这里以mysql为例
dbHandler,err:=sql.Open("mysql","username:password@tcp(127.0.0.1:3306)/jazzrecords")
//这个DSN：dataSoruceName string一般可以由特定的sql-driver生成。
//比如在MySQL里面的的config
// Specify connection properties.
cfg := mysql.Config{
    User:   username,
    Passwd: password,
    Net:    "tcp",
    Addr:   "127.0.0.1:3306",
    DBName: "jazzrecords",
}

// Get a database handle.
db, err = sql.Open("mysql", cfg.FormatDSN())
if err != nil {
    log.Fatal(err)
}

3、Confirming a connection
db, err = sql.Open("mysql", connString)

// Confirm a successful connection.
if err := db.Ping(); err != nil {
    log.Fatal(err)
}
```



把用户/密码放在文件里面可能会有安全问题，可通过环境变量解决。

```go
username:=os.Getenv("DB_USER")
passwd:=os.GetEnv("DB_PASS")

//in bash
$ export DB_USER=username
$ export DB_PASS=password
```

虽然我们不需要在意数据库连接，但有些时候需要释放不需要的资源。          

Those can include resources held by an `sql.Rows` representing data returned from a query or an `sql.Stmt` representing a prepared statement.

```go
rows, err := db.Query("SELECT * FROM album WHERE artist = ?", artist)
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

// Loop through returned rows.
```





#### Executing SQL statements that don't return data

When you perform database actions that don’t return data, use an `Exec` or `ExecContext` method from the `database/sql` package. SQL statements you’d execute this way include `INSERT`, `DELETE`, and `UPDATE`.





When your query might return rows, use a `Query` or `QueryContext` method instead. For more, see [Querying a database](https://go.dev/doc/database/querying).





An `ExecContext` method works as an `Exec` method does, but with an additional `context.Context` argument, as described in [Canceling in-progress operations](https://go.dev/doc/database/cancel-operations).

```go
//eg insert
func AddAlbum(album []Album)(int64,error){
    
}
```

https://go.dev/doc/database/change-data
