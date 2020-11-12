package service

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/guid"
	captcha "github.com/mojocn/base64Captcha"
)

var Captcha = new(captchaService)

type captchaService struct{}

var (
	captchaStore  = captcha.DefaultMemStore
	captchaDriver = newDriver()
)

func newDriver() *captcha.DriverString {
	driver := &captcha.DriverString{
		Height:          44,
		Width:           120,
		NoiseCount:      5,
		ShowLineOptions: captcha.OptionShowSineLine | captcha.OptionShowSlimeLine | captcha.OptionShowHollowLine,
		Length:          4,
		Source:          "1234567890",
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	return driver.ConvertFonts()
}

// 创建验证码，直接输出验证码图片内容到HTTP Response.
func (s *captchaService) NewAndStore(r *ghttp.Request, name string) error {
	c := captcha.NewCaptcha(captchaDriver, captchaStore)
	_, content, answer := c.Driver.GenerateIdQuestionAnswer()
	item, _ := c.Driver.DrawCaptcha(content)
	captchaStoreKey := guid.S()
	r.Session.Set(name, captchaStoreKey)
	c.Store.Set(captchaStoreKey, answer)
	_, err := item.WriteTo(r.Response.Writer)
	return err
}

// 校验验证码，并清空缓存的验证码信息
func (s *captchaService) VerifyAndClear(r *ghttp.Request, name string, value string) bool {
	defer r.Session.Remove(name)
	captchaStoreKey := r.Session.GetString(name)
	return captchaStore.Verify(captchaStoreKey, value, true)
}
