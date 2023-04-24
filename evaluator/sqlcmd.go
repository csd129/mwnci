package evaluator

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"mwnci/object"
	"mwnci/typing"
)

func SqlCmd(args ...object.Object) object.Object {
	if err := typing.Check(
		"sqlcmd", args,
		typing.ExactArgs(6),
	); err != nil {
		return newError(err.Error())
	}

	database := args[0].(*object.String).Value
	user := args[1].(*object.String).Value
	passwd := args[2].(*object.String).Value
	server := args[3].(*object.String).Value
	port := args[4].(*object.String).Value
	sqlcommand := args[5].(*object.String).Value

	port = fmt.Sprintf("%v", port)
	connect_string := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, passwd, server, port, database)
	db, err := sql.Open("mysql", connect_string)
	if err != nil {
		return newError(err.Error())
	}
	defer db.Close()

	res, err := db.Exec(sqlcommand)

	if err != nil {
		return newError(err.Error())
	}
	lastId, err := res.LastInsertId()
	return &object.Integer{Value: int64(lastId)}
}

func init() {
	RegisterBuiltin("sqlcmd",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (SqlCmd(args...))
		})
}
