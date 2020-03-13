package main

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

// User contains user information
type SomeStruct struct {
	IntervalGte time.Duration `validate:"gtedfield=1s"`
	IntervalLte time.Duration `validate:"ltedfield=5s"`
	IntervalGt  time.Duration `validate:"gtdfield=1s"`
	IntervalLt  time.Duration `validate:"ltdfield=5s"`
	IntervalEq  time.Duration `validate:"eqdfield=5s"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {

	validate = validator.New()

	validateStruct()
}

func validateStruct() {

	some := &SomeStruct{
		IntervalGte: 1 * time.Second,
		IntervalLte: 5 * time.Second,
		IntervalGt:  2 * time.Second,
		IntervalLt:  4 * time.Second,
		IntervalEq:  5 * time.Second,
	}

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(some)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}
}
