package base

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	//	enTranslations "github.com/go-playground/validator/v10/translations/en"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func NewServer(c *Config) *gin.Engine {
	if !c.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //chinese
		enT := en.New() //english
		uni := ut.New(enT, zhT, enT)

		trans, _ = uni.GetTranslator("zh")
		_ = chTranslations.RegisterDefaultTranslations(v, trans)
	}
	e := gin.Default()
	e.Use(cors.Default())
	e.Use(gzip.Gzip(gzip.DefaultCompression))
	return e
}
