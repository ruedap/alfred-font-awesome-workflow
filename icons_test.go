package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func testIconsHelper_Find(terms []string) Icons {
	return NewIcons().Find(terms)
}

func TestIcons_iconsYamlPath_TestEnv(t *testing.T) {
	actual := iconsYamlPath()
	expected := "workflow/icons.yml"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIcons_iconsYamlPath_ProductionEnv(t *testing.T) {
	resetEnv := setTestEnvHelper("FAW_ICONS_YAML_PATH", "")
	defer resetEnv()

	actual := iconsYamlPath()
	expected := "icons.yml"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIcons_iconsReadYaml(t *testing.T) {
	path := "workflow/icons.yml"
	actual, _ := iconsReadYaml(path)

	expected, _ := ioutil.ReadFile(path)
	if !reflect.DeepEqual(actual, expected) {
		t.Error("failed to read file")
	}
}

func TestIcons_iconsReadYaml_Error(t *testing.T) {
	path := ""
	_, err := iconsReadYaml(path)

	if err == nil {
		t.Error("expected error, but nil")
	}
}

func TestIcons_Find_AllIcons(t *testing.T) {
	fi := testIconsHelper_Find([]string{""})

	actual := len(fi)
	expected := 519
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIcons_Find_ZeroIcon(t *testing.T) {
	fi := testIconsHelper_Find([]string{"foo-bar-baz"})

	actual := len(fi)
	expected := 0
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIcons_Find_OneIcon(t *testing.T) {
	fi := testIconsHelper_Find([]string{"github-square"})

	actual := len(fi)
	expected := 1
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIcons_Find_TwoIcons(t *testing.T) {
	fi := testIconsHelper_Find([]string{"github-"})

	actual := len(fi)
	expected := 2
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIcons_Find_FirstIcon(t *testing.T) {
	fi := testIconsHelper_Find([]string{""})

	actual := fi[0].ID
	expected := "adjust"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIcons_Find_LastIcon(t *testing.T) {
	fi := testIconsHelper_Find([]string{""})

	actual := fi[len(fi)-1].ID
	expected := "youtube-square"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIcons_Find_TaxiIcon(t *testing.T) {
	fi := testIconsHelper_Find([]string{"taxi"})

	actual := fi[0].Name
	expected := "Taxi"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	actual = fi[0].ID
	expected = "taxi"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	actual = fi[0].Unicode
	expected = "f1ba"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	actual = fi[0].Created
	expected = "4.1"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	actual = fi[0].Aliases[0]
	expected = "cab"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	actual = fi[0].Filter[0]
	expected = "vehicle"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	actual = fi[0].Categories[0]
	expected = "Web Application Icons"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIcons_Find_Aliases(t *testing.T) {
	fi := testIconsHelper_Find([]string{"navicon"})

	actual := fi[0].ID
	expected := "bars"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	if len(fi) != 1 {
		t.Errorf("expected %v to eq %v", len(fi), 1)
	}
}
