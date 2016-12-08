package main 

import (
	"ethos/syscall"
	"ethos/ethos"
	ethosLog "ethos/log"
	"ethos/efmt"
	"math"

    	"os"
    	"log"
)

func New(x int64, y int64) MyType{
	return MyType{x,y}
}

func (data1 MyType) Distance(data2 MyType) float64{
	first := math.Pow(float64(data2.x - data1.x),2)
	second := math.Pow(float64(data2.y - data1.y),2)
	return math.Sqrt(first + second)
} 

func main () {
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
	status := ethosLog.RedirectToLog("myProgram")
	
	if status != syscall.StatusOk {
		efmt.Fprintf (syscall.Stderr, "Error opening %v: %v\n", path, status)
	syscall.Exit(syscall.StatusOk)
	}


//comment

	b, err := ioutil.ReadFile("one.txt") // just pass the file name
    if err != nil {
        efmt.Print(err)
    }

    efmt.Println(b) // print the content as 'bytes'

    str := string(b) // convert content to a 'string'

    efmt.Println(str) // print the content as a 'string'
	
	fd, status := ethos.OpenDirectoryPath(path)
	p1.Write(fd)
	p1.WriteVar(path +"foobar")
	efmt.Println("The desired operations have been performed")
}
