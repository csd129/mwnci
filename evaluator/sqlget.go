
package evaluator

import (
	"fmt"
	"database/sql"
	"strings"
	
	"mwnci/object"
	"mwnci/typing"

	_ "github.com/go-sql-driver/mysql"
)

func SqlGet(args ...object.Object) object.Object {
	if err := typing.Check(
		"sqlget", args,
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
		
	connect_string := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, passwd, server, port, database)
	db, err := sql.Open("mysql", connect_string)
	if err != nil {
		return newError(err.Error())
	}
	defer db.Close()
	rows, err := db.Query(sqlcommand)
	if err != nil {
		return newError(err.Error())
	}
	columns, err := rows.Columns()
	if err != nil {
		return newError(err.Error())
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var alldata string
	var build string
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for _, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			build = fmt.Sprintf("%v\t%v",build, value)
		}
		build = strings.Trim(build, "\t")
		alldata = fmt.Sprintf("%s%s%s",alldata,build, "\n")
		build=""
	}
	if err = rows.Err(); err != nil {
		return newError(err.Error())
	}
	return &object.String{Value: strings.Trim(alldata, "\n")}
}

func init() {
	RegisterBuiltin("sqlget",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (SqlGet(args...))
		})
}
