package osrelease

import (
	"testing"
)

func TestParse(t *testing.T) {
	content := `
# A basic os-release file
ID=linux
NAME="Linux Distro"
`
	parsed, e := Parse(content)
	id := parsed["ID"]
	name := parsed["NAME"]
	pretty := parsed["PRETTY_NAME"]

	if (e != nil) {
		t.Error("Expected nil error but got", e)
	}

	if (id != "linux") {
		t.Error("Expected value of ID to be linux but got", id)
	}

	if (name != "Linux Distro") {
		t.Error("Expected value of NAME to be Linux Distro but got", name)
	}

	if (pretty != "") {
		t.Error("Expected nil value for PRETTY_NAME but got", pretty)
	}

}
func TestParseLine(t *testing.T) {
	valid := "ID=\"linux\"  "
	invalid := "foobar"
	comment := "  # Operator information"
	empty := "   "

	k, v, _ := parseLine(valid)
	if ( !(k == "ID" && v == "linux") ) {
		t.Error("Expected (ID, linux) to be returned but got:", k, ",", v)
	}

	k, v, e := parseLine(invalid)
	if ( !(k == "" && v == "" && e != nil) ) {
		t.Error("Expected no k or v vales with an error but got:", k, ",", v, ",", e)
	}
	e = nil

	k, v, e = parseLine(comment)
	if ( !(k == "" && v == "" && e == nil) ) {
		t.Error("Expected empty k, v values without errors but got:", k, ",", v, ",", e)
	}
	e = nil

	k, v, e = parseLine(empty)
	if ( !(k == "" && v == "" && e == nil) ) {
		t.Error("Expected empty k, v values without errors but got:", k, ",", v, ",", e)
	}
}
