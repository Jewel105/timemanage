package response

type PageResponse struct {
	Page int         `json:"page"`
	Size int         `json:"size"`
	Data interface{} `json:"data"`
}
