package main

import (
	"ethos/syscall"
	"ethos/ethos"
//	ethosLog "ethos/log"
	"ethos/efmt"
	"log"
	
)


func main () {
	me := syscall.GetUser()
	path:= "/user/"+me +"/myDir/"
	fd, statusw := ethos.OpenDirectoryPath(path)
	if statusw!= syscall.StatusOk{
		log.Fatalf ("Error %v %v", fd, statusw)
	}	
	
	var data MyType
	status := data.CreateDirectory("/user/said3","")
	efmt.Println("Sai",status)
	
	data.Field1 = "How"
	data.Field2 = "Are"

	efmt.Println("Sai, Input", data)

	status = data.WriteVar("/user/said3/foo.txt")
	efmt.Println("Sai, status", status)
	
	d:= MyType{"bar","aa"}
	status = d.ReadVar("/user/said3/foo.txt")
	
		
	efmt.Println("Sai" , status)
	efmt.Println("Sai Read out", d.Field1)

	if d.Field1 == "is"{
	efmt.Println("Sai found the word")
	}else
	{
	efmt.Println("Sai not found")
	}

	}
