package request

import (
	"github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en2 "gopkg.in/go-playground/validator.v9/translations/en"
)

type BaseRequest struct {
	Validate *validator.Validate
	Uni      *ut.UniversalTranslator
	Trans    ut.Translator
}

func New() *BaseRequest {
	en := en_US.New()
	validate := validator.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	en2.RegisterDefaultTranslations(validate, trans)

	baseRequest := &BaseRequest{
		Validate: validate,
		Uni:      uni,
		Trans:    trans,
	}

	return baseRequest
}
