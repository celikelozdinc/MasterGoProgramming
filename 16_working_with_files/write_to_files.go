package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	byteSlice := []byte("I will write 21 bytes")
	fd, err := os.OpenFile("new.log", os.O_CREATE|os.O_WRONLY, 0666)
	defer fd.Close()
	if err != nil {
		log.Fatal(err)
	}

	bytesWritten, err := fd.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("#%d bytes are written", bytesWritten)
	}

	writeErr := ioutil.WriteFile("newest.log", byteSlice, 0666)
	if writeErr != nil {
		log.Fatal(err)
	}

	log.Println(strings.Repeat("@", 20))
	fd, _ = os.OpenFile("buffered_writer.log", os.O_CREATE|os.O_WRONLY, 0666)
	defer fd.Close()

	bufferedWriter := bufio.NewWriter(fd)

	bs := []byte{97, 98, 99}

	bytesAvail := bufferedWriter.Available()
	log.Printf("#%d bytes are avail before writing\n", bytesAvail)

	bytesWritten, _ = bufferedWriter.Write(bs)
	log.Printf("#%d bytes are written to buffer, not to file\n", bytesWritten)

	bytesAvail = bufferedWriter.Available()
	log.Printf("#%d bytes are avail after writing\n", bytesAvail)

	unflushedBufSize := bufferedWriter.Buffered()
	log.Printf("Unflushed buffer size (# bytes waiting to be written to disk) : %d\n", unflushedBufSize)

	bufferedWriter.Flush()
}

