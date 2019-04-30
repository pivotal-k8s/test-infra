package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCiCreator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CiCreator Suite")
}
