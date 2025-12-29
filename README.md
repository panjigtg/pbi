# Evermos Social Commerce Service (PBI API)

## 📌 Latar Belakang
Evermos adalah platform **social commerce reseller** yang berfokus pada penjualan produk-produk Muslim di Indonesia.  
Platform ini menyediakan berbagai fitur utama seperti **katalog barang**, **toko online**, dan **sistem distributor**.

Proyek ini bertujuan untuk membangun **layanan backend sederhana** menggunakan **Golang** dan **MySQL** yang mampu menangani alur transaksi penjualan secara **efisien**, **aman**, dan **terukur**.

---

## 🚀 Fitur Utama
Sistem ini dibangun dengan mengikuti standar industri dan kebutuhan spesifik platform social commerce:

### 🔐 Autentikasi & Keamanan
- Login dan registrasi menggunakan **JWT (JSON Web Token)**

### 👤 Manajemen User & Toko
- Akun pengguna otomatis terintegrasi dengan pembuatan **toko** saat registrasi

### 📦 Manajemen Produk
- CRUD produk
- Upload file (foto produk)
- Filtering & pagination

### 🌍 Sistem Wilayah
- Integrasi data **Provinsi** dan **Kota**
- Validasi alamat pengiriman

### 🗂️ Manajemen Kategori
- Khusus **Admin**
- Pembatasan akses berbasis role

### 💰 Transaksi & Log Produk
- Proses transaksi menggunakan **database transaction**
- Pencatatan **Log Produk (snapshot data)** untuk menjaga konsistensi riwayat transaksi

### 🛡️ Proteksi Data
- **Ownership Checker**
- Mencegah user mengakses atau memodifikasi data milik user lain

---

## 🛠️ Teknologi yang Digunakan

| Komponen | Teknologi |
|--------|----------|
| Bahasa Pemrograman | Go `v1.25.5` |
| Web Framework | Fiber v2 |
| ORM | GORM (MySQL Driver) |
| Konfigurasi | Viper |
| Dokumentasi API | Swagger (Swag) |
| Logging | Zerolog |
| Validasi | Validator v10 |
| Utility | gosimple/slug |

---

## ⚙️ Persiapan & Instalasi

## ⚙️ Environment Setup

Salin file environment contoh berikut:

```bash
cp .env.example .env

