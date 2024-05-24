package packets

import (
	"encoding/json"
	"time"
)

type Barang struct {
	Id             int              `json:"id" gorm:"primaryKey"`
	NamaBarang     string           `json:"nama_barang" gorm:"type:varchar(50);not null"`
	JumlahStok     uint             `json:"jumlah_stok"`
	Harga          float64          `json:"harga" gorm:"type:float;not null"`
	Gambar         string           `json:"gambar" gorm:"type:text"`
	Deskripsi      string           `json:"deskripsi" gorm:"type:text;not null"`
	CreatedAt      time.Time        `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	KategoriBarang []KategoriBarang `json:"kategoi_barang,omitempty" gorm:"many2many:kategorisasi"`
}

type KategoriBarang struct {
	Id           int       `json:"id" gorm:"PrimaryKey"`
	NamaKategori string    `json:"nama_kategori_barang" gorm:"type:varchar(50);not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Barang       []Barang  `json:"barang,omitempty" gorm:"many2many:kategorisasi"`
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
