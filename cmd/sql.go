package cmd

import (
	"github.com/go-programming-tour-book/tour/internal/sql2struct"
	"github.com/spf13/cobra"
	"log"
)

// 声明cmd全局变量，用于接收外部的命令行参数
var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	// 声明sql子命令
	Use: "sql",
	Short: "sql转换和处理",
	Long: "sql转换和处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var sql2structCmd = &cobra.Command{
	// 声明sql子命令对应的子命令struct，
	// 并在sql2structCmd中完成了对数据库的查询、模板对象的组装、渲染等动作的流转，
	// 相当于一个控制器。
	Use: "struct",
	Short: "sql转换",
	Long: "sql转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType: dbType,
			Host: host,
			Username: username,
			Password: password,
			Charset: charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v\n", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v\n", err)
		}

		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v\n", err)
		}
	},
}

// 进行默认的cmd初始化动作和命令行参数的绑定
func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "请输入数据库账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "请输入数据库密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库IP:PORT")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库编码")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入表名称")
}