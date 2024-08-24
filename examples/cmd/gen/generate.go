package main

import (
	"github.com/runkits/gorm/examples/conf"
	"github.com/runkits/gorm/examples/dal"
	"gorm.io/gen"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()

	prepare(dal.DB) // prepare table for generate
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
	})

	g.UseDB(dal.DB)

	// generate all table from database
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
