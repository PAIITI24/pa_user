package packets

import (
	"encoding/json"
	"time"
)

type Barang struct {
	Id             int              `json:"id"`
	NamaBarang     string           `json:"nama_barang"`
	JumlahStok     uint             `json:"jumlah_stok"`
	Harga          float64          `json:"harga"`
	Gambar         string           `json:"gambar"`
	Deskripsi      string           `json:"deskripsi"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
	KategoriBarang []KategoriBarang `json:"kategoi_barang,omitempty"`
}

type KategoriBarang struct {
	Id                 int       `json:"id"`
	NamaKategoriBarang string    `json:"nama_kategori_barang"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Barang             []Barang  `json:"barang,omitempty"`
}

func (B *KategoriBarang) MarshalJSON() ([]byte, error) {
	type Alias KategoriBarang

	return json.Marshal(&struct {
		CreatedAt string `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
		UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
		*Alias
	}{
		CreatedAt: B.CreatedAt.Format("02/01/2006"),
		UpdatedAt: B.CreatedAt.Format("02/01/2006"),
		Alias:     (*Alias)(B),
	})
}

func (B *Barang) MarshalJSON() ([]byte, error) {
	type Alias Barang

	return json.Marshal(&struct {
		CreatedAt string `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
		UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
		*Alias
	}{
		CreatedAt: B.CreatedAt.Format("02/01/2006"),
		UpdatedAt: B.CreatedAt.Format("02/01/2006"),
		Alias:     (*Alias)(B),
	})
}
