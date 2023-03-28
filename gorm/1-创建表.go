package main

type Student struct {
	Id   int
	Name string
	Age  int
}

func main() {
	//
	//dsn := "root:fowWNObs5AlS1aHx@tcp(123.60.77.193:3306)/mexico_microloan_admin"
	//mysqlConfig := mysql.Config{
	//	DSN:                       dsn,   // DSN data source name
	//	DefaultStringSize:         191,   // string 类型字段的默认长度
	//	DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	//	DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	//	DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	//	SkipInitializeWithVersion: false, // 根据版本自动配置
	//}
	//if db, err := gorm.Open("mysql", &mysqlConfig); err != nil {
	//	panic("failed to connect database")
	//	os.Exit(0)
	//}
	//
	////借助gorm创建表
	//panic(db.AutoMigrate(new(Student)).Error())
}
