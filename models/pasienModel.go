package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/1amkaizen/go_crud/config"
	"github.com/1amkaizen/go_crud/entities"
)

type pasienModel struct {
	conn *sql.DB
}

func NewPasienModel() *pasienModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &pasienModel{
		conn: conn,
	}
}
func (p *pasienModel) FindAll() ([]entities.Pasien, error) {

	rows, err := p.conn.Query("select * from pasien")
	if err != nil {
		return []entities.Pasien{}, err
	}
	defer rows.Close()

	var dataPasien []entities.Pasien
	for rows.Next() {
		var pasien entities.Pasien
		rows.Scan(
			&pasien.Id,
			&pasien.NamaLengkap,
			&pasien.NIK,
			&pasien.JenisKelamin,
			&pasien.TempatLahir,
			&pasien.TanggalLahir,
			&pasien.Alamat,
			&pasien.NoHP)

		if pasien.JenisKelamin == "1" {
			pasien.JenisKelamin = "Laki-Laki"
		} else {
			pasien.JenisKelamin = "Perempuan"
		}

		tgl_lahir, _ := time.Parse("2006-01-02", pasien.TanggalLahir)
		pasien.TanggalLahir = tgl_lahir.Format("02-01-2006")

		dataPasien = append(dataPasien, pasien)
	}
	return dataPasien, nil

}

func (p *pasienModel) Create(pasien entities.Pasien) bool {

	result, err := p.conn.Exec("Insert into pasien (nama_lengkap, nik, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat, nohp) values(?,?,?,?,?,?,?)", pasien.NamaLengkap, pasien.NIK, pasien.JenisKelamin, pasien.TempatLahir, pasien.TanggalLahir, pasien.Alamat, pasien.NoHP)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()
	return lastInsertId > 0
}

func (p *pasienModel) Find(id int64, pasien *entities.Pasien) error {
	return p.conn.QueryRow("select * from pasien where id = ?", id).
		Scan(
			&pasien.Id,
			&pasien.NamaLengkap,
			&pasien.NIK,
			&pasien.JenisKelamin,
			&pasien.TempatLahir,
			&pasien.TanggalLahir,
			&pasien.Alamat,
			&pasien.NoHP,
		)
}

func (p *pasienModel) Update(pasien entities.Pasien) error {
	_, err := p.conn.Exec(
		"Update pasien set  nama_lengkap = ?, nik = ?, jenis_kelamin = ?, tempat_lahir = ?, tanggal_lahir = ?, alamat = ?, nohp = ? where id = ?",
		pasien.NamaLengkap, pasien.NIK, pasien.JenisKelamin, pasien.TempatLahir, pasien.TanggalLahir, pasien.Alamat, pasien.NoHP, pasien.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func (p *pasienModel) Delete(id int64) {
	p.conn.Exec("Delete from pasien where id = ?", id)
}
