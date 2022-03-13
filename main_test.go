package main

import (
	"os"
	"testing"

)

func TestRemovePreviousPipelineRun(t *testing.T) {
	if _, e := os.Stat("./testing"); os.IsNotExist(e) {
		println("Testing directory doesn't exist. Adding it")
		e := os.Mkdir("./testing", 0755)
		if e != nil {
			println("Error creating testing directory: ", e)
		}
		_, e = os.Create("./testing/test.txt")
		if e != nil {
			println("Error creating test file: ", e)
		}
	} else {
		println("Testing directory exists. Running test")
	}

	removePreviousPipelineRun()

	if _, e := os.Stat("./testing"); os.IsNotExist(e) {
		println("Testing directory is gone")
	} else {
		println("Testing directory is still there")
		t.Fail()
	}
}
