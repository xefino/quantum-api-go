package data

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
)

// Create a new test runner we'll use to test all the
// modules in the data package
func TestData(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Data Suite")
}

var _ = Describe("Data.ClientType Marshal/Unmarshal Tests", func() {

	// Test that attempting to deserialize a Data.ClientType will fail and
	// return an error if the value canno be deserialized from a JSON value to a string
	It("UnmarshalJSON fails - Error", func() {

		// Attempt to convert a non-parseable string value into a Data.ClientType
		// This should return an error
		enum := new(ClientType)
		err := enum.UnmarshalJSON([]byte("derp"))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.ClientType"))
	})

	// Test that attempting to deserialize a Data.ClientType will fail and
	// return an error if the value cannot be converted to either the name value or integer
	// value of the enum option
	It("UnmarshalJSON - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a Data.ClientType
		// This should return an error
		enum := new(ClientType)
		err := enum.UnmarshalJSON([]byte("\"derp\""))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.ClientType"))
	})

	// Test the conditions under which values should be convertible to a Data.ClientType
	DescribeTable("UnmarshalJSON Tests",
		func(value string, shouldBe ClientType) {

			// Attempt to convert the string value into a Data.ClientType
			// This should not fail
			var enum ClientType
			err := enum.UnmarshalJSON([]byte(value))

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Empty string - Works", "\"\"", ClientType_Invalid),
		Entry("web - Works", "\"web\"", ClientType_Web),
		Entry("cli - Works", "\"cli\"", ClientType_CLI),
		Entry("other - Works", "\"other\"", ClientType_Other),
		Entry("Invalid - Works", "\"Invalid\"", ClientType_Invalid),
		Entry("Web - Works", "\"Web\"", ClientType_Web),
		Entry("CLI - Works", "\"CLI\"", ClientType_CLI),
		Entry("Other - Works", "\"Other\"", ClientType_Other),
		Entry("0 - Works", "\"0\"", ClientType_Invalid),
		Entry("1 - Works", "\"1\"", ClientType_Web),
		Entry("2 - Works", "\"2\"", ClientType_CLI),
		Entry("3 - Works", "\"3\"", ClientType_Other))
})

