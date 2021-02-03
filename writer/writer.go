package writer

import (
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