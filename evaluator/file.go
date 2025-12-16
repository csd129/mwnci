//go:build (linux || openbsd)
package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"os"
	"syscall"
)

// File ...
func File(args ...object.Object) object.Object {
	if err := typing.Check(
		"file", args,
		typing.RangeOfArgs(1, 2),
		typing.WithTypes(object.STRING_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError("%s", err.Error())
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
