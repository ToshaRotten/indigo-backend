package storage

import (
	"context"

	"main/db/ent"
	"main/db/ent/product"
	"main/db/entity"
)

func NewProductStorage(client *ent.Client) productStorage {
	return productStorage{
		Client: client,
	}
}

type productStorage struct {
	Client *ent.Client
}

func (p productStorage) Get(ctx context.Context, productID int) (entity.ProductModel, error) {
	product, err := p.Client.Product.Query().Where(product.IDEQ(productID)).Only(ctx)
	if err != nil {
		return entity.ProductModel{}, err
	}

	productData := entity.ProductModel{
		Name:               product.Name,
		Weight:             product.Weight,
		ProductComposition: product.ProductComposition,
		MinStorageTemp:     product.MinStorageTemp,
		MaxStorageTemp:     product.MaxStorageTemp,
		ShelfLife:          product.ShelfLife,
		Picture:            product.Picture,
	}

	return productData, nil
}

func (p productStorage) GetAll(ctx context.Context) ([]entity.ProductModel, error) {
	products, err := p.Client.Product.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var productData []entity.ProductModel
	for _, product := range products {
		productData = append(productData, entity.ProductModel{
			Name:               product.Name,
			Weight:             product.Weight,
			ProductComposition: product.ProductComposition,
			MinStorageTemp:     product.MinStorageTemp,
			MaxStorageTemp:     product.MaxStorageTemp,
			ShelfLife:          product.ShelfLife,
			Picture:            product.Picture,
		})
	}

	return productData, nil
}

func (p productStorage) GetByID(ctx context.Context, id int) (entity.ProductModel, error) {
	product, err := p.Client.Product.Get(ctx, id)
	if err != nil {
		return entity.ProductModel{}, err
	}

	productData := entity.ProductModel{
		Name:               product.Name,
		Weight:             product.Weight,
		ProductComposition: product.ProductComposition,
		MinStorageTemp:     product.MinStorageTemp,
		MaxStorageTemp:     product.MaxStorageTemp,
		ShelfLife:          product.ShelfLife,
	}

	return productData, nil
}

func (p productStorage) Create(ctx context.Context, product entity.ProductModel) error {
	create := p.Client.Product.Create().
		SetMaxStorageTemp(product.MaxStorageTemp).
		SetMaxStorageTemp(product.MaxStorageTemp).
		SetName(product.Name).
		SetProductCategoryID(product.ProductCategoryID)

	return create.Exec(ctx)
}

func (p productStorage) Update(ctx context.Context, product entity.ProductModel) error {
	update := p.Client.Product.UpdateOneID(product.ID).
		SetMaxStorageTemp(product.MaxStorageTemp).
		SetMaxStorageTemp(product.MaxStorageTemp).
		SetName(product.Name).
		SetProductCategoryID(product.ProductCategoryID)

	return update.Exec(ctx)
}

func (p productStorage) Delete(ctx context.Context, id int) error {
	delete := p.Client.Product.DeleteOneID(id)

	return delete.Exec(ctx)
}
