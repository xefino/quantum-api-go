package data

import (
	"github.com/xefino/protobuf-gen-go/utils"
)

// ClientTypeAlternates contains alternative values for the DynamoDB.ClientType enum
var ClientTypeAlternates = map[string]ClientType{
	"":      ClientType_Invalid,
	"web":   ClientType_Web,
	"cli":   ClientType_CLI,
	"other": ClientType_Other,
}

// UnmarshalJSON converts JSON data into a Data.ClientType
func (enum *ClientType) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalValue(data, ClientType_value, ClientTypeAlternates, enum)
}
