package helpers

import (
	"fmt"
	"reflect"
)

var people = []Employee{
	{id: 0, nama: "Ahmad", alamat: "Aceh", pekerjaan: "Agen Rahasia", alasan: "Alasan Ahmad"},
	{id: 1, nama: "Beni", alamat: "Bandung", pekerjaan: "Backend", alasan: "Alasan Beni"},
	{id: 2, nama: "Chanif", alamat: "Colomadu", pekerjaan: "Chef", alasan: "Alasan Chanif"},
	{id: 3, nama: "Didik", alamat: "Demak", pekerjaan: "Dokter", alasan: "Alasan Didik"},
	{id: 4, nama: "Eko", alamat: "Empat Lawang", pekerjaan: "Embedded", alasan: "Alasan Eko"},
}

func (p Employee) Show() {
	// show Employee data

	// fmt.Printf("ID: %d\n", p.id)
	// fmt.Printf("nama: %s\n", p.nama)
	// fmt.Printf("alamat: %s\n", p.alamat)
	// fmt.Printf("pekerjaan: %s\n", p.pekerjaan)
	// fmt.Printf("alasan: %s\n", p.alasan)

	ValueOf := reflect.ValueOf(p)
	typeOf := reflect.TypeOf(p)

	for i := 0; i < ValueOf.NumField(); i++ {
		fmt.Printf("%+v : %+v\n", typeOf.Field(i).Name, ValueOf.Field(i))
	}
}

func Shortbyname(name string, p *Employee) bool {
	// return true if desire name is found, otherwise false
	for _, value := range people {
		if value.nama == name {
			*p = value
			return true
		}
	}
	return false
}

func ShortbyID(id int, p *Employee) bool {
	// return true if desire id is found, otherwise false
	for _, value := range people {
		if value.id == id {
			*p = value
			return true
		}
	}
	return false
}
