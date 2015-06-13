package listing_test

import (
	. "github.com/socialnews/reddit-feeder/listing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("listing", func() {
	var (
		listing Listing
	)

	Describe("Listing object", func() {
		Context("At the beginning of it all", func() {
			It("should exist", func() {
				Expect(listing).NotTo(Equal(nil))
			})
		})
	})
})
