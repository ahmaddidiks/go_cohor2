package request

import "mime/multipart"

type BookRequest struct {
	Title  string                `form:"title" binding:"required"`
	Author string                `form:"author" binding:"required"`
	Stock  int                   `form:"stock" binding:"required"`
	Image  *multipart.FileHeader `form:"file"`
}
