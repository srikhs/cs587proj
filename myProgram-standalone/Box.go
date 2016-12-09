package main

import (
	"ethos/syscall"
	"ethos/ethos"
//	ethosLog "ethos/log"
	"ethos/efmt"
	"log"
	
)

func CreateData(x string, y string) MyType{
	return MyType{x,y}
}


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
	
	p1:= CreateData("Great","Human")
	p2:= CreateData("Humor","Evil")

	efmt.Println("Sai, Input", p1)
	efmt.Println("Sai, Input2", p2)

	status = p1.WriteVar("/user/said3/foo.txt")
	p2.WriteVar("/user/said3/foo1.txt")

	efmt.Println("Sai, status", status)
	
	d:= MyType{"bar","aa"}
	status = d.ReadVar("/user/said3/foo.txt")
	e:= MyType{"s","aa"}
	e.ReadVar("/user/said3/foo1.txt")
		
	efmt.Println("Sai" , status)
	efmt.Println("Sai Read out", d)
	efmt.Println("Sai Read out2",e)

	if d.Field1 == "is"{
	efmt.Println("Sai found the word")
	}else
	{
	efmt.Println("Sai not found")
	}	

	
	}
