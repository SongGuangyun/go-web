package global

import (
	"github.com/Songguangyun/go-web/config"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	ConfigServer   config.Server
	Logger         *logrus.Logger
	DB             *gorm.DB
	Validator      *validator.Validate
	ValidatorTrans ut.Translator
)
