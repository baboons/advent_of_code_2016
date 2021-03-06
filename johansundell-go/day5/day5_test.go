package main

import "testing"

func Test_Ex1(t *testing.T) {
	input := "abc3231929"
	if hash(input) != "1" {
		t.Fail()
	}
}

func Test_Ex2(t *testing.T) {
	input := "abc5017308"
	if hash(input) != "8" {
		t.Fail()
	}
}

func Test_Ex3(t *testing.T) {
	input := "abc5278568"
	if hash(input) != "f" {
		t.Fail()
	}
}

func Test_Ex4(t *testing.T) {
	input := "abc"
	if findPassword(input) != "18f47a30" {
		t.Fail()
	}
}

func Test_Ex5(t *testing.T) {
	input := "abc3231929"
	b, pos, _ := hashVersion2(input)
	if string(b) != "5" || pos != 1 {
		t.Fail()
	}
}

func Test_Ex6(t *testing.T) {
	input := "abc5017308"
	str, pos, err := hashVersion2(input)

	if str != 0 || pos != 0 || err == nil {
		t.Fail()
	}
}

func Test_Ex7(t *testing.T) {
	input := "abc5357525"
	b, pos, _ := hashVersion2(input)
	if string(b) != "e" || pos != 4 {
		t.Fail()
	}
}

func Test_Ex8(t *testing.T) {
	input := "abc"
	correct := "05ace8e3"
	result := findPasswordVersion2(input)
	if result != correct {
		t.Error("expected " + correct + " got " + result)
	}
}