var _ = Describe("data.Frequency Marshal/Unmarshal Tests", func() {

	// Test that converting the data.Frequency enum to JSON works for all values
	DescribeTable("MarshalJSON Tests",
		func(enum Frequency, value string) {
			data, err := json.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("InvalidFrequency - Works", Frequency_InvalidFrequency, "\"\""),
		Entry("Second - Works", Frequency_Second, "\"s\""),
		Entry("Minute - Works", Frequency_Minute, "\"m\""),
		Entry("Hour - Works", Frequency_Hour, "\"h\""),
		Entry("Day - Works", Frequency_Day, "\"d\""),
		Entry("Week - Works", Frequency_Week, "\"w\""),
		Entry("Month - Works", Frequency_Month, "\"mn\""),
		Entry("Quarter - Works", Frequency_Quarter, "\"q\""),
		Entry("Year - Works", Frequency_Year, "\"y\""))

	// Test that converting the data.Frequency enum to a CSV column works for all values
	DescribeTable("MarshalCSV Tests",
		func(enum Frequency, value string) {
			data, err := enum.MarshalCSV()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("InvalidFrequency - Works", Frequency_InvalidFrequency, ""),
		Entry("Second - Works", Frequency_Second, "s"),
		Entry("Minute - Works", Frequency_Minute, "m"),
		Entry("Hour - Works", Frequency_Hour, "h"),
		Entry("Day - Works", Frequency_Day, "d"),
		Entry("Week - Works", Frequency_Week, "w"),
		Entry("Month - Works", Frequency_Month, "mn"),
		Entry("Quarter - Works", Frequency_Quarter, "q"),
		Entry("Year - Works", Frequency_Year, "y"))

	// Test that converting the data.Frequency enum to a YAML node works for all values
	DescribeTable("MarshalYAML - Works",
		func(enum Frequency, value string) {
			data, err := enum.MarshalYAML()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("InvalidFrequency - Works", Frequency_InvalidFrequency, ""),
		Entry("Second - Works", Frequency_Second, "s"),
		Entry("Minute - Works", Frequency_Minute, "m"),
		Entry("Hour - Works", Frequency_Hour, "h"),
		Entry("Day - Works", Frequency_Day, "d"),
		Entry("Week - Works", Frequency_Week, "w"),
		Entry("Month - Works", Frequency_Month, "mn"),
		Entry("Quarter - Works", Frequency_Quarter, "q"),
		Entry("Year - Works", Frequency_Year, "y"))

	// Test that converting the data.Frequency enum to a DynamoDB AttributeVAlue works for all values
	DescribeTable("MarshalDynamoDBAttributeValue - Works",
		func(enum Frequency, value string) {
			data, err := attributevalue.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data.(*types.AttributeValueMemberS).Value).Should(Equal(value))
		},
		Entry("InvalidFrequency - Works", Frequency_InvalidFrequency, ""),
		Entry("Second - Works", Frequency_Second, "s"),
		Entry("Minute - Works", Frequency_Minute, "m"),
		Entry("Hour - Works", Frequency_Hour, "h"),
		Entry("Day - Works", Frequency_Day, "d"),
		Entry("Week - Works", Frequency_Week, "w"),
		Entry("Month - Works", Frequency_Month, "mn"),
		Entry("Quarter - Works", Frequency_Quarter, "q"),
		Entry("Year - Works", Frequency_Year, "y"))

	// Test that converting the data.Frequency enum to an SQL value for all values
	DescribeTable("Value Tests",
		func(enum Frequency, value string) {
			data, err := enum.Value()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("InvalidFrequency - Works", Frequency_InvalidFrequency, ""),
		Entry("Second - Works", Frequency_Second, "s"),
		Entry("Minute - Works", Frequency_Minute, "m"),
		Entry("Hour - Works", Frequency_Hour, "h"),
		Entry("Day - Works", Frequency_Day, "d"),
		Entry("Week - Works", Frequency_Week, "w"),
		Entry("Month - Works", Frequency_Month, "mn"),
		Entry("Quarter - Works", Frequency_Quarter, "q"),
		Entry("Year - Works", Frequency_Year, "y"))

	// Test that attempting to deserialize a data.Frequency will fail and return an error if the value
	// canno be deserialized from a JSON value to a string
	It("UnmarshalJSON fails - Error", func() {

		// Attempt to convert a non-parseable string value into a data.Frequency; this should return an error
		enum := new(Frequency)
		err := enum.UnmarshalJSON([]byte("derp"))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Frequency"))
	})

	// Test that attempting to deserialize a data.Frequency will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalJSON - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a data.Frequency; this should return an error
		enum := new(Frequency)
		err := enum.UnmarshalJSON([]byte("\"derp\""))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Frequency"))
	})

	// Test the conditions under which values should be convertible to a data.Frequency
	DescribeTable("UnmarshalJSON Tests",
		func(value string, shouldBe Frequency) {

			// Attempt to convert the string value into a data.Frequency; this should not fail
			var enum Frequency
			err := enum.UnmarshalJSON([]byte(value))

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Empty string - Works", "\"\"", Frequency_InvalidFrequency),
		Entry("s - Works", "\"s\"", Frequency_Second),
		Entry("second - Works", "\"second\"", Frequency_Second),
		Entry("m - Works", "\"m\"", Frequency_Minute),
		Entry("minute - Works", "\"minute\"", Frequency_Minute),
		Entry("h - Works", "\"h\"", Frequency_Hour),
		Entry("hour - Works", "\"hour\"", Frequency_Hour),
		Entry("d - Works", "\"d\"", Frequency_Day),
		Entry("day - Works", "\"day\"", Frequency_Day),
		Entry("w - Works", "\"w\"", Frequency_Week),
		Entry("week - Works", "\"week\"", Frequency_Week),
		Entry("mn - Works", "\"mn\"", Frequency_Month),
		Entry("month - Works", "\"month\"", Frequency_Month),
		Entry("q - Works", "\"q\"", Frequency_Quarter),
		Entry("quarter - Works", "\"quarter\"", Frequency_Quarter),
		Entry("y - Works", "\"y\"", Frequency_Year),
		Entry("year - Works", "\"year\"", Frequency_Year),
		Entry("InvalidFrequency - Works", "\"InvalidFrequency\"", Frequency_InvalidFrequency),
		Entry("S - Works", "\"S\"", Frequency_Second),
		Entry("Second - Works", "\"Second\"", Frequency_Second),
		Entry("M - Works", "\"M\"", Frequency_Minute),
		Entry("Minute - Works", "\"Minute\"", Frequency_Minute),
		Entry("H - Works", "\"H\"", Frequency_Hour),
		Entry("Hour - Works", "\"Hour\"", Frequency_Hour),
		Entry("D - Works", "\"D\"", Frequency_Day),
		Entry("Day - Works", "\"Day\"", Frequency_Day),
		Entry("W - Works", "\"W\"", Frequency_Week),
		Entry("Week - Works", "\"Week\"", Frequency_Week),
		Entry("MN - Works", "\"MN\"", Frequency_Month),
		Entry("Month - Works", "\"Month\"", Frequency_Month),
		Entry("Q - Works", "\"Q\"", Frequency_Quarter),
		Entry("Quarter - Works", "\"Quarter\"", Frequency_Quarter),
		Entry("Y - Works", "\"Y\"", Frequency_Year),
		Entry("Year - Works", "\"Year\"", Frequency_Year),
		Entry("0 - Works", "\"0\"", Frequency_InvalidFrequency),
		Entry("1 - Works", "\"1\"", Frequency_Second),
		Entry("2 - Works", "\"2\"", Frequency_Minute),
		Entry("3 - Works", "\"3\"", Frequency_Hour),
		Entry("4 - Works", "\"4\"", Frequency_Day),
		Entry("5 - Works", "\"5\"", Frequency_Week),
		Entry("6 - Works", "\"6\"", Frequency_Month),
		Entry("7 - Works", "\"7\"", Frequency_Quarter),
		Entry("8 - Works", "\"8\"", Frequency_Year))

	// Test the conditions under which values should be convertible to a data.Frequency
	DescribeTable("UnmarshalCSV Tests",
		func(value string, shouldBe Frequency) {

			// Attempt to convert the value into a data.Frequency; this should not fail
			var enum Frequency
			err := enum.UnmarshalCSV(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Empty string - Works", "", Frequency_InvalidFrequency),
		Entry("s - Works", "s", Frequency_Second),
		Entry("second - Works", "second", Frequency_Second),
		Entry("m - Works", "m", Frequency_Minute),
		Entry("minute - Works", "minute", Frequency_Minute),
		Entry("h - Works", "h", Frequency_Hour),
		Entry("hour - Works", "hour", Frequency_Hour),
		Entry("d - Works", "d", Frequency_Day),
		Entry("day - Works", "day", Frequency_Day),
		Entry("w - Works", "w", Frequency_Week),
		Entry("week - Works", "week", Frequency_Week),
		Entry("mn - Works", "mn", Frequency_Month),
		Entry("month - Works", "month", Frequency_Month),
		Entry("q - Works", "q", Frequency_Quarter),
		Entry("quarter - Works", "quarter", Frequency_Quarter),
		Entry("y - Works", "y", Frequency_Year),
		Entry("year - Works", "year", Frequency_Year),
		Entry("InvalidFrequency - Works", "InvalidFrequency", Frequency_InvalidFrequency),
		Entry("S - Works", "S", Frequency_Second),
		Entry("Second - Works", "Second", Frequency_Second),
		Entry("M - Works", "M", Frequency_Minute),
		Entry("Minute - Works", "Minute", Frequency_Minute),
		Entry("H - Works", "H", Frequency_Hour),
		Entry("Hour - Works", "Hour", Frequency_Hour),
		Entry("D - Works", "D", Frequency_Day),
		Entry("Day - Works", "Day", Frequency_Day),
		Entry("W - Works", "W", Frequency_Week),
		Entry("Week - Works", "Week", Frequency_Week),
		Entry("MN - Works", "MN", Frequency_Month),
		Entry("Month - Works", "Month", Frequency_Month),
		Entry("Q - Works", "Q", Frequency_Quarter),
		Entry("Quarter - Works", "Quarter", Frequency_Quarter),
		Entry("Y - Works", "Y", Frequency_Year),
		Entry("Year - Works", "Year", Frequency_Year),
		Entry("0 - Works", "0", Frequency_InvalidFrequency),
		Entry("1 - Works", "1", Frequency_Second),
		Entry("2 - Works", "2", Frequency_Minute),
		Entry("3 - Works", "3", Frequency_Hour),
		Entry("4 - Works", "4", Frequency_Day),
		Entry("5 - Works", "5", Frequency_Week),
		Entry("6 - Works", "6", Frequency_Month),
		Entry("7 - Works", "7", Frequency_Quarter),
		Entry("8 - Works", "8", Frequency_Year))

	// Test that attempting to deserialize a data.Frequency will fail and return an error if the YAML
	// node does not represent a scalar value
	It("UnmarshalYAML - Node type is not scalar - Error", func() {
		enum := new(Frequency)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.AliasNode})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("YAML node had an invalid kind (expected scalar value)"))
	})

	// Test that attempting to deserialize a data.Frequency will fail and return an error if the YAML
	// node value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalYAML - Parse fails - Error", func() {
		enum := new(Frequency)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: "derp"})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Frequency"))
	})

	// Test the conditions under which YAML node values should be convertible to a data.Frequency
	DescribeTable("UnmarshalYAML Tests",
		func(value string, shouldBe Frequency) {
			var enum Frequency
			err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: value})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Empty string - Works", "", Frequency_InvalidFrequency),
		Entry("s - Works", "s", Frequency_Second),
		Entry("second - Works", "second", Frequency_Second),
		Entry("m - Works", "m", Frequency_Minute),
		Entry("minute - Works", "minute", Frequency_Minute),
		Entry("h - Works", "h", Frequency_Hour),
		Entry("hour - Works", "hour", Frequency_Hour),
		Entry("d - Works", "d", Frequency_Day),
		Entry("day - Works", "day", Frequency_Day),
		Entry("w - Works", "w", Frequency_Week),
		Entry("week - Works", "week", Frequency_Week),
		Entry("mn - Works", "mn", Frequency_Month),
		Entry("month - Works", "month", Frequency_Month),
		Entry("q - Works", "q", Frequency_Quarter),
		Entry("quarter - Works", "quarter", Frequency_Quarter),
		Entry("y - Works", "y", Frequency_Year),
		Entry("year - Works", "year", Frequency_Year),
		Entry("InvalidFrequency - Works", "InvalidFrequency", Frequency_InvalidFrequency),
		Entry("S - Works", "S", Frequency_Second),
		Entry("Second - Works", "Second", Frequency_Second),
		Entry("M - Works", "M", Frequency_Minute),
		Entry("Minute - Works", "Minute", Frequency_Minute),
		Entry("H - Works", "H", Frequency_Hour),
		Entry("Hour - Works", "Hour", Frequency_Hour),
		Entry("D - Works", "D", Frequency_Day),
		Entry("Day - Works", "Day", Frequency_Day),
		Entry("W - Works", "W", Frequency_Week),
		Entry("Week - Works", "Week", Frequency_Week),
		Entry("MN - Works", "MN", Frequency_Month),
		Entry("Month - Works", "Month", Frequency_Month),
		Entry("Q - Works", "Q", Frequency_Quarter),
		Entry("Quarter - Works", "Quarter", Frequency_Quarter),
		Entry("Y - Works", "Y", Frequency_Year),
		Entry("Year - Works", "Year", Frequency_Year),
		Entry("0 - Works", "0", Frequency_InvalidFrequency),
		Entry("1 - Works", "1", Frequency_Second),
		Entry("2 - Works", "2", Frequency_Minute),
		Entry("3 - Works", "3", Frequency_Hour),
		Entry("4 - Works", "4", Frequency_Day),
		Entry("5 - Works", "5", Frequency_Week),
		Entry("6 - Works", "6", Frequency_Month),
		Entry("7 - Works", "7", Frequency_Quarter),
		Entry("8 - Works", "8", Frequency_Year))

	// Test that attempting to deserialize a data.Frequency will fail and return an error if the value
	// canno be deserialized from a param value to a string
	It("UnmarshalParam fails - Error", func() {

		// Attempt to convert a non-parseable string value into a data.Frequency; this should return an error
		enum := new(Frequency)
		err := enum.UnmarshalParam("derp")

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Frequency"))
	})

	// Test that attempting to deserialize a data.Frequency will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalParam - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a data.Frequency; this should return an error
		enum := new(Frequency)
		err := enum.UnmarshalParam("derp")

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Frequency"))
	})

	// Test the conditions under which values should be convertible to a data.Frequency
	DescribeTable("UnmarshalParam Tests",
		func(value string, shouldBe Frequency) {

			// Attempt to convert the string value into a data.Frequency; this should not fail
			var enum Frequency
			err := enum.UnmarshalParam(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Empty string - Works", "", Frequency_InvalidFrequency),
		Entry("s - Works", "s", Frequency_Second),
		Entry("second - Works", "second", Frequency_Second),
		Entry("m - Works", "m", Frequency_Minute),
		Entry("minute - Works", "minute", Frequency_Minute),
		Entry("h - Works", "h", Frequency_Hour),
		Entry("hour - Works", "hour", Frequency_Hour),
		Entry("d - Works", "d", Frequency_Day),
		Entry("day - Works", "day", Frequency_Day),
		Entry("w - Works", "w", Frequency_Week),
		Entry("week - Works", "week", Frequency_Week),
		Entry("mn - Works", "mn", Frequency_Month),
		Entry("month - Works", "month", Frequency_Month),
		Entry("q - Works", "q", Frequency_Quarter),
		Entry("quarter - Works", "quarter", Frequency_Quarter),
		Entry("y - Works", "y", Frequency_Year),
		Entry("year - Works", "year", Frequency_Year),
		Entry("InvalidFrequency - Works", "InvalidFrequency", Frequency_InvalidFrequency),
		Entry("S - Works", "S", Frequency_Second),
		Entry("Second - Works", "Second", Frequency_Second),
		Entry("M - Works", "M", Frequency_Minute),
		Entry("Minute - Works", "Minute", Frequency_Minute),
		Entry("H - Works", "H", Frequency_Hour),
		Entry("Hour - Works", "Hour", Frequency_Hour),
		Entry("D - Works", "D", Frequency_Day),
		Entry("Day - Works", "Day", Frequency_Day),
		Entry("W - Works", "W", Frequency_Week),
		Entry("Week - Works", "Week", Frequency_Week),
		Entry("MN - Works", "MN", Frequency_Month),
		Entry("Month - Works", "Month", Frequency_Month),
		Entry("Q - Works", "Q", Frequency_Quarter),
		Entry("Quarter - Works", "Quarter", Frequency_Quarter),
		Entry("Y - Works", "Y", Frequency_Year),
		Entry("Year - Works", "Year", Frequency_Year),
		Entry("0 - Works", "0", Frequency_InvalidFrequency),
		Entry("1 - Works", "1", Frequency_Second),
		Entry("2 - Works", "2", Frequency_Minute),
		Entry("3 - Works", "3", Frequency_Hour),
		Entry("4 - Works", "4", Frequency_Day),
		Entry("5 - Works", "5", Frequency_Week),
		Entry("6 - Works", "6", Frequency_Month),
		Entry("7 - Works", "7", Frequency_Quarter),
		Entry("8 - Works", "8", Frequency_Year))

	// Tests that, if the attribute type submitted to UnmarshalDynamoDBAttributeValue is not one we
	// recognize, then the function will return an error
	It("UnmarshalDynamoDBAttributeValue - AttributeValue type invalid - Error", func() {
		enum := new(Frequency)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberBOOL{Value: true}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("Attribute value of *types.AttributeValueMemberBOOL could not be converted to a data.Frequency"))
	})

	// Tests that, if time parsing fails, then calling UnmarshalDynamoDBAttributeValue will return an error
	It("UnmarshalDynamoDBAttributeValue - Parse fails - Error", func() {
		enum := new(Frequency)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberS{Value: "derp"}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Frequency"))
	})

	// Tests the conditions under which UnmarshalDynamoDBAttributeValue is called and no error is generated
	DescribeTable("UnmarshalDynamoDBAttributeValue - AttributeValue Conditions",
		func(value types.AttributeValue, expected Frequency) {
			var enum Frequency
			err := attributevalue.Unmarshal(value, &enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(expected))
		},
		Entry("Value is []bytes, Empty string - Works",
			&types.AttributeValueMemberB{Value: []byte("")}, Frequency_InvalidFrequency),
		Entry("Value is []bytes, s - Works",
			&types.AttributeValueMemberB{Value: []byte("s")}, Frequency_Second),
		Entry("Value is []bytes, second - Works",
			&types.AttributeValueMemberB{Value: []byte("second")}, Frequency_Second),
		Entry("Value is []bytes, m - Works",
			&types.AttributeValueMemberB{Value: []byte("m")}, Frequency_Minute),
		Entry("Value is []bytes, minute - Works",
			&types.AttributeValueMemberB{Value: []byte("minute")}, Frequency_Minute),
		Entry("Value is []bytes, h - Works",
			&types.AttributeValueMemberB{Value: []byte("h")}, Frequency_Hour),
		Entry("Value is []bytes, hour - Works",
			&types.AttributeValueMemberB{Value: []byte("hour")}, Frequency_Hour),
		Entry("Value is []bytes, d - Works",
			&types.AttributeValueMemberB{Value: []byte("d")}, Frequency_Day),
		Entry("Value is []bytes, day - Works",
			&types.AttributeValueMemberB{Value: []byte("day")}, Frequency_Day),
		Entry("Value is []bytes, w - Works",
			&types.AttributeValueMemberB{Value: []byte("w")}, Frequency_Week),
		Entry("Value is []bytes, week - Works",
			&types.AttributeValueMemberB{Value: []byte("week")}, Frequency_Week),
		Entry("Value is []bytes, mn - Works",
			&types.AttributeValueMemberB{Value: []byte("mn")}, Frequency_Month),
		Entry("Value is []bytes, month - Works",
			&types.AttributeValueMemberB{Value: []byte("month")}, Frequency_Month),
		Entry("Value is []bytes, q - Works",
			&types.AttributeValueMemberB{Value: []byte("q")}, Frequency_Quarter),
		Entry("Value is []bytes, quarter - Works",
			&types.AttributeValueMemberB{Value: []byte("quarter")}, Frequency_Quarter),
		Entry("Value is []bytes, y - Works",
			&types.AttributeValueMemberB{Value: []byte("y")}, Frequency_Year),
		Entry("Value is []bytes, year - Works",
			&types.AttributeValueMemberB{Value: []byte("year")}, Frequency_Year),
		Entry("Value is []bytes, InvalidFrequency - Works",
			&types.AttributeValueMemberB{Value: []byte("InvalidFrequency")}, Frequency_InvalidFrequency),
		Entry("Value is []bytes, S - Works",
			&types.AttributeValueMemberB{Value: []byte("S")}, Frequency_Second),
		Entry("Value is []bytes, Second - Works",
			&types.AttributeValueMemberB{Value: []byte("Second")}, Frequency_Second),
		Entry("Value is []bytes, M - Works",
			&types.AttributeValueMemberB{Value: []byte("M")}, Frequency_Minute),
		Entry("Value is []bytes, Minute - Works",
			&types.AttributeValueMemberB{Value: []byte("Minute")}, Frequency_Minute),
		Entry("Value is []bytes, H - Works",
			&types.AttributeValueMemberB{Value: []byte("H")}, Frequency_Hour),
		Entry("Value is []bytes, Hour - Works",
			&types.AttributeValueMemberB{Value: []byte("Hour")}, Frequency_Hour),
		Entry("Value is []bytes, D - Works",
			&types.AttributeValueMemberB{Value: []byte("D")}, Frequency_Day),
		Entry("Value is []bytes, Day - Works",
			&types.AttributeValueMemberB{Value: []byte("Day")}, Frequency_Day),
		Entry("Value is []bytes, W - Works",
			&types.AttributeValueMemberB{Value: []byte("W")}, Frequency_Week),
		Entry("Value is []bytes, Week - Works",
			&types.AttributeValueMemberB{Value: []byte("Week")}, Frequency_Week),
		Entry("Value is []bytes, MN - Works",
			&types.AttributeValueMemberB{Value: []byte("MN")}, Frequency_Month),
		Entry("Value is []bytes, Month - Works",
			&types.AttributeValueMemberB{Value: []byte("Month")}, Frequency_Month),
		Entry("Value is []bytes, Q - Works",
			&types.AttributeValueMemberB{Value: []byte("Q")}, Frequency_Quarter),
		Entry("Value is []bytes, Quarter - Works",
			&types.AttributeValueMemberB{Value: []byte("Quarter")}, Frequency_Quarter),
		Entry("Value is []bytes, Y - Works",
			&types.AttributeValueMemberB{Value: []byte("Y")}, Frequency_Year),
		Entry("Value is []bytes, Year - Works",
			&types.AttributeValueMemberB{Value: []byte("Year")}, Frequency_Year),
		Entry("Value is []bytes, 0 - Works",
			&types.AttributeValueMemberB{Value: []byte("0")}, Frequency_InvalidFrequency),
		Entry("Value is []bytes, 1 - Works",
			&types.AttributeValueMemberB{Value: []byte("1")}, Frequency_Second),
		Entry("Value is []bytes, 2 - Works",
			&types.AttributeValueMemberB{Value: []byte("2")}, Frequency_Minute),
		Entry("Value is []bytes, 3 - Works",
			&types.AttributeValueMemberB{Value: []byte("3")}, Frequency_Hour),
		Entry("Value is []bytes, 4 - Works",
			&types.AttributeValueMemberB{Value: []byte("4")}, Frequency_Day),
		Entry("Value is []bytes, 5 - Works",
			&types.AttributeValueMemberB{Value: []byte("5")}, Frequency_Week),
		Entry("Value is []bytes, 6 - Works",
			&types.AttributeValueMemberB{Value: []byte("6")}, Frequency_Month),
		Entry("Value is []bytes, 7 - Works",
			&types.AttributeValueMemberB{Value: []byte("7")}, Frequency_Quarter),
		Entry("Value is []bytes, 8 - Works",
			&types.AttributeValueMemberB{Value: []byte("8")}, Frequency_Year),
		Entry("Value is int, 0 - Works",
			&types.AttributeValueMemberN{Value: "0"}, Frequency_InvalidFrequency),
		Entry("Value is int, 1 - Works",
			&types.AttributeValueMemberN{Value: "1"}, Frequency_Second),
		Entry("Value is int, 2 - Works",
			&types.AttributeValueMemberN{Value: "2"}, Frequency_Minute),
		Entry("Value is int, 3 - Works",
			&types.AttributeValueMemberN{Value: "3"}, Frequency_Hour),
		Entry("Value is int, 4 - Works",
			&types.AttributeValueMemberN{Value: "4"}, Frequency_Day),
		Entry("Value is int, 5 - Works",
			&types.AttributeValueMemberN{Value: "5"}, Frequency_Week),
		Entry("Value is int, 6 - Works",
			&types.AttributeValueMemberN{Value: "6"}, Frequency_Month),
		Entry("Value is int, 7 - Works",
			&types.AttributeValueMemberN{Value: "7"}, Frequency_Quarter),
		Entry("Value is int, 8 - Works",
			&types.AttributeValueMemberN{Value: "8"}, Frequency_Year),
		Entry("Value is NULL - Works", new(types.AttributeValueMemberNULL), Frequency(0)),
		Entry("Value is string, Empty string - Works",
			&types.AttributeValueMemberS{Value: ""}, Frequency_InvalidFrequency),
		Entry("Value is string, s - Works",
			&types.AttributeValueMemberS{Value: "s"}, Frequency_Second),
		Entry("Value is string, second - Works",
			&types.AttributeValueMemberS{Value: "second"}, Frequency_Second),
		Entry("Value is string, m - Works",
			&types.AttributeValueMemberS{Value: "m"}, Frequency_Minute),
		Entry("Value is string, minute - Works",
			&types.AttributeValueMemberS{Value: "minute"}, Frequency_Minute),
		Entry("Value is string, h - Works",
			&types.AttributeValueMemberS{Value: "h"}, Frequency_Hour),
		Entry("Value is string, hour - Works",
			&types.AttributeValueMemberS{Value: "hour"}, Frequency_Hour),
		Entry("Value is string, d - Works",
			&types.AttributeValueMemberS{Value: "d"}, Frequency_Day),
		Entry("Value is string, day - Works",
			&types.AttributeValueMemberS{Value: "day"}, Frequency_Day),
		Entry("Value is string, w - Works",
			&types.AttributeValueMemberS{Value: "w"}, Frequency_Week),
		Entry("Value is string, week - Works",
			&types.AttributeValueMemberS{Value: "week"}, Frequency_Week),
		Entry("Value is string, mn - Works",
			&types.AttributeValueMemberS{Value: "mn"}, Frequency_Month),
		Entry("Value is string, month - Works",
			&types.AttributeValueMemberS{Value: "month"}, Frequency_Month),
		Entry("Value is string, q - Works",
			&types.AttributeValueMemberS{Value: "q"}, Frequency_Quarter),
		Entry("Value is string, quarter - Works",
			&types.AttributeValueMemberS{Value: "quarter"}, Frequency_Quarter),
		Entry("Value is string, y - Works",
			&types.AttributeValueMemberS{Value: "y"}, Frequency_Year),
		Entry("Value is string, year - Works",
			&types.AttributeValueMemberS{Value: "year"}, Frequency_Year),
		Entry("Value is string, InvalidFrequency - Works",
			&types.AttributeValueMemberS{Value: "InvalidFrequency"}, Frequency_InvalidFrequency),
		Entry("Value is string, S - Works",
			&types.AttributeValueMemberS{Value: "S"}, Frequency_Second),
		Entry("Value is string, Second - Works",
			&types.AttributeValueMemberS{Value: "Second"}, Frequency_Second),
		Entry("Value is string, M - Works",
			&types.AttributeValueMemberS{Value: "M"}, Frequency_Minute),
		Entry("Value is string, Minute - Works",
			&types.AttributeValueMemberS{Value: "Minute"}, Frequency_Minute),
		Entry("Value is string, H - Works",
			&types.AttributeValueMemberS{Value: "H"}, Frequency_Hour),
		Entry("Value is string, Hour - Works",
			&types.AttributeValueMemberS{Value: "Hour"}, Frequency_Hour),
		Entry("Value is string, D - Works",
			&types.AttributeValueMemberS{Value: "D"}, Frequency_Day),
		Entry("Value is string, Day - Works",
			&types.AttributeValueMemberS{Value: "Day"}, Frequency_Day),
		Entry("Value is string, W - Works",
			&types.AttributeValueMemberS{Value: "W"}, Frequency_Week),
		Entry("Value is string, Week - Works",
			&types.AttributeValueMemberS{Value: "Week"}, Frequency_Week),
		Entry("Value is string, MN - Works",
			&types.AttributeValueMemberS{Value: "MN"}, Frequency_Month),
		Entry("Value is string, Month - Works",
			&types.AttributeValueMemberS{Value: "Month"}, Frequency_Month),
		Entry("Value is string, Q - Works",
			&types.AttributeValueMemberS{Value: "Q"}, Frequency_Quarter),
		Entry("Value is string, Quarter - Works",
			&types.AttributeValueMemberS{Value: "Quarter"}, Frequency_Quarter),
		Entry("Value is string, Y - Works",
			&types.AttributeValueMemberS{Value: "Y"}, Frequency_Year),
		Entry("Value is string, Year - Works",
			&types.AttributeValueMemberS{Value: "Year"}, Frequency_Year),
		Entry("Value is string, 0 - Works",
			&types.AttributeValueMemberS{Value: "0"}, Frequency_InvalidFrequency),
		Entry("Value is string, 1 - Works",
			&types.AttributeValueMemberS{Value: "1"}, Frequency_Second),
		Entry("Value is string, 2 - Works",
			&types.AttributeValueMemberS{Value: "2"}, Frequency_Minute),
		Entry("Value is string, 3 - Works",
			&types.AttributeValueMemberS{Value: "3"}, Frequency_Hour),
		Entry("Value is string, 4 - Works",
			&types.AttributeValueMemberS{Value: "4"}, Frequency_Day),
		Entry("Value is string, 5 - Works",
			&types.AttributeValueMemberS{Value: "5"}, Frequency_Week),
		Entry("Value is string, 6 - Works",
			&types.AttributeValueMemberS{Value: "6"}, Frequency_Month),
		Entry("Value is string, 7 - Works",
			&types.AttributeValueMemberS{Value: "7"}, Frequency_Quarter),
		Entry("Value is string, 8 - Works",
			&types.AttributeValueMemberS{Value: "8"}, Frequency_Year))

	// Test that attempting to deserialize a data.Frequency will fial and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("Scan - Value is nil - Error", func() {

		// Attempt to convert a fake string value into a data.Frequency; this should return an error
		var enum *Frequency
		err := enum.Scan(nil)

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of %!q(<nil>) had an invalid type of <nil>"))
		Expect(enum).Should(BeNil())
	})

	// Test the conditions under which values should be convertible to a data.Frequency
	DescribeTable("Scan Tests",
		func(value interface{}, shouldBe Frequency) {

			// Attempt to convert the value into a data.Frequency; this should not fail
			var enum Frequency
			err := enum.Scan(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Empty string - Works", "", Frequency_InvalidFrequency),
		Entry("s - Works", "s", Frequency_Second),
		Entry("second - Works", "second", Frequency_Second),
		Entry("m - Works", "m", Frequency_Minute),
		Entry("minute - Works", "minute", Frequency_Minute),
		Entry("h - Works", "h", Frequency_Hour),
		Entry("hour - Works", "hour", Frequency_Hour),
		Entry("d - Works", "d", Frequency_Day),
		Entry("day - Works", "day", Frequency_Day),
		Entry("w - Works", "w", Frequency_Week),
		Entry("week - Works", "week", Frequency_Week),
		Entry("mn - Works", "mn", Frequency_Month),
		Entry("month - Works", "month", Frequency_Month),
		Entry("q - Works", "q", Frequency_Quarter),
		Entry("quarter - Works", "quarter", Frequency_Quarter),
		Entry("y - Works", "y", Frequency_Year),
		Entry("year - Works", "year", Frequency_Year),
		Entry("InvalidFrequency - Works", "InvalidFrequency", Frequency_InvalidFrequency),
		Entry("S - Works", "S", Frequency_Second),
		Entry("Second - Works", "Second", Frequency_Second),
		Entry("M - Works", "M", Frequency_Minute),
		Entry("Minute - Works", "Minute", Frequency_Minute),
		Entry("H - Works", "H", Frequency_Hour),
		Entry("Hour - Works", "Hour", Frequency_Hour),
		Entry("D - Works", "D", Frequency_Day),
		Entry("Day - Works", "Day", Frequency_Day),
		Entry("W - Works", "W", Frequency_Week),
		Entry("Week - Works", "Week", Frequency_Week),
		Entry("MN - Works", "MN", Frequency_Month),
		Entry("Month - Works", "Month", Frequency_Month),
		Entry("Q - Works", "Q", Frequency_Quarter),
		Entry("Quarter - Works", "Quarter", Frequency_Quarter),
		Entry("Y - Works", "Y", Frequency_Year),
		Entry("Year - Works", "Year", Frequency_Year),
		Entry("0 - Works", 0, Frequency_InvalidFrequency),
		Entry("1 - Works", 1, Frequency_Second),
		Entry("2 - Works", 2, Frequency_Minute),
		Entry("3 - Works", 3, Frequency_Hour),
		Entry("4 - Works", 4, Frequency_Day),
		Entry("5 - Works", 5, Frequency_Week),
		Entry("6 - Works", 6, Frequency_Month),
		Entry("7 - Works", 7, Frequency_Quarter),
		Entry("8 - Works", 8, Frequency_Year))
})
