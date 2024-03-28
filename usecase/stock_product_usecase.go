package usecase

import (
	"errors"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	"medioker-bank/repository"
)

type StockProductUseCase interface {
	CreateStockProduct(payload dto.StockProductCreateDto) (model.StockProduct, error)
	FindStockProductById(id string) (model.StockProduct, error)
	FindAllStockProducts() ([]model.StockProduct, error)
	FindStockProductsByQuery(payload dto.StockProductSearchByQueryDto) ([]model.StockProduct, error)
	UpdateStockProducts(id string, payload dto.StockProductUpdateDto) (model.StockProduct, error)
	DeleteStockProducts(id string) (model.StockProduct, error)
}

type stockProductUseCase struct {
	repo repository.StockProductRepository
}

func (s *stockProductUseCase) CreateStockProduct(payload dto.StockProductCreateDto) (model.StockProduct, error) {
	riskLists := []string{"low", "medium", "high", "very high"}
	riskValidate := false
	for _, risk := range riskLists {
		if risk == payload.Risk {
			riskValidate = true
		}
	}
	if !riskValidate {
		return model.StockProduct{}, errors.New("allowed risk: 'low', 'medium', 'high', 'very high'")
	}
	stockProduct, err := s.repo.Create(payload)
	if err != nil {
		return model.StockProduct{}, err
	}
	return stockProduct, nil
}

func (s *stockProductUseCase) FindStockProductById(id string) (model.StockProduct, error) {
	stockProduct, err := s.repo.FindById(id)
	if err != nil {
		return model.StockProduct{}, err
	}
	return stockProduct, nil
}

func (s *stockProductUseCase) FindAllStockProducts() ([]model.StockProduct, error) {
	var stockProducts []model.StockProduct
	var err error
	stockProducts, err = s.repo.FindAll()
	if err != nil {
		return []model.StockProduct{}, err
	}
	return stockProducts, nil
}

func (s *stockProductUseCase) FindStockProductsByQuery(payload dto.StockProductSearchByQueryDto) ([]model.StockProduct, error) {
	var stockProducts []model.StockProduct
	var err error
	stockProducts, err = s.repo.FindByQuery(payload)
	if err != nil {
		return []model.StockProduct{}, err
	}
	return stockProducts, nil
}

func (s *stockProductUseCase) UpdateStockProducts(id string, payload dto.StockProductUpdateDto) (model.StockProduct, error) {
	riskLists := []string{"low", "medium", "high", "very high"}
	riskValidate := false
	for _, risk := range riskLists {
		if risk == payload.Risk {
			riskValidate = true
		}
	}
	if !riskValidate {
		return model.StockProduct{}, errors.New("allowed risk: 'low', 'medium', 'high', 'very high'")
	}
	if payload.Rating <= 0 || payload.Rating > 5 {
		return model.StockProduct{}, errors.New("allowed rating: 1 - 5")
	}
	stockProduct, err := s.repo.Update(id, payload)
	if err != nil {
		return model.StockProduct{}, err
	}
	return stockProduct, nil
}

func (s *stockProductUseCase) DeleteStockProducts(id string) (model.StockProduct, error) {
	stockProduct, err := s.repo.Delete(id)
	if err != nil {
		return model.StockProduct{}, err
	}
	return stockProduct, nil
}

func NewStockProductUseCase(repo repository.StockProductRepository) StockProductUseCase {
	return &stockProductUseCase{repo: repo}
}
