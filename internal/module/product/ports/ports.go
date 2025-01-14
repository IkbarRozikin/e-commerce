package ports

import (
	"context"
	"product-service/internal/module/product/entity"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req *entity.CreateProductRequest) (entity.UpsertProductResponse, error)
	GetProducts(ctx context.Context, req *entity.GetProductsRequest) (entity.GetProductsResponse, error)
	UpdateProduct(ctx context.Context, req *entity.UpdateProductRequest) (entity.UpsertProductResponse, error)
	DeleteProduct(ctx context.Context, req *entity.DeleteProductRequest) error
	GetProductDetails(ctx context.Context, req *entity.GetProductDetailsByIdRequest) (entity.GetProductDetailsByIdResponse, error)
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, req *entity.CreateProductRequest) (entity.UpsertProductResponse, error)
	GetProducts(ctx context.Context, req *entity.GetProductsRequest) (entity.GetProductsResponse, error)
	UpdateProduct(ctx context.Context, req *entity.UpdateProductRequest) (entity.UpsertProductResponse, error)
	DeleteProduct(ctx context.Context, req *entity.DeleteProductRequest) error
	GetProductDetailsById(ctx context.Context, req *entity.GetProductDetailsByIdRequest) (entity.GetProductDetailsByIdResponse, error)

	IsShopOwner(ctx context.Context, userId, shopId string) (bool, error)
	IsProductOwner(ctx context.Context, userId, productId string) (bool, error)
}
