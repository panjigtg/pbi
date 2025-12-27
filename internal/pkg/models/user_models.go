package models

type (
	RegisterRequest struct {
		Nama        string `json:"nama"`
		Email       string `json:"email"`
		NoTelp      string `json:"notelp"`
		KataSandi   string `json:"kata_sandi"`
		IDProvinsi  string `json:"id_provinsi"`
		IDKota      string `json:"id_kota"`
		TanggalLahir string `json:"tanggal_lahir"` // YYYY-MM-DD
		JenisKelamin string `json:"jenis_kelamin"`
		Tentang      string `json:"tentang"`
		Pekerjaan    string `json:"pekerjaan"`
	}

	LoginRequest struct {
		Email     string `json:"email"`
		KataSandi string `json:"kata_sandi"`
	}

	LoginResponse struct {
		Nama         string           `json:"nama"`
		NoTelp       string           `json:"no_telp"`
		TanggalLahir string           `json:"tanggal_Lahir"`
		Tentang      string           `json:"tentang"`
		Pekerjaan    string           `json:"pekerjaan"`
		Email        string           `json:"email"`
		IDProvinsi   ProvinceResponse `json:"id_provinsi"`
		IDKota       CityResponse     `json:"id_kota"`
		Token        string           `json:"token"`
	}
	UserResponse struct {
		Nama   			string 			`json:"nama"`
		Email  			string 			`json:"email"`
		NoTelp 			string 			`json:"notelp"`
		TanggalLahir	string 			`json:"tanggal_lahir"`
		JenisKelamin 	string 			`json:"jenis_kelamin"`
		Tentang      	string 			`json:"tentang"`
		Pekerjaan   	string 			`json:"pekerjaan"`
		IDProvinsi    	ProvinceResponse `json:"id_provinsi"`
		IDKota       	CityResponse     `json:"id_kota"`
	}

	UpdateProfileRequest struct {
		Nama   			*string 			`json:"nama"`
		Email  			*string 			`json:"email"`
		KataSandi   	*string 			`json:"kata_sandi"`
		NoTelp 			*string 			`json:"notelp"`
		TanggalLahir	*string 			`json:"tanggal_lahir"`
		JenisKelamin 	*string 			`json:"jenis_kelamin"`
		Tentang      	*string 			`json:"tentang"`
		Pekerjaan   	*string 			`json:"pekerjaan"`
		IDProvinsi     	*string 			`json:"id_provinsi"`
		IDKota          *string 			`json:"id_kota"`
	}

	Profile struct {
		Nama         string           `json:"nama"`
		Email        string           `json:"email"`
		NoTelp       string           `json:"no_telp"`
		TanggalLahir string           `json:"tanggal_lahir"`
		JenisKelamin string           `json:"jenis_kelamin"`
		Tentang      string           `json:"tentang"`
		Pekerjaan    string           `json:"pekerjaan"`
		IDProvinsi   ProvinceResponse `json:"id_provinsi"`
		IDKota       CityResponse     `json:"id_kota"`
	}

)

