package api

import (
	"focus/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	captcha "github.com/mojocn/base64Captcha"
	"time"
)

var Captcha = new(CaptchaApi)

type CaptchaApi struct{}

var store = captcha.DefaultMemStore

func NewDriver() *captcha.DriverString {
	driver := new(captcha.DriverString)
	driver.Height = 44
	driver.Width = 120
	driver.NoiseCount = 5
	driver.ShowLineOptions = captcha.OptionShowSineLine | captcha.OptionShowSlimeLine | captcha.OptionShowHollowLine
	driver.Length = 4
	driver.Source = "1234567890"
	driver.Fonts = []string{"wqy-microhei.ttc"}
	return driver
}

func (a *CaptchaApi) Get(r *ghttp.Request) {
	var driver = NewDriver().ConvertFonts()

	c := captcha.NewCaptcha(driver, store)

	_, content, answer := c.Driver.GenerateIdQuestionAnswer()

	item, _ := c.Driver.DrawCaptcha(content)

	microsecond := gconv.String(time.Now().UnixNano())
	g.Dump(microsecond)
	_ = r.Session.Set("captcha", microsecond)

	c.Store.Set(microsecond, answer)

	_, _ = item.WriteTo(r.Response.Writer)
}

func (a *CaptchaApi) Verify(r *ghttp.Request) {

	code := r.GetQueryString("code")
	if store.Verify(r.Session.GetString("captcha"), code, true) {
		r.Session.Remove("captcha")
		response.JsonExit(r, 1, "ok", g.Map{"status": true})
	} else {
		response.JsonExit(r, 1, "ok", g.Map{"status": false})
	}
}
