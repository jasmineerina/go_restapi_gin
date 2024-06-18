package models

// menyimpan table product

type Product struct {
	Id          int64  `gorm:"PrimaryKey" json:id`
	NamaProduct string `gorm:"type:varchar(300)" json:nama_product`
	Deskripsi   string `gorm:"type:text" json:deskripsi`
}
