package types

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCategories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Categories Suite")
}

var _ = Describe("Categories", func() {
	Describe("JSON-deserialising", func() {
		Context("An object with string keys", func() {
			It("should return a valid Categories object", func() {
				var c Categories
				b := []byte(`{
				  "0": "zero",
				  "1": "one",
				  "2": "two",
				  "-8": "minus eight"
				}`)
				err := json.Unmarshal(b, &c)
				Expect(err).NotTo(HaveOccurred())
				Expect(c).To(Equal(Categories{
					0:  "zero",
					1:  "one",
					2:  "two",
					-8: "minus eight",
				}))
			})
		})
		Context("A JSON null literal", func() {
			It("should return nil", func() {
				var c Categories
				b := []byte("null")
				err := json.Unmarshal(b, &c)
				Expect(err).NotTo(HaveOccurred())
				Expect(c).To(BeNil())
			})
		})
		Context("Anything else", func() {
			It("should return an error", func() {
				var c Categories
				b := []byte("[]")
				err := json.Unmarshal(b, &c)
				Expect(err).To(HaveOccurred())
				Expect(c).To(BeNil())
			})
		})
	})
})
