package evaluator

import (
	"fmt"
	"log"
	"log/syslog"
	"mwnci/object"
	"mwnci/typing"
	"strings"
)

func Syslog(args ...object.Object) object.Object {
	if err := typing.Check(
		"syslog", args,
		typing.ExactArgs(3),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	var (
		logwriter *syslog.Writer
	)
	Level := strings.ToUpper(args[0].(*object.String).Value)
	Message := args[1].(*object.String).Value
	Tag := args[2].(*object.String).Value
	TypeTag := fmt.Sprintf("%s %s", Level, Tag)
	switch Level {
	case "ALERT":
		logwriter, _ = syslog.New(syslog.LOG_ALERT, TypeTag)
	case "CRIT":
		logwriter, _ = syslog.New(syslog.LOG_CRIT, TypeTag)
	case "DEBUG":
		logwriter, _ = syslog.New(syslog.LOG_DEBUG, TypeTag)
	case "EMERG":
		logwriter, _ = syslog.New(syslog.LOG_EMERG, TypeTag)
	case "ERR":
		logwriter, _ = syslog.New(syslog.LOG_ERR, TypeTag)
	case "INFO":
		logwriter, _ = syslog.New(syslog.LOG_INFO, TypeTag)
	case "NOTICE":
		logwriter, _ = syslog.New(syslog.LOG_NOTICE, TypeTag)
	case "WARNING":
		logwriter, _ = syslog.New(syslog.LOG_WARNING, TypeTag)
	}

	log.SetOutput(logwriter)
	log.Print(Message)
	return NULL
}
