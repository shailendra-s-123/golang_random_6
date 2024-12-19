package main  
import (  
    "bufio"
    "fmt"
    "io"
    "os"
)

func main() {  
    // Read data from a file into a slice
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    var data []byte

    // Read data in chunks and append it to the slice
    for {
        chunk := make([]byte, 1024)
        n, err := reader.Read(chunk)
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println("Error reading file:", err)
            return
        }
        data = append(data, chunk[:n]...)
    }

    // Write data from the slice to a file
    outFile, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer outFile.Close()

    writer := bufio.NewWriter(outFile)
    _, err = writer.Write(data)
    if err != nil {
        fmt.Println("Error writing file:", err)
        return
    }
    writer.Flush()
}  