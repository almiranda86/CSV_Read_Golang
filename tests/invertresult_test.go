package main

import (
	"core/main/helper"
	"core/main/service"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestInvertResult(t *testing.T) {
	file, err := os.Open("../matrix.csv")

	if err != nil {
		t.Errorf("Error while reading file during tests %q", err)
	}
	defer file.Close()

	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		t.Errorf("Error while reading file rows during tests %q", err)
	}

	var want string
	var tranposed = helper.Transpose(rows)

	for _, row := range tranposed {
		want = fmt.Sprintf("%s%s\n", want, strings.Join(row, ","))
	}

	response := service.InvertResult(rows)

	if response != want {
		t.Errorf("got %q want %q", response, want)
	}
}
