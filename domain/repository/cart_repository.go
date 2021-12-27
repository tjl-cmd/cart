package repository

import (
	"cart/domain/model"
	"errors"
	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	InitTable() error
	FindCartByID(int64) (*model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int64) ([]model.Cart, error)
	ClearCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CateRepository{mysqlDb: db}
}

type CateRepository struct {
	mysqlDb *gorm.DB
}

func (c *CateRepository) InitTable() error {
	return c.mysqlDb.CreateTable(&model.Cart{}).Error
}

func (c *CateRepository) FindCartByID(i int64) (cart *model.Cart, err error) {
	cart = &model.Cart{}
	return cart, c.mysqlDb.First(cart, i).Error
}

func (c *CateRepository) CreateCart(cart *model.Cart) (int64, error) {
	db := c.mysqlDb.FirstOrCreate(cart, model.Cart{ProductId: cart.ProductId, SizeId: cart.SizeId, UserId: cart.UserId})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("购物车插入失败")
	}
	return cart.Id, nil
}

func (c *CateRepository) DeleteCartByID(i int64) error {
	return c.mysqlDb.Where("id = ?", i).Delete(&model.Cart{}).Error
}

func (c *CateRepository) UpdateCart(cart *model.Cart) error {
	return c.mysqlDb.Model(cart).Update(cart).Error
}

func (c *CateRepository) FindAll(userId int64) (cartAll []model.Cart, err error) {
	return cartAll, c.mysqlDb.Where("user_id =?", userId).Find(&cartAll).Error
}

func (c *CateRepository) ClearCart(i int64) error {
	return c.mysqlDb.Where("user_id = ?", i).Delete(&model.Cart{}).Error
}

func (c *CateRepository) IncrNum(id int64, num int64) error {
	cart := &model.Cart{
		Id: id,
	}
	return c.mysqlDb.Model(cart).UpdateColumn("num + ?", gorm.Expr("num + ?", num)).Error
}

func (c *CateRepository) DecrNum(CartId int64, num int64) error {
	cart := &model.Cart{Id: CartId}
	db := c.mysqlDb.Model(cart).Where("num >= ?", num).UpdateColumn("num", gorm.Expr("num - ?", num))
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("减少失败")
	}
	return nil
}
