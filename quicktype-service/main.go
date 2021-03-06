package main

import (
	"comm/define"
	"comm/logger"
	"comm/service/web"
	"quicktype-service/handler"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	hdl := handler.Handler{}
	srv := web.New(web.Name("quicktype"))
	srv.HandleFunc("/", hdl.Index)
	srv.HandleFunc("/logout", hdl.Logout)
	srv.HandleFunc("/type/tables", hdl.Tables)
	srv.HandleFunc("/type/table2go", hdl.Table2Go)
	srv.HandleFunc("/type/table2proto", hdl.Table2Proto)
	srv.HandleFunc("/type/table2handler", hdl.Table2Handler)
	srv.HandleFunc("/type/table2rw", hdl.Table2RW)
	srv.HandleFunc("/type/table2db", hdl.Table2DB)
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
