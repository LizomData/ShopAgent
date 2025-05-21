package commodityController

type QueryRequest struct {
	Page *int    `json:"page" binding:"required"`
	Size *int    `json:"size" binding:"required"`
	Name *string `json:"name"`
}

type DeleteRequest struct {
	ID int64 `json:"id,string" binding:"required"`
}
