package chatgpt

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func (client *Client) SignIn(ctx context.Context) error {
	if client.config.Login == "" || client.config.Password == "" {
		return fmt.Errorf("login or Password not provided")
	}
	log.Println("signin...")

	err := chromedp.Run(ctx, chromedp.Tasks{
		client.uaEmulation,
		chromedp.Navigate(`https://chat.openai.com/auth/login?next=/chat`),
		chromedp.Sleep(5 * time.Second),

		chromedp.Click(`//div[text()='Log in']/ancestor::button`, chromedp.NodeVisible),

		chromedp.SendKeys(`input[name="username"]`, client.config.Login, chromedp.NodeVisible),
		chromedp.Click(`button[name="action"]`, chromedp.NodeVisible),
		chromedp.SendKeys(`input[name="password"]`, client.config.Password, chromedp.NodeVisible),
		chromedp.Click(`button[name="action"]`, chromedp.NodeVisible),

		//chromedp.Sleep(5 * time.Second),
		//chromedp.ActionFunc(clickNextBtn),
		//chromedp.ActionFunc(clickNextBtn),
		//chromedp.ActionFunc(clickNextBtn),

		chromedp.WaitReady(`textarea+button:not([disabled])`),
	})
	if err != nil {
		return err
	}

	return nil
}

func clickNextBtn(ctx context.Context) error {
	nextBtnSel := `//div[text()="Next" or text()="Done"]/ancestor::button`
	var nodes []*cdp.Node
	if err := chromedp.Nodes(nextBtnSel, &nodes, chromedp.AtLeast(0)).Do(ctx); err != nil {
		return err
	}

	if len(nodes) > 0 {
		chromedp.Click(nextBtnSel, chromedp.NodeVisible)
	}
	return nil
}
