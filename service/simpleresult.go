package service

import (
	"fmt"
	"strings"
)

//This method has the function of read the file
//and return it as is, in string format
func SimpleResult(records [][]string) string {

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	return response
}
