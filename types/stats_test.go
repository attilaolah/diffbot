package types

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStats(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stats Suite")
}

var _ = Describe("Stats", func() {
	Describe("JSON-deserialising", func() {
		Context("An object with confidence set as a string", func() {
			It("should return a valid Stats object", func() {
				var s Stats
				b := []byte(`{
				  "confidence": "0.850"
				}`)
				err := json.Unmarshal(b, &s)
				Expect(err).NotTo(HaveOccurred())
				Expect(s).To(Equal(Stats{
					Confidence: .85,
				}))
			})
		})
		Context("A JSON null literal", func() {
			It("should return an a Stats object with no confidence info", func() {
				var s Stats
				b := []byte("null")
				err := json.Unmarshal(b, &s)
				Expect(err).NotTo(HaveOccurred())
				Expect(s).To(Equal(Stats{}))
			})
		})
		Context("Anything else", func() {
			It("should return an error", func() {
				var s Stats
				b := []byte("[]")
				err := json.Unmarshal(b, &s)
				Expect(err).To(HaveOccurred())
				Expect(s).To(Equal(Stats{}))
			})
		})
	})
})
