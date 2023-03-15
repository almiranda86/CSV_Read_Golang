package service

import (
	"core/main/helper"
	"fmt"
	"strings"
)

//This method has the function of read the file
//and invert the matrix and then return it as string
func InvertResult(records [][]string) string {
	var tranposed = helper.Transpose(records)

	var response string
	for _, row := range tranposed {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}
