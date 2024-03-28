package repository

import (
	"database/sql"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	"medioker-bank/utils/common"
	"time"
)

type StockProductRepository interface {
	Create(payload dto.StockProductCreateDto) (model.StockProduct, error)
	FindById(id string) (model.StockProduct, error)
	FindAll() ([]model.StockProduct, error)
	FindByQuery(payload dto.StockProductSearchByQueryDto) ([]model.StockProduct, error)
	Update(id string, payload dto.StockProductUpdateDto) (model.StockProduct, error)
	Delete(id string) (model.StockProduct, error)
}

type stockProductRepository struct {
	db *sql.DB
}

func (s *stockProductRepository) Create(payload dto.StockProductCreateDto) (model.StockProduct, error) {
	var stockProduct model.StockProduct

	err := s.db.QueryRow(common.CreateStockProduct, payload.CompanyName, 1, payload.Risk, time.Now(), time.Now()).Scan(
		&stockProduct.Id, &stockProduct.CompanyName, &stockProduct.Rating, &stockProduct.Risk, &stockProduct.CreatedAt, &stockProduct.UpdatedAt,
	)
	if err != nil {
		return model.StockProduct{}, err
	}

	return stockProduct, nil
}

func (s *stockProductRepository) FindById(id string) (model.StockProduct, error) {
	var stockProduct model.StockProduct

	err := s.db.QueryRow(common.FindStockProductById, id).Scan(
		&stockProduct.Id, &stockProduct.CompanyName, &stockProduct.Rating, &stockProduct.Risk, &stockProduct.CreatedAt, &stockProduct.UpdatedAt,
	)
	if err != nil {
		return model.StockProduct{}, err
	}

	return stockProduct, nil
}

func (s *stockProductRepository) FindAll() ([]model.StockProduct, error) {
	var stockProducts []model.StockProduct

	rows, err := s.db.Query(common.FindAllStockProducts)
	if err != nil {
		return []model.StockProduct{}, err
	}

	for rows.Next() {
		var stockProduct model.StockProduct
		err := rows.Scan(
			&stockProduct.Id, &stockProduct.CompanyName, &stockProduct.Rating, &stockProduct.Risk, &stockProduct.CreatedAt, &stockProduct.UpdatedAt,
		)

		if err != nil {
			return []model.StockProduct{}, err
		}

		stockProducts = append(stockProducts, stockProduct)
	}

	return stockProducts, nil
}

func (s *stockProductRepository) FindByQuery(payload dto.StockProductSearchByQueryDto) ([]model.StockProduct, error) {
	var stockProducts []model.StockProduct
	var stmt *sql.Stmt
	var err error
	var args []any

	if payload.Rating == 0 {
		stmt, err = s.db.Prepare(common.FindStockProductsByQueryRisk)
		if err != nil {
			return []model.StockProduct{}, err
		}
		args = append(args, payload.Risk)
	} else if payload.Risk == "" {
		stmt, err = s.db.Prepare(common.FindStockProductsByQueryRating)
		if err != nil {
			return []model.StockProduct{}, err
		}
		args = append(args, payload.Rating)
	} else {
		stmt, err = s.db.Prepare(common.FindStockProductsByQueryBoth)
		if err != nil {
			return []model.StockProduct{}, err
		}
		args = append(args, payload.Rating, payload.Risk)
	}

	rows, err := stmt.Query(args...)
	if err != nil {
		return []model.StockProduct{}, err
	}

	for rows.Next() {
		var stockProduct model.StockProduct
		err := rows.Scan(
			&stockProduct.Id, &stockProduct.CompanyName, &stockProduct.Rating, &stockProduct.Risk, &stockProduct.CreatedAt, &stockProduct.UpdatedAt,
		)

		if err != nil {
			return []model.StockProduct{}, err
		}

		stockProducts = append(stockProducts, stockProduct)
	}

	return stockProducts, nil
}

func (s *stockProductRepository) Update(id string, payload dto.StockProductUpdateDto) (model.StockProduct, error) {
	var stockProduct model.StockProduct

	err := s.db.QueryRow(common.UpdateStockProduct, payload.CompanyName, payload.Rating, payload.Risk, time.Now(), id).Scan(
		&stockProduct.Id, &stockProduct.CompanyName, &stockProduct.Rating, &stockProduct.Risk, &stockProduct.CreatedAt, &stockProduct.UpdatedAt,
	)
	if err != nil {
		return model.StockProduct{}, err
	}

	return stockProduct, nil
}

func (s *stockProductRepository) Delete(id string) (model.StockProduct, error) {
	var stockProduct model.StockProduct

	err := s.db.QueryRow(common.DeleteStockProduct, id).Scan(
		&stockProduct.Id, &stockProduct.CompanyName, &stockProduct.Rating, &stockProduct.Risk, &stockProduct.CreatedAt, &stockProduct.UpdatedAt,
	)
	if err != nil {
		return model.StockProduct{}, err
	}

	return stockProduct, nil
}

func NewStockProductRepository(db *sql.DB) StockProductRepository {
	return &stockProductRepository{db: db}
}
