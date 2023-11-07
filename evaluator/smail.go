package evaluator

import (
	"crypto/tls"
	"mwnci/object"
	"mwnci/typing"
	"fmt"
	gomail "gopkg.in/mail.v2"
)

func Smail(args ...object.Object) object.Object {
	if err := typing.Check(
		"smail", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.HASH_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	Mail := gomail.NewMessage()
	Mhash := args[0].(*object.Hash)
	Message := args[1].Inspect()
	BodyType:=""
	Server:=""
	for _, v := range Mhash.Pairs {
		HKey := fmt.Sprintf("%v", v.Key)
		HVal := fmt.Sprintf("%v", v.Value)
		switch HKey {
		case "To":
			Mail.SetHeader(HKey, HVal)
		case "From":
			Mail.SetHeader(HKey, HVal)
		case "Subject":
			Mail.SetHeader(HKey, HVal)
		case "Server":
			Server = HVal
		case "Type":
			BodyType = HVal
		}
	}
	Mail.SetBody(BodyType, Message)
	d := gomail.NewDialer(Server, 25, "afry", "Yny5ddu")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(Mail); err != nil {
		return newError(err.Error())
	}
	return NULL
}

func init() {
	RegisterBuiltin("smail",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Smail(args...))
		})
}
