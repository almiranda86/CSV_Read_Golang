package service

import (
	"strings"
)

//This method has the function of read the file
//and return all it's in line
func FlattenResult(records [][]string) string {
	var response []string
	for _, row := range records {
		s := strings.Join(row, ",")
		response = append(response, s)
	}
	s := strings.Join(response, ",")
	return s
}
