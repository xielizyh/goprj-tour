package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DBModel 数据库连接对象
type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

// DBInfo 存储连接数据库的基本信息
type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

// TableColumn 存储COLUMNS表中需要的字段
type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

// DBTypeToStructType 数据表字段类型到go结构体类型映射
var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}

// NewDBModel 新建模型对象
func NewDBModel(dbInfo *DBInfo) *DBModel {
	return &DBModel{DBInfo: dbInfo}
}

// Connect 连接数据库
func (m *DBModel) Connect() error {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local",
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	// 打开information_schema数据库
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}

// GetColumns 获取tableName所有列
func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := "SELECT " +
		"COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
	// 查询数据库information_schema中columns列，筛选条件为待转换dbName中的数据表tableName，类似执行sql语句
	// select column_name from information_schema.columns where table_schema="blog_service" and table_name="blog_tag";
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	// if rows == nil {
	// 	return nil, errors.New("没有数据")
	// }
	defer rows.Close()

	var columns []*TableColumn
	// 遍历所有行
	for rows.Next() {
		var column TableColumn
		// 拷贝需要的字段
		err := rows.Scan(
			&column.ColumnName, &column.DataType,
			&column.ColumnKey, &column.IsNullable,
			&column.ColumnType, &column.ColumnComment,
		)
		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	// 判断列数是否为0
	if len(columns) == 0 {
		return nil, errors.New("没有数据")
	}

	return columns, nil
}
