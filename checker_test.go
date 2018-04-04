package checker_test

import (
	"testing"

	"github.com/lukerodham/email-checker"
)

func TestIsGeneric(t *testing.T) {
	tests := []struct {
		Email       string
		IsGeneric   bool
		ShouldError bool
	}{
		{
			Email:       "example@gmail.com",
			IsGeneric:   true,
			ShouldError: false,
		},
		{
			Email:       "example@hotmail.com",
			IsGeneric:   true,
			ShouldError: false,
		},
		{
			Email:       "example@yahoo.com",
			IsGeneric:   true,
			ShouldError: false,
		},
		{
			Email:       "example@google.com",
			IsGeneric:   false,
			ShouldError: false,
		},
		{
			Email:       "yahoo.com",
			IsGeneric:   false,
			ShouldError: true,
		},
	}

	for _, test := range tests {
		is, err := checker.IsGeneric(test.Email)

		if err != nil && test.ShouldError == false {
			t.Error(err)
		}

		if test.IsGeneric != is {
			t.Errorf("expect %v to equal %v", test.IsGeneric, is)
		}
	}
}
