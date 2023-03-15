package service

import (
	"core/main/helper"
	"fmt"
	"strconv"
)

//This method has the function of read the file
//and return all it's values sumed
// func SumAllResult(r *http.Request) string {
// 	var system_messages SystemMessages
// 	system_messages.fill_defaults()

func SumAllResult(records [][]string) string {
	var system_messages helper.SystemMessages
	system_messages.Fill_defaults()

	var response int
	for _, row := range records {
		for _, num := range row {
			ret, err := strconv.Atoi(num)
			if err != nil {
				return fmt.Sprintf(system_messages.ErrorWhileReadingValue, err.Error())
			}
			response += ret
		}
	}

	return strconv.Itoa(response)
}
