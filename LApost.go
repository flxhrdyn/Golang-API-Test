package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type mahasiswa struct {
	KodeMatkul string `json:"kode_mk"`
	NamaMatkul string    `json:"nama_matkul"`
	SKS int `json:"sks"`
}

func PostMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Mhs mahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Mhs); err != nil {
				log.Fatal(err)
			}
		} else {
			
			kode_mk := r.PostFormValue("kode_mk")
			nama_matkul := r.PostFormValue("nama_matkul")
			getSks := r.PostFormValue("sks")
			sks, _ := strconv.Atoi(getSks)
			Mhs = mahasiswa{
				KodeMatkul: kode_mk,
				NamaMatkul: nama_matkul,
				SKS:        sks,
				
			}
		}
		dataMahasiswa, _ := json.Marshal(Mhs) 
		w.Write(dataMahasiswa)                
		return
	}
	http.Error(w, "MAAF ERROR", http.StatusNotFound)
	return
}

func main() {
	http.HandleFunc("/post_mahasiswa", PostMahasiswa)
	fmt.Println("server running...")
	if err := http.ListenAndServe(":7777", nil); err != nil {
		log.Fatal(err)
	}
}
