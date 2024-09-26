package main

import (
	"sync"
	"testing"
)

// muMockAppDataDir prevents mocking from causing issues during parallel tests
var muMockAppDataDir sync.Mutex

// withMockAppDataDir replaces getAppDataDir to return dir until unmock() is called
func withMockAppDataDir(t *testing.T) (unmock func()) {
	muMockAppDataDir.Lock()

	f := getAppDataDir
	getAppDataDir = func() (string, error) {
		return t.TempDir(), nil
	}

	return func() {
		getAppDataDir = f
		muMockAppDataDir.Unlock()
	}
}
