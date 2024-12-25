package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
)

func readMemoryMappedFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}

	fileSize := fileStat.Size()
	if fileSize == 0 {
		return fmt.Errorf("file size is zero")
	}

	// Map the file into memory
	mmap, err := syscall.Mmap(int(file.Fd()), 0, int(fileSize),
		syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_PRIVATE)
	if err != nil {
		return err
	}
	defer syscall.Munmap(mmap)

	// Access the mapped memory
	fmt.Printf("Mapped file content:\n%s\n", string(mmap))

	return nil
}

func main() {
	filePath := "example.txt" // Replace with your file path
	err := readMemoryMappedFile(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading memory-mapped file:", err)
	} else {
		fmt.Println("File content read successfully.")
	}
}