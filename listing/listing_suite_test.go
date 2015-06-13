package listing_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestListing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Listing Suite")
}
