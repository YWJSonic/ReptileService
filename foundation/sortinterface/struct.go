package sortinterface

// Int64Slice Sotr int64 interface
type Int64Slice []int64

func (i64 Int64Slice) Len() int           { return len(i64) }
func (i64 Int64Slice) Swap(i, j int)      { i64[i], i64[j] = i64[j], i64[i] }
func (i64 Int64Slice) Less(i, j int) bool { return i64[i] < i64[j] }
