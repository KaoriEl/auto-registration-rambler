package registration

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/fatih/color"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"io/ioutil"
	"main/internal/chrome"
	"main/internal/chrome/tasks"
	"main/internal/server/services/api"
	"main/internal/structures"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func Registration(i structures.AccInfo, rdb *redis.Client) {
	ctx, cancel := chrome.ChromeConfiguration()
	defer cancel()
	filePrefix, _ := filepath.Abs("/var/www/investments-auto-registration-rambler/captcha/")
	i.Password = GenerateString(20)
	var count int

	var args = structures.Args{
		I:          i,
		Prefix:     GenerateString(20),
		FilePrefix: filePrefix,
	}

	//Подготавливаю инфу для реги.
	i = Plan(args.I)

	var b []byte

	BaseStep(b, ctx, args)

	fmt.Println("REDIS_LIST: " + "investments-goroutine" + strconv.Itoa(int(i.ClientId)))

	//Чистка редиса, онлайн и без регистрации
	rdb.Del("investments-goroutine" + strconv.Itoa(int(i.ClientId)))

	api.SendCaptcha("Пожалуйста, введите символы с картинки. Через 1 минуту капча станет недействительна.", i, args.Prefix)

	color.New(color.FgHiYellow).Add(color.BgBlack).Println("Gorutina plunged into the eternal tsukuyomi...")

	//Демон который редис контролит на момент капчи
out:
	for j := 0; j < 61; j++ {
		time.Sleep(time.Second)
		captcha, _ := rdb.LRange("investments-goroutine"+strconv.Itoa(int(i.ClientId)), 0, -1).Result()
		if len(captcha) > 0 {
			break out
		}
		if j == 60 {
			cancel()
			api.CaptchaStatus("Вы не решили капчу за отведенное время, для повторной попытки, нажмите на 'Продать аккаунт Coinlist' в главном меню", i)
			return
		}
	}

	captcha, _ := rdb.LPop("investments-goroutine" + strconv.Itoa(int(i.ClientId))).Result()
	InputCaptcha(b, captcha, ctx, args, rdb, cancel, count)
	AccountInfo(b, ctx, args)

}

func BaseStep(b []byte, ctx context.Context, args structures.Args) {
	if err := chromedp.Run(ctx, tasks.RamblerBaseStep(os.Getenv("RamblerUrl"), &b, args.I.Password, args.I)); err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Couldn't launch chrome browser"))
	}

	if err := ioutil.WriteFile(args.FilePrefix+"/"+args.Prefix+"-captcha.jpg", b, 0755); err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Couldn't save screenshot"))
	} else {
		color.New(color.FgHiWhite).Add(color.Bold).Println("Complete, save screenshot. Exit func...")
	}
}

func InputCaptcha(b []byte, captcha string, ctx context.Context, args structures.Args, rdb *redis.Client, cancel context.CancelFunc, count int) {
	var res string
	if err := chromedp.Run(ctx, tasks.RamblerCaptchaInput(captcha, &res)); err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Couldn't launch chrome browser"))
	}

	if res != "" {
		switch res {
		case "Неверные символы":
			color.New(color.FgRed).Add(color.Underline).Println("failed captcha")
			FailedCaptcha(b, ctx, args.Prefix, args.FilePrefix)
			count++
			if count <= 5 {
				api.SendCaptcha("Вы неправильно ввели капчу, пожалуйста, введите её еще раз.", args.I, args.Prefix)
			} else {
				cancel()
				api.CaptchaStatus("Закончились попытки решить капчу, пожалуйста, попробуйте еще раз через 5-10 минут", args.I)
				return
			}

			rdb.Del("investments-goroutine" + strconv.Itoa(int(args.I.ClientId)))
		out:
			for j := 0; j < 61; j++ {
				time.Sleep(time.Second)
				captcha, _ := rdb.LRange("investments-goroutine"+strconv.Itoa(int(args.I.ClientId)), 0, -1).Result()
				if len(captcha) > 0 {
					break out
				}
				if j == 60 {
					cancel()
					api.CaptchaStatus("Вы не решили капчу за отведенное время, для повторной попытки, нажмите на 'Продать аккаунт Coinlist' в главном меню", args.I)
					return
				}
			}
			captcha, _ := rdb.LPop("investments-goroutine" + strconv.Itoa(int(args.I.ClientId))).Result()
			InputCaptcha(b, captcha, ctx, args, rdb, cancel, count)
		}
	} else {
		api.CaptchaStatus("Капча введена верно", args.I)
	}
}

func AccountInfo(b []byte, ctx context.Context, args structures.Args) {
	err := os.Remove(args.FilePrefix + "/" + args.Prefix + "-captcha.jpg")
	if err != nil {
		fmt.Println(err)
	}
	words := strings.Fields(args.I.Name)

	var name string
	var lastname string

	if len(words) > 1 {
		name = words[0]
		lastname = words[1]
	} else {
		name = words[0]
		lastname = words[0]
	}

	location := args.I.City

	if err := chromedp.Run(ctx, tasks.AccountInfo(100, &b, name, lastname, location)); err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Couldn't launch chrome browser"))
		return
	}

	api.ChangeStepUser(args.I)
	api.RegisterPassword(args.I)
	api.CaptchaStatus("Почта зарегистрирована", args.I)

}

func FailedCaptcha(b []byte, ctx context.Context, prefix string, filePrefix string) {
	if err := chromedp.Run(ctx, tasks.FailedCaptcha(&b)); err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Couldn't launch chrome browser"))
	}

	if err := ioutil.WriteFile(filePrefix+"/"+prefix+"-captcha.jpg", b, 0755); err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Couldn't save screenshot"))
	} else {
		color.New(color.FgHiWhite).Add(color.Bold).Println("Complete, save screenshot. Exit func...")
	}
}
