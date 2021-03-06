# gdbc-mysql
GDBC Mysql Driver - It is based on [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

[![Go Report Card](https://goreportcard.com/badge/github.com/go-gdbc/gdbc-mysql)](https://goreportcard.com/report/github.com/go-gdbc/gdbc-mysql)
[![codecov](https://codecov.io/gh/go-gdbc/gdbc-mysql/branch/main/graph/badge.svg?token=7UNHBOILSV)](https://codecov.io/gh/go-gdbc/gdbc-mysql)
[![Build Status](https://travis-ci.com/go-gdbc/gdbc-mysql.svg?branch=main)](https://travis-ci.com/go-gdbc/gdbc-mysql)

# Usage
```go
dataSource, err := gdbc.GetDataSource("gdbc:mysql://username:password@localhost:3000/testdb?charset=utf8mb4")
if err != nil {
    panic(err)
}

var connection *sql.DB
connection, err = dataSource.GetConnection()
if err != nil {
    panic(err)
}
```

MySQL GDBC URL takes one of the following forms:

```
gdbc:mysql://host:port/database-name?arg1=value1
gdbc:mysql://host/database-name?arg1=value1
gdbc:mysql:database-name?arg1=value1
gdbc:mysql:?arg1=value1
gdbc:mysql://username:password@host:port/database-name?arg1=value1
```

Using Socket:

You have to pass the argument **socket**, the argument does not belong to the driver.
```
gdbc:mysql:/database-name?socket=/tmp/mysql.sock&args1=value1
```

Default Values:
* **Host** : localhost
* **Port** : 3306
* **User** : root
* **Password** : 

Checkout [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) for arguments details.
