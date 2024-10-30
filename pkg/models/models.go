package models

type TestCase struct {
	Name           string
	Args           []string
	ExpectedOutput interface{}
}
