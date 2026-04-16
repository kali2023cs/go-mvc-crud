package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// InitCustomValidators registers custom rules for the validator engine
func InitCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Custom Rule: name must not be a reserved word
		v.RegisterValidation("not-reserved", func(fl validator.FieldLevel) bool {
			reserved := []string{"admin", "test", "root", "null"}
			name := strings.ToLower(fl.Field().String())
			for _, r := range reserved {
				if name == r {
					return false
				}
			}
			return true
		})
	}
}

// FormatError converts validator.ValidationErrors into a readable map
func FormatError(err error) map[string]string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make(map[string]string)
		for _, fe := range ve {
			out[fe.Field()] = getErrorMsg(fe)
		}
		return out
	}
	return map[string]string{"message": err.Error()}
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "numeric":
		return "This field must be a number"
	case "gt":
		return fmt.Sprintf("This field must be greater than %s", fe.Param())
	case "not-reserved":
		return "This name is reserved and cannot be used"
	}
	return "Invalid input"
}
