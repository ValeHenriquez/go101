package models

type Response struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Data    []byte `json:"data"`
}

type Headers struct {
	Authorization string `json:"Authorization"`
}
