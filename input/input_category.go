package input

type CategoryInput struct {
	Name string `json:"name" binding:"required"`
}

type GetinputCategoryID struct {
	ID int `uri:"id" binding:"required"`
}

// type GetCategoryID struct {
// 	ID int `uri:"categoryID" binding:"required"`
// }
