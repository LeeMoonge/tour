package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
	// ！！注意：在程序中必须导入github.com/go-sql-driver/mysql包进行Mysql驱动程序的初始化，否则会报错
	_ "github.com/go-sql-driver/mysql"
)

// 生成结构体

// 数据库连接核心对象
type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

// 存储连接Mysql的一些基本信息
type DBInfo struct {
	DBType   string
	Host     string
	Username string
	Password string
	Charset  string
}

// 存储Columns表中需要的一些字段
type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

// 表字段关系映射
var DBTypeTostructType = map[string]string{
	"int": "int32",
	"tinyint": "int8",
	"smallint": "int",
	"mediumint": "int64",
	"bigint": "int64",
	"bit": "int",
	"bool": "bool",
	"enum": "string",
	"set": "string",
	"varchar": "string",
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

// 连接Mysql数据库的具体方法
func (m *DBModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?" +
		"charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		s,
		m.DBInfo.Username,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset)
	// 连接Mysql数据库时使用的是标准库database/sql的Open方法，第一个参数为驱动名称，第二个参数是数据库的连接信息
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}

// 获取表中列的消息
func (m *DBModel) GetColumns(dbname, tablename string) ([]*TableColumn, error) {
	query := "SELECT COLUMN_NAME, DATA_NAME, COLUMN_KEY, " +
		"IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT" +
		"FROM COLUMN WHERE TABLE_SCHRMA = ? AND TABLE_NAME = ? "
	rows, err := m.DBEngine.Query(query, dbname, tablename)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()

	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType,
			&column.ColumnKey, &column.IsNullable, &column.ColumnType,
			&column.ColumnComment)
		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	return columns, nil
}
