package utils

import "database/sql"

func GetConn() *sql.DB {
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/etraca?charset=utf8")
	CkeckError(err, "open db conn error ")
	return db
}

func PrepareStmt(sql string,db *sql.DB) *sql.Stmt {
	stmt, err := db.Prepare(sql)
	CkeckError(err,"prepare stmt error")
	return stmt
}