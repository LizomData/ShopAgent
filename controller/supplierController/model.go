package supplierController

type QueryRequest struct {
	Page *int `json:"page" binding:"required"`
	Size *int `json:"size" binding:"required"`
}

type DeleteRequest struct {
	ID int64 `json:"id" binding:"required"`
}
