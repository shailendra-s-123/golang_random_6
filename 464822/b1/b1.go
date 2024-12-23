package main  
import (  
    "bufio"  
    "fmt"  
    "io"  
    "os"  
)  
const (  
    bufferSize = 64 * 1024  
)  
func main() {  
    // Open the input file for reading  
    inputFile, err := os.Open("input.txt")  
    if err != nil {  
        fmt.Println("Error opening input file:", err)  
        return  
    }  
    defer inputFile.Close()  
    // Create a buffered reader for efficient reading  
    reader := bufio.NewReaderSize(inputFile, bufferSize)  
    // Open the output file for writing  
    outputFile, err := os.Create("output.txt")  
    if err != nil {  
        fmt.Println("Error opening output file:", err)  
        return  
    }  
    defer outputFile.Close()  
    // Create a buffered writer for efficient writing  
    writer := bufio.NewWriterSize(outputFile, bufferSize)  
    // Read and write data in chunks  
    for {  
        buf := make([]byte, bufferSize)  
        n, err := reader.Read(buf)  
        if err != nil {  
            if err == io.EOF {  
                break  
            }  
            fmt.Println("Error reading from file:", err)  
            return  
        }  
        _, err = writer.Write(buf[:n])  
        if err != nil {  
            fmt.Println("Error writing to file:", err)  
            return  
        }  
    }  
    // Flush the writer to ensure all data is written to the file  
    err = writer.Flush()  
    if err != nil {  
        fmt.Println("Error flushing writer:", err)  
        return  
    }  
    fmt.Println("File I/O operation completed successfully.")  
}  