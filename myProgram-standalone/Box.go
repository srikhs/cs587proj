package main 

import (
	"ethos/syscall"
	"ethos/ethos"
	ethosLog "ethos/log"
	"ethos/efmt"
)

func main () {
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
	status := ethosLog.RedirectToLog("myProgram")
	
	if status != syscall.StatusOk {
		efmt.Fprintf (syscall.Stderr, "Error opening %v: %v\n", path, status)
	syscall.Exit(syscall.StatusOk)
	}


	data    := MyType {}
	data.x = 4
	data.y = 2
	data.x1 = 10
	data.y1 = 4
	efmt.Println("The input coordinates are") 
	fd, status := ethos.OpenDirectoryPath(path)
	data.Write(fd)
//	data.WriteVar(path +"foobar")
	efmt.Println(data)
//	ypoints := data.y1 - data.y
//	xpoints := data.x1 - data.x
	//slope := ypoints/xpoints
	efmt.Println(data.y1 - data.y/data.x1 - data.x)
//	efmt.Println("saiiiiiiiiiii")
}
