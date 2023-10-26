package request

import "mime/multipart"


type BookRequest struct {
	Title string
	Author string
	Stcok int
	Image *multipart.FileHeader
}