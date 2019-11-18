package miindex

// Info 每日收盤行情查詢
type Info struct {
	Date      string        `json:"date"`
	Stat      string        `json:"stat"`
	Title     string        `json:"title"`
	Params    miindexParams `json:"params,omitempty"`
	Fields    []string      `json:"fields,omitempty"`
	Fields1   []string      `json:"fields1,omitempty"`
	Fields2   []string      `json:"fields2,omitempty"`
	Fields3   []string      `json:"fields3,omitempty"`
	Fields4   []string      `json:"fields4,omitempty"`
	Fields5   []string      `json:"fields5,omitempty"`
	Fields6   []string      `json:"fields6,omitempty"`
	Fields7   []string      `json:"fields7,omitempty"`
	Fields8   []string      `json:"fields8,omitempty"`
	Fields9   []string      `json:"fields9,omitempty"`
	Subtitle  string        `json:"subtitle,omitempty"`
	Subtitle1 string        `json:"subtitle1,omitempty"`
	Subtitle2 string        `json:"subtitle2,omitempty"`
	Subtitle3 string        `json:"subtitle3,omitempty"`
	Subtitle4 string        `json:"subtitle4,omitempty"`
	Subtitle5 string        `json:"subtitle5,omitempty"`
	Subtitle6 string        `json:"subtitle6,omitempty"`
	Subtitle7 string        `json:"subtitle7,omitempty"`
	Subtitle8 string        `json:"subtitle8,omitempty"`
	Subtitle9 string        `json:"subtitle9,omitempty"`
	Data      []interface{} `json:"data,omitempty"`
	Data1     []interface{} `json:"data1,omitempty"`
	Data2     []interface{} `json:"data2,omitempty"`
	Data3     []interface{} `json:"data3,omitempty"`
	Data4     []interface{} `json:"data4,omitempty"`
	Data5     []interface{} `json:"data5,omitempty"`
	Data6     []interface{} `json:"data6,omitempty"`
	Data7     []interface{} `json:"data7,omitempty"`
	Data8     []interface{} `json:"data8,omitempty"`
	Data9     []interface{} `json:"data9,omitempty"`
	Notes     []string      `json:"notes,omitempty"`
	Notes1    []string      `json:"notes1,omitempty"`
	Notes2    []string      `json:"notes2,omitempty"`
	Notes3    []string      `json:"notes3,omitempty"`
	Notes4    []string      `json:"notes4,omitempty"`
	Notes5    []string      `json:"notes5,omitempty"`
	Notes6    []string      `json:"notes6,omitempty"`
	Notes7    []string      `json:"notes7,omitempty"`
	Notes8    []string      `json:"notes8,omitempty"`
	Notes9    []string      `json:"notes9,omitempty"`
}

type miindexParams struct {
	Response   string `json:"response"`
	Date       string `json:"date"`
	Type       string `json:"type"`
	Time       string `json:"_"`
	Controller string `json:"controller"`
	Format     string `json:"format"`
	Action     string `json:"action"`
	Lang       string `json:"lang"`
}
