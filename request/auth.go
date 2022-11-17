package request

type LoginRequest struct {
	PhoneNumber string `json:"phoneNumber" label:"手机号码" validate:"required,mobile"` // 手机号码
	Captcha     string `json:"captcha" label:"验证码" validate:"required,gt=10"`       // 验证码
}
