package prediction

import (
	"fmt"

	"github.com/YWJSonic/ReptileService/foundation"
)

// Analysis ...
func Analysis(Data InputData) {

	var AvgGroup []float32
	var MGroup [][]float32
	var AvgMGroup []float32

	AvgGroup = append(AvgGroup, (Data.DayOp+Data.DayCL)/2)
	AvgGroup = append(AvgGroup, MathAvg(Data.Day3))
	AvgGroup = append(AvgGroup, MathAvg(Data.Day7))
	AvgGroup = append(AvgGroup, MathAvg(Data.Day14))
	AvgGroup = append(AvgGroup, MathAvg(Data.Day30))
	AvgGroup = append(AvgGroup, MathAvg(Data.Day60))
	AvgGroup = append(AvgGroup, MathAvg(Data.Day90))
	AvgGroup = append(AvgGroup, MathAvg(Data.Day120))
	AvgGroup = append(AvgGroup, MathAvg(Data.Day180))
	AvgGroup = append(AvgGroup, MathAvg(Data.Day240))
	AvgGroup = append(AvgGroup, MathAvg(Data.Year1))

	MGroup = append(MGroup, MathM([]float32{Data.DayOp, Data.DayCL}))
	MGroup = append(MGroup, MathM(Data.Day3))
	MGroup = append(MGroup, MathM(Data.Day7))
	MGroup = append(MGroup, MathM(Data.Day14))
	MGroup = append(MGroup, MathM(Data.Day30))
	MGroup = append(MGroup, MathM(Data.Day60))
	MGroup = append(MGroup, MathM(Data.Day90))
	MGroup = append(MGroup, MathM(Data.Day120))
	MGroup = append(MGroup, MathM(Data.Day180))
	MGroup = append(MGroup, MathM(Data.Day240))
	MGroup = append(MGroup, MathM(Data.Year1))

	AvgMGroup = append(AvgMGroup, MGroup[0][0])
	AvgMGroup = append(AvgMGroup, MathAvg(MGroup[1]))
	AvgMGroup = append(AvgMGroup, MathAvg(MGroup[2]))
	AvgMGroup = append(AvgMGroup, MathAvg(MGroup[3]))
	AvgMGroup = append(AvgMGroup, MathAvg(MGroup[4]))
	AvgMGroup = append(AvgMGroup, MathAvg(MGroup[5]))
	AvgMGroup = append(AvgMGroup, MathAvg(MGroup[6]))
	AvgMGroup = append(AvgMGroup, MathAvg(MGroup[7]))
	AvgMGroup = append(AvgMGroup, MathAvg(MGroup[8]))
	AvgMGroup = append(AvgMGroup, MathAvg(MGroup[9]))
	AvgMGroup = append(AvgMGroup, MathAvg(MGroup[10]))

	Strs := foundation.Float32ArrayColor("red", AvgGroup...)
	fmt.Printf("Day:%v Day3:%v Day7:%v Day14:%v Day30:%v\n Day60:%v Day90:%v\n Day120:%v Day180:%v Day240:%v Year1:%v\n", Strs...)
	fmt.Println("------------------------------------------------------------------------")
	// fmt.Println(MGroup)
	// fmt.Println("------------------------------------------------------------------------")
	Strs = foundation.Float32ArrayColor("red", AvgMGroup...)
	fmt.Printf("Day:%v Day3:%v Day7:%v Day14:%v Day30:%v\n Day60:%v Day90:%v\n Day120:%v Day180:%v Day240:%v Year1:%v\n", Strs...)
	fmt.Println("#############################################################################################")
}

// MathAvg ...
func MathAvg(Data []float32) float32 {

	var AvgSumTemp float32
	for index, count := 0, len(Data); index < count; index++ {
		AvgSumTemp += Data[index]
	}
	return AvgSumTemp / (float32)(len(Data))
}

// MathM ...
func MathM(Data []float32) []float32 {
	var result []float32

	for index, count := 0, len(Data)-1; index < count; index++ {
		result = append(result, Data[index+1]-Data[index])
	}

	return result
}
