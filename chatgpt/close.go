package chatgpt

import (
	"context"
	"github.com/chromedp/chromedp"
)

func (client *Client) Close(ctx context.Context) error {
	return chromedp.Cancel(ctx)
}
