package mysql

import (
	"github.com/go-gdbc/gdbc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getDSN(t *testing.T, dataSourceUrl string) (string, error) {
	adapter := gdbc.GetDataSourceNameAdapter("mysql")
	dataSource, err := gdbc.GetDataSource(dataSourceUrl)
	assert.Nil(t, err)
	return adapter.GetDataSourceName(dataSource)
}

func getDSNWithUser(t *testing.T, dataSourceUrl string, username string, password string) (string, error) {
	adapter := gdbc.GetDataSourceNameAdapter("mysql")
	dataSource, err := gdbc.GetDataSource(dataSourceUrl, gdbc.Username(username), gdbc.Password(password))
	assert.Nil(t, err)
	return adapter.GetDataSourceName(dataSource)
}

func TestMySQLDataSourceNameAdapter_GetDataSourceNameWithWrongDatabaseName(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:mysql://localhost:3000/testdb/wrong?charset=utf8mb4")
	assert.NotNil(t, err)
	assert.Equal(t, "database name format is wrong : testdb/wrong", err.Error())
	assert.Empty(t, dsn)
}

func TestMySQLDataSourceNameAdapter_GetDataSourceNameWithoutUser(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:mysql://localhost:3000/testdb?charset=utf8mb4")
	assert.Nil(t, err)
	assert.Equal(t, dsn, DefaultUsername+"@tcp(localhost:3000)/testdb?charset=utf8mb4")
}

func TestMySQLDataSourceNameAdapter_GetDataSourceNameWithoutArguments(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:mysql://localhost:3000/testdb")
	assert.Nil(t, err)
	assert.Equal(t, dsn, DefaultUsername+"@tcp(localhost:3000)/testdb")
}

func TestMySQLDataSourceNameAdapter_GetDataSourceNameWithoutUserAndPort(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:mysql://localhost/testdb?charset=utf8mb4")
	assert.Nil(t, err)
	assert.Equal(t, dsn, DefaultUsername+"@tcp(localhost:"+DefaultPort+")/testdb?charset=utf8mb4")
}

func TestMySQLDataSourceNameAdapter_GetDataSourceNameWithoutUserAndHostAndPort(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:mysql:testdb?charset=utf8mb4")
	assert.Nil(t, err)
	assert.Equal(t, dsn, DefaultUsername+"@tcp("+DefaultHost+":"+DefaultPort+")/testdb?charset=utf8mb4")
}

func TestMySQLDataSourceNameAdapter_GetDataSourceNameWithoutUserAndHostAndPortAndDatabase(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:mysql:?charset=utf8mb4")
	assert.Nil(t, err)
	assert.Equal(t, dsn, DefaultUsername+"@tcp("+DefaultHost+":"+DefaultPort+")/?charset=utf8mb4")
}

func TestMySQLDataSourceNameAdapter_GetDataSourceNameWithFullFormat(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:mysql://username:password@localhost:3000/testdb?charset=utf8mb4")
	assert.Nil(t, err)
	assert.Equal(t, dsn, "username:password@tcp(localhost:3000)/testdb?charset=utf8mb4")
}

func TestPostgresDataSourceNameAdapter_GetDataSourceNameWithSocket(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:mysql:/testdb?socket=/tmp/mysql.sock&charset=utf8mb4")
	assert.Nil(t, err)
	assert.Equal(t, dsn, "unix(/tmp/mysql.sock)/testdb?charset=utf8mb4")
}

func TestPostgresDataSourceNameAdapter_GetDataSourceNameWithSocketUsingUserAndPassword(t *testing.T) {
	dsn, err := getDSNWithUser(t, "gdbc:mysql:/testdb?socket=/tmp/mysql.sock&charset=utf8mb4", "username", "password")
	assert.Nil(t, err)
	assert.Equal(t, dsn, "username:password@unix(/tmp/mysql.sock)/testdb?charset=utf8mb4")
}

func TestPostgresDataSourceNameAdapter_GetDataSourceNameWithSocketUsingUser(t *testing.T) {
	dsn, err := getDSNWithUser(t, "gdbc:mysql:/testdb?socket=/tmp/mysql.sock&charset=utf8mb4", "username", "")
	assert.Nil(t, err)
	assert.Equal(t, dsn, "username@unix(/tmp/mysql.sock)/testdb?charset=utf8mb4")
}

func TestMySQLDataSourceNameAdapter_GetDataSourceNameWithoutUsernameUsingSocket(t *testing.T) {
	dsn, err := getDSNWithUser(t, "gdbc:mysql:/testdb?socket=/tmp/mysql.sock&charset=utf8mb4", "", "test")
	assert.NotNil(t, err)
	assert.Equal(t, "user must not be empty when a password is specified", err.Error())
	assert.Empty(t, dsn)
}
