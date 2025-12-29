package entity

type DetailTrxWithJoin struct {
    IDTrx       int
    Kuantitas   int
    HargaTotal  int

    IDLogProduk   int
    NamaProduk    string
    Slug          string
    HargaReseller int
    HargaKonsumen int
    Deskripsi     string
    IDCategory    int 

    IDToko   int
    NamaToko string
    URLFoto  string
    
    NamaCategory string
}