package flags

import (
	"fmt"
	"log"
	"strings"
)

type Required struct {
	Name  string
	Value string
	Usage string
}

// AssertRequired panics if a given flag is not set
func AssertRequired(requiredFlags []*Required) {
	if errors := getValidationErrors(requiredFlags); len(errors) > 0 {
		var messages []string
		for _, error := range errors {
			messages = append(messages, error.Error())
		}
		log.Fatal(strings.Join(messages, ", "))
	}
}

func getValidationErrors(r []*Required) []error {
	var errors []error
	for _, requiredFlag := range r {
		if error := requiredFlag.validate(); error != nil {
			errors = append(errors, error)
		}
	}
	return errors
}

func (r *Required) validate() error {
	if r.Value == "" {
		return fmt.Errorf("%s is required", r.Name)
	}
	return nil
}
