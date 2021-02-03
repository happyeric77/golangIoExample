package reader

import (
	"io/ioutil"
	"strings"
	"fmt"
	"log"
	"io"
	"os"
)

// FileReaderExample uses File type which implement Reader interface as an example to show how Reader.Read() works.
func FileReaderExample() {
	f, err := os.Open("sampleTXT.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 50)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			log.Println("End of file")
			break
		}
		if err != nil {
			log.Println("FILE READING ERROR", err)
		}
		fmt.Println(string(buf[:n]))
	}
}


// ReaderReadVsIoutilReadAllExample uses strings.NewReader to generate a Reader interface type.
func ReaderReadVsIoutilReadAllExample() {	
	
	//NewReader() returns a Reader interface type which implements .Read() method
	sReader := strings.NewReader("The example to generate Reader which contains a piece of string") 

	// Use .Read() method to read through the sReader.
	buf := make([]byte, 10)
	for {
		n, err := sReader.Read(buf)
		if err == io.EOF {
			log.Println("End of file")
			break
		}
		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Println(string(buf[:n]))
	}

	sReader = strings.NewReader("The example to generate Reader which contains a piece of string")

	// Use ioutil.ReadAll() to read the sReader more easily
	content, err := ioutil.ReadAll(sReader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
	
}

