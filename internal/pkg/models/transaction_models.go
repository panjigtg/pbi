package models


type TokoBasicInfo struct {
    NamaToko string `json:"nama_toko"`
    URLFoto  string `json:"url_foto"`
}

type TokoInfo struct {
    ID       int    `json:"id"`
    NamaToko string `json:"nama_toko"`
    URLFoto  string `json:"url_foto"`
}

type CategoryInfo struct {
    ID           int    `json:"id"`
    NamaCategory string `json:"nama_category"`
}

type PhotoInfo struct {
    ID        int    `json:"id"`
    ProductID int    `json:"product_id"`
    URL       string `json:"url"`
}

type (

	ProductDetailInTransaction struct {
		ID            int           `json:"id"`
		NamaProduk    string        `json:"nama_produk"`
		Slug          string        `json:"slug"`
		HargaReseller int           `json:"harga_reseler"` 
		HargaKonsumen int           `json:"harga_konsumen"`
		Deskripsi     string        `json:"deskripsi"`
		Toko          TokoBasicInfo `json:"toko"`     
		Category      CategoryInfo  `json:"category"` 
		Photos        []PhotoInfo   `json:"photos"`   
	}

	CreateTrxItemRequest struct {
		ProductID int `json:"product_id" validate:"required"`
		Kuantitas int `json:"kuantitas" validate:"required,min=1"`
	}

	CreateTrxRequest struct {
		MethodBayar string                 `json:"method_bayar" validate:"required,oneof=ovo dana gopay cod"`
		AlamatKirim int                    `json:"alamat_kirim" validate:"required"`
		DetailTrx   []CreateTrxItemRequest `json:"detail_trx" validate:"required,min=1,dive"`
	}

	TransactionListResponse struct {
		ID           	int    						`json:"id"`
		HargaTotal   	float64     						`json:"harga_total"`
		KodeInvoice  	string 						`json:"kode_invoice"`
		MethodBayar 	string 						`json:"method_bayar"`
		AlamatKirim 	DestinationResponse 		`json:"alamat_kirim"`
		DetailTrx    	[]TransactionDetailResponse `json:"detail_trx"`
	}

	TransactionDetailResponse struct {
		Product    ProductDetailInTransaction `json:"product"`
		Toko       TokoInfo                   `json:"toko"`
		Kuantitas  int                        `json:"kuantitas"`
		HargaTotal int                        `json:"harga_total"`
	}

	TransactionDetailByIDResponse struct {
		ID          int                         `json:"id"`
		HargaTotal  float64                     `json:"harga_total"`
		KodeInvoice string                      `json:"kode_invoice"`
		MethodBayar string                      `json:"method_bayar"`
		AlamatKirim DestinationResponse         `json:"alamat_kirim"`
		DetailTrx   []TransactionDetailResponse `json:"detail_trx"`
	}
	
)

type TransactionListResponseWrapper struct {
    Data  []TransactionListResponse `json:"data"`
    Page  int                       `json:"page"`
    Limit int                       `json:"limit"`
}