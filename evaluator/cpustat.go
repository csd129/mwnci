package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

func TwoPoint(Value float64) float64 {
	Stringy := fmt.Sprintf("%.2f", Value)
	Floaty, _ := strconv.ParseFloat(Stringy, 64)
	return Floaty
}

func Cpustat(args ...object.Object) object.Object {
	if err := typing.Check(
		"cpustat", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError("%s", err.Error())
	}

	Cpuinfo1, _ := cpu.Times(true)
	time.Sleep(time.Duration(1000) * time.Millisecond)
	Cpuinfo2, _ := cpu.Times(true)
	elements := make([]object.Object, 0)
	for Index := range Cpuinfo1 {
		newHash := make(map[object.HashKey]object.HashPair)
		CPUName := Cpuinfo1[Index].CPU
		UserC := Cpuinfo2[Index].User - Cpuinfo1[Index].User
		SystemC := Cpuinfo2[Index].System - Cpuinfo1[Index].System
		StealC := Cpuinfo2[Index].Steal - Cpuinfo1[Index].Steal
		SoftirqC := Cpuinfo2[Index].Softirq - Cpuinfo1[Index].Softirq
		NiceC := Cpuinfo2[Index].Nice - Cpuinfo1[Index].Nice
		IrqC := Cpuinfo2[Index].Irq - Cpuinfo1[Index].Irq
		IowaitC := Cpuinfo2[Index].Iowait - Cpuinfo1[Index].Iowait
		IdleC := Cpuinfo2[Index].Idle - Cpuinfo1[Index].Idle
		GuestNiceC := Cpuinfo2[Index].GuestNice - Cpuinfo1[Index].GuestNice
		GuestC := Cpuinfo2[Index].Guest - Cpuinfo1[Index].Guest
		TotalJiffies := UserC + NiceC + SystemC + IdleC + IowaitC + IrqC + SoftirqC + StealC
		BusyJiffies := TotalJiffies - IdleC
		AvCpu := TwoPoint((BusyJiffies / TotalJiffies) * 100)
		UserA := TwoPoint((UserC / TotalJiffies) * 100)
		SystemA := TwoPoint((SystemC / TotalJiffies) * 100)
		StealA := TwoPoint((StealC / TotalJiffies) * 100)
		SoftirqA := TwoPoint((SoftirqC / TotalJiffies) * 100)
		NiceA := TwoPoint((NiceC / TotalJiffies) * 100)
		IrqA := TwoPoint((IrqC / TotalJiffies) * 100)
		IowaitA := TwoPoint((IowaitC / TotalJiffies) * 100)
		IdleA := TwoPoint((IdleC / TotalJiffies) * 100)
		GuestNiceA := TwoPoint((GuestNiceC / TotalJiffies) * 100)
		GuestA := TwoPoint((GuestC / TotalJiffies) * 100)

		key := &object.String{Value: "cpu"}
		val := &object.String{Value: CPUName}
		newHashPair := object.HashPair{Key: key, Value: val}
		newHash[key.HashKey()] = newHashPair
		key = &object.String{Value: "user"}
		fval := &object.Float{Value: TwoPoint(UserA)}
		newHashPair = object.HashPair{Key: key, Value: fval}
		newHash[key.HashKey()] = newHashPair
		key = &object.String{Value: "system"}
		fval = &object.Float{Value: TwoPoint(SystemA)}
		newHashPair = object.HashPair{Key: key, Value: fval}
		newHash[key.HashKey()] = newHashPair
		key = &object.String{Value: "steal"}
		fval = &object.Float{Value: TwoPoint(StealA)}
		newHashPair = object.HashPair{Key: key, Value: fval}
		newHash[key.HashKey()] = newHashPair
		key = &object.String{Value: "softirq"}
		fval = &object.Float{Value: TwoPoint(SoftirqA)}
		newHashPair = object.HashPair{Key: key, Value: fval}
		newHash[key.HashKey()] = newHashPair
		key = &object.String{Value: "nice"}
		fval = &object.Float{Value: TwoPoint(NiceA)}
		newHashPair = object.HashPair{Key: key, Value: fval}
		newHash[key.HashKey()] = newHashPair
		key = &object.String{Value: "irq"}
		fval = &object.Float{Value: TwoPoint(IrqA)}
		newHashPair = object.HashPair{Key: key, Value: fval}
		newHash[key.HashKey()] = newHashPair
		key = &object.String{Value: "iowait"}
		fval = &object.Float{Value: TwoPoint(IowaitA)}
		newHashPair = object.HashPair{Key: key, Value: fval}
		newHash[key.HashKey()] = newHashPair
		key = &object.String{Value: "idle"}
		fval = &object.Float{Value: TwoPoint(IdleA)}
		newHashPair = object.HashPair{Key: key, Value: fval}
		newHash[key.HashKey()] = newHashPair
		key = &object.String{Value: "guestnice"}
		fval = &object.Float{Value: TwoPoint(GuestNiceA)}
		newHashPair = object.HashPair{Key: key, Value: fval}
		newHash[key.HashKey()] = newHashPair
		key = &object.String{Value: "guest"}
		fval = &object.Float{Value: TwoPoint(GuestA)}
		newHashPair = object.HashPair{Key: key, Value: fval}
		newHash[key.HashKey()] = newHashPair
		key = &object.String{Value: "busy"}
		fval = &object.Float{Value: TwoPoint(AvCpu)}
		newHashPair = object.HashPair{Key: key, Value: fval}
		newHash[key.HashKey()] = newHashPair

		elements = append(elements, &object.Hash{Pairs: newHash})
	}
	return &object.Array{Elements: elements}
}
