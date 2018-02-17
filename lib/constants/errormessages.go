package constants

import (
	"errors"
)

var (
	ErrLimitHit       = errors.New("429: Breaking a request rate limit")
	ErrIPAddrBanner   = errors.New("418: IP address banned for continuing to send requests after receiveng 429 code")
	ErrOnExchangeSide = errors.New("504: Problem on an exchange side")
)
