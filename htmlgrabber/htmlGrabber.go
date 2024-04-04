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

func GrabGroupIronGraphPage() (*[]byte, error) {
	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// Create a context with timeout for chromedp.Run
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second) // 5-second timeout
	defer cancel()

	path := "http://localhost:7749/graph"
	fmt.Println(path)

	// capture screenshot of an element
	var buf []byte
	err := chromedp.Run(ctx, elementScreenshot(path, &buf, 1, "#chart-container > div > canvas", "#chart-container"))
	if err != nil {
		fmt.Println("ERROR", err)
		return nil, err
	}

	// log the size of the img
	fmt.Printf("screenshot size: %d\n", len(buf))

	return &buf, nil
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
	err := chromedp.Run(ctx, elementScreenshot(path, &buf, 4, "#parent-box", "#parent-box"))
	if err != nil {
		fmt.Println("ERROR", err)
		return nil, err
	}

	// log the size of the img
	fmt.Printf("screenshot size: %d\n", len(buf))

	return &buf, nil
}

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr string, res *[]byte, scale float64, waitVisibleElement string, screenShotElement string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(waitVisibleElement),
		chromedp.ScreenshotScale(screenShotElement, scale, res, chromedp.NodeVisible),
	}
}
