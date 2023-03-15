package main

import (
	"core/main/service"
	"encoding/csv"
	"os"
	"testing"
)

func TestMultiplyAllResult(t *testing.T) {
	file, err := os.Open("../matrix.csv")

	if err != nil {
		t.Errorf("Error while reading file during tests %q", err)
	}
	defer file.Close()

	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		t.Errorf("Error while reading file rows during tests %q", err)
	}

	response := service.MultiplyAllResult(rows)

	want := "362880"

	if response != want {
		t.Errorf("got %q want %q", response, want)
	}
}
