package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	translator := zh.New()
	uni = ut.New(translator, translator)
	trans, _ = uni.GetTranslator("zh")
	validate := binding.Validator.Engine().(*validator.Validate)
	_ = zh_translations.RegisterDefaultTranslations(validate, trans)
}

func Translate(err error) string {
	var result string

	errors := err.(validator.ValidationErrors)

	for _, err := range errors {
		//typ := err.Type()
		//fieldName := err.StructField()
		//field, _ := typ.FieldByName(fieldName)
		//display := field.Tag.Get("display")
		//errMessage := strings.ReplaceAll(err.Translate(trans), err.Field(), display)
		errMessage := err.Translate(trans)
		result += errMessage + ";"
	}
	return result
}
