//go:build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	binaryName = "tetris"
	mainPath   = "./cmd"
)

// Build builds the binary
func Build() error {
	err := sh.Run("go", "build", "-o", binaryName, mainPath)
	if err != nil {
		fmt.Println("❌ Build failed!")
		return err
	}
	fmt.Println("✅ Build succeeded!")
	return nil
}

// Clean removes built artifacts
func Clean() error {
	if _, err := os.Stat(binaryName); err == nil {
		if err := os.Remove(binaryName); err != nil {
			fmt.Println("❌ Clean failed!")
			return err
		}
		fmt.Println("✅ Clean succeeded!")
		return nil
	}
	fmt.Println("✅ Nothing to clean!")
	return nil
}

// Run builds and runs the game
func Run() error {
	mg.Deps(Build)
	binaryPath, err := filepath.Abs(binaryName)
	if err != nil {
		fmt.Println("❌ Run failed!")
		return err
	}
	err = sh.Run(binaryPath)
	if err != nil {
		fmt.Println("❌ Run failed!")
		return err
	}
	fmt.Println("✅ Run succeeded!")
	return nil
}

// Test runs unit tests
func Test() error {
	out, err := sh.Output("go", "test", "./...")
	if err != nil {
		fmt.Println(string(out))
		fmt.Println("❌ Tests failed!")
		return err
	}
	fmt.Println("✅ Tests passed!")
	return nil
}

// Lint runs golangci-lint
func Lint() error {
	if err := sh.Run("which", "golangci-lint"); err != nil {
		fmt.Println("⚠️ golangci-lint not found.")
		fmt.Println("Install it with: mage installLint")
		return nil
	}
	err := sh.RunV("golangci-lint", "run")
	if err != nil {
		fmt.Println("❌ Lint failed!")
		return err
	}
	fmt.Println("✅ Lint passed!")
	return nil
}

// InstallLint installs golangci-lint
func InstallLint() error {
	err := sh.RunV("go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint@latest")
	if err != nil {
		fmt.Println("❌ InstallLint failed!")
		return err
	}
	fmt.Println("✅ InstallLint succeeded!")
	return nil
}

// Format formats the code
func Format() error {
	err := sh.Run("go", "fmt", "./...")
	if err != nil {
		fmt.Println("❌ Format failed!")
		return err
	}
	fmt.Println("✅ Format succeeded!")
	return nil
}

// Vet runs go vet
func Vet() error {
	err := sh.RunV("go", "vet", "./...")
	if err != nil {
		fmt.Println("❌ Vet failed!")
		return err
	}
	fmt.Println("✅ Vet passed!")
	return nil
}

// Check runs formatting, vetting, linting, and tests
func Check() error {
	mg.Deps(Format, Vet, Lint, Test)
	fmt.Println("✅ All checks complete!")
	return nil
}

// All runs the full build pipeline
func All() error {
	mg.Deps(Clean, Check, Build)
	fmt.Println("✅ All steps complete!")
	return nil
}

// Default target
var Default = Build
