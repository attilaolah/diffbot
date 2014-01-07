package types

import (
	"encoding/json"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Time Suite")
}

var _ = Describe("Time", func() {
	Describe("JSON-deserialising", func() {
		Context("An RFC3339-encoded timestamp", func() {
			It("should return a valid Time object", func() {
				var t Time
				b := []byte(`"2014-01-07T20:02:24Z"`)
				err := json.Unmarshal(b, &t)
				Expect(err).NotTo(HaveOccurred())
				Expect(t.Time().Equal(time.Date(2014, time.January, 7, 20, 2, 24, 0, time.UTC))).To(BeTrue())
			})
		})
		Context("An RFC1123-encoded timestamp", func() {
			It("should return a valid Time object", func() {
				var t Time
				b := []byte(`"Tue, 7 Jan 2014 20:02:24 ICT"`)
				err := json.Unmarshal(b, &t)
				Expect(err).NotTo(HaveOccurred())
				ict, err := time.LoadLocation("Asia/Bangkok")
				Expect(err).NotTo(HaveOccurred())
				Expect(t.Time().Equal(time.Date(2014, time.January, 7, 20, 2, 24, 0, ict))).To(BeTrue())
			})
		})
	})
})
