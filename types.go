package geetest

import "github.com/vimbing/http"

type Solver struct {
	httpClient   *http.Client
	captchaId    string
	callbackName string
}

type em struct {
	Ph int    `json:"ph"`
	Cp int    `json:"cp"`
	Ek string `json:"ek"`
	Wd int    `json:"wd"`
	Nt int    `json:"nt"`
	Si int    `json:"si"`
	Sc int    `json:"sc"`
}

type loadData struct {
	w            payloadW
	payload      string
	processToken string
}
