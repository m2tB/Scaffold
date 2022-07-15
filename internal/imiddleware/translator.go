package imiddleware

import (
	"GhortLinks/internal/initialize/icommon"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var trans ut.Translator

func Translator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//设置支持语言
		enT := en.New()
		zhT := zh.New()

		//设置国际化翻译器
		uni := ut.New(zhT, zhT, enT)
		val := validator.New()

		//根据参数取翻译器实例
		language := ctx.DefaultQuery("locale", icommon.DEFAULT_LOCAL)
		trans, _ = uni.GetTranslator(language)

		//翻译器注册到validator
		switch language {
		case "en":
			_ = enTranslations.RegisterDefaultTranslations(val, trans)
			// 注册一个函数,将struct tag里添加的"en_label"作为备用名
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("en_label")
			})
			break
		default:
			_ = zhTranslations.RegisterDefaultTranslations(val, trans)
			// 注册一个函数,将struct tag里添加的"zh_label"作为备用名
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("zh_label")
			})

			//自定义验证方法
			customVerification(val)
			//自定义翻译器
			customTranslate(val)
			break
		}
		ctx.Set(icommon.CURRENT_TRANSLATOR, trans)
		ctx.Set(icommon.CURRENT_VALIDATOR, val)
		ctx.Next()
	}
}

// 自定义验证方法
func customVerification(val *validator.Validate) {
	//// 验证登录账号是否符合规范
	//_ = val.RegisterValidation("isValidateAccount", func(fl validator.FieldLevel) bool {
	//	return strings.HasSuffix(fl.Field().String(), common.AccountFormat)
	//})
	//// 验证密码格式是否符合规范
	//_ = val.RegisterValidation("checkConformPwdFormat", func(fl validator.FieldLevel) bool {
	//	// 保证为8位以上的密码以及大写开头
	//	currentPassword := fl.Field().String()
	//	flag := true
	//	if len(currentPassword) < 8 {
	//		flag = false
	//	}
	//	matched, _ := regexp.MatchString(`^(?!\d+$)(?![a-z]+$)(?![A-Z]+$)(?![!#$%^&*]+$)[\da-zA-z!#$%^&*]{8,16}$`, currentPassword)
	//	if !matched {
	//		flag = false
	//	}
	//	return flag
	//})
	////  验证服务数据名称是否满足要求
	//_ = val.RegisterValidation("checkConformServiceNameFormat", func(fl validator.FieldLevel) bool {
	//	matched, _ := regexp.Match(`^\w{6,128}$`, istring.String2Byte(fl.Field().String()))
	//	return matched
	//})
	//// 验证IP格式是否满足要求
	//_ = val.RegisterValidation("checkConformServiceRoundIpsFormat", func(fl validator.FieldLevel) bool {
	//	for _, each := range strings.Split(fl.Field().String(), ",") {
	//		match, _ := regexp.MatchString(`^(((25[0-5]|2[0-4]d|((1\d{2})|([1-9]?\d)))\.){3}(25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))):(\d|[1-9]\d{1,3}|[1-5]\d{4}|6[0-4]\d{4}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])$`, each)
	//		if !match {
	//			return false
	//		}
	//	}
	//	return true
	//})
	//// 验证权重列表
	//_ = val.RegisterValidation("checkConformServiceWeightListFormat", func(fl validator.FieldLevel) bool {
	//	for _, each := range strings.Split(fl.Field().String(), ",") {
	//		match, _ := regexp.Match(`^\d+$`, []byte(each))
	//		if !match {
	//			return false
	//		}
	//	}
	//	return true
	//})
	////  验证URL重写规则
	//_ = val.RegisterValidation("checkConformServiceUrlRewriteFormat", func(fl validator.FieldLevel) bool {
	//	if fl.Field().String() == "" {
	//		return true
	//	}
	//	for _, each := range strings.Split(fl.Field().String(), ",") {
	//		if len(strings.Split(each, " ")) != 2 {
	//			return false
	//		}
	//	}
	//	return true
	//})
	////  验证header转换规则
	//_ = val.RegisterValidation("checkConformServiceHeaderTransferFormat", func(fl validator.FieldLevel) bool {
	//	if fl.Field().String() == "" {
	//		return true
	//	}
	//	for _, each := range strings.Split(fl.Field().String(), ",") {
	//		if len(strings.Split(each, " ")) != 3 {
	//			return false
	//		}
	//	}
	//	return true
	//})
	////  验证port格式
	//_ = val.RegisterValidation("checkConformPortFormat", func(fl validator.FieldLevel) bool {
	//	if fl.Field().String() == "" {
	//		return false
	//	}
	//	for _, each := range strings.Split(fl.Field().String(), ",") {
	//		match, _ := regexp.Match(`^(\d|[1-9]\d{1,3}|[1-5]\d{4}|6[0-4]\d{4}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])$`, istring.String2Byte(each))
	//		if !match {
	//			return false
	//		}
	//		port, err := strconv.Atoi(each)
	//		if err != nil || port < 10310 {
	//			return false
	//		}
	//	}
	//	return true
	//})
}

