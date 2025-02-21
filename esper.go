package esper

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"blekksprut.net/natto/gemini"
	"blekksprut.net/natto/spartan"
	"blekksprut.net/yofukashi/nex"
)

const Version = "0.0.1"

var Cards = []string{
	"🂡", "🂱", "🃁", "🃑",
	"🂢", "🂲", "🃂", "🃒",
	"🂣", "🂳", "🃃", "🃓",
	"🂤", "🂴", "🃄", "🃔",
	"🂥", "🂵", "🃅", "🃕",
	"🂦", "🂶", "🃆", "🃖",
	"🂧", "🂷", "🃇", "🃗",
	"🂨", "🂸", "🃈", "🃘",
	"🂩", "🂹", "🃉", "🃙",
	"🂪", "🂺", "🃊", "🃚",
	"🂫", "🂻", "🃋", "🃛",
	"🂬", "🂼", "🃌", "🃜",
	"🂭", "🂽", "🃍", "🃝",
	"🂮", "🂾", "🃎", "🃞",
	"🃟",
}

var Animals = []string{
	"🐕", "🐺", "🦊", "🐩", "🐈", "🐆", "🐅", "🦍", "🐒", "🐖", "🐎",
	"🐐", "🐏", "🐑", "🦏", "🐘", "🐪", "🐄", "🐂", "🐃", "🐿", "🐇",
	"🐓", "🕊", "🦅", "🦉", "🦆", "🐧", "🐢", "🐙", "🦀", "🦐", "🦈",
	"🐟", "🐠", "🐋", "🐡", "🐍", "🐊", "🦎", "🐛", "🐜", "🐌", "🐝",
	"🐞", "🦋",
}

func Try(ref, name string) error {
	u, err := url.Parse(ref)
	if err != nil {
		return fmt.Errorf("invalid url")
	}

	ctx := context.Background()
	switch u.Scheme {
	case "gemini":
		r, err := gemini.Request(ctx, u.String())
		if err != nil {
			return err
		}
		if r.Status != gemini.Success {
			return fmt.Errorf("%d", r.Status)
		}
		return nil
	case "spartan":
		r, err := spartan.Request(ctx, u.String(), spartan.Data{})
		if err != nil {
			return err
		}
		if r.Status != spartan.Success {
			return fmt.Errorf("%d", r.Status)
		}
		return nil
	case "nex":
		_, err := nex.Request(ctx, u.String())
		return err
	case "http", "https":
		r, err := http.Get(u.String())
		if err != nil {
			return err
		}
		if r.Status[0] != '2' {
			return fmt.Errorf("%d", r.StatusCode)
		}
		return nil
	default:
		return fmt.Errorf("?")
	}
	return nil
}
