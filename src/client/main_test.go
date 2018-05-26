package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	str1 := "6+30"
	x1, y1, operator1, err1 := parse(&str1)
	if err1 != nil {
		t.Fail()
	}

	if x1 != 6 {
		t.Error("x1 is false")
	}

	if y1 != 30 {
		t.Error("y1 is false")
	}

	if operator1 != "+" {
		t.Error("operator1 is false")
	}

	// test spaces
	str2 := " 6 + 30 "
	x2, y2, operator2, err2 := parse(&str2)
	if err2 != nil {
		t.Fail()
	}

	if x2 != 6 {
		t.Error("x2 is false")
	}

	if y2 != 30 {
		t.Error("y2 is false")
	}

	if operator2 != "+" {
		t.Error("operator2 is false")
	}

	// test wrong expresions
	str3 := " 6 +- 30 "
	_, _, _, err3 := parse(&str3)
	if err3 == nil {
		t.Fail()
	}

	// test *
	str4 := " 10 * 5 "
	x4, y4, operator4, err4 := parse(&str4)
	if err4 != nil {
		t.Fail()
	}

	if x4 != 10 {
		t.Error("x4 is false")
	}

	if y4 != 5 {
		t.Error("y4 is false")
	}

	if operator4 != "*" {
		t.Error("operator4 is false")
	}

	// test /
	str5 := "08 / 2 "
	x5, y5, operator5, err5 := parse(&str5)
	if err5 != nil {
		t.Fail()
	}

	if x5 != 8 {
		t.Error("x5 is false")
	}

	if y5 != 2 {
		t.Error("y5 is false")
	}

	if operator5 != "/" {
		t.Error("operator5 is false")
	}

	// test -
	str6 := "9999999 - 77777 "
	x6, y6, operator6, err6 := parse(&str6)
	if err6 != nil {
		t.Fail()
	}

	if x6 != 9999999 {
		t.Error("x6 is false")
	}

	if y6 != 77777 {
		t.Error("y6 is false")
	}

	if operator6 != "-" {
		t.Error("operator6 is false")
	}

}