// 自定义翻译器
func customTranslate(val *validator.Validate) {
	//// 验证账号格式是否符合规范 - 翻译器
	//_ = val.RegisterTranslation("isValidateAccount", trans, func(ut ut.Translator) error {
	//	return ut.Add("isValidateAccount", iresponse.AccountFormat.Msg(), true)
	//}, func(ut ut.Translator, fe validator.FieldError) string {
	//	t, _ := ut.T("isValidateAccount", fe.Field())
	//	return t
	//})
	//// 验证密码格式是否符合规范 - 翻译器
	//_ = val.RegisterTranslation("checkConformPwdFormat", trans, func(ut ut.Translator) error {
	//	return ut.Add("checkConformPwdFormat", iresponse.PasswordFormat.Msg(), true)
	//}, func(ut ut.Translator, fe validator.FieldError) string {
	//	t, _ := ut.T("checkConformPwdFormat", fe.Field())
	//	return t
	//})
	//_ = val.RegisterTranslation("checkConformServiceNameFormat", trans, func(ut ut.Translator) error {
	//	return ut.Add("checkConformServiceNameFormat", iresponse.ServiceNameFormat.Msg(), true)
	//}, func(ut ut.Translator, fe validator.FieldError) string {
	//	t, _ := ut.T("checkConformServiceNameFormat", fe.Field())
	//	return t
	//})
	//_ = val.RegisterTranslation("checkConformServiceRoundIpsFormat", trans, func(ut ut.Translator) error {
	//	return ut.Add("checkConformServiceRoundIpsFormat", iresponse.ServiceNameFormat.Msg(), true)
	//}, func(ut ut.Translator, fe validator.FieldError) string {
	//	t, _ := ut.T("checkConformServiceRoundIpsFormat", fe.Field())
	//	return t
	//})
	//_ = val.RegisterTranslation("checkConformServiceWeightListFormat", trans, func(ut ut.Translator) error {
	//	return ut.Add("checkConformServiceWeightListFormat", iresponse.IpWeightNotEqual.Msg(), true)
	//}, func(ut ut.Translator, fe validator.FieldError) string {
	//	t, _ := ut.T("checkConformServiceWeightListFormat", fe.Field())
	//	return t
	//})
	//_ = val.RegisterTranslation("checkConformServiceUrlRewriteFormat", trans, func(ut ut.Translator) error {
	//	return ut.Add("checkConformServiceUrlRewriteFormat", iresponse.UrlRewriteFormat.Msg(), true)
	//}, func(ut ut.Translator, fe validator.FieldError) string {
	//	t, _ := ut.T("checkConformServiceUrlRewriteFormat", fe.Field())
	//	return t
	//})
	//_ = val.RegisterTranslation("checkConformServiceHeaderTransferFormat", trans, func(ut ut.Translator) error {
	//	return ut.Add("checkConformServiceHeaderTransferFormat", iresponse.HeaderTransferFormat.Msg(), true)
	//}, func(ut ut.Translator, fe validator.FieldError) string {
	//	t, _ := ut.T("checkConformServiceHeaderTransferFormat", fe.Field())
	//	return t
	//})
	//_ = val.RegisterTranslation("checkConformPortFormat", trans, func(ut ut.Translator) error {
	//	return ut.Add("checkConformPortFormat", iresponse.PortFormat.Msg(), true)
	//}, func(ut ut.Translator, fe validator.FieldError) string {
	//	t, _ := ut.T("checkConformPortFormat", fe.Field())
	//	return t
	//})
}
