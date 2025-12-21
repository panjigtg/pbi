
CREATE TABLE category (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama_category VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(255),
    kata_sandi VARCHAR(255),
    notelp VARCHAR(255) UNIQUE,
    tanggal_lahir DATE,
    jenis_kelamin VARCHAR(255),
    tentang TEXT,
    pekerjaan VARCHAR(255),
    email VARCHAR(255),
    id_provinsi VARCHAR(255),
    id_kota VARCHAR(255),
    is_admin BOOLEAN DEFAULT FALSE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE alamat (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_user INT NOT NULL,
    judul_alamat VARCHAR(255),
    nama_penerima VARCHAR(255),
    no_telp VARCHAR(255),
    detail_alamat VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_user) REFERENCES user(id)
);


CREATE TABLE toko (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_user INT NOT NULL,
    nama_toko VARCHAR(255),
    url_foto VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_user) REFERENCES user(id)
);


CREATE TABLE produk (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_toko INT NOT NULL,
    id_category INT NOT NULL,
    nama_produk VARCHAR(255),
    slug VARCHAR(255),
    harga_reseller VARCHAR(255),
    harga_konsumen VARCHAR(255),
    stok INT,
    deskripsi TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_toko) REFERENCES toko(id),
    FOREIGN KEY (id_category) REFERENCES category(id)
);

CREATE TABLE foto_produk (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_produk INT NOT NULL,
    url VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_produk) REFERENCES produk(id)
);

CREATE TABLE log_produk (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_produk INT NOT NULL,
    id_toko INT NOT NULL,
    id_category INT NOT NULL,
    nama_produk VARCHAR(255),
    slug VARCHAR(255),
    harga_reseller VARCHAR(255),
    harga_konsumen VARCHAR(255),
    deskripsi TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_produk) REFERENCES produk(id),
    FOREIGN KEY (id_toko) REFERENCES toko(id),
    FOREIGN KEY (id_category) REFERENCES category(id)
);

CREATE TABLE trx (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_user INT NOT NULL,
    alamat_pengiriman INT NOT NULL,
    harga_total INT,
    kode_invoice VARCHAR(255),
    metode_bayar VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_user) REFERENCES user(id),
    FOREIGN KEY (alamat_pengiriman) REFERENCES alamat(id)
);

CREATE TABLE detail_trx (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_trx INT NOT NULL,
    id_log_produk INT NOT NULL,
    id_toko INT NOT NULL,
    kuantitas INT,
    harga_total INT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_trx) REFERENCES trx(id),
    FOREIGN KEY (id_log_produk) REFERENCES log_produk(id),
    FOREIGN KEY (id_toko) REFERENCES toko(id)
);
