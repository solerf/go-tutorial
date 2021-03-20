package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }

    /** 
    * := operator is a shortcut for declaring and initializing a variable in one line
    * 
    * Or...
    *
    * var message string
	* message = fmt.Sprintf("Hi, %v. Welcome!", name)
	*/
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}