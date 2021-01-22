package mysql

import (
	"errors"
	"github.com/go-gdbc/gdbc"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const DefaultHost = "localhost"
const DefaultPort = "3306"

const DefaultUsername = "root"
const DefaultPassword = ""

func init() {
	gdbc.Register("mysql", "mysql", &MySQLDataSourceNameAdapter{})
}

type MySQLDataSourceNameAdapter struct {
}

func (dsnAdapter MySQLDataSourceNameAdapter) GetDataSourceName(dataSource gdbc.DataSource) (string, error) {
	dsn := ""
	host := DefaultHost
	port := DefaultPort
	user := DefaultUsername
	password := DefaultPassword
	databaseName := ""

	dataSourceUrl := dataSource.GetURL()

	arguments := dataSourceUrl.Query()
	socketArgument := ""
	if arguments != nil {
		socketArgument = arguments.Get("socket")
		if socketArgument != "" {
			user = ""
			arguments.Del("socket")
		}
	}

	if dataSourceUrl.Opaque != "" {
		databaseName = dataSourceUrl.Opaque
	} else {
		if dataSourceUrl.Hostname() != "" {
			host = dataSourceUrl.Hostname()
		}

		if dataSourceUrl.Port() != "" {
			port = dataSourceUrl.Port()
		}

		if dataSourceUrl.User != nil {
			if dataSourceUrl.User.Username() != "" {
				user = dataSourceUrl.User.Username()
			}
			userPassword, _ := dataSourceUrl.User.Password()
			if userPassword != "" {
				password = userPassword
			}
		} else {
			if dataSource.GetUsername() != "" {
				user = dataSource.GetUsername()
			}
			if dataSource.GetPassword() != "" {
				password = dataSource.GetPassword()
			}
		}

		if dataSourceUrl.Path != "" {
			databaseName = dataSourceUrl.Path
		}
	}

	if strings.HasPrefix(databaseName, "/") {
		databaseName = databaseName[1:]
	}

	if strings.Contains(databaseName, "/") {
		return "", errors.New("database name format is wrong : " + databaseName)
	}

	if user != "" {
		if password != "" {
			dsn = user + ":" + password
		} else if password == "" {
			dsn = user
		}
	} else {
		if password != "" {
			return "", errors.New("user must not be empty when a password is specified")
		}
	}

	if dsn != "" {
		dsn = dsn + "@"
	}

	if socketArgument != "" {
		dsn = dsn + "unix(" + socketArgument + ")/" + databaseName
	} else {
		dsn = dsn + "tcp(" + host + ":" + port + ")/" + databaseName
	}

	if len(arguments) == 0 {
		return dsn, nil
	}

	dsn = dsn + "?"
	for argumentName, values := range arguments {
		dsn = dsn + argumentName + "=" + values[0] + "&"
	}
	dsn = dsn[:len(dsn)-1]

	return dsn, nil

}
