package data

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Data.Frequency Extension Tests", func() {

	// Tests the conditions under which the ForDate function will return true/false
	DescribeTable("ForDate - Conditions",
		func(enum Frequency, result bool) {
			Expect(enum.ForDate()).Should(Equal(result))
		},
		Entry("Invalid - False", Frequency_InvalidFrequency, false),
		Entry("Second - False", Frequency_Second, false),
		Entry("Minute - False", Frequency_Minute, false),
		Entry("Hour - False", Frequency_Hour, false),
		Entry("Day - True", Frequency_Day, true),
		Entry("Week - True", Frequency_Week, true),
		Entry("Month - True", Frequency_Month, true),
		Entry("Quarter - True", Frequency_Quarter, true),
		Entry("Year - True", Frequency_Year, true))

	// Tests the conditions under which the ForTime function will return true/false
	DescribeTable("ForTime - Conditions",
		func(enum Frequency, result bool) {
			Expect(enum.ForTime()).Should(Equal(result))
		},
		Entry("Invalid - False", Frequency_InvalidFrequency, false),
		Entry("Second - True", Frequency_Second, true),
		Entry("Minute - True", Frequency_Minute, true),
		Entry("Hour - True", Frequency_Hour, true),
		Entry("Day - False", Frequency_Day, false),
		Entry("Week - False", Frequency_Week, false),
		Entry("Month - False", Frequency_Month, false),
		Entry("Quarter - False", Frequency_Quarter, false),
		Entry("Year - False", Frequency_Year, false))
})
