package main

import (
	"bytes"
	"errors"
	"testing"

	"calc-lib/calc"
	"calc-lib/handler"
)

func TestTooFewArguments(t *testing.T) {
	myHandler := handler.NewHandler(&bytes.Buffer{}, calc.Addition{})
	err := myHandler.Handle(nil)
	if !errors.Is(err, handler.ErrTooFewArgs) {
		t.Errorf("expected %v, got %v", handler.ErrTooFewArgs, err)

		t.Error("Expected an error")
	}
}

func TestMalformedFirstArgument(t *testing.T) {
	myHandler := handler.NewHandler(&bytes.Buffer{}, calc.Addition{})
	err := myHandler.Handle([]string{"NaN", "1"})
	if !errors.Is(err, handler.ErrMalformedArgs) {
		t.Errorf("expected %v, got %v", handler.ErrMalformedArgs, err)
	}
}

func TestMalformedSecondArgument(t *testing.T) {
	myHandler := handler.NewHandler(&bytes.Buffer{}, calc.Addition{})
	err := myHandler.Handle([]string{"1", "NaN"})
	if !errors.Is(err, handler.ErrMalformedArgs) {
		t.Errorf("expected %v, got %v", handler.ErrMalformedArgs, err)
	}
}

func TestOutputWriterError(t *testing.T) {
	taco := errors.New("tacos")
	myHandler := handler.NewHandler(&ErringWriter{taco}, calc.Addition{})
	err := myHandler.Handle([]string{"1", "2"})
	if !errors.Is(err, handler.ErrOutputWriter) {
		t.Errorf("expected %v, got %v", handler.ErrOutputWriter, err)
	}

	if !errors.Is(err, taco) {
		t.Errorf("expected %v, got %v", taco, err)
	}

}

func TestHappyPath(t *testing.T) {
	myBuffer := new(bytes.Buffer)
	myHandler := handler.NewHandler(myBuffer, calc.Addition{})
	err := myHandler.Handle([]string{"1", "2"})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	result := myBuffer.String()

	if len(result) != 2 {
		t.Errorf("expected myBuffer to be of length 1, got %v", myBuffer.String())
	}
	if result != "3\n" {
		t.Errorf("expected addition result to be 3, got %v", myBuffer.String())
	}

}

type ErringWriter struct {
	err error
}

func (this *ErringWriter) Write(p []byte) (n int, err error) {
	return 0, this.err
}
