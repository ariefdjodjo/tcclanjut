package migration

import "time"

type Jurusan struct {
	IdJurusan 	int 	`gorm:"primary_key:yes;column:id_jurusan;auto_increment:yes"`
	NamaJurusan	string	`gorm:"column:nama_jurusan"`
}

func (Jurusan) TableName() string {
	return "jurusan"
}

type Mahasiswa struct {
	Nim			int 	`gorm:"primary_key:yes;column:nim;auto_increment:yes"`
	Nama 		string 	`gorm:"column:nama"`
	Jk			int 	`gorm:"column:jk"`
	Alamat 		string 	`gorm:"column:alamat"`
	TempatLahir	string 	`gorm:"column:tempat_lahir"`
	TglLahir 	date 	`gorm:"column:tgl_lahir"`
	Agama 		int 	`gorm:"column:agama"`
}

func (Mahasiswa) TableName() string {
	return "mahasiswa"
}