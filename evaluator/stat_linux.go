//go:build (linux || openbsd)
package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"os"
	"syscall"
)

func Stat(args ...object.Object) object.Object {
	if err := typing.Check(
		"", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	file := args[0].(*object.String).Value

	fileStat, err := os.Lstat(file)
	if err != nil {
		return NULL
	}
	fileSys := fileStat.Sys()
	fileMtime := &object.Integer{Value: int64(fileSys.(*syscall.Stat_t).Mtim.Sec)}
	fileUid := &object.Integer{Value: int64(fileSys.(*syscall.Stat_t).Uid)}
	fileGid := &object.Integer{Value: int64(fileSys.(*syscall.Stat_t).Gid)}
	fileCtime := &object.Integer{Value: int64(fileSys.(*syscall.Stat_t).Ctim.Sec)}
	fileAtime := &object.Integer{Value: int64(fileSys.(*syscall.Stat_t).Atim.Sec)}
	fileMode := fileSys.(*syscall.Stat_t).Mode & 07777
	fileLinks := &object.Integer{Value: int64(fileSys.(*syscall.Stat_t).Nlink)}
	fileIno := &object.Integer{Value: int64(fileSys.(*syscall.Stat_t).Ino)}
	fileBlkSiz := &object.Integer{Value: int64(fileSys.(*syscall.Stat_t).Blksize)}
	fileBlocks := &object.Integer{Value: int64(fileSys.(*syscall.Stat_t).Blocks)}
	fileSize := &object.Integer{Value: int64(fileStat.Size())}
	fileRdev := fileSys.(*syscall.Stat_t).Rdev
	fileDev:= fileSys.(*syscall.Stat_t).Dev
	devMajor := &object.Integer{Value: int64(fileRdev / 256)}
	devMinor := &object.Integer{Value: int64(fileRdev % 256)}
	fileMajor := &object.Integer{Value: int64(fileDev / 256)}
	fileMinor := &object.Integer{Value: int64(fileDev % 256)}
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
	Ftype := &object.String{Value: fileType}
	Mode := &object.String{Value: fmt.Sprintf("%04o", fileMode)}
	newHash := make(map[object.HashKey]object.HashPair)
	key := &object.String{Value: "mtime"}
	newHashPair := object.HashPair{Key: key, Value: fileMtime}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "uid"}
	newHashPair = object.HashPair{Key: key, Value: fileUid}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "mode"}
	newHashPair = object.HashPair{Key: key, Value: Mode}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "gid"}
	newHashPair = object.HashPair{Key: key, Value: fileGid}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "inode"}
	newHashPair = object.HashPair{Key: key, Value: fileIno}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "type"}
	newHashPair = object.HashPair{Key: key, Value: Ftype}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "ctime"}
	newHashPair = object.HashPair{Key: key, Value: fileCtime}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "atime"}
	newHashPair = object.HashPair{Key: key, Value: fileAtime}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "mtime"}
	newHashPair = object.HashPair{Key: key, Value: fileMtime}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "nlinks"}
	newHashPair = object.HashPair{Key: key, Value: fileLinks}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "blksize"}
	newHashPair = object.HashPair{Key: key, Value: fileBlkSiz}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "blocks"}
	newHashPair = object.HashPair{Key: key, Value: fileBlocks}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "size"}
	newHashPair = object.HashPair{Key: key, Value: fileSize}
	newHash[key.HashKey()] = newHashPair
	if fileType != "BLOCK" && fileType != "CHR" {
		key = &object.String{Value: "major"}
		newHashPair = object.HashPair{Key: key, Value: fileMajor}
		newHash[key.HashKey()] = newHashPair
		key = &object.String{Value: "minor"}
		newHashPair = object.HashPair{Key: key, Value: fileMinor}
		newHash[key.HashKey()] = newHashPair
	} else {
		key = &object.String{Value: "major"}
		newHashPair = object.HashPair{Key: key, Value: devMajor}
		newHash[key.HashKey()] = newHashPair
		key = &object.String{Value: "minor"}
		newHashPair = object.HashPair{Key: key, Value: devMinor}
		newHash[key.HashKey()] = newHashPair
	}
	return &object.Hash{Pairs: newHash}
}
