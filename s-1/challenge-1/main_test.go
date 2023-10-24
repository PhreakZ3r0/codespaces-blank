package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	var wg sync.WaitGroup
	var testString = "TEST_STRING"
	wg.Add(1)

	go updateMessage(testString, &wg)
	wg.Wait()

	if !strings.Contains(msg, testString) {
		t.Error("Error, msg value is not what it is expected: " + testString)
	}

	
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "TESTING_STRING"

	printMessage()

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "TESTING_STRING") {
		t.Error("output of printMessage expected 'TESTING_STRING', recieved something else")
	}
}

func Test_main(t *testing.T) {

	orderedOutput := make([]string, 3)
	orderedOutput[0] = "Hello, universe!"
	orderedOutput[1] = "Hello, cosmos!"
	orderedOutput[2] = "Hello, world!"

	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	var expectedOutput string
	for _, value := range orderedOutput {
		expectedOutput += value + "/n"
	}

	if !strings.Contains(output, expectedOutput) {
		t.Error("Expected: " + expectedOutput + "\n\nrecieved:" + output)
	}

}
