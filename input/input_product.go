package input

type ProductInput struct {
	Name        string `json:"name" binding:"required"`
	RentCost    int    `json:"rent_cost" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID  int    `json:"category_id" binding:"required"`
}

// type GetinputProductID struct {
// 	ID int `uri:"id" binding:"required"`
// }

// type GetCategoryID struct {
// 	ID int `uri:"categoryID" binding:"required"`
// }
