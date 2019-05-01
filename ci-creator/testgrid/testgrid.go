package testgrid

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type TestgridConfig struct {
	// TestGroups []TestGroup `yaml:"test_groups"`
	Dashboards      []Dashboard      `yaml:"dashboards"`
	DashboardGroups []DashboardGroup `yaml:"dashboard_groups"`
}

type Dashboard struct {
	Name         string         `yaml:"name"`
	DashboardTab []DashboardTab `yaml:"dashboard_tab"`
}

type DashboardTab struct {
	Name          string `yaml:"name"`
	Description   string `yaml:"description,omitempty"`
	TestGroupName string `yaml:"test_group_name,omitempty"`
}

type DashboardGroup struct {
	Name           string
	DashboardNames []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GenerateTestgrid(upcomingRelease string) (TestgridConfig, error) {

	data, err := ioutil.ReadFile("../../testgrid/config.yaml")
	check(err)

	var testgridConfig TestgridConfig

	err = yaml.Unmarshal([]byte(data), &testgridConfig)
	check(err)

	dashboardForUpcomingRelease := Dashboard{
		Name: fmt.Sprintf("sig-release-%s-blocking", upcomingRelease),
	}

	testgridConfig.Dashboards = append(testgridConfig.Dashboards, dashboardForUpcomingRelease)

	return testgridConfig, err
}
