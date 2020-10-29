package onepassword

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func getExtension() string {
	if runtime.GOOS == "windows" {
		return ".exe"
	}
	return ""
}

func buildMockOnePassword() (string, error) {
	cmd := exec.Command(
		"go",
		"install",
		"github.com/ContainerSolutions/onepassword/cmd/mock-op")

	if output, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("failed to build mock op program: %s\n%s", err, output)
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = filepath.Join(os.Getenv("HOME"), "go")
	}

	programPath := filepath.Join(
		filepath.SplitList(gopath)[0],
		"bin",
		"mock-op"+getExtension())

	return programPath, nil
}

// randomString generates a string of size N including lower case letters and numbers
func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
