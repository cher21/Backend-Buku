package controller

import (
	"encoding/json" // package untuk enkode dan mendekode json menjadi struct dan sebaliknya
	"fmt"
	"strconv" // package yang digunakan untuk mengubah string menjadi tipe int

	"log"
	"net/http" // digunakan untuk mengakses objek permintaan dan respons dari api

	"crud/models" //models package dimana Buku didefinisikan

	"github.com/gorilla/mux" // digunakan untuk mendapatkan parameter dari router
	_ "github.com/lib/pq"    // postgres golang driver
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []models.Buku `json:"data"`
}

// func TmbhBuku(w http.ResponseWriter, r *http.Request) {
// 	var buku models.Buku

// 	// Parse data formulir multipart
// 	err := r.ParseMultipartForm(10 << 20) // Batasan 10 MB untuk ukuran file gambar
// 	if err != nil {
// 		log.Fatalf("Tidak dapat mem-parsing formulir: %v", err)
// 	}

// 	// Dapatkan file formulir, abaikan error jika tidak ada file
// 	file, handler, err := r.FormFile("Image")
// 	if err == nil {
// 		defer file.Close()

// 		// Buat file baru di folder "uploads"
// 		f, err := os.OpenFile("uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
// 		if err != nil {
// 			log.Fatalf("Tidak dapat membuka file: %v", err)
// 		}
// 		defer f.Close()

// 		// Salin file ke lokasi baru
// 		_, err = io.Copy(f, file)
// 		if err != nil {
// 			log.Fatalf("Tidak dapat menyalin file: %v", err)
// 		}

// 		buku.Image = handler.Filename
// 	}

// 	// Isi field lainnya
// 	buku.Judul_buku = r.FormValue("Judul_buku")
// 	buku.Penulis = r.FormValue("Penulis")

// 	// Panggil fungsi models untuk memasukkan data buku
// 	insertID := models.TambahBuku(buku)

// 	// Format objek respons
// 	res := response{
// 		ID:      insertID,
// 		Message: "Data buku telah ditambahkan",
// 	}

// 	// Kirim respons
// 	json.NewEncoder(w).Encode(res)
// }

func TmbhBuku(w http.ResponseWriter, r *http.Request) {
	var buku models.Buku

	// Baca data JSON dari body permintaan
	err := json.NewDecoder(r.Body).Decode(&buku)
	if err != nil {
		http.Error(w, "Gagal memparsing JSON", http.StatusBadRequest)
		return
	}

	// Panggil fungsi models untuk memasukkan data buku
	insertID := models.TambahBuku(buku)

	// Format objek respons
	res := response{
		ID:      insertID,
		Message: "Data buku telah ditambahkan",
	}

	// Kirim respons
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// AmbilBuku mengambil single data dengan parameter id
func AmbilBuku(w http.ResponseWriter, r *http.Request) {
	// kita set headernya
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// dapatkan idbuku dari parameter request, keynya adalah "id"
	params := mux.Vars(r)

	// konversi id dari tring ke int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// memanggil models ambilsatubuku dengan parameter id yg nantinya akan mengambil single data
	buku, err := models.AmbilSatuBuku(int64(id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data buku. %v", err)
	}

	// kirim response
	json.NewEncoder(w).Encode(buku)
}

// Ambil semua data buku
func AmbilSemuaBuku(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// memanggil models AmbilSemuaBuku
	bukus, err := models.AmbilSemuaBuku()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response Response
	response.Status = 1
	response.Message = "Success"
	response.Data = bukus

	// kirim semua response
	json.NewEncoder(w).Encode(response)
}

func UpdateBuku(w http.ResponseWriter, r *http.Request) {

	// kita ambil request parameter idnya
	params := mux.Vars(r)

	// konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// buat variable buku dengan type models.Buku
	var buku models.Buku

	// decode json request ke variable buku
	err = json.NewDecoder(r.Body).Decode(&buku)

	if err != nil {
		log.Fatalf("Tidak bisa decode request body.  %v", err)
	}

	// // Parse data formulir multipart
	// err = r.ParseMultipartForm(10 << 20)
	// if err != nil {
	// 	log.Fatalf("Tidak dapat mem-parsing formulir: %v", err)
	// }

	// file, handler, err := r.FormFile("Image")
	// if err != nil {
	// 	log.Fatalf("Error mengambil file: %v", err)
	// }
	// defer file.Close()

	// if handler != nil {
	// 	// Buat file baru di folder "uploads"
	// 	f, err := os.OpenFile("uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	// 	if err != nil {
	// 		log.Fatalf("Tidak dapat membuka file: %v", err)
	// 	}
	// 	defer f.Close()

	// 	// Salin file ke lokasi baru
	// 	_, err = io.Copy(f, file)
	// 	if err != nil {
	// 		log.Fatalf("Tidak dapat menyalin file: %v", err)
	// 	}

	// 	buku.Image = handler.Filename
	// }

	// // Isi field lainnya
	// buku.Judul_buku = r.FormValue("Judul_buku")
	// buku.Penulis = r.FormValue("Penulis")

	// panggil updatebuku untuk mengupdate data
	updatedRows := models.UpdateBuku(int64(id), buku)

	// ini adalah format message berupa string
	msg := fmt.Sprintf("Buku telah berhasil diupdate. Jumlah yang diupdate %v rows/record", updatedRows)

	// ini adalah format response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// kirim berupa response
	json.NewEncoder(w).Encode(res)
}

func HapusBuku(w http.ResponseWriter, r *http.Request) {

	// kita ambil request parameter idnya
	params := mux.Vars(r)

	// konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// panggil fungsi hapusbuku , dan convert int ke int64
	deletedRows := models.HapusBuku(int64(id))

	// ini adalah format message berupa string
	msg := fmt.Sprintf("buku sukses di hapus. Total data yang dihapus %v", deletedRows)

	// ini adalah format reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}
