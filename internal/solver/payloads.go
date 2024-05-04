package solver

type payloadW struct {
	DeviceID     string  `json:"device_id"`
	Em           em      `json:"em"`
	Ep           string  `json:"ep"`
	GeeGuard     any     `json:"gee_guard"`
	Geetest      string  `json:"geetest"`
	Jkvg         string  `json:"jkvg"`
	Lang         string  `json:"lang"`
	LotNumber    string  `json:"lot_number"`
	Passtime     int     `json:"passtime"`
	PowMsg       string  `json:"pow_msg"`
	PowSign      string  `json:"pow_sign"`
	SetLeft      int     `json:"setLeft"`
	Userresponse float64 `json:"userresponse"`
	Yeg6         string  `json:"yeg6"`
}
