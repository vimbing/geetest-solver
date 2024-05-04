package solver

import (
	"fmt"
	"geetest/internal/utils"

	"github.com/vimbing/http"
)

func New(captchaId string) *Solver {
	return &Solver{
		httpClient:   http.Init(),
		captchaId:    captchaId,
		callbackName: fmt.Sprintf("geetest_%d", utils.GetRandom()),
	}
}
