package models

type Lookup struct {
	LookupId    int    `json:"lookupid"`
	Custom1     string `json:"custom1,omitempty"`
	Custom1Bool bool   `json:"custom1_bool,omitempty"`
	Custom2     string `json:"custom2,omitempty"`
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Sequence    int    `json:"sequence,omitempty"`
	Value2      string `json:"value2,omitempty"`
	Value2Bool  bool   `json:"value2_bool,omitempty"`
	Value3      string `json:"value3,omitempty"`
	Value3Bool  bool   `json:"value3_bool,omitempty"`
	Value4      string `json:"value4,omitempty"`
	Value4Bool  bool   `json:"value4_bool,omitempty"`
	Value5      string `json:"value5,omitempty"`
	Value5Bool  bool   `json:"value5_bool,omitempty"`
	Value6      string `json:"value6,omitempty"`
	Value6Bool  bool   `json:"value6_bool,omitempty"`
	Value7      string `json:"value7,omitempty"`
	Value7Bool  bool   `json:"value7_bool,omitempty"`
	Valueint1   int    `json:"valueint1,omitempty"`
}
