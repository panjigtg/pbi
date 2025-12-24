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
		// User  UserResponse `json:"user"`
		ID    int          `json:"id"`
		Nama  string       `json:"nama"`
		Token string       `json:"token"`
	}

	UserResponse struct {
		ID     int    `json:"id"`
		Nama   string `json:"nama"`
		Email  string `json:"email"`
		NoTelp string `json:"notelp"`
		TanggalLahir string `json:"tanggal_lahir"`
		JenisKelamin string `json:"jenis_kelamin"`
		Tentang      string `json:"tentang"`
		Pekerjaan    string `json:"pekerjaan"`
	}
)
