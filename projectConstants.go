package main

// This error message enum defines the types of error messages
// that have text in the ErrorMessages map
type ErrorMessage int

const (
	MissingCommandLineArg ErrorMessage = iota
)

var ErrorMessages = map[ErrorMessage]string{
	MissingCommandLineArg: "Please provide an argument as in\n\tgo run . textToMatch",
}
var DoneMessage = "Done!"
