package testgrid_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTestgrid(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testgrid Suite")
}
