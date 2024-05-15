package packets

import "time"

type Produk struct {
	Id             int              `json:"id"`
	NamaProduk     string           `json:"nama_produk"`
	Harga          float64          `json:"harga"`
	Gambar         string           `json:"gambar"`
	Deskripsi      string           `json:"deskripsi"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
	KategoriProduk []KategoriProduk `json:"kategori_produk,omitempty"`
}

type KategoriProduk struct {
	Id           int       `json:"id"`
	NamaKategori string    `json:"nama_kategori_produk"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Produk       []Produk  `json:"produk,omitempty"`
}
