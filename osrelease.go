package osrelease

import (
	"errors"
	"strings"
)

//const EtcOsRelease string = "/etc/os-release"
//const UsrLibOsRelease string = "/usr/lib/os-release"

// Parse accepts the contents of an os-release file as
// a string and returns a map containing the parsed values
func Parse(contents string) (map[string]string, error) {
	return nil, errors.New("Not Implemented")
}

// Accepts a string representing a single line of
// an os-release file and returns the line in
// key, value form. If the line is empty or a
// comment, the key string will be empty.
// If the line cannot be parsed, an error is
// returned.
func parseLine(line string) (string, string, error) {
	line = strings.Trim(line, " \t\n")
	line = string(line)

	// Skip parsing empty or comment lines
	if (len(line) < 1 || line[0] == '#') {
		return "", "", nil
	}


	parsed := strings.SplitN(line, "=", 2)

	// Failing to split the line is a parsing error
	if (len(parsed) != 2) {
		return "", "", errors.New("parse: Error splitting line into key/value pair")
	}
	return parsed[0], parsed[1], nil
}
