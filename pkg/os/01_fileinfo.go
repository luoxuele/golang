package main

import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stat("test/readme.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(info.Name(),info.Size(),info.Mode(),info.IsDir())
	//fmt.Printf("%#v \n",info)

	
}


/*
fileinfo
// A FileInfo describes a file and is returned by Stat.
type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	Mode() FileMode     // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() interface{}   // underlying data source (can return nil)
}
*/

