package commodityController

type QueryRequest struct {
	Page *int `json:"page" binding:"required"`
	Size *int `json:"size" binding:"required"`
}

type UpdateRequest struct {
	ID             int64    `json:"id" binding:"required"`
	Name           *string  `json:"name" binding:"required"`
	Type           *string  `json:"type" binding:"required"`
	Specifications *string  `json:"specifications" binding:"required"`
	Quantity       *int     `json:"quantity" binding:"required"`
	Price          *float64 `json:"price" binding:"required"`
	CartImage      *string  `json:"cart_image" binding:"required"`
}

type DeleteRequest struct {
	ID int64 `json:"id" binding:"required"`
}
