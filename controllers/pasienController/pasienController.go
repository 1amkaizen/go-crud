package pasienController

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/1amkaizen/go_crud/entities"
	"github.com/1amkaizen/go_crud/liblaries"
	"github.com/1amkaizen/go_crud/models"
)

var pasienModel = models.NewPasienModel()
var validation = liblaries.NewValidation()

func Index(response http.ResponseWriter, request *http.Request) {

	pasien, _ := pasienModel.FindAll()

	data := map[string]interface{}{
		"pasien": pasien,
	}

	temp, err := template.ParseFiles("views/pasien/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(response, data)

}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/pasien/add.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var pasien entities.Pasien
		pasien.NamaLengkap = request.Form.Get("nama_lengkap")
		pasien.NIK = request.Form.Get("nik")
		pasien.JenisKelamin = request.Form.Get("jenis_kelamin")
		pasien.TempatLahir = request.Form.Get("tempat_lahir")
		pasien.TanggalLahir = request.Form.Get("tanggal_lahir")
		pasien.Alamat = request.Form.Get("alamat")
		pasien.NoHP = request.Form.Get("nohp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(pasien)
		if vErrors != nil {
			data["pasien"] = pasien
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data pasien berhasil di simpan"
			pasienModel.Create(pasien)

		}

		temp, _ := template.ParseFiles("views/pasien/add.html")
		temp.Execute(response, data)

	}
}

func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var pasien entities.Pasien
		pasienModel.Find(id, &pasien)

		data := map[string]interface{}{
			"pasien": pasien,
		}

		temp, err := template.ParseFiles("views/pasien/edit.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(response, data)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var pasien entities.Pasien
		pasien.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		pasien.NamaLengkap = request.Form.Get("nama_lengkap")
		pasien.NIK = request.Form.Get("nik")
		pasien.JenisKelamin = request.Form.Get("jenis_kelamin")
		pasien.TempatLahir = request.Form.Get("tempat_lahir")
		pasien.TanggalLahir = request.Form.Get("tanggal_lahir")
		pasien.Alamat = request.Form.Get("alamat")
		pasien.NoHP = request.Form.Get("nohp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(pasien)
		if vErrors != nil {
			data["pasien"] = pasien
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data pasien berhasil di perbarui"
			pasienModel.Update(pasien)

		}

		temp, _ := template.ParseFiles("views/pasien/edit.html")
		temp.Execute(response, data)

	}
}
func Delete(response http.ResponseWriter, request *http.Request) {
	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	pasienModel.Delete(id)

	http.Redirect(response, request, "/pasien", http.StatusSeeOther)

}
