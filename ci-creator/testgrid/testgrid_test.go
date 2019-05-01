package testgrid_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "k8s.io/test-infra/ci-creator/testgrid"
)

var (
	testgridConfig TestgridConfig
	err            error
)

var _ = Describe("Testgrid", func() {

	BeforeEach(func() {
		testgridConfig, err = GenerateTestgrid("1.15")
	})

	It("succeeds", func() {
		Expect(err).NotTo(HaveOccurred())
	})

	It("adds an entry for the new release under dashboard_groups['sig-release']", func() {
		Expect(findInSigReleaseDashboardGroups("sig-release-1.15-blocking")).To(Succeed())
	})

	It("adds a dashboard for the new release under dashboards", func() {
		Expect(findInDashboards("sig-release-1.15-blocking")).To(Succeed())
	})

	It("contains the same tests as master, but named after the release name", func() {

	})

})

func findInDashboards(name string) error {
	for _, dashboard := range testgridConfig.Dashboards {
		if dashboard.Name == name {
			return nil
		}
	}
	return errors.New("")
}

func findInSigReleaseDashboardGroups(name string) error {
	for _, group := range testgridConfig.DashboardGroups {
		if group.Name == "sig-release" {
			for _, dashboardName := range group.DashboardNames {
				if dashboardName == name {
					break
				}
			}
		}
	}
	return errors.New("")
}
