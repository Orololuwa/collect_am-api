package main

import "testing"

func TestRun(t *testing.T){
	_, _, err := run();
	if err != nil {
		t.Error("failed run")
	}
}