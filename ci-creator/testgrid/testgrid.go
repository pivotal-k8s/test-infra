package testgrid

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type TestgridConfig struct {
	// TestGroups []TestGroup `yaml:"test_groups"`
	Dashboards []Dashboard `yaml:"dashboards"`
}

// type TestGroup struct {
// 	Name      string `yaml:"name"`
// 	GcsPrefix string `yaml:"gcs_prefix"`
// }

type Dashboard struct {
	Name         string       `yaml:"name"`
	DashboardTab DashboardTab `yaml:"dashboard_tab,omitempty"`
}

type DashboardTab struct {
	Name          string `yaml:"name"`
	Description   string `yaml:"description,omitempty"`
	TestGroupName string `yaml:"test_group_name,omitempty"`
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

	return testgridConfig, err
}
