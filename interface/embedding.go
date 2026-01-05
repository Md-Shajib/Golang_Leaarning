package interfaces

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	content string
}

func (f *File) Read() string {
	return f.content
}

func (f *File) Write(data string) {
	f.content = data
}

func Process(rw ReadWriter){
	rw.Write("This is writing test")
	fmt.Println(rw.Read())
}

func EmbeddingMain(){
	fmt.Println("--------------Embedding Interface-------------------")
	f := &File{}
	Process(f)
}