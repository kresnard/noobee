package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

func RegisterServiceProduct(router fiber.Router, db *gorm.DB, dbSqlx *sqlx.DB) {
	repo := NewRepositoryGormDB(db)
	// repo := NewPostgresSQLXRepository(dbSqlx)

	service := NewService(repo)
	handler := NewHandler(service)

	var productRouter = router.Group("products")
	{
		productRouter.Post("", handler.CreateProduct)
		productRouter.Get("", handler.GetProducts)
		productRouter.Get(":id", handler.GetProduct)
		productRouter.Put(":id", handler.UpdateProduct)
		productRouter.Delete(":id", handler.DeleteProduct)
	}
}
