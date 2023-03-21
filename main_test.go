package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// setup
	os.Exit(m.Run())
	// teardown
}

func TestCombineLog(t *testing.T) {
	os.Setenv("LOGFILEDIR", "testlogs")
	os.Setenv("LOGFILEREGEX", ".*leetprog.*")
	os.Setenv("LOGLINEREGEX2", "impThing.*")
	os.Setenv("LOGLINEREGEX3", "erbe.*")
	os.Setenv("LOGFILEOUTPUT", "testlogs/combined.log")
	combineLog()
}
