package geetest

type responseGeneric struct {
	Status string `json:"status"`
}

type responseLoad struct {
	Status string `json:"status"`
	Data   struct {
		LotNumber   string `json:"lot_number"`
		CaptchaType string `json:"captcha_type"`
		Slice       string `json:"slice"`
		Bg          string `json:"bg"`
		Ypos        int    `json:"ypos"`
		Arrow       string `json:"arrow"`
		Js          string `json:"js"`
		CSS         string `json:"css"`
		StaticPath  string `json:"static_path"`
		GctPath     string `json:"gct_path"`
		ShowVoice   bool   `json:"show_voice"`
		Feedback    string `json:"feedback"`
		Logo        bool   `json:"logo"`
		Pt          string `json:"pt"`
		CaptchaMode string `json:"captcha_mode"`
		Guard       bool   `json:"guard"`
		CheckDevice bool   `json:"check_device"`
		Language    string `json:"language"`
		CustomTheme struct {
			Style      string `json:"_style"`
			Color      string `json:"_color"`
			Gradient   string `json:"_gradient"`
			Hover      string `json:"_hover"`
			Brightness string `json:"_brightness"`
			Radius     string `json:"_radius"`
		} `json:"custom_theme"`
		PowDetail struct {
			Version  string `json:"version"`
			Bits     int    `json:"bits"`
			Datetime string `json:"datetime"`
			Hashfunc string `json:"hashfunc"`
		} `json:"pow_detail"`
		Payload         string `json:"payload"`
		ProcessToken    string `json:"process_token"`
		PayloadProtocol int    `json:"payload_protocol"`
	} `json:"data"`
}

type responseVerify struct {
	Status string `json:"status"`
	Data   struct {
		LotNumber string `json:"lot_number"`
		Result    string `json:"result"`
		FailCount int    `json:"fail_count"`
		Seccode   struct {
			CaptchaID     string `json:"captcha_id"`
			LotNumber     string `json:"lot_number"`
			PassToken     string `json:"pass_token"`
			GenTime       string `json:"gen_time"`
			CaptchaOutput string `json:"captcha_output"`
		} `json:"seccode"`
		Score           string `json:"score"`
		Payload         string `json:"payload"`
		ProcessToken    string `json:"process_token"`
		PayloadProtocol int    `json:"payload_protocol"`
	} `json:"data"`
}
