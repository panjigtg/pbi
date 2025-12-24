package models

type (
	TokoRequest struct {
		IDUser   int    `json:"id_user"`
		NamaToko string `json:"nama_toko"`
		UrlFoto  string `json:"url_foto"`
	}
)
