package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Info struct {
	SynonymInfos string    `json:"synonym_infos"`
	SchemaInfos  []Columns `json:"schema_infos"`
	TableId      string    `json:"table_id"`
	TableDesc    string    `json:"table_desc"`
}

type Columns struct {
	ColCaption string `json:"col_caption"`
	ColName    string `json:"col_name"`
}

func main() {
	// 数据库连接信息
	dataSourceName := "root:fowWNObs5AlS1aHx@tcp(123.60.77.193:3306)/mexico_microloan"

	// 连接数据库
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 测试连接是否成功
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("成功连接到数据库")

	// 执行查询操作
	rows, err := db.Query("SELECT table_name,table_comment FROM information_schema.tables WHERE table_schema = 'mexico_microloan'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 遍历表信息
	for rows.Next() {
		var tableName, tableComment string
		if err := rows.Scan(&tableName, &tableComment); err != nil {
			log.Fatal(err)
		}

		// 查询每个表的字段信息
		fields, err := db.Query("SELECT column_name, column_comment FROM information_schema.columns WHERE table_schema = 'mexico_microloan' AND table_name = ?", tableName)
		if err != nil {
			log.Fatal(err)
		}
		defer fields.Close()

		// 遍历字段信息
		var col []Columns
		for fields.Next() {
			var columns Columns
			var columnName, columnComment string
			if err := fields.Scan(&columnName, &columnComment); err != nil {
				log.Fatal(err)
			}
			columns.ColName = columnName
			columns.ColCaption = columnComment
			col = append(col, columns)
		}
		var info Info
		info.TableDesc = tableComment
		info.SynonymInfos = tableName
		info.SchemaInfos = col

		jsonData, err := json.Marshal(info)
		if err != nil {
			fmt.Println("转换为JSON时出错:", err)
			return
		}
		//打印 json 数据
		fmt.Println(string(jsonData))
		fmt.Println("---------------------------------------------------------------------------------")

	}

}
