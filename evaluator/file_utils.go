package evaluator

import (
	"fmt"
	"io"
	"mwnci/object"
	"mwnci/typing"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"syscall"
	"time"
)

// Change a mode of a file - note the second argument is a string
// to emphasise octal.
func chmodFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"chmod", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	path := args[0].Inspect()
	mode := ""

	switch args[1].(type) {
	case *object.String:
		mode = args[1].(*object.String).Value
	default:
		return newError("Second argument must be STRING, got %v", args[1])
	}

	// convert from octal -> decimal
	result, err := strconv.ParseInt(mode, 8, 64)
	if err != nil {
		//		return newError("Unable to create permissions with %v", args[1])
		fmt.Fprintf(os.Stderr, "Error calling chmod(): Unable to create permissions with %v\n", args[1])
		return FALSE
	}

	// Change the mode.
	err = os.Chmod(path, os.FileMode(result))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling chmod(): %v\n", err.Error())
		return FALSE
	}
	return TRUE
}

func Chown(args ...object.Object) object.Object {
	if err := typing.Check(
		"chown", args,
		typing.ExactArgs(3),
		typing.WithTypes(object.STRING_OBJ, object.INTEGER_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	path := args[0].Inspect()
	owner := int(args[1].(*object.Integer).Value)
	group := int(args[2].(*object.Integer).Value)

	err := os.Chown(path, owner, group)
	if err != nil {
		return newError(err.Error())
	}
	return TRUE
}

func FCp(args ...object.Object) object.Object {
	if err := typing.Check(
		"cp", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	orig := args[0].Inspect()
	dest := args[1].Inspect()

	source, err := os.Open(orig)
	if err != nil {
		return newError(err.Error())
	}
	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		return newError(err.Error())
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	if err != nil {
		return newError(err.Error())
	}
	return &object.Integer{Value: nBytes}
}

func mkdirFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"mkdir", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	path := args[0].(*object.String).Value
	mode, err := strconv.ParseInt("755", 8, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling chmod(): Unable to create permissions with 755\n")
		return FALSE
	}
	err = os.Mkdir(path, os.FileMode(mode))
	if err != nil && !os.IsExist(err) {
		fmt.Fprintf(os.Stderr, "Error calling mkdir(): %v\n", err.Error())
		return FALSE
	} else {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling mkdir(): %v\n", err.Error())
			return FALSE
		}
	}
	return TRUE
}

// Mkdir...
func Mkdirhier(args ...object.Object) object.Object {
	if err := typing.Check(
		"mkdirhier", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	directory := args[0].(*object.String).Value
	perms := args[1].(*object.String).Value
	mode, err := strconv.ParseInt(perms, 8, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling mkdirhier(): Unable to create permissions with %v\n", perms)
		return FALSE
	}
	err = os.MkdirAll(directory, os.FileMode(mode))
	if err != nil && !os.IsExist(err) {
		fmt.Fprintf(os.Stderr, "Error calling mkdirhier(): %v\n", err.Error())
		return FALSE
	}
	return TRUE
}

func Mkfifo(args ...object.Object) object.Object {
	if err := typing.Check(
		"mkfifo", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	path := args[0].(*object.String).Value
	err := syscall.Mkfifo(path, 0600)
	if err != nil && !os.IsExist(err) {
		return FALSE
	}
	return TRUE
}

func IsReadable(args ...object.Object) object.Object {
	if err := typing.Check(
		"isreadable", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	file := args[0].(*object.String).Value

	_, err := os.Open(file)
	if err != nil {
		return FALSE
	}
	return TRUE
}

func IsWriteable(args ...object.Object) object.Object {
	if err := typing.Check(
		"iswriteable", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	file := args[0].(*object.String).Value

	f, err := os.OpenFile(file, os.O_RDWR, 0666)
	if err != nil {
		return FALSE
	}
	f.Close()
	return TRUE
}

func Touch(args ...object.Object) object.Object {
	if err := typing.Check(
		"touch", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	path := args[0].Inspect()

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling touch(): %v\n", err)
			return FALSE
		}
		defer file.Close()
	} else {
		currentTime := time.Now().Local()
		err = os.Chtimes(path, currentTime, currentTime)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling touch(): %v\n", err)
			return FALSE
		}
	}
	return TRUE
}

func FindFile(args ...object.Object) object.Object {
	var files []string
	if err := typing.Check(
		"findfile", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	rootpath := args[0].(*object.String).Value
	fileInfo, err := os.Stat(rootpath)
	if err != nil {
		return newError("Unable to get status of %s", rootpath)
	}
	if fileInfo.IsDir() {
		filesearch := args[1].(*object.String).Value
		r, err := regexp.Compile(filesearch)
		if err != nil {
			return newError("failed to compile regexp %s: %s", filesearch, err)
		}

		filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
			basefile := filepath.Base(path)
			if r.MatchString(basefile) {
				files = append(files, path)
			}
			return nil
		})
		elements := make([]object.Object, len(files))
		if len(files) != 0 {
			for i, token := range files {
				elements[i] = &object.String{Value: token}
			}
		}
		return &object.Array{Elements: elements}
	}
	return newError("%s is not a directory", rootpath)
}

// File ...
func File(args ...object.Object) object.Object {
	if err := typing.Check(
		"file", args,
		typing.RangeOfArgs(1, 2),
		typing.WithTypes(object.STRING_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	file := args[0].(*object.String).Value

	fileStat, err := os.Lstat(file)
	if err != nil {
		return NULL
	}
	fileSys := fileStat.Sys()
	fileMtime := fileSys.(*syscall.Stat_t).Mtim.Sec
	if len(args) == 1 {
		return &object.Integer{Value: int64(fileMtime)}
	}
	fileUid := fileSys.(*syscall.Stat_t).Uid
	fileGid := fileSys.(*syscall.Stat_t).Gid
	fileCtime := fileSys.(*syscall.Stat_t).Ctim.Sec
	fileAtime := fileSys.(*syscall.Stat_t).Atim.Sec
	fileMode := fileSys.(*syscall.Stat_t).Mode & 07777
	fileLinks := fileSys.(*syscall.Stat_t).Nlink
	fileSize := fileStat.Size()
	var (
		fileType string
	)
	switch fileSys.(*syscall.Stat_t).Mode & syscall.S_IFMT {
	case syscall.S_IFREG:
		fileType = "FILE"
	case syscall.S_IFDIR:
		fileType = "DIR"
	case syscall.S_IFBLK:
		fileType = "BLOCK"
	case syscall.S_IFCHR:
		fileType = "CHR"
	case syscall.S_IFIFO:
		fileType = "FIFO"
	case syscall.S_IFLNK:
		fileType = "LINK"
	case syscall.S_IFSOCK:
		fileType = "SOCKET"
	default:
		fileType = "UNKNOWN"
	}
	value := fmt.Sprintf("%d %d %d %04o %d %d %s %d %d", fileMtime, fileAtime, fileCtime, fileMode, fileSize, fileLinks, fileType, fileUid, fileGid)
	return &object.String{Value: value}
}

func Glob(args ...object.Object) object.Object {
	if err := typing.Check(
		"glob", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	checkpath := args[0].(*object.String).Value
	files, err := filepath.Glob(checkpath)
	if err != nil {
		return newError(err.Error())
	}
	elements := make([]object.Object, len(files))

	for i, token := range files {
		elements[i] = &object.String{Value: token}
	}
	return &object.Array{Elements: elements}
}

func unlinkFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"unlink", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	path := args[0].Inspect()

	err := os.Remove(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling unlink(): %v\n", err.Error())
		return FALSE
	}
	return TRUE
}
