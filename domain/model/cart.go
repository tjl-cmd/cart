package model

type Cart struct {
	Id        int64 `json:"id" gorm:"primary_key;not_null;auto_increment"`
	ProductId int64 `json:"product_id" gorm:"not_null"`
	Num       int64 `json:"num" gorm:"not_null"`
	SizeId    int64 `json:"size_id" gorm:"not_null"`
	UserId    int64 `json:"user_id" gorm:"not_null"`
}
