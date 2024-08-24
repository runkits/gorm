package main

import (
	"context"
	"fmt"

	"github.com/runkits/gorm/examples/biz"
	"github.com/runkits/gorm/examples/conf"
	"github.com/runkits/gorm/examples/dal"
	"github.com/runkits/gorm/examples/dal/query"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()
}

func main() {
	// start your project here
	fmt.Println("hello world")
	defer fmt.Println("bye~")

	query.SetDefault(dal.DB)
	biz.Query(context.Background())
}
