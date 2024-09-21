package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// thanks chat gipity
func caseBlindBinSearch(arr []string, target string) int {
	target = strings.ToLower(target)

	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		midValue := strings.ToLower(arr[mid])

		if midValue == target {
			return mid
		}
		if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func checkSetup(configDir string) bool {
	_, err := os.ReadDir(configDir)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s does not exist, attempting to create...\n", configDir)

			os.Mkdir(configDir, os.ModePerm)

			fmt.Printf("Created %s!\n", configDir)

			return false
		}

		fmt.Printf("Error while checking for installed templates: %s\n", err)

		return false
	}

	return true
}

func setup(configDir string) {
	fmt.Println("Setting up...")

	dirs := make([]string, 0, 2)

	err := filepath.WalkDir(configDir, func(path string, d fs.DirEntry, err error) error {
		if d.Type().IsDir() {
			if len(path) > len(configDir) {
				dirs = append(dirs, path)
			}

			return nil
		}

		if !strings.HasSuffix(path, ".gitignore") {
			return nil
		}

		relative := path[len(configDir)+1:]

		c := 0

		for i := 0; i < len(relative) && c == 0; i++ {
			if relative[i] == '/' {
				c++
			}
		}

		if c > 0 {
			os.Rename(path, configDir+"/"+d.Name())
		}

		return nil
	})

	for _, dir := range dirs {
		os.RemoveAll(dir)
	}

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Println("Finished setting up!")
}

func download(configDir string) {
	cmd := exec.Command(
		"git",
		"clone",
		"https://github.com/github/gitignore",
		configDir,
	)

	fmt.Println("Attempting to download templates...")

	err := cmd.Run()

	if err != nil {
		fmt.Printf("Error while cloning templates: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully downloaded templates!")
}

// thanks chat gipity
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destinationFile.Close()

	if _, err := io.Copy(destinationFile, sourceFile); err != nil {
		return fmt.Errorf("failed to copy file content: %w", err)
	}

	return nil
}

func main() {
	configDir, err := os.UserConfigDir()

	if err != nil {
		fmt.Println("Home folder could not be determine for configuration, exiting...")
		os.Exit(1)
	}

	configDir += "/genignore"

	if !checkSetup(configDir) {
		download(configDir)
		setup(configDir)
	}

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No argument provided, help menu coming soon!")
		os.Exit(0)
	}

	template := args[0]

	files := getFiles()
	actual := caseBlindBinSearch(files, template)

	if actual > -1 {
		cwd, err := os.Getwd()

		if err != nil {
			fmt.Printf("Error while finding cwd: %s, exiting...\n", err)
			os.Exit(1)
		}

		fname := cwd + "/.gitignore"

		fmt.Printf("Found %s, copying to %s...\n", files[actual], fname)

		err = copyFile(configDir+"/"+files[actual]+".gitignore", fname)

		if err != nil {
			fmt.Printf("Error while copying file: %s\n", err)
		} else {
			fmt.Println("Copied!")
		}
	} else {
		fmt.Printf("A .gitignore file for %s was not found :(\n", template)
	}
}
