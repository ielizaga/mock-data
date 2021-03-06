package core

import (
	"regexp"
	"fmt"
	"strconv"
	"strings"
	"math"
	"os"
	"time"
	"path/filepath"
	"log"
	"bufio"
)


// Extract the current time now.
func TimeNow() string {
	return time.Now().Format("20060102150405")
}

// Create a file ( if not exists ), append the content and then close the file
func WriteToFile(filename string, message string) error {

	// open files r, w mode
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY,0600)
	if err != nil {
		return err
	}

	// Close the file
	defer file.Close()

	// Append the message or content to be written
	if _, err = file.WriteString(message); err != nil {
		return err
	}

	return nil
}

// List all the backup sql file to recreate the constraints
func ListFile(dir, suffix string) ([]string, error) {
	return filepath.Glob(filepath.Join(dir, suffix))
}

// Read the file content and send it across
func ReadFile(filename string) ([]string, error) {

	var contentSaver []string

	// Open th file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contentSaver = append(contentSaver, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return contentSaver, err
	}
	return contentSaver, nil
}

// Is the value string or integer
func IsIntorString(v string) bool {
	_, err := strconv.Atoi(v)
	if err != nil {
		return false
	}
	return true
}

// Ignore Error strings matches
func IgnoreErrorString(errmsg string, ignoreErr []string) bool {
	for _, ignore := range ignoreErr {
		if strings.HasSuffix(errmsg, ignore) || strings.HasPrefix(errmsg, ignore) {
			return true
		}
	}
	return false
}

// Built a method to find if the values exits with a slice
func StringContains(item string, slice []string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

// Build a method to find if the value starts with specific word within a slice
func StringHasPrefix(item string, slice []string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		if strings.HasPrefix(item, s) {
			set[item] = struct{}{}
		}
	}
	_, ok := set[item]
	return ok
}

// Extract total characters that the datatype char can store.
func CharLen(dt string) (int, error) {
	var rgx = regexp.MustCompile(`\((.*?)\)`)
	var returnValue int
	var err error
	rs := rgx.FindStringSubmatch(dt)
	if len(rs) > 0 { // If the datatypes has number of value defined
		returnValue, err = strconv.Atoi(rs[1])
		if err != nil {
			return 0, err
		}
	} else {
		returnValue = 1
	}
	return returnValue, nil
}

// Column Extractor from the provided constraint key
func ColExtractor(conkey,regExp string) (string, error) {
	var rgx = regexp.MustCompile(regExp)
	rs := rgx.FindStringSubmatch(conkey)
	if len(rs) > 0 {
		return rs[0], nil
	} else {
		return "", fmt.Errorf("Unable to extract the columns from the constraint key")
	}
	return "", nil
}

// If given a datatype see if it has a bracket or not.
func BracketsExists(dt string) bool {
	var rgx = regexp.MustCompile(`\(.*\)`)
	rs := rgx.FindStringSubmatch(dt)
	if len(rs) > 0 {
		return true
	} else {
		return false
	}
}

// Extract Float precision from the float datatypes
func FloatPrecision(dt string) (int, int, error) {

	// check if brackets exists, if it doesn't then add some virtual values
	if !BracketsExists(dt) && strings.HasSuffix(dt, "[]") {
		dt = strings.Replace(dt, "[]", "", 1) + "(5,3)[]"
	} else if !BracketsExists(dt) && !strings.HasSuffix(dt, "[]") {
		dt = dt + "(5,3)"
	}
	// Get the ranges in the brackets
	var rgx = regexp.MustCompile(`\((.*?)\)`)
	rs := rgx.FindStringSubmatch(dt)
	split := strings.Split(rs[1], ",")
	m, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Float Precision (min): %v", err)
	}
	p, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Float Precision (precision): %v", err)
	}
	return m, p, nil
}

// If the random value of numeric datatype is greater than specifed, it ends up with
// i.e error "numeric field overflow"
// The below helper helps to reduce the size of the value
func TruncateFloat(f float64, max, precision int) float64 {
	stringFloat := strconv.FormatFloat(f, 'f', precision, 64)
	if len(stringFloat) > max {
		f = math.Log10(f)
	}
	return f
}