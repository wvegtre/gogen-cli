package detector

import (
	"context"
	"database/sql"
	"fmt"

	"gogen-cli/aigo/config"
	"gogen-cli/aigo/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type MysqlDetector struct {
	database string
	dbClient *sql.DB
}

func NewMysqlDetector(dc *config.DBAuth) *MysqlDetector {
	dbClient, err := dialMysql(dc)
	if err != nil {
		panic(err)
	}
	return &MysqlDetector{
		database: dc.Database,
		dbClient: dbClient,
	}
}

func dialMysql(c *config.DBAuth) (*sql.DB, error) {
	// {user}:{pwd}@tcp({ip}:{port})/{dbClient}?charset=utf8
	dsn := `%s:%s@tcp(%s:%s)/%s?charset=%s`
	dsn = fmt.Sprintf(dsn, c.UserName, c.Password, c.IP, c.Port, c.Database, "utf8")
	client, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	err = client.Ping()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return client, nil
}

func (m *MysqlDetector) GetAllTables() ([]string, error) {
	var tableNames []string

	rows, err := m.dbClient.Query(`SHOW TABLES`)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		tableNames = append(tableNames, tableName)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return tableNames, nil
}

func (m *MysqlDetector) GetTableFields(ctx context.Context, tableName string) (model.TableFields, error) {
	// query and get table columns from mysql
	query := `SELECT COLUMN_NAME, DATA_TYPE, COLUMN_COMMENT 
		FROM information_schema.COLUMNS 
		WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?`
	rows, err := m.dbClient.QueryContext(ctx, query, m.database, tableName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()
	// parse table columns
	var fields []model.TableField
	for rows.Next() {
		var field model.TableField
		err := rows.Scan(&field.FieldName, &field.FieldType, &field.FieldDescription)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		fields = append(fields, field)
	}
	return fields, nil
}
