package main

import "os"

const (
	removeDirectories = []string{"testing"}
)

func main() {
	removePreviousPipelineRun()
}

func removePreviousPipelineRun() {
	println("Removing User space of previous pipeline runs")
	if err := os.RemoveAll("./testing"); err != nil {
		println("Error removing user space: ", err)
	}
}
