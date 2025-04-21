package utils_test

import (
	"errors"
	"github.com/Varshi292/RoastWear/internal/utils"
	"testing"
)

func TestInitializeValidator(t *testing.T) {
	type testCase struct {
		name           string
		initializeFunc func() error
		errorExpected  bool
	}

	tests := []testCase{
		{
			name: "Valid: Initialization Succeeds",
			initializeFunc: func() error {
				return utils.InitializeValidator()
			},
			errorExpected: false,
		},
		{
			name: "Valid: Double Initialization",
			initializeFunc: func() error {
				_ = utils.InitializeValidator()
				return utils.InitializeValidator()
			},
			errorExpected: false,
		},
		{
			name: "Valid: Validator Instance Is Set",
			initializeFunc: func() error {
				if err := utils.InitializeValidator(); err != nil {
					return err
				}
				if utils.Validate == nil {
					return errors.New("validator instance was not set")
				}
				return nil
			},
			errorExpected: false,
		},
		{
			name: "Valid: Port Validator Is Successfully Registered",
			initializeFunc: func() error {
				if err := utils.InitializeValidator(); err != nil {
					return err
				}
				type MockPort struct {
					Port string `validate:"required,port"`
				}
				return utils.Validate.Struct(MockPort{Port: "3000"})
			},
			errorExpected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.initializeFunc()
			if tc.errorExpected && err == nil {
				t.Errorf("Test %q: expected expected initialization to fail but it passed: %v", tc.name, err)
			}
			if !tc.errorExpected && err != nil {
				t.Errorf("Test %q: expected expected initialization to pass but it failed: %v", tc.name, err)
			}
		})
	}
}

func TestPortValidation(t *testing.T) {
	type MockPort struct {
		Port string `validate:"required,port"`
	}

	if err := utils.InitializeValidator(); err != nil {
		t.Fatal(err)
	}

	type testCase struct {
		name          string
		input         MockPort
		errorExpected bool
	}

	tests := []testCase{
		{
			name:          "Valid: Normal Port",
			input:         MockPort{Port: "3000"},
			errorExpected: false,
		},
		{
			name:          "Invalid: Empty Port",
			input:         MockPort{Port: ""},
			errorExpected: true,
		},
		{
			name:          "Valid: Lower Bound Port",
			input:         MockPort{Port: "1"},
			errorExpected: false,
		},
		{
			name:          "Valid: Upper Bound Port",
			input:         MockPort{Port: "65535"},
			errorExpected: false,
		},
		{
			name:          "Invalid: Port Too Low",
			input:         MockPort{Port: "0"},
			errorExpected: true,
		},
		{
			name:          "Invalid: Port Too High",
			input:         MockPort{Port: "3000000"},
			errorExpected: true,
		},
		{
			name:          "Invalid: Negative Port",
			input:         MockPort{Port: "-3000"},
			errorExpected: true,
		},
		{
			name:          "Invalid: Decimal Port",
			input:         MockPort{Port: "3000.0"},
			errorExpected: true,
		},
		{
			name:          "Invalid: Hexadecimal Port",
			input:         MockPort{Port: "0xBB8"},
			errorExpected: true,
		},
		{
			name:          "Invalid: Non-numeric Port",
			input:         MockPort{Port: "port"},
			errorExpected: true,
		},
		{
			name:          "Invalid: Port with Whitespace",
			input:         MockPort{Port: " 3000 "},
			errorExpected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := utils.Validate.Struct(tc.input)
			if tc.errorExpected && err == nil {
				t.Errorf("Test %q: expected validation to fail but it passed (input: %+v)", tc.name, tc.input)
			}
			if !tc.errorExpected && err != nil {
				t.Errorf("Test %q: expected validation to pass but it failed: %v (input: %+v)", tc.name, err, tc.input)
			}
		})
	}
}
