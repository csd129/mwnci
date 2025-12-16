package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"github.com/shirou/gopsutil/v4/mem"
)

func Memstat(args ...object.Object) object.Object {
	if err := typing.Check(
		"memstat", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}

	newHash := make(map[object.HashKey]object.HashPair)
	v, _ := mem.VirtualMemory()
	key := &object.String{Value: "total"}
	val := &object.Integer{Value: int64(v.Total)}
	newHashPair := object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "available"}
	val = &object.Integer{Value: int64(v.Available)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "used"}
	val = &object.Integer{Value: int64(v.Used)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "usedpercent"}
	fval := &object.Float{Value: float64(v.UsedPercent)}
	newHashPair = object.HashPair{Key: key, Value: fval}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "free"}
	val = &object.Integer{Value: int64(v.Free)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "active"}
	val = &object.Integer{Value: int64(v.Active)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "inactive"}
	val = &object.Integer{Value: int64(v.Inactive)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "wired"}
	val = &object.Integer{Value: int64(v.Wired)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "laundry"}
	val = &object.Integer{Value: int64(v.Laundry)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "buffers"}
	val = &object.Integer{Value: int64(v.Buffers)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "cached"}
	val = &object.Integer{Value: int64(v.Cached)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "writeback"}
	val = &object.Integer{Value: int64(v.WriteBack)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "writebacktmp"}
	val = &object.Integer{Value: int64(v.WriteBackTmp)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "dirty"}
	val = &object.Integer{Value: int64(v.Dirty)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "shared"}
	val = &object.Integer{Value: int64(v.Shared)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "slab"}
	val = &object.Integer{Value: int64(v.Slab)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "sreclaimable"}
	val = &object.Integer{Value: int64(v.Sreclaimable)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "sunreclaim"}
	val = &object.Integer{Value: int64(v.Sunreclaim)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "pagetables"}
	val = &object.Integer{Value: int64(v.PageTables)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "swapcached"}
	val = &object.Integer{Value: int64(v.SwapCached)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "commitlimit"}
	val = &object.Integer{Value: int64(v.CommitLimit)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "hightotal"}
	val = &object.Integer{Value: int64(v.HighTotal)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "highfree"}
	val = &object.Integer{Value: int64(v.HighFree)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "lowtotal"}
	val = &object.Integer{Value: int64(v.LowTotal)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "lowfree"}
	val = &object.Integer{Value: int64(v.LowFree)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "swaptotal"}
	val = &object.Integer{Value: int64(v.SwapTotal)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "swapfree"}
	val = &object.Integer{Value: int64(v.SwapFree)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "mapped"}
	val = &object.Integer{Value: int64(v.Mapped)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "vmalloctotal"}
	val = &object.Integer{Value: int64(v.VmallocTotal)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "vmallocused"}
	val = &object.Integer{Value: int64(v.VmallocUsed)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "vmallocchunk"}
	val = &object.Integer{Value: int64(v.VmallocChunk)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "hugepagestotal"}
	val = &object.Integer{Value: int64(v.HugePagesTotal)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "hugepagesfree"}
	val = &object.Integer{Value: int64(v.HugePagesFree)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "hugepagesrsvd"}
	val = &object.Integer{Value: int64(v.HugePagesRsvd)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "hugepagessurp"}
	val = &object.Integer{Value: int64(v.HugePagesSurp)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "hugepagesize"}
	val = &object.Integer{Value: int64(v.HugePageSize)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair
	key = &object.String{Value: "anonhugepages"}
	val = &object.Integer{Value: int64(v.AnonHugePages)}
	newHashPair = object.HashPair{Key: key, Value: val}
	newHash[key.HashKey()] = newHashPair

	return &object.Hash{Pairs: newHash}
}
