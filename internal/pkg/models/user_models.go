package models

type (
	RegisterRequest struct {
		Nama       string `json:"nama" validate:"required"`
		Email      string `json:"email" validate:"required,email"`
		NoTelp     string `json:"notelp" validate:"required"`
		KataSandi  string `json:"kata_sandi" validate:"required,min=6"`
		IDProvinsi string `json:"id_provinsi" validate:"required"`
		IDKota     string `json:"id_kota" validate:"required"`
	}

	LoginRequest struct {
		Email     string `json:"email"`
		KataSandi string `json:"kata_sandi"`
	}

	LoginResponse struct {
		User  UserResponse `json:"user"`
		Token string       `json:"token"`
	}

	UserResponse struct {
		ID     int    `json:"id"`
		Nama   string `json:"nama"`
		Email  string `json:"email"`
		NoTelp string `json:"notelp"`
	}
)
