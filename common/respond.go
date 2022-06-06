package common

type successRes struct {
	Data		interface{} `json:"data"`
}

func NewSuccessRespone(data interface{}) *successRes {
	return &successRes{Data: data}
}

func SimpleSuccessRespone(data interface{}) *successRes {
	return NewSuccessRespone(data)
}