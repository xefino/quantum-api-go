package data

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Create a new test runner we'll use to test all the
// modules in the data package
func TestData(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Data Suite")
}

var _ = Describe("DynamoDB.ClientType Marshal/Unmarshal Tests", func() {

	// Test that converting a DynamoDB.ClientType to a AttributeValue works for all values
	DescribeTable("MarshalDynamoDBAttributeValue Tests",
		func(enum ClientType, value string) {
			data, err := enum.MarshalDynamoDBAttributeValue()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data.(*types.AttributeValueMemberS).Value).Should(Equal(value))
		},
		Entry("Web - Works", ClientType_Web, "Web"),
		Entry("CLI - Works", ClientType_CLI, "CLI"),
		Entry("Other - Works", ClientType_Other, "Other"))

	// Tests that, if the UnmarshalDynamoDBAttributeValue function is called with an invalid AttributeValue
	// type, then the function will return an error
	It("UnmarshalDynamoDBAttributeValue - Type invalid - Error", func() {
		var enum *ClientType
		err := enum.UnmarshalDynamoDBAttributeValue(&types.AttributeValueMemberBOOL{Value: *aws.Bool(false)})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("Attribute value of *types.AttributeValueMemberBOOL could not be converted to a dynamodb.ClientType"))
	})

	// Tests that, if UnmarshalDynamoDBAttributeValue is called with a AttributeValueMemberNULL,
	// then the dynamodb.ClientType will not be modified and instead will be returned as nil
	It("UnmarshalDynamoDBAttributeValue - Value is NULL - Works", func() {
		var enum *ClientType
		err := enum.UnmarshalDynamoDBAttributeValue(&types.AttributeValueMemberNULL{})
		Expect(err).ShouldNot(HaveOccurred())
		Expect(enum).Should(BeNil())
	})

	// Tests that the UnmarshalDynamoDBAttributeValue works with various AttributeValue types
	DescribeTable("UnmarshalDynamoDBAttributeValue - Conditions",
		func(attr types.AttributeValue, expected ClientType) {
			enum := new(ClientType)
			err := enum.UnmarshalDynamoDBAttributeValue(attr)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(*enum).Should(Equal(expected))
		},
		Entry("Value is web, []byte - Works",
			&types.AttributeValueMemberB{Value: []byte("web")}, ClientType_Web),
		Entry("Value is cli, []byte - Works",
			&types.AttributeValueMemberB{Value: []byte("cli")}, ClientType_CLI),
		Entry("Value is other, []byte - Works",
			&types.AttributeValueMemberB{Value: []byte("other")}, ClientType_Other),
		Entry("Value is 0, []byte - Works",
			&types.AttributeValueMemberB{Value: []byte("0")}, ClientType_Web),
		Entry("Value is 1, []byte - Works",
			&types.AttributeValueMemberB{Value: []byte("1")}, ClientType_CLI),
		Entry("Value is 2, []byte - Works",
			&types.AttributeValueMemberB{Value: []byte("2")}, ClientType_Other),
		Entry("Value is Web, []byte - Works",
			&types.AttributeValueMemberB{Value: []byte("Web")}, ClientType_Web),
		Entry("Value is CLI, []byte - Works",
			&types.AttributeValueMemberB{Value: []byte("CLI")}, ClientType_CLI),
		Entry("Value is Other, []byte - Works",
			&types.AttributeValueMemberB{Value: []byte("Other")}, ClientType_Other),
		Entry("Value is 0, number - Works",
			&types.AttributeValueMemberN{Value: "0"}, ClientType_Web),
		Entry("Value is 1, number - Works",
			&types.AttributeValueMemberN{Value: "1"}, ClientType_CLI),
		Entry("Value is 2, number - Works",
			&types.AttributeValueMemberN{Value: "2"}, ClientType_Other),
		Entry("Value is web, string - Works",
			&types.AttributeValueMemberS{Value: "web"}, ClientType_Web),
		Entry("Value is cli, string - Works",
			&types.AttributeValueMemberS{Value: "cli"}, ClientType_CLI),
		Entry("Value is other, string - Works",
			&types.AttributeValueMemberS{Value: "other"}, ClientType_Other),
		Entry("Value is Web, string - Works",
			&types.AttributeValueMemberS{Value: "Web"}, ClientType_Web),
		Entry("Value is CLI, string - Works",
			&types.AttributeValueMemberS{Value: "CLI"}, ClientType_CLI),
		Entry("Value is Other, string - Works",
			&types.AttributeValueMemberS{Value: "Other"}, ClientType_Other),
		Entry("Value is 0, string - Works",
			&types.AttributeValueMemberS{Value: "0"}, ClientType_Web),
		Entry("Value is 1, string - Works",
			&types.AttributeValueMemberS{Value: "1"}, ClientType_CLI),
		Entry("Value is 2, string - Works",
			&types.AttributeValueMemberS{Value: "2"}, ClientType_Other))
})
