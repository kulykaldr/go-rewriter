package chatgpt

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/nochso/gomd/eol"
	"regexp"
	"strings"
	"time"
)

func (client *Client) Rewrite(ctx context.Context, text string, command string) (string, error) {
	if command == "" {
		command = "REWRITE"
	}

	if !client.config.IsLogin {
		if err := chromedp.Run(ctx, chromedp.Tasks{
			client.uaEmulation,
			chromedp.Navigate(`https://chat.openai.com/chat`),
		}); err != nil {
			return "", err
		}
	}

	re, _ := regexp.Compile(`[\w\W\s]{0,1000}[^.!? ]+[.!?]+`)
	el, _ := eol.Detect(text)
	tempStr := strings.Replace(text, el.String(), " ", -1)

	var results []string
	for {
		textPart := re.FindString(tempStr)
		tempStr = strings.Replace(tempStr, textPart, "", 1)

		if len(strings.TrimSpace(textPart)) == 0 {
			break
		}

		var answer string
		sendButtonSel := `textarea+button:not([disabled])`
		var isContentPolicy bool
		if err := chromedp.Run(ctx, chromedp.Tasks{
			client.uaEmulation,
			chromedp.SetValue(`textarea`, ""),
			chromedp.SendKeys(`textarea`, fmt.Sprintf("%s. %s", command, textPart)),
			chromedp.Sleep(2 * time.Second),
			chromedp.Click(sendButtonSel, chromedp.NodeVisible),
			chromedp.WaitReady(sendButtonSel),
			chromedp.TextContent(`//div[contains(@class, 'dark:bg-[#444654]')][last()]//div[contains(@class, 'markdown')]`, &answer),
			chromedp.ActionFunc(func(ctx context.Context) error {
				nextBtnSel := `//a[contains(text(), 'content policy')]`
				var nodes []*cdp.Node
				if err := chromedp.Nodes(nextBtnSel, &nodes, chromedp.AtLeast(0)).Do(ctx); err != nil {
					return err
				}

				if len(nodes) > 0 {
					isContentPolicy = true
				}
				return nil
			}),
		}); err != nil {
			return "", err
		}

		if isContentPolicy {
			return "", fmt.Errorf("article violate content policy")
		}

		results = append(results, answer)

		if len(strings.TrimSpace(tempStr)) == 0 {
			break
		}

		time.Sleep(2 * time.Second)
	}

	textResult := strings.Join(results, el.String()+el.String())

	return textResult, nil
}
