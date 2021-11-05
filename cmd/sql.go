package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xielizyh/goprj-tour/internal/sql2struct"
)

var (
	username  string
	password  string
	host      string
	charset   string
	dbType    string
	dbName    string
	tableName string
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转换和处理",
	Long:  "sql转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql转换",
	Long:  "sql转换",
	Run: func(cmd *cobra.Command, args []string) {
		// 数据库基本信息
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		// 新建数据库模型
		dbModel := sql2struct.NewDBModel(dbInfo)
		// 连接数据库information_schema
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect error: %v", err)
		}
		// 获取拟转换数据库中数据表的所有列
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns error: %v", err)
		}
		// 新建转换模板
		template := sql2struct.NewStructTemplate()
		// 转换为结构体信息
		templateColumns := template.AssemblyColumns(columns)
		// 执行渲染
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate error: %v", err)
		}
	},
}

func init() {
	// 注册命令
	sqlCmd.AddCommand(sql2structCmd)
	// 添加选项
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "请输入数据库的账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "请输入数据库的密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库的HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库的编码")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入表名称")
}
