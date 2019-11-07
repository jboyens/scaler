// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Build

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "scaler", ".")
	return cmd.Run()
}

func Test() error {
	mg.Deps(InstallDeps)
	fmt.Println("Running tests...")
	return sh.RunV("go", "run", "github.com/onsi/ginkgo/ginkgo", "--randomizeAllSpecs", "--randomizeSuites", "--failOnPending", "--cover", "--trace", "--race", "--progress", "-r")
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	cmd := exec.Command("go", "get")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("scaler")
}
