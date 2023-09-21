package checker

import (
	"checker/internal/config"
	"encoding/json"
	"os"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)


var (
	c 		= config.Config
	logger 	= config.Logger

	file, _ = os.OpenFile("valid.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
)

func Check(code, proxy string) {
	client := fasthttp.Client{
		Dial: fasthttp.DialFunc(fasthttpproxy.FasthttpHTTPDialer(proxy)),
		MaxConnsPerHost:     10,
		MaxIdleConnDuration: 20 * time.Second,
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI("https://discord.com/api/v9/invites/" + code + "?with_counts=true")
	req.Header.SetMethod("GET")

	req.Header.Set("User-Agent", "Mozilla/5.0")
	

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := client.Do(req, resp)
	if err != nil {
		logger.Error().Str("Code", code).Err(err).Msg("Error");return
	}

	if resp.StatusCode() != 200 {
		logger.Error().Str("Code", code).Int("Status", resp.StatusCode()).Msg("Invalid Invite");return
	}

	var res InviteRes

	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		logger.Error().Str("Code", code).Err(err).Msg("Error");return
	}

	ID 		:= res.GuildID
	boosts 	:= res.Guild.PremiumSubscriptionCount
	Members := res.ApproximateMemberCount
	Online 	:= res.ApproximatePresenceCount


	if boosts >= c.MinBoosts && Members >= c.MinMembers && Online >= c.MinOnline {
		logger.Info().Str("ID", ID).Int("Members", Members).Int("Online", Online).Msg("W Invite")
		file.WriteString(code + "\n")
	} else {
		logger.Error().Str("ID", ID).Int("Members", Members).Int("Online", Online).Msg("L Invite")
	}
}