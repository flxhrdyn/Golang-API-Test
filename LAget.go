package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Mahasiswa struct {
	KodeMatkul string `json:"kode_mk"`
	NamaMatkul string    `json:"nama_matkul"`
	SKS int `json:"sks"`
}

func NewMahasiswa() []Mahasiswa {
	mhs := []Mahasiswa{
		Mahasiswa{
			KodeMatkul: "AP123",
			NamaMatkul: "Algoritma dan Pemrograman",
			SKS : 3,
		},
		Mahasiswa{
			KodeMatkul: "MI123",
			NamaMatkul: "Matematika Informatika",
			SKS: 2,
		},
		Mahasiswa{
			KodeMatkul: "SD123",
			NamaMatkul: "Struktur Data",
			SKS: 2,
		},
	}
	return mhs
}

func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		mhs := NewMahasiswa()
		datamahasiswa, err := json.Marshal(mhs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(datamahasiswa)
		return
	}
	http.Error(w, "MAAF ERROR", http.StatusNotFound)
}
func main() {
	http.HandleFunc("/mahasiswa", GetMahasiswa)
	fmt.Println("server running...")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}
}