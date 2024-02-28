package htmlgrabber

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/chromedp/chromedp"
)

func NewChromeDpCtx() (context.Context, context.CancelFunc) {
	return chromedp.NewContext(
		context.Background(),
		chromedp.WithDebugf(log.Printf),
	)
}

func GrabGroupIronSkillsPage(name string) (*[]byte, error) {
	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// Create a context with timeout for chromedp.Run
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second) // 5-second timeout
	defer cancel()

	path := url.PathEscape(name)
	path = "http://localhost:7749/skills/" + path
	fmt.Println(path)
	
	// capture screenshot of an element
	var buf []byte
	err := chromedp.Run(ctx, elementScreenshot(path, &buf))
	if err != nil {
		fmt.Println("ERROR", err)
		return nil, err
	}

	return &buf, nil
}

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible("#parent-box"),
		chromedp.ScreenshotScale("#parent-box", 4, res, chromedp.NodeVisible),
	}
}
