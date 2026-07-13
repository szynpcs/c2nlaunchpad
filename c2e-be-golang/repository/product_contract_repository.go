package repository

import (
	"c2n/api/request"
	"c2n/model"
	"fmt"

	"gorm.io/gorm"
)

type ProductContractRepository struct {
	DB *gorm.DB
}

func NewProductContractRepository(db *gorm.DB) *ProductContractRepository {
	return &ProductContractRepository{DB: db}
}

func (r *ProductContractRepository) GetById(productId string) (*model.ProductContract, error) {
	var productContract model.ProductContract

	if err := r.DB.Where("id = ?", productId).Find(&productContract).Error; err != nil {
		return nil, err
	}
	return &productContract, nil
}

func (r *ProductContractRepository) List() ([]model.ProductContract, error) {
	var productContracts []model.ProductContract

	fmt.Printf("Executing List query on table: product_contract\n")

	// 先测试一个简单的 count 查询
	var count int64
	if err := r.DB.Table("product_contract").Count(&count).Error; err != nil {
		fmt.Printf("Count query error: %v\n", err)
		return nil, err
	}
	fmt.Printf("Table product_contract has %d records\n", count)

	// 使用 Model 方法，让 GORM 自动识别表名
	result := r.DB.Model(&model.ProductContract{}).Find(&productContracts)
	if result.Error != nil {
		fmt.Printf("List query error: %v\n", result.Error)
		return nil, result.Error
	}

	fmt.Printf("List query returned %d records, RowsAffected: %d\n", len(productContracts), result.RowsAffected)

	// 如果查询成功但没有扫描到数据，尝试使用原生 SQL
	if len(productContracts) == 0 && count > 0 {
		fmt.Printf("Warning: Query succeeded but no records scanned. Trying raw SQL...\n")
		if err := r.DB.Raw("SELECT * FROM product_contract").Scan(&productContracts).Error; err != nil {
			fmt.Printf("Raw SQL query error: %v\n", err)
		} else {
			fmt.Printf("Raw SQL query returned %d records\n", len(productContracts))
		}
	}

	return productContracts, nil
}

func (r *ProductContractRepository) Update(productContractUpdateRequest *request.ProductContractUpdateRequest) error {
	if err := r.DB.Model(&model.ProductContract{}).Where("id = ?", productContractUpdateRequest.ID).Select(
		"SaleContractAddress",
		"TokenAddress",
		"TokenPriceInPT",
		"TotalTokensSold",
		"SaleEnd",
		"UnlockTime",
		"RegistrationTimeStarts",
		"RegistrationTimeEnds",
		"SaleStart",
	).Updates(model.ProductContract{
		SaleContractAddress:    productContractUpdateRequest.SaleAddress,
		TokenAddress:           productContractUpdateRequest.SaleToken,
		TokenPriceInPT:         productContractUpdateRequest.TokenPriceInPT,
		TotalTokensSold:        productContractUpdateRequest.TotalTokens,
		SaleEnd:                productContractUpdateRequest.SaleEndTime.Time(),
		UnlockTime:             productContractUpdateRequest.TokensUnlockTime.Time(),
		RegistrationTimeStarts: productContractUpdateRequest.RegistrationStart.Time(),
		RegistrationTimeEnds:   productContractUpdateRequest.RegistrationEnd.Time(),
		SaleStart:              productContractUpdateRequest.SaleStartTime.Time(),
	}).Error; err != nil {
		return err
	}
	return nil
}
