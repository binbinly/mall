package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"strings"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "internal/dal",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		WithUnitTest: true,
	})

	gormdb, _ := gorm.Open(mysql.Open("root:root@(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	var dataMap = map[string]func(gorm.ColumnType) (dataType string){
		// int mapping
		"int": func(columnType gorm.ColumnType) (dataType string) {
			if n, ok := columnType.Nullable(); ok && n {
				return "*int"
			}
			return "int"
		},

		// bool mapping
		"tinyint": func(columnType gorm.ColumnType) (dataType string) {
			ct, _ := columnType.ColumnType()
			if strings.HasPrefix(ct, "tinyint(1)") {
				return "bool"
			}
			return "int8"
		},
	}

	g.WithDataTypeMap(dataMap)

	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
