package solver

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"geetest/internal/utils"
	"image/png"
	"math/rand"
	"regexp"

	"github.com/corona10/goimagehash"
	"github.com/google/uuid"
	fhttp "github.com/vimbing/fhttp"
	"github.com/vimbing/gorand"
	"github.com/vimbing/http"
)

func (s *Solver) unmarshalRes(res *http.Response, out any) {
	parsedBody := s.parseGeetestBody(res.BodyString())
	json.Unmarshal([]byte(parsedBody), &out)
}

func (s *Solver) responseValid(res *http.Response) error {
	if res.StatusCode() != 200 {
		return ErrResponseInvalid
	}

	var resJson responseGeneric
	s.unmarshalRes(res, &resJson)

	if resJson.Status == "success" {
		return nil
	}

	return ErrResponseInvalid
}

func (s *Solver) parseGeetestBody(body string) string {
	re := regexp.MustCompile(`geetest_\d+\(`)
	part1 := re.ReplaceAllString(body, "")
	return part1[:len(part1)-1]
}

func (s *Solver) headers() fhttp.Header {
	return fhttp.Header{
		"accept":             {"*/*"},
		"accept-language":    {"pl-PL,pl;q=0.9,en-US;q=0.8,en;q=0.7,la;q=0.6,de;q=0.5"},
		"dnt":                {"1"},
		"sec-ch-ua":          {`"Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"`},
		"sec-ch-ua-mobile":   {"?0"},
		"sec-ch-ua-platform": {`"Linux"`},
		"sec-fetch-dest":     {"script"},
		"sec-fetch-mode":     {"no-cors"},
		"sec-fetch-site":     {"cross-site"},
		"user-agent":         {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"},
	}
}

func (s *Solver) urlWithParams(url string, params map[string]string) string {
	return fmt.Sprintf("%s?%s", url, utils.Params(params))
}

func (s *Solver) getImageSetLeft(bg string) (int, error) {
	res, err := s.httpClient.Get(fmt.Sprintf("https://static.geetest.com/%s", bg), s.headers())

	if err != nil {
		return 0, err
	}

	img, err := png.Decode(bytes.NewReader(res.Body))

	if err != nil {
		return 0, err
	}

	hash, err := goimagehash.PerceptionHash(img)

	if err != nil {
		return 0, err
	}

	if len(hash.ToString()) < 3 {
		return 0, ErrImageNotRecognized
	}

	setLeft, ok := imageIndex[hash.ToString()[2:]]

	if !ok {
		return 0, ErrImageNotRecognized
	}

	return (setLeft - 41), nil
}

func (s *Solver) load() (loadData, error) {
	res, err := s.httpClient.Get(
		s.urlWithParams(
			"https://gcaptcha4.geetest.com/load",
			map[string]string{
				"captcha_id":  s.captchaId,
				"challenge":   uuid.NewString(),
				"client_type": "web",
				"lang":        "pl",
				"callback":    s.callbackName,
			},
		),
		s.headers(),
	)

	if err != nil {
		return loadData{}, err
	}

	if err := s.responseValid(res); err != nil {
		return loadData{}, err
	}

	var resJson responseLoad
	s.unmarshalRes(res, &resJson)

	powMsg := fmt.Sprintf(
		"1|0|md5|%s|%s|%s||%s",
		resJson.Data.PowDetail.Datetime,
		s.captchaId,
		resJson.Data.LotNumber,
		utils.GetGuid(),
	)

	imgSetLeft, err := s.getImageSetLeft(resJson.Data.Bg)

	if err != nil {
		return loadData{}, err
	}

	return loadData{
		w: payloadW{
			DeviceID: utils.Md5(fmt.Sprint(rand.Float64())),
			Em: em{
				Ph: 0,
				Cp: 0,
				Ek: "11",
				Wd: 1,
				Nt: 0,
				Si: 0,
				Sc: 0,
			},
			Ep:           "123",
			GeeGuard:     nil,
			Geetest:      "captcha",
			Jkvg:         "342414482",
			Lang:         "zh",
			LotNumber:    resJson.Data.LotNumber,
			Passtime:     gorand.RandomInt(500, 700),
			PowMsg:       powMsg,
			PowSign:      utils.Md5(powMsg),
			SetLeft:      imgSetLeft,
			Userresponse: float64(imgSetLeft) + rand.Float64(),
			Yeg6:         "d6w9",
		},
		payload:      resJson.Data.Payload,
		processToken: resJson.Data.ProcessToken,
	}, nil
}

func (s *Solver) verify(lData loadData) (string, error) {
	encryptedW, err := utils.EncryptPayload(utils.HardMarshal[string](lData.w))

	if err != nil {
		return "", err
	}

	res, err := s.httpClient.Get(
		s.urlWithParams(
			"https://gcaptcha4.geetest.com/verify",
			map[string]string{
				"callback":         s.callbackName,
				"captcha_id":       s.captchaId,
				"client_type":      "web",
				"lot_number":       lData.w.LotNumber,
				"payload":          lData.payload,
				"process_token":    lData.processToken,
				"payload_protocol": "1",
				"pt":               "1",
				"w":                encryptedW,
			},
		),
		s.headers(),
	)

	if err != nil {
		return "", err
	}

	if err := s.responseValid(res); err != nil {
		return "", err
	}

	var resJson responseVerify
	s.unmarshalRes(res, &resJson)

	secCodeString, err := json.Marshal(resJson.Data.Seccode)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(secCodeString), nil
}

func (s *Solver) Solve() (string, error) {
	wPayload, err := s.load()

	if err != nil {
		return "", err
	}

	return s.verify(wPayload)
}
