package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	urlBase   = "https://metroretro.io"
	boardPath = "/api/v1/boards/%s/export?format=json&dl=1"
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

	defer resp.Body.Close()

	parsedResp := new(APIResp)
	if err = json.NewDecoder(resp.Body).Decode(&parsedResp); err != nil {
		panic(err)
	}

	return parsedResp, nil
}
