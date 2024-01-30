package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestFiltro(t *testing.T) {
	RunTest(t, []string{"e", "3", "r", "4", "Y", "3"}, "eeerrrrYYY")
	RunTest(t, []string{"R", "3", "A", "1", "B", "10"}, "RRRABBBBBBBBBB")
	RunTest(t, []string{"W", "2"}, "WW")
	RunTest(t, []string{"r", "3", "P", "4"}, "rrrPPPP")
}

func RunTest(t *testing.T, input []string, result string) {
	tempFile, err := ioutil.TempFile("", "out")
	if err != nil {
		t.Error("Error while creating temporary file")
	}
	defer tempFile.Close()

	tempFile, os.Stdout = os.Stdout, tempFile

	os.Args = append([]string{"esercizio_filtro.go"}, input...)

	main()

	_, err = os.Stdout.Seek(0, 0)
	if err != nil {
		t.Error("Error while resetting file pointer temporary file")
	}
	output, err := ioutil.ReadAll(os.Stdout)
	if err != nil {
		t.Error("Error while reading output temporary file")
	}
	if tOut, tRes := strings.Trim(string(output), "\n"), strings.Trim(result, "\n"); tOut != tRes {
		t.Error(fmt.Sprintf("\nCodifica compressa stringa: %s\nExpected result: <%s>\nGiven result: <%s>\n", strings.Join(os.Args[1:], " "), tRes, tOut))
	}

	tempFile, os.Stdout = os.Stdout, tempFile
}
