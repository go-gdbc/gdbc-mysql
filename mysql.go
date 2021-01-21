package mysql

import (
	"github.com/go-gdbc/gdbc"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	gdbc.Register("mysql", "mysql", &MySQLDataSourceNameAdapter{})
}

type MySQLDataSourceNameAdapter struct {
}

func (dsnAdapter MySQLDataSourceNameAdapter) GetDataSourceName(dataSource gdbc.DataSource) (string, error) {
	dsn := ""

	return dsn, nil
}
