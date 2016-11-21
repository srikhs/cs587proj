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
	data.x = 2
	data.y = 3
	data.x1 = 6
	data.y1 = 8
	efmt.Println("The input coordinates are") 
	fd, status := ethos.OpenDirectoryPath(path)
	data.Write(fd)
//	data.WriteVar(path +"foobar")
	efmt.Println(data)
//	efmt.Println("saiiiiiiiiiii")
}
