package onepassword

import (
	"crypto/rand"
	"fmt"
	"math/big"
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

// generateRandomNumber generates a random int number from 0 to maxNumber
func generateRandomNumber(maxNumber int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(maxNumber)))
	if err != nil {
		return 0, err
	}

	return int(n.Int64()), nil
}

// generateOpDevice generates a string N chars long, only lower case letters and numbers
func generateOpDevice(n int) (string, error) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	s := make([]rune, n)
	for i := range s {
		randomNumber, err := generateRandomNumber(len(letters))
		if err != nil {
			return "", err
		}

		s[i] = letters[randomNumber]
	}

	return string(s), nil
}
