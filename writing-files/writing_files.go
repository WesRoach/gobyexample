package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello, world!\n")
	/* creates file, writes, closes file */
	err := ioutil.WriteFile("./dat1.txt", d1, os.ModeAppend)
	check(err)

	/* immediately creates file and filehandler */
	file2, err := os.Create("./dat2.txt")
	check(err)

	/* runs at end of func main() */
	defer file2.Close()

	byte_string := []byte{115, 111, 109, 101, 10} /* "some" */

	var bytes_written int

	/* writes data into file2 handler */
	bytes_written, err = file2.Write(byte_string)
	check(err)
	fmt.Printf("Wrote %d bytes\n", bytes_written)

	bytes_written, err = file2.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", bytes_written)

	file2.Sync()

	writer := bufio.NewWriter(file2)
	bytes_written, err = writer.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", bytes_written)

	writer.Flush()

}
