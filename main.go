package main

import (
	"log"
	"reflect"
)

type ListRecord struct {
	WListId     uint32
	PListId     uint32
	ListName    string
	Description string
	MonitorName string
}

func main() {
	list := ListRecord{
		WListId:     1,
		PListId:     2,
		ListName:    "a",
		Description: "b",
		MonitorName: "c",
	}
	t := reflect.TypeOf(list)
	v := reflect.ValueOf(list)
	for i := 0; i < t.NumField(); i++ {
		log.Printf("%s--%v--%v", t.Field(i).Name, t.Field(i).Type.Name(), v.Field(i).Interface())
	}

}
