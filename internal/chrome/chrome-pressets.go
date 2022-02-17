package chrome

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/fatih/color"
)

func ChromeConfiguration() (context.Context, context.CancelFunc) {
	color.New(color.FgHiWhite).Add(color.Bold).Println("Сhrome configuration options...")
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
		chromedp.WindowSize(1920, 1080),
		chromedp.Flag("blink-settings", "imagesEnabled=true"),
	)
	color.New(color.FgHiWhite).Add(color.Bold).Println("Сhrome NewExecAllocator...")

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	color.New(color.FgHiWhite).Add(color.Bold).Println("Сhrome context generate...")
	ctx, cancel = chromedp.NewContext(
		ctx,
	)
	return ctx, cancel
}
