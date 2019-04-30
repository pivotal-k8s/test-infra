package testgrid_test

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "k8s.io/test-infra/ci-creator/testgrid"
)

var (
	testgridConfig TestgridConfig
	err            error
)

var _ = Describe("Testgrid", func() {

	BeforeSuite(func() {
		testgridConfig, err = GenerateTestgrid("1.15")
	})

	It("succeeds", func() {
		Expect(err).NotTo(HaveOccurred())
	})

	It("adds a dashboard for the new release", func() {
		Expect(findInDashboards("sig-release-master-blocking")).To(Succeed())
	})

})

func findInDashboards(name string) error {
	for _, dashboard := range testgridConfig.Dashboards {
		fmt.Println(dashboard)
		if dashboard.Name == name {
			return nil
		}
	}
	return errors.New("")
}
