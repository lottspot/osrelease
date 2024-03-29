package osrelease

import (
	"errors"
	"strings"
	"io"
	"bufio"
	"os"
)

// EtcOsRelease provides a path to the well known location /etc/os-release
const EtcOsRelease string = "/etc/os-release"

// UsrLibOsRelease provides a path to the well known location /usr/lib/os-release
const UsrLibOsRelease string = "/usr/lib/os-release"

// Load will attempt to automatically load and
// parse os-release information from the two
// well-known os-release locations
func Load() (map[string]string, error) {
	parsed, e := LoadPath(EtcOsRelease)
	if (e != nil) {
		parsed, e := LoadPath(UsrLibOsRelease)
		if (e != nil) {
			return nil, e
		}
		return parsed, nil
	}

	return parsed, nil
}

// LoadPath will open a specified path and pass the opened
// file descriptor to the Read function
func LoadPath(path string) (map[string]string, error) {

	fd, eOpen := os.Open(path)
	if (eOpen != nil) {
		return nil, eOpen
	}
	defer fd.Close()

	parsed, eRead := Read(fd)
	if (eRead != nil) {
		return nil, eRead
	}

	return parsed, nil
}

// Read accepts an io.Reader pointing to an os-release
// file, reads the contents, and returns a map of parsed
// values
func Read(src io.Reader) (map[string]string, error) {
	scanner := bufio.NewScanner(src)
	data    := ""

	for scanner.Scan() {
		data += scanner.Text()
	}
	eRead := scanner.Err()
	if (eRead != nil) {
		return nil, errors.New("read: " + eRead.Error())
	}

	parsed, eParse := Parse(data)
	if (eParse != nil) {
		return nil, eParse
	}

	return parsed, nil
}

// Parse accepts the contents of an os-release file as
// a string and returns a map containing the parsed values
func Parse(contents string) (map[string]string, error) {
	parsed := make(map[string]string)
	lines  := strings.Split(contents, "\n")

	for i, line := range lines {
		k, v, e := parseLine(line)
		if (e != nil) {
			return nil, errors.New("line " + string(i+1) + ": " + e.Error())
		}
		parsed[k] = v
	}

	return parsed, nil
}

// Accepts a string representing a single line of
// an os-release file and returns the line in
// key, value form. If the line is empty or a
// comment, the key string will be empty.
// If the line cannot be parsed, an error is
// returned.
func parseLine(line string) (string, string, error) {
	line = strings.Trim(line, " \t\n")

	// Skip parsing empty or comment lines
	if (len(line) < 1 || line[0] == '#') {
		return "", "", nil
	}

	parsed := strings.SplitN(line, "=", 2)

	// Failing to split the line is a parsing error
	if (len(parsed) != 2) {
		return "", "", errors.New("parse: Error splitting line into key/value pair")
	}

	k := parsed[0]
	v := strings.Trim(parsed[1], "'\"")

	return k, v, nil
}
