package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Mahasiswa struct {
	ID	 int	  `json:"id"`
	NPM  int	  `json:"npm"`
	Nama string   `json:"nama"`
}

func NewMahasiswa() []Mahasiswa {
	mhs := []Mahasiswa{
		Mahasiswa{
			ID:	  1,
			NPM:  50421957,
			Nama: "Abut",
		},
		Mahasiswa{
			ID:   2,
			NPM:  50414022,
			Nama: "Abuy",
		},
		Mahasiswa{
			ID:   3,
			NPM:  50421234,
			Nama: "Abah",
		},
	}
	return mhs
}
func GetMahasiwswa(w http.ResponseWriter, 
  r *http.Request) {
 	if r.Method == "GET" {
		mhs := NewMahasiswa()
		datamahasiswa, err := json.Marshal(mhs)
		if err != nil {
			http.Error(w, err.Error(),
		    http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/son")
		w.WriteHeader(http.StatusOK)
		w.Write(datamahasiswa)
		return
	}
	http.Error(w, "kamu nanyeaa", 
	http.StatusNotFound)
  }

  func main() {
	http.HandleFunc("/mahasiswa", GetMahasiwswa)
	fmt.Println("server running...")
	if err := 
	http.ListenAndServe(":8000", nil);
	err != nil {
		log.Fatal(err)
	}
  }