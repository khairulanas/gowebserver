package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("ok")
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type Output struct {
	nama                  string
	kemungkinanTahunLahir int
	jumlahKata            int
	panjangNama           int
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./static/index.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.FormValue("name")
		umur := r.FormValue("umur")
		umurint, err := strconv.ParseInt(umur, 6, 12)
		tglLahir := 2020 - umurint
		if name == "" || err != nil {
			if name == "" {
				fmt.Fprintf(w, "nama tidak boleh kosong")
			}
			if err != nil {
				fmt.Fprintf(w, "umur harus angka")
			}
		} else {
			fmt.Fprintf(w, "{nama: %s, kemungkinanTahunLahir: %o, jumlahKata: %o, panjangNama: %o}", name, tglLahir, len(strings.Split(name, " ")), len(name))
		}

		// output := &Output{nama: name, kemungkinanTahunLahir: tglLahir, jumlahKata: 2, panjangNama: 15}
		// js, err := json.Marshal(output)
		// if err != nil {
		// 	fmt.Fprintf(w, "js error")
		// }

		// w.Header().Set("Content-Type", "application/json")
		// w.Write(js)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
