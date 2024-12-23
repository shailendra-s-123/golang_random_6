package main  
import (  
    "fmt"
    "io"
    "os"
)

const bufferSize = 1 << 20 // 1 MB buffer size

func main() {
    sourceFile := "source_file.txt"
    destinationFile := "destination_file.txt"
    // Read data from the source file using a slice with a buffer size
    file, err := os.Open(sourceFile)
    if err != nil {
        fmt.Println("Error opening source file:", err)
        return
    }
    defer file.Close()

    data := make([]byte, bufferSize)
    writer, err := os.Create(destinationFile)
    if err != nil {
        fmt.Println("Error creating destination file:", err)
        return
    }
    defer writer.Close()

    for {
        n, err := file.Read(data)
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println("Error reading data:", err)
            return
        }
        _, err = writer.Write(data[:n])
        if err != nil {
            fmt.Println("Error writing data:", err)
            return
        }
    }

    // Optionally, you can flush the writer to ensure all data is written to the file
    err = writer.Sync()
    if err != nil {
        fmt.Println("Error syncing writer:", err)
    }

    fmt.Println("Data copied successfully!")
}  