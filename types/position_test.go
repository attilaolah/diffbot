package types

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPosition(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Position Suite")
}

var _ = Describe("Position", func() {
	Describe("JSON-deserialising", func() {
		Context("An array of two ints", func() {
			It("should return a valid Position object", func() {
				var p Position
				b := []byte(`[
				  1,
				  2
				]`)
				err := json.Unmarshal(b, &p)
				Expect(err).NotTo(HaveOccurred())
				Expect(p).To(Equal(Position{
					X: 1,
					Y: 2,
				}))
			})
		})
		Context("An array of three ints", func() {
			It("should return an error", func() {
				var p Position
				b := []byte(`[
				  1,
				  2,
				  3
				]`)
				err := json.Unmarshal(b, &p)
				Expect(err).To(HaveOccurred())
			})
		})
		Context("An array of only one int", func() {
			It("should return an error", func() {
				var p Position
				b := []byte(`[
				  1
				]`)
				err := json.Unmarshal(b, &p)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
