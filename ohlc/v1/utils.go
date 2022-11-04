package v1

import utils "github.com/xefino/protobuf-gen-go/utils"

// FrequencyAlternates contains alternative values for the OHLC.v1.Frequency enum
var FrequencyAlternates = map[string]Frequency{
	"":        Frequency_InvalidFrequency,
	"s":       Frequency_Second,
	"S":       Frequency_Second,
	"second":  Frequency_Second,
	"m":       Frequency_Minute,
	"M":       Frequency_Minute,
	"minute":  Frequency_Minute,
	"h":       Frequency_Hour,
	"H":       Frequency_Hour,
	"hour":    Frequency_Hour,
	"d":       Frequency_Day,
	"D":       Frequency_Day,
	"day":     Frequency_Day,
	"w":       Frequency_Week,
	"W":       Frequency_Week,
	"week":    Frequency_Week,
	"mn":      Frequency_Month,
	"MN":      Frequency_Month,
	"month":   Frequency_Month,
	"q":       Frequency_Quarter,
	"Q":       Frequency_Quarter,
	"quarter": Frequency_Quarter,
	"y":       Frequency_Year,
	"Y":       Frequency_Year,
	"year":    Frequency_Year,
}

// UnmarshalJSON converts JSON data into an OHLC.v1.Frequency
func (enum *Frequency) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalValue(data, Frequency_value, FrequencyAlternates, enum)
}

// UnmarshalParam converts a form or query parameter into an OHLC.v1.Frequency
func (enum *Frequency) UnmarshalParam(param string) error {
	return utils.UnmarshalString(param, Frequency_value, FrequencyAlternates, enum)
}
