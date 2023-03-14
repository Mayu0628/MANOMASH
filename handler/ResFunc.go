package handler

import (
	"MANOMASH/model"
)

func ResFlgCreate(ReSt int, ReRe string, ReId uint) (model.ResFlg) {
	var ResData model.ResFlg
	ResData.Status = ReSt
	ResData.Result = ReRe
	ResData.Id = ReId
	return ResData
}