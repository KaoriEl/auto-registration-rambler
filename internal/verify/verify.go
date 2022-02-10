package verify

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/fatih/color"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"main/internal/chrome"
	"main/internal/chrome/tasks"
	"main/internal/registration"
	"main/internal/structures"
	"os"
	"path/filepath"
)

func Verify(i structures.AccInfo, rdb *redis.Client) string {
	ctx, cancel := chrome.ChromeConfiguration()
	defer cancel()
	var b []byte
	filePrefix, _ := filepath.Abs("/var/www/investments-auto-registration-rambler/captcha/")

	var args = structures.Args{
		I:          i,
		Prefix:     registration.GenerateString(20),
		FilePrefix: filePrefix,
	}

	FirstStep(ctx, args)
	SecondStep(ctx)
	return ThirdStep(b, ctx, args)
}

func FirstStep(ctx context.Context, args structures.Args) {
	var b []byte
	if err := chromedp.Run(ctx, tasks.RamblerFirstStep(os.Getenv("RamblerLoginUrl"), args.I, &b)); err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Couldn't launch chrome browser"))
	}
}

func SecondStep(ctx context.Context) {
	var res string
	var b []byte
	if err := chromedp.Run(ctx, tasks.RamblerSecondStep(&res, &b)); err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Couldn't launch chrome browser"))
	}

}

func ThirdStep(b []byte, ctx context.Context, args structures.Args) string {
	var res string
	if err := chromedp.Run(ctx, tasks.RamblerThirdStep(100, &b, &res)); err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Couldn't launch chrome browser"))
	}

	return res

}
