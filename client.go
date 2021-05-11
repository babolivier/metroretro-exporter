package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

const (
	urlBase   = "https://metroretro.io"
	boardPath = "/api/v1/boards/%s/export?format=json&dl=1"
)

var (
	ErrAuthFailed = errors.New("Failed to authenticate against MetroRetro with the provided credentials")
)

type metroretroClient struct {
	c http.Client
}

func newClient(config *MetroretroConfig) (*metroretroClient, error) {
	mc := new(metroretroClient)

	u, err := url.Parse(urlBase)
	if err != nil {
		return nil, err
	}

	cookies := []*http.Cookie{
		{
			Name:  "metret.sess",
			Value: config.Session,
		},
		{
			Name:  "metret.sess.sig",
			Value: config.SessionSig,
		},
	}

	// cookiejar.New can't return an error, so it's safe to ignore.
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(u, cookies)

	mc.c = http.Client{
		Jar: jar,
	}

	return mc, nil
}

func (mc *metroretroClient) getBoard(id string) (*APIResp, error) {
	u := fmt.Sprintf(urlBase+boardPath, id)

	resp, err := mc.c.Get(u)
	if err != nil {
		return nil, err
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/json") {
		return nil, ErrAuthFailed
	}

	defer resp.Body.Close()

	parsedResp := new(APIResp)
	if err = json.NewDecoder(resp.Body).Decode(&parsedResp); err != nil {
		panic(err)
	}

	return parsedResp, nil
}
