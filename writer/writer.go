package writer

import (
	"io/ioutil"
	"net"
	"encoding/json"
	"fmt"
	"os"
)

// WriterExample is an Example using Write() method to write content into the Writer interface object (os.File)
func WriterExample() {
	f, err := os.OpenFile("writerSampleTXT.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.Write([]byte("Write this string into os.File(Writer)\n"))
	if err != nil {
		panic(err)
	}
}

// FmtFprintVsWriterWrite uses fmt.Fprint(Writer) to do the same functionality as Writer.Write()
func FmtFprintVsWriterWrite() {

	f, err := os.OpenFile("FprintSampleTXT.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	_, err = fmt.Fprintln(f, "Write this string into os.File(Writer) by Fprint method")
	if err != nil {
		panic(err)
	}
}

func WriterAdvancedExample() {
	f, err := os.OpenFile("WriterAdvancedSampleTXT.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	foo := Foo{"Eric", "18"}
	err = json.NewEncoder(f).Encode(foo)
	if err != nil {
		panic(err)
	}
}

type Foo struct {
	Name string `json:"name"`
	Age string `json:age` 
}

func NetConnReaderWriter() {
	// net.Dail() returns a net.Conn type which implement Read() and Write()
	netConn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		panic(err)
	}
	defer netConn.Close()

	// We can use Write() or Fprint to write the net command which we want to implement such as "GET" or "POST"
	netConn.Write([]byte("GET HTTP 1.0 \r\n\r\n")) // we can also use better way fmt.Fprint(netConn, "GET HTTP 1.0 \r\n\r\n")

	// Create a File by os.OpenFile() as a Writer
	f, err := os.OpenFile("ReadWriteSumExample.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Pass net.Conn as Reader into ioutil.ReadAll. It returns the data in the net.Conn
	connContent, err := ioutil.ReadAll(netConn)
	if err != nil {
		panic(err)
	}

	// Write the connContent as []byte into File as Writer.
	_, err = f.Write(connContent)
	if err != nil {
		panic(err)
	}

}