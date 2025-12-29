Evermos Social Commerce Service (PBI API)
📌 Latar Belakang
Evermos adalah platform social commerce reseller yang berfokus pada penjualan produk-produk Muslim di Indonesia. Platform ini menyediakan berbagai fitur utama seperti katalog barang, toko online, dan sistem distributor. Proyek ini bertujuan untuk membangun layanan backend sederhana menggunakan Golang dan MySQL yang mampu menangani alur transaksi penjualan secara efisien, aman, dan terukur.

🚀 Fitur Utama
Sistem ini dibangun dengan mengikuti standar industri dan kebutuhan spesifik platform social commerce:

Autentikasi & Keamanan: Implementasi login dan registrasi menggunakan JWT (JSON Web Token).

Manajemen User & Toko: Akun pengguna terintegrasi dengan pembuatan toko secara otomatis saat berhasil mendaftar.

Manajemen Produk: Pengelolaan katalog produk lengkap dengan fitur upload file (foto produk) serta sistem filtering dan pagination.

Sistem Wilayah: Integrasi data wilayah (Provinsi dan Kota) untuk validasi alamat pengiriman.

Manajemen Kategori: Pengelolaan kategori produk yang dibatasi khusus untuk pengguna dengan peran Admin.

Transaksi & Log Produk: Proses transaksi yang aman menggunakan database transaction. Setiap transaksi akan mencatat Log Produk (snapshot data produk saat transaksi terjadi) untuk menjaga riwayat data.

Proteksi Data: Menerapkan Ownership Checker untuk memastikan pengguna tidak dapat mengakses atau mengubah data milik pengguna lain.

🛠️ Teknologi yang Digunakan
Bahasa Pemrograman: Go (v1.25.5).

Web Framework: Gofiber/fiber v2.

ORM: GORM dengan driver MySQL.

Konfigurasi: Viper untuk manajemen .env.

Dokumentasi API: Swagger (Swag).

Lainnya: Zerolog (Logging), Validator v10, dan Gosimple/slug.

⚙️ Persiapan & Instalasi
1. Konfigurasi Environment
Salin konten berikut ke dalam file .env di direktori akar dan sesuaikan dengan kredensial database Anda:

Cuplikan kode

APP_NAME=Evermos-Service
APP_PORT=3000

DB_USER=root
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=evermos_db
DB_MAX_OPEN_CONN=10
DB_MAX_IDLE_CONN=5
DB_CONN_MAX_LIFETIME=3600
DB_CONN_MAX_IDLE_TIME=1800

JWT_SECRET=your_secret_key
JWT_EXPIRE_MINUTES=1440
2. Perintah Makefile
Gunakan Makefile yang tersedia untuk mempermudah operasional:

Menjalankan Aplikasi:

Bash

make run
Migrasi Database (Up):

Bash

make migrate-up
Migrasi Database (Down):

Bash

make migrate-down
📖 Dokumentasi API
Setelah aplikasi berjalan, Anda dapat mengakses dokumentasi API interaktif melalui Swagger di: http://localhost:3000/swagger/index.html

🏗️ Arsitektur Proyek
Proyek ini menerapkan Clean Architecture untuk memastikan kode mudah dipelihara dan diuji:

Controller: Menangani request HTTP dan validasi input awal.

UseCase: Berisi logika bisnis inti (misal: perhitungan harga, validasi stok).

Repository: Berinteraksi langsung dengan database melalui GORM.

Entity & Models: Definisi skema database dan struktur data transfer (DTO).

Middleware: Menangani otorisasi (Admin/User) dan validasi JWT.
