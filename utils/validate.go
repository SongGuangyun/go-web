package utils

import (
	"fmt"
	"github.com/Songguangyun/go-web/internal/global"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhs "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
)

var (
	chinese = zh.New()                 // 獲取中文翻譯器
	uni     = ut.New(chinese, chinese) // 設置成中文翻譯器
)

// InitValidator 初始化通用验证器
func InitValidator() {
	global.Validator = validator.New()
	// 注册一个函数，获取struct tag里自定义的label作为字段名
	global.Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
	// 注册手机号验证器
	global.Validator.RegisterValidation("mobile", MobileValidation)
	// 初始化中文翻译器
	err := InitValidatorTrans()
	if err != nil {
		return
	}
	// 注册手机号验证翻译器
	_ = global.Validator.RegisterTranslation("mobile", global.ValidatorTrans, func(ut ut.Translator) error {
		return ut.Add("mobile", "{0}格式错误", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("mobile", fe.Field())
		return t
	})
}

// InitValidatorTrans 注册中文翻译器
func InitValidatorTrans() (err error) {
	if trans, ok := uni.GetTranslator("zh"); ok {
		if err := zhs.RegisterDefaultTranslations(global.Validator, trans); err != nil {
			return err
		}
		global.ValidatorTrans = trans
	} else {
		return fmt.Errorf("uni.GetTranslator(%s) failed", "zh")
	}
	return
}

// 手机号码验证
func MobileValidation(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	pattern := `^1[3-9]\d{9}$`
	if ok, _ := regexp.MatchString(pattern, mobile); ok {
		return true
	}
	return false
}
