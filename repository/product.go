package repository

import(
	"fmt"
	"mall/model"
	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

type ProductRepoInterface interface {
	Get(Product model.Product) (*model.Product, error)
	Exist(Product model.Product) *model.Product
	ExistByProductID(id string) *model.Product
	Add(Product model.Product) (*model.Product, error)
	Edit(Product model.Product) (bool, error)
	Delete(u model.Product) (bool, error)
}

//查询商品
func (repo *ProductRepository) Get(product model.Product) (*model.Product, error) {
	if err := repo.DB.Where(&product).Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

//根据商品名查询商品
func (repo *ProductRepository) Exist(product model.Product) *model.Product {
	if product.ProductName != "" {
		var temp model.Product
		repo.DB.Where("product_name= ?", product.ProductName).First(&temp)
		return &temp
	}
	return nil
}

//根据商品ID查询商品
func (repo *ProductRepository) ExistByProductID(id string) *model.Product {
	var p model.Product
	repo.DB.Where("product_id = ?", id).First(&p)
	return &p
}

//添加商品
func (repo *ProductRepository) Add(product model.Product) (*model.Product, error) {
	exist := repo.Exist(product)
	if exist != nil && exist.ProductName != "" {
		return &product, fmt.Errorf("商品已存在")
	}
	err := repo.DB.Create(product).Error
	if err != nil {
		return nil, fmt.Errorf("商品添加失败")
	}
	return &product, nil
}

//修改商品
func (repo *ProductRepository) Edit(product model.Product) (bool, error) {
	if product.ProductId == "" {
		return false, fmt.Errorf("请传入更新 ID")
	}
	p := &model.Product{}
	err := repo.DB.Model(p).Where("product_id=?", product.ProductId).Updates(map[string]interface{}{
		"product_name": product.ProductName,
		"product_intro": product.ProductIntro,
		"price": product.Price,
		"stock_num": product.StockNum,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

//删除商品
func (repo *ProductRepository) Delete(product model.Product) (bool, error) {

	err := repo.DB.Model(&product).Where("product_id = ?", product.ProductId).Delete(&product).Error
	if err != nil {
		return false, err
	}
	return true, nil
}