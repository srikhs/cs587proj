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
	p3:= CreateData("Hello","World")
	p4:= CreateData("World","Fun")

	efmt.Println("Sai, Input", p1)
	efmt.Println("Sai, Input2", p2)
	efmt.Println("Sai, Input3", p3)
	efmt.Println("Sai, Input4", p4)
	
	status = p1.WriteVar("/user/said3/foo.txt")
	p2.WriteVar("/user/said3/foo1.txt")
	p3.WriteVar("/user/said3/foo2.txt")
	p4.WriteVar("/user/said3/foo3.txt")
	
	efmt.Println("Sai, status", status)
	
	d:= MyType{"bar","aa"}
	status = d.ReadVar("/user/said3/foo.txt")
	e:= MyType{"s","aa"}
	e.ReadVar("/user/said3/foo1.txt")
	f:= MyType{"s","aa"}
	f.ReadVar("/user/said3/foo2.txt")
	g:= MyType{"s","aa"}
	g.ReadVar("/user/said3/foo2.txt")
		
	efmt.Println("Sai" , status)
	efmt.Println("Sai Read out", d)
	efmt.Println("Sai Read out2", e)
	efmt.Println("Sai Read out4", f)
	efmt.Println("Sai Read out4", g)

	if d.Field1; d.Field2 == "Great"{
	efmt.Println("Sai found the word")
	}else
	{
	efmt.Println("Sai not found")
	}	

	
	}
