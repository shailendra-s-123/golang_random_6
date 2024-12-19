package main  
import (  
  "bufio"
  "fmt"
  "io"
  "os"
  "time"
  // "github.com/kelindar/permutation" - uncomment if using this third-party library
)

const largeBufferSize = 1 << 20 // 1 MB buffer size

func main() {  
  start := time.Now()
  readWriteEfficiently()
  elapsed := time.Since(start)
  fmt.Println("File I/O with slice efficiency completed in:", elapsed)
}
func readWriteEfficiently() {
  var largeBuffer []byte
  //Reader:
  file, err := os.Open("input_large.txt")
  if err != nil {
    fmt.Println("Error opening file for reading:", err)
    return
  }
  defer file.Close()
  reader := bufio.NewReader(file)

  //Writer:
  outFile, err := os.Create("output_large_efficient.txt")
  if err != nil {
    fmt.Println("Error creating file for writing:", err)
    return
  }
  defer outFile.Close()
  writer := bufio.NewWriter(outFile)

  for {
    largeBuffer = make([]byte, 0, largeBufferSize) //reuse this slice
    if n, err := reader.Read(largeBuffer[:cap(largeBuffer)]); err != nil {
      if err == io.EOF {
        break // Reached end of file
      }
      fmt.Println("Error reading data:", err)
      return
    }
    //Resize the slice to actual bytes read to avoid wasting space
    largeBuffer = largeBuffer[:n]
    //Process the buffer (if needed) and write