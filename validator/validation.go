package validator

import (
	"core/main/helper"
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"net/http"
)

//a struct to perform common validations related to the business functionality
type CustomValidationModel struct {
	Result bool
	Err    error
}

// constructor function
func (cvm *CustomValidationModel) fill_defaults() {
	cvm.Err = nil
	cvm.Result = true
}

//verify if the file is a valid one
func ValidateFile(r *http.Request) *CustomValidationModel {
	var validationResponse = new(CustomValidationModel)
	validationResponse.fill_defaults()

	_, _, err := r.FormFile("file")

	if err != nil {
		validationResponse.Err = err
		validationResponse.Result = false
	}

	return validationResponse
}

//verify if the records are in the file
func ValidateRecordsfile(file multipart.File) *CustomValidationModel {
	var validationResponse = new(CustomValidationModel)
	validationResponse.fill_defaults()

	_, err := csv.NewReader(file).ReadAll()

	if err != nil {
		validationResponse.Err = err
		validationResponse.Result = false
	}

	return validationResponse
}

//verify if matrix is valid
//if the len is equal to zero it means that the matrix isn't a quadratic one
func IsMatrixValid(s [][]string) bool {
	return len(s) == 0
}

func ValidateRequest(r *http.Request) (string, [][]string) {
	var system_messages helper.SystemMessages
	system_messages.Fill_defaults()

	file, _, _ := r.FormFile("file")
	defer file.Close()
	records, _ := csv.NewReader(file).ReadAll()

	fileValidationResult := ValidateFile(r)

	if !fileValidationResult.Result {
		return fmt.Sprintf(system_messages.ErrorWhileValidating, fileValidationResult.Err.Error()), nil
	}

	recordsValidationResult := ValidateRecordsfile(file)

	if !recordsValidationResult.Result {
		return fmt.Sprintf(system_messages.ErrorWhileReadingRecord, recordsValidationResult.Err.Error()), nil
	}

	if IsMatrixValid(records) {
		return system_messages.ValueInformedIncorret, nil
	}

	return system_messages.ValidationOK, records
}
