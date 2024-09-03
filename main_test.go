package main

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	"calc-lib/calc"
	"calc-lib/handler"
)

func assertEqual(t *testing.T, expected, actual any) {
	if !reflect.DeepEqual(expected, actual) {
		t.Helper()
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func assertError(t *testing.T, expected, actual error) {
	if !errors.Is(expected, actual) {
		t.Helper() //don't pay attention to this as part of the call stack.
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestTooFewArguments(t *testing.T) {
	myHandler := handler.NewHandler(&bytes.Buffer{}, calc.Addition{})
	err := myHandler.Handle(nil)
	assertError(t, handler.ErrTooFewArgs, err)
}

func TestMalformedFirstArgument(t *testing.T) {
	myHandler := handler.NewHandler(&bytes.Buffer{}, calc.Addition{})
	err := myHandler.Handle([]string{"NaN", "1"})
	assertError(t, handler.ErrMalformedArgs, err)
}

func TestMalformedSecondArgument(t *testing.T) {
	myHandler := handler.NewHandler(&bytes.Buffer{}, calc.Addition{})
	err := myHandler.Handle([]string{"1", "NaN"})
	assertError(t, handler.ErrMalformedArgs, err)
}

func TestOutputWriterError(t *testing.T) {
	taco := errors.New("tacos")
	myHandler := handler.NewHandler(&ErringWriter{taco}, calc.Addition{})
	err := myHandler.Handle([]string{"1", "2"})

	assertError(t, handler.ErrOutputWriter, err)
	assertError(t, taco, err)
}

func TestHappyPath(t *testing.T) {
	myBuffer := new(bytes.Buffer)
	myHandler := handler.NewHandler(myBuffer, calc.Addition{})
	err := myHandler.Handle([]string{"1", "2"})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	result := myBuffer.String()

	assertEqual(t, 2, len(result))
	//if  len(result) != 2 {
	//	t.Errorf("expected myBuffer to be of length 1, got %v", myBuffer.String())
	//}
	assertEqual(t, "3\n", result)
	//if result != "3\n" {
	//	t.Errorf("expected addition result to be 3, got %v", myBuffer.String())
	//}
}

type ErringWriter struct {
	err error
}

func (this *ErringWriter) Write(p []byte) (n int, err error) {
	return 0, this.err
}
