package packets

import "time"

type KategoriObat struct {
	ID               int       `json:"id,omitempty"`
	NamaKategoriObat string    `json:"nama_kategori_obat"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Obat             []Obat    `json:"kategori,omitempty"`
}

type Obat struct {
	ID            int            `json:"id"`
	NamaObat      string         `json:"nama_obat"`
	DosisObat     string         `json:"dosis_obat"`
	BentukSediaan string         `json:"bentuk_sediaan"`
	Harga         float32        `json:"harga"`
	Gambar        string         `json:"gambar"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	KategoriObat  []KategoriObat `json:"kategori,omitempty"`
}
