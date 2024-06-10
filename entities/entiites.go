package entities

import "time"

type User struct {
	ID    int
	Nama_buku   string
	Waktu_pengambilan time.Time
	Nama_peminjam string
}