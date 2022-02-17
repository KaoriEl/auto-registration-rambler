package tasks

import (
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"main/internal/structures"
	"time"
)

const (
	CheckCaptcha = `if (document.querySelector("#__next > div > div > div.styles_popup__hP12r > div > div > div > div.styles_leftColumn__GopPD > form > section:nth-child(6) > div > div > div.styles_value__5sF2f > div > div.rui-FieldStatus-message") != null){ document.querySelector("#__next > div > div > div.styles_popup__hP12r > div > div > div > div.styles_leftColumn__GopPD > form > section:nth-child(6) > div > div > div.styles_value__5sF2f > div > div.rui-FieldStatus-message").textContent }else{ "200" }`
)

func RamblerBaseStep(url string, buffer *[]byte, password string, i structures.AccInfo) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Sleep(1 * time.Second),
		chromedp.Navigate(url),
		chromedp.Sleep(5 * time.Second),
		chromedp.WaitVisible(`#login`),
		chromedp.Click(`#login`, chromedp.NodeVisible),
		chromedp.SendKeys(`#login`, i.Email),
		chromedp.WaitVisible(`#newPassword`),
		chromedp.Click(`#newPassword`, chromedp.NodeVisible),
		chromedp.SendKeys(`#newPassword`, password),
		chromedp.WaitVisible(`#confirmPassword`),
		chromedp.Click(`#confirmPassword`, chromedp.NodeVisible),
		chromedp.SendKeys(`#confirmPassword`, password),
		chromedp.Sleep(1 * time.Second),
		chromedp.Click(`//*[@id="__next"]/div/div/div[2]/div/div/div/div[1]/form/section[4]/div/div/div/div/div/div/input`, chromedp.NodeVisible),
		chromedp.SendKeys(`//*[@id="__next"]/div/div/div[2]/div/div/div/div[1]/form/section[4]/div/div/div/div/div/div/input`, kb.ArrowDown+kb.Enter),
		chromedp.Sleep(1 * time.Second),
		chromedp.Click(`#answer`, chromedp.NodeVisible),
		chromedp.SendKeys(`#answer`, "630073"),
		chromedp.Screenshot("//*[@id=\"__next\"]/div/div/div[2]/div/div/div/div[1]/form/section[6]/div/div/div[1]/img", buffer, chromedp.NodeVisible),
	}
}

func RamblerCaptchaInput(captcha string, res *string, quality int, buffer *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Sleep(2 * time.Second),
		chromedp.WaitVisible(`#__next > div > div > div.styles_popup__hP12r > div > div > div > div.styles_leftColumn__GopPD > form > section:nth-child(6) > div > div > div.styles_image__syjy5 > img`),
		chromedp.Click(`//*[@id="__next"]/div/div/div[2]/div/div/div/div[1]/form/section[6]/div/div/div[2]/div/div[1]/input`, chromedp.NodeVisible),
		chromedp.SendKeys(`//*[@id="__next"]/div/div/div[2]/div/div/div/div[1]/form/section[6]/div/div/div[2]/div/div[1]/input`, captcha),
		chromedp.Click(`//*[@id="__next"]/div/div/div[2]/div/div/div/div[1]/form/button`, chromedp.NodeVisible),
		chromedp.Sleep(6 * time.Second),
		chromedp.Evaluate(CheckCaptcha, &res),
	}
}

func AccountInfo(quality int, buffer *[]byte, name string, lastname string, location string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Sleep(1 * time.Second),
		chromedp.Click(`#birthday`, chromedp.NodeVisible),
		chromedp.Sleep(3 * time.Second),
		chromedp.SendKeys(`#birthday`, kb.ArrowDown+kb.ArrowDown+kb.ArrowDown+kb.Enter),
		chromedp.Sleep(2 * time.Second),
		chromedp.Click(`//*[@id="__next"]/div/div/div[2]/div/div/div[2]/div[1]/form/section[4]/div/div/div[1]/div[2]/div/div/div/input`, chromedp.NodeVisible),
		chromedp.Sleep(2 * time.Second),
		chromedp.SendKeys(`//*[@id="__next"]/div/div/div[2]/div/div/div[2]/div[1]/form/section[4]/div/div/div[1]/div[2]/div/div/div/input`, kb.ArrowDown+kb.ArrowDown+kb.ArrowDown+kb.Enter),
		chromedp.Sleep(1 * time.Second),
		chromedp.Click(`//*[@id="__next"]/div/div/div[2]/div/div/div[2]/div[1]/form/section[4]/div/div/div[1]/div[3]/div/div/div/input`, chromedp.NodeVisible),
		chromedp.Sleep(2 * time.Second),
		chromedp.SendKeys(
			`//*[@id="__next"]/div/div/div[2]/div/div/div[2]/div[1]/form/section[4]/div/div/div[1]/div[3]/div/div/div/input`,
			kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.ArrowDown+
				kb.Enter,
		),
		chromedp.Sleep(1 * time.Second),
		chromedp.Click(`//*[@id="geoid"]`, chromedp.NodeVisible),
		chromedp.Sleep(1 * time.Second),
		chromedp.SendKeys(`//*[@id="geoid"]`, location),
		chromedp.Sleep(1 * time.Second),
		chromedp.Click(`//*[@id="geoid"]`, chromedp.NodeVisible),
		chromedp.Sleep(1 * time.Second),
		chromedp.SendKeys(`//*[@id="geoid"]`, kb.ArrowDown+kb.Enter),
		chromedp.Sleep(1 * time.Second),

		chromedp.WaitVisible(`//*[@id="firstname"]`),
		chromedp.SendKeys(`//*[@id="firstname"]`, name),
		chromedp.WaitVisible(`//*[@id="lastname"]`),
		chromedp.SendKeys(`//*[@id="lastname"]`, lastname),
		chromedp.Sleep(1 * time.Second),
		chromedp.Click(`//*[@id="gender"]`, chromedp.NodeVisible),
		chromedp.SendKeys(`//*[@id="gender"]`, kb.ArrowDown+kb.ArrowDown+kb.Enter),
		chromedp.Sleep(1 * time.Second),
		chromedp.FullScreenshot(buffer, quality),
	}
}

func FailedCaptcha(buffer *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Sleep(2 * time.Second),
		chromedp.Screenshot("//*[@id=\"__next\"]/div/div/div[2]/div/div/div/div[1]/form/section[6]/div/div/div[1]/img", buffer, chromedp.NodeVisible),
	}
}
