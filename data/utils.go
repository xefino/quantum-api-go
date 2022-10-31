package data

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/xefino/quantum-api-go/utils"
)

// ClientTypeAlternates contains alternative values for the DynamoDB.ClientType enum
var ClientTypeAlternates = map[string]ClientType{
	"web":   ClientType_Web,
	"cli":   ClientType_CLI,
	"other": ClientType_Other,
}

// Marshaler converts a ClientType to a DynamoDB attribute value
func (enum *ClientType) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{
		Value: enum.String(),
	}, nil
}

// UnmarshalDynamoDBAttributeValue converts a DynamoDB attribute value to a ClientType
func (enum *ClientType) UnmarshalDynamoDBAttributeValue(value types.AttributeValue) error {
	switch casted := value.(type) {
	case *types.AttributeValueMemberB:
		return utils.UnmarshalValue(casted.Value, ClientType_value, ClientTypeAlternates, enum)
	case *types.AttributeValueMemberN:
		return utils.UnmarshalString(casted.Value, ClientType_value, ClientTypeAlternates, enum)
	case *types.AttributeValueMemberNULL:
		return nil
	case *types.AttributeValueMemberS:
		return utils.UnmarshalString(casted.Value, ClientType_value, ClientTypeAlternates, enum)
	default:
		return fmt.Errorf("Attribute value of %T could not be converted to a dynamodb.ClientType", value)
	}
}
