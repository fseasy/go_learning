package main

import (
    "os"
    "io"
    "bufio"
    "log"
    "fmt"
)

func main() {

    const path string = "main.go"
    inputF, err := os.Open(path)
    if err != nil {
        log.Fatal(err)    
    }
    defer func(){
        if err := inputF.Close(); err != nil {
            log.Fatal(err)
        }
    }()

    // read by buffered (binary read)
    reader := bufio.NewReader(inputF)
    for {
        buf := make([]byte, 10)
        readLen, err := reader.Read(buf)
        if err == io.EOF && readLen == 0 {
            break
        }
        fmt.Printf("%s", string(buf[:readLen]))
    }

    inputF.Seek(0, 0)
    fmt.Println()
    content := make([]string, 0)
    // read by line
    scanner := bufio.NewScanner(inputF)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        line := scanner.Text()
        content = append(content, line)
        fmt.Println(line)
    }

    outFile, err := os.OpenFile("tmp.out", os.O_CREATE, 0755)
    if err != nil {
        log.Fatal(err)
    }
    defer func(){
        if err := outFile.Close(); err != nil {
            log.Fatal(err)
        }
    }()

    writer := bufio.NewWriter(outFile)
    for i := range content {
        line := content[i]
        writer.WriteString(line)
        fmt.Fprintln(writer)
    }
    writer.Flush()

}