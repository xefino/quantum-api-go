package v1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Create a new test runner we'll use to test all the
// modules in the OHLC package
func TestOHLC(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OHLC Suite")
}

var _ = Describe("OHLC.v1.Frequency Marshal/Unmarshal Tests", func() {

	// Test that attempting to deserialize a OHLC.v1.Frequency will fail and
	// return an error if the value canno be deserialized from a JSON value to a string
	It("UnmarshalJSON fails - Error", func() {

		// Attempt to convert a non-parseable string value into a OHLC.v1.Frequency
		// This should return an error
		enum := new(Frequency)
		err := enum.UnmarshalJSON([]byte("derp"))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a v1.Frequency"))
	})

	// Test that attempting to deserialize a OHLC.v1.Frequency will fail and
	// return an error if the value cannot be converted to either the name value or integer
	// value of the enum option
	It("UnmarshalJSON - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a OHLC.v1.Frequency
		// This should return an error
		enum := new(Frequency)
		err := enum.UnmarshalJSON([]byte("\"derp\""))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a v1.Frequency"))
	})

	// Test the conditions under which values should be convertible to a OHLC.v1.Frequency
	DescribeTable("UnmarshalJSON Tests",
		func(value string, shouldBe Frequency) {

			// Attempt to convert the string value into a OHLC.v1.Frequency
			// This should not fail
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

	// Test that attempting to deserialize a OHLC.v1.Frequency will fail and
	// return an error if the value canno be deserialized from a param value to a string
	It("UnmarshalParam fails - Error", func() {

		// Attempt to convert a non-parseable string value into a OHLC.v1.Frequency
		// This should return an error
		enum := new(Frequency)
		err := enum.UnmarshalParam("derp")

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a v1.Frequency"))
	})

	// Test that attempting to deserialize a OHLC.v1.Frequency will fail and
	// return an error if the value cannot be converted to either the name value or integer
	// value of the enum option
	It("UnmarshalParam - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a OHLC.v1.Frequency
		// This should return an error
		enum := new(Frequency)
		err := enum.UnmarshalParam("derp")

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a v1.Frequency"))
	})

	// Test the conditions under which values should be convertible to a OHLC.v1.Frequency
	DescribeTable("UnmarshalParam Tests",
		func(value string, shouldBe Frequency) {

			// Attempt to convert the string value into a OHLC.v1.Frequency
			// This should not fail
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
		Entry("h - Works", "\"h\"", Frequency_Hour),
		Entry("hour - Works", "hour", Frequency_Hour),
		Entry("d - Works", "\"d\"", Frequency_Day),
		Entry("day - Works", "day", Frequency_Day),
		Entry("w - Works", "\"w\"", Frequency_Week),
		Entry("week - Works", "week", Frequency_Week),
		Entry("mn - Works", "\"mn\"", Frequency_Month),
		Entry("month - Works", "month", Frequency_Month),
		Entry("q - Works", "\"q\"", Frequency_Quarter),
		Entry("quarter - Works", "quarter", Frequency_Quarter),
		Entry("y - Works", "\"y\"", Frequency_Year),
		Entry("year - Works", "year", Frequency_Year),
		Entry("InvalidFrequency - Works", "InvalidFrequency", Frequency_InvalidFrequency),
		Entry("S - Works", "S", Frequency_Second),
		Entry("Second - Works", "Second", Frequency_Second),
		Entry("M - Works", "\"M\"", Frequency_Minute),
		Entry("Minute - Works", "Minute", Frequency_Minute),
		Entry("H - Works", "\"H\"", Frequency_Hour),
		Entry("Hour - Works", "Hour", Frequency_Hour),
		Entry("D - Works", "\"D\"", Frequency_Day),
		Entry("Day - Works", "Day", Frequency_Day),
		Entry("W - Works", "\"W\"", Frequency_Week),
		Entry("Week - Works", "Week", Frequency_Week),
		Entry("MN - Works", "\"MN\"", Frequency_Month),
		Entry("Month - Works", "Month", Frequency_Month),
		Entry("Q - Works", "\"Q\"", Frequency_Quarter),
		Entry("Quarter - Works", "Quarter", Frequency_Quarter),
		Entry("Y - Works", "\"Y\"", Frequency_Year),
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
})
