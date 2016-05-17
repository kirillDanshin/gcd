package gcd

import "testing"

func TestCalc(t *testing.T) {
	if Calc(1071, 462) != 21 {
		t.Fatalf("Calc(1071, 462) is not 21")
	}
        if Calc(1000, 42) != 2 {
                t.Fatalf("Calc(1000, 42) is not 2")
        }
        if Calc(1931, 522) != 1 {
                t.Fatalf("Calc(1931, 522) is not 1")
        }

}
