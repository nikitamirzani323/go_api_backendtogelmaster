package helpers

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
	Time    string      `json:"time"`
}
type ResponseListKeluaran struct {
	Status       int         `json:"status"`
	Message      string      `json:"message"`
	Record       interface{} `json:"record"`
	Time         string      `json:"time"`
	Totalwinlose int         `json:"totalwinlose"`
}
type ResponseListPermainan struct {
	Status         int         `json:"status"`
	Message        string      `json:"message"`
	Record         interface{} `json:"record"`
	Time           string      `json:"time"`
	Totalbet       int         `json:"totalbet"`
	Subtotal       int         `json:"subtotal"`
	Subtotalcancel int         `json:"subtotalcancel"`
	Subtotalwin    int         `json:"subtotalwin"`
}

type ErrorResponse struct {
	Field string
	Tag   string
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
