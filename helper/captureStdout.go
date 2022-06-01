package helper

import (
	"io/ioutil"
	"os"
	"testing"
)

// source:
// https://groups.google.com/g/golang-nuts/c/hVUtoeyNL7Y
// from mfried...@gmail.com

func dieOn(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

func CatchStdOut(t *testing.T, runnable func()) string {

	realStdout := os.Stdout
	defer func() { os.Stdout = realStdout }()

	r, fakeStdout, err := os.Pipe()
	dieOn(err, t)
	os.Stdout = fakeStdout
	runnable()
	// need to close here, otherwise ReadAll never gets "EOF".
	dieOn(fakeStdout.Close(), t)
	newOutBytes, err := ioutil.ReadAll(r)
	dieOn(err, t)
	dieOn(r.Close(), t)
	return string(newOutBytes)
}
