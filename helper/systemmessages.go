package helper

//a struct to store all the system messages
type SystemMessages struct {
	ErrorWhileValidating    string
	ErrorWhileReadingRecord string
	ValueInformedIncorret   string
	ErrorWhileReadingValue  string
	ValidationOK            string
}

// constructor function
func (sm *SystemMessages) Fill_defaults() {
	sm.ErrorWhileValidating = "Error while validating the file %s"
	sm.ErrorWhileReadingRecord = "Error while reading the record %s"
	sm.ValueInformedIncorret = "Values informed in file are in a incorrect format"
	sm.ErrorWhileReadingValue = "Error while reading a value%s"
	sm.ValidationOK = "OK"
}
