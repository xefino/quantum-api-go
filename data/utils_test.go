package data

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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
