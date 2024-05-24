package packets

import (
	"encoding/json"
	"time"
)

type KategoriObat struct {
	ID               int       `json:"id,omitempty"`
	NamaKategoriObat string    `json:"nama_kategori_obat"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Obat             []Obat    `json:"obat,omitempty"`
}

type Obat struct {
	ID            int            `json:"id"`
	NamaObat      string         `json:"nama_obat"`
	JumlahStok    uint           `json:"jumlah_stok"`
	DosisObat     string         `json:"dosis_obat"`
	BentukSediaan string         `json:"bentuk_sediaan"`
	Harga         float32        `json:"harga"`
	Gambar        string         `json:"gambar"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	KategoriObat  []KategoriObat `json:"kategori,omitempty"`
}

func (K *KategoriObat) MarshalJSON() ([]byte, error) {
	type Alias KategoriObat

	return json.Marshal(&struct {
		ID               int    `json:"id,omitempty" gorm:"primary_key;auto_increment"`
		NamaKategoriObat string `json:"nama_kategori_obat" gorm:"type:varchar(50)"`
		CreatedAt        string `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
		UpdatedAt        string `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
		Obat             []Obat `json:"kategori,omitempty" gorm:"many2many:kategorisasi"`
	}{
		ID:               K.ID,
		NamaKategoriObat: K.NamaKategoriObat,
		CreatedAt:        K.CreatedAt.Format("02-01-2006"),
		UpdatedAt:        K.UpdatedAt.Format("02-01-2006"),
		Obat:             K.Obat,
	})
}

func (O *Obat) MarshalJSON() ([]byte, error) {
	type Alias Obat
	return json.Marshal(&struct {
		ID            int            `json:"id" gorm:"primaryKey,autoIncrement"`
		NamaObat      string         `json:"nama_obat" gorm:"type:varchar(50)"`
		JumlahStok    uint           `json:"jumlah_stok" gorm:"type:int"`
		DosisObat     string         `json:"dosis_obat" gorm:"type:varchar(50)"`
		BentukSediaan string         `json:"bentuk_sediaan" gorm:"type:varchar(50)"`
		Harga         float32        `json:"harga" gorm:"type:float"`
		Gambar        string         `json:"gambar" gorm:"type:text"`
		CreatedAt     string         `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
		UpdatedAt     string         `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
		KategoriObat  []KategoriObat `json:"kategori,omitempty" gorm:"many2many:kategorisasi"`
	}{
		ID:            O.ID,
		NamaObat:      O.NamaObat,
		JumlahStok:    O.JumlahStok,
		DosisObat:     O.DosisObat,
		BentukSediaan: O.BentukSediaan,
		Harga:         O.Harga,
		Gambar:        O.Gambar,
		CreatedAt:     O.CreatedAt.Format("02-01-2006"),
		UpdatedAt:     O.UpdatedAt.Format("02-01-2006"),
		KategoriObat:  O.KategoriObat,
	})
}
