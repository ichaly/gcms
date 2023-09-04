package base

import (
	"context"
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

type Validate struct {
	validate  *validator.Validate
	translate ut.Translator
}

func NewValidate() *Validate {
	z := zh.New()
	uni := ut.New(z, z)
	t, _ := uni.GetTranslator("zh")
	v := validator.New()
	_ = zhTranslations.RegisterDefaultTranslations(v, t)
	return &Validate{validate: v, translate: t}
}

func (my *Validate) Struct(s interface{}) error {
	err := my.validate.StructCtx(context.Background(), s)
	if err != nil {
		var errs validator.ValidationErrors
		errors.As(err, &errs)
		return errors.New(removeStructName(errs.Translate(my.translate)))
	}
	return err
}

func removeStructName(fields map[string]string) string {
	var errMsg string
	result := map[string]string{}
	for field, err := range fields {
		result[field[strings.Index(field, ".")+1:]] = err
	}
	for _, set := range result {
		errMsg = errMsg + set + " "
	}
	return errMsg
}
