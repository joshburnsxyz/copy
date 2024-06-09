package main

import (
	"github.com/schollz/progressbar/v3"
	"io"
	"os"
	"path/filepath"
)

func copyFile(src string, dest string, progress bool) error {
	// Open source file
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create target destination
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	file, err := srcFile.Stat()
	if err != nil {
		return err
	}

	// Build progress bar
	var bar *progressbar.ProgressBar
	if progress {
		bar = progressbar.DefaultBytes(
			file.Size(),
			"copying",
		)
	}

	// Use buffer to copy byte chunks to destination
	buf := make([]byte, 1024)
	for {
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if progress {
			bar.Add(n)
		}

		if _, err := destFile.Write(buf[:n]); err != nil {
			return err
		}
	}

	return nil
}

func copyDir(src string, dest string, progress bool) error {
	// Use the filepath.Walk function to iterate over the source directory and
	// offload each found path to the copyFile function
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)

		if err != nil {
			return err
		}

		destPath := filepath.Join(dest, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		return copyFile(path, destPath, progress)
	})
}
