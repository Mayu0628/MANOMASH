package handler

import (
	"MANOMASH/model"
)

func ResFlgCreate(ReSt int, ReRe string) (model.ResFlg) {
	var ResData model.ResFlg
	ResData.Status = ReSt
	ResData.Result = ReRe
	return ResData
}