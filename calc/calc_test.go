package calc

import "testing"

func TestAddition_Calculate(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"addition calculation", args{1, 2}, 3},
		{"addition calculation", args{2, -1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ad := Addition{}
			if got := ad.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddition(t *testing.T) {
	var calculator Addition
	assertEqual(t, 0, calculator.Calculate(0, 0))
	assertEqual(t, 2, calculator.Calculate(1, 1))
	assertEqual(t, -4, calculator.Calculate(0, -4))
	assertEqual(t, 5, calculator.Calculate(-2, 7))
	assertEqual(t, 2, calculator.Calculate(1, 1))

}

func assertEqual(t *testing.T, expected, actual any) {
	if actual != expected {
		t.Helper() //so that the stack trace shows the actual test that failed
		t.Errorf("Expected = %v, Got = %v", expected, actual)
	}
}
