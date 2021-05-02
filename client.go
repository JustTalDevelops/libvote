package libvote

import (
	"fmt"
	"github.com/corpix/uarand"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"errors"
	"strings"
	"time"
)

// Client is a vote bot client. It can be used to vote bot on servers.
type Client struct {
	http.Client
}

// NewJar generates a new cookie jar for Minecraft Pocket Servers.
func NewJar() (*cookiejar.Jar, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	urlStruct, err := url.Parse(BaseUrl)
	if err != nil {
		return nil, err
	}
	jar.SetCookies(urlStruct, []*http.Cookie{{Name: "vote_privacy_accept", Value: "1"}})
	return jar, nil
}

// NewClient creates a new client for Minecraft Pocket Servers.
func NewClient() *Client {
	return &Client{Client: http.Client{Timeout: 30 * time.Second}}
}

// Vote votes for a server key using a username and captcha code.
func (c *Client) Vote(serverKey int, username string, code string) (bool, error) {
	body := url.Values{}
	body.Set("steam_login", "0")
	body.Set("token", "")
	body.Set("nickname", username)
	body.Set("accept", "1")
	body.Set("g-recaptcha-response", code)
	body.Set("h-captcha-response", code)
	req, err := http.NewRequest("POST", fmt.Sprintf(VoteEndpoint, serverKey), strings.NewReader(body.Encode()))
	if err != nil {
		return false, err
	}
	req.Header.Set("Origin", BaseUrl)
	req.Header.Set("Referer", "https://minecraftpocket-servers.com/server/" + strconv.Itoa(serverKey) + "/vote/")
	req.Header.Set("User-Agent", uarand.GetRandom())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	c.Jar, err = NewJar()
	if err != nil {
		return false, err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	start := time.Now()
	var data []byte
	for {
		// Wait up to 5 seconds for a response. If we still don't get a response, it probably errored.
		if time.Now().Sub(start).Seconds() >= 5 {
			break
		}
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}
		if strings.Contains(string(data), "Thank you") {
			return true, nil
		}
	}

	return false, errors.New(string(data))
}
