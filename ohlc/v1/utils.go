package v1

import utils "github.com/xefino/protobuf-gen-go/utils"

// FrequencyAlternates contains alternative values for the OHLC.v1.Frequency enum
var FrequencyAlternates = map[string]Frequency{
	"":        Frequency_InvalidFrequency,
	"second":  Frequency_Second,
	"minute":  Frequency_Minute,
	"hour":    Frequency_Hour,
	"day":     Frequency_Day,
	"week":    Frequency_Week,
	"month":   Frequency_Month,
	"quarter": Frequency_Quarter,
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
