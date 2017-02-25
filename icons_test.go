package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func testIconsFindHelper(terms []string) icons {
	return newIcons().find(terms)
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

func TestIcons_iconsUnmarshalYaml(t *testing.T) {
	b := []byte(`
icons:
  - name:       Glass
    id:         glass
    unicode:    f000
    created:    1.0
    filter:
      - martini
`)
	actual, _ := iconsUnmarshalYaml(b)

	icon := icon{
		Name:    "Glass",
		ID:      "glass",
		Unicode: "f000",
		Created: "1.0",
		Filter:  []string{"martini"},
	}
	expected := iconsYaml{icons{icon}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIcons_find_AllIcons(t *testing.T) {
	fi := testIconsFindHelper([]string{""})

	actual := len(fi)
	expected := 675
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIcons_find_ZeroIcon(t *testing.T) {
	fi := testIconsFindHelper([]string{"foo-bar-baz"})

	actual := len(fi)
	expected := 0
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIcons_find_OneIcon(t *testing.T) {
	fi := testIconsFindHelper([]string{"github-square"})

	actual := len(fi)
	expected := 1
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIcons_find_TwoIcons(t *testing.T) {
	fi := testIconsFindHelper([]string{"github-"})

	actual := len(fi)
	expected := 2
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIcons_find_FirstIcon(t *testing.T) {
	fi := testIconsFindHelper([]string{""})

	actual := fi[0].ID
	expected := "500px"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIcons_find_LastIcon(t *testing.T) {
	fi := testIconsFindHelper([]string{""})

	actual := fi[len(fi)-1].ID
	expected := "youtube-square"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIcons_find_TaxiIcon(t *testing.T) {
	fi := testIconsFindHelper([]string{"taxi"})

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

func TestIcons_find_Aliases(t *testing.T) {
	fi := testIconsFindHelper([]string{"navicon"})

	actual := fi[0].ID
	expected := "bars"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	if len(fi) != 1 {
		t.Errorf("expected %v to eq %v", len(fi), 1)
	}
}

func TestIcons_findByUnicode(t *testing.T) {
	fi := newIcons().findByUnicode("f067")

	actual := fi[0].ID
	expected := "plus"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	if len(fi) != 1 {
		t.Errorf("expected %v to eq %v", len(fi), 1)
	}
}
