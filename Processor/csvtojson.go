package csvtojson

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"
)

// Convert reads from the provided io.Reader (expected to be CSV data) and returns its JSON representation.
func Convert(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)

	// Read headers
	if !scanner.Scan() {
		return "", errors.New("failed to read headers from CSV data")
	}
	headers := splitCSV(scanner.Text())

	var buffer bytes.Buffer
	buffer.WriteString("[\n")

	// Read and process each data row
	for scanner.Scan() {
		row := splitCSV(scanner.Text())

		if len(row) != len(headers) {
			return "", errors.New("data row does not match headers in number of columns")
		}

		buffer.WriteString("  {\n")
		for i, header := range headers {
			buffer.WriteString("    \"" + header + "\": \"" + row[i] + "\"")
			if i < len(headers)-1 {
				buffer.WriteString(",\n")
			}
		}
		buffer.WriteString("\n  },\n")
	}

	// Remove trailing comma and close the JSON array
	jsonStr := buffer.String()
	if strings.HasSuffix(jsonStr, ",\n") {
		jsonStr = jsonStr[:len(jsonStr)-2]
	}
	jsonStr += "\n]"

	return jsonStr, nil
}

// splitCSV splits a CSV string into columns, taking into account quoted delimiters.
func splitCSV(data string) []string {
	var result []string
	var currentField strings.Builder
	insideQuotes := false

	for _, ch := range data {
		switch {
		case ch == ',' && !insideQuotes:
			result = append(result, currentField.String())
			currentField.Reset()
		case ch == '"':
			insideQuotes = !insideQuotes
		default:
			currentField.WriteRune(ch)
		}
	}

	// Append the last field
	result = append(result, currentField.String())
	return result
}
