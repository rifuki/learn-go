package main

import (
	"fmt"
	"strings"
)

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{
			Message: "an validation error occured.",
		}
	}

	if id != "rifuki" {
		return &notFoundError{
			Message: "data not found!",
		}
	}

	return nil
}

func main() {
	err1 := SaveData("", nil) /* will return validation error */

	if err1 != nil {
		if validationErr, ok := err1.(*validationError); ok { /* <- type assertions returns 2 value (result conversion and result is conversion is succes) */
			fmt.Println("validation error1:", validationErr.Error())
		} else if notFoundErr, ok := err1.(*validationError); ok {
			fmt.Println("not found error1:", notFoundErr.Error())
		} else {
			fmt.Println("unknown error1.")
		}
	} else {
		fmt.Println("success1")
	}

	err2 := SaveData("John", nil) /* <- will return not found error */
	if err2 != nil {
		if validationErr, ok := err2.(*validationError); ok { /* <- type assertions returns 2 value (result conversion and result is conversion is succes) */
			fmt.Println("validation error2:", validationErr.Error())
		} else if notFoundErr, ok := err2.(*notFoundError); ok {
			fmt.Println("not found error2:", notFoundErr.Error())
		} else {
			fmt.Println("unknown error2.")
		}
	} else {
		fmt.Println("success2")
	}

	err3 := SaveData("Fate", nil) /* <- will return not found error */
	if err3 != nil {
		switch finalError := err3.(type) {
		case *validationError:
			fmt.Println("validation error3:", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error3:", finalError.Error())
		default:
			fmt.Println("unknown error3.")
		}
	} else {
		fmt.Println("success2")
	}

	err4 := SaveData(strings.ToLower("Rifuki"), nil) /* <- will return success */
	if err4 != nil {
		switch finalError := err4.(type) {
		case *validationError:
			fmt.Println("validation error4:", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error4:", finalError.Error())
		default:
			fmt.Println("unknown error4.")
		}
	} else {
		fmt.Println("success4")
	}
}
