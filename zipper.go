// +build zipper

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	// "runtime"
)

func buildExecutable(goos, goarch, outputName string) error {
	fmt.Printf("Building for %s/%s...\n", goos, goarch)

	// Set environment variables for cross-compilation
	cmd := exec.Command("go", "build", "-o", outputName, "main.go")
	cmd.Env = append(os.Environ(),
		"GOOS="+goos,
		"GOARCH="+goarch,
	)

	// Run the build command
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build for %s/%s: %w", goos, goarch, err)
	}
	fmt.Printf("Successfully built %s\n", outputName)
	return nil
}

func zipFile(sourceFile, destinationZip string) error {
	// Create a new zip file
	zipFile, err := os.Create(destinationZip)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %w", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = addFileToZip(zipWriter, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to add file to zip: %w", err)
	}

	fmt.Printf("Created %s successfully\n", destinationZip)
	return nil
}

func addFileToZip(zipWriter *zip.Writer, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return fmt.Errorf("failed to create zip header: %w", err)
	}
	header.Name = filepath.Base(filePath)
	header.Method = zip.Store             

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return fmt.Errorf("failed to create zip header writer: %w", err)
	}

	_, err = io.Copy(writer, file)
	if err != nil {
		return fmt.Errorf("failed to write file to zip: %w", err)
	}

	return nil
}

func main() {
	targets := []struct {
		goos, goarch, output, zipName string
	}{
		{"windows", "amd64", "torri-windows.exe", "torri-windows.zip"},
		{"linux", "amd64", "torri-linux", "torri-linux.zip"},
		{"darwin", "amd64", "torri-macos", "torri-macos.zip"},
	}

	for _, target := range targets {
		err := buildExecutable(target.goos, target.goarch, target.output)
		if err != nil {
			fmt.Printf("Error building %s: %v\n", target.output, err)
			continue
		}

		err = zipFile(target.output, target.zipName)
		if err != nil {
			fmt.Printf("Error zipping %s: %v\n", target.output, err)
			continue
		}

		err = os.Remove(target.output)
		if err != nil {
			fmt.Printf("Error removing %s after zipping: %v\n", target.output, err)
		}
	}
}
