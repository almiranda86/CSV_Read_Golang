package service

import (
	"core/main/helper"
	"fmt"
	"strconv"
)

func MultiplyAllResult(records [][]string) string {
	var system_messages helper.SystemMessages
	system_messages.Fill_defaults()

	var response = 1
	for _, row := range records {
		for _, num := range row {
			ret, err := strconv.Atoi(num)
			if err != nil {
				return fmt.Sprintf(system_messages.ErrorWhileReadingValue, err.Error())
			}
			response *= ret
		}
	}
	return strconv.Itoa(response)
}
