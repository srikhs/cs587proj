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
	data.x = 1
	data.y = 2
	data.x1 = 3
	data.y1 = 4
	efmt.Println("The input coordinates are") 
	fd, status := ethos.OpenDirectoryPath(path)
	data.Write(fd)
//	data.WriteVar(path +"foobar")
	efmt.Println(data)
//	efmt.Println("saiiiiiiiiiii")
}
