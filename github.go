package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	githubUrlBase  = "https://api.github.com"
	gistCreatePath = "/gists"
)

type githubClient struct {
	cfg *GithubConfig
	c   http.Client
}

type gistCreateReq struct {
	Public bool                         `json:"public"`
	Files  map[string]gistCreateReqFile `json:"files"`
}

type gistCreateReqFile struct {
	Content string `json:"content"`
}

type gistCreateResp struct {
	HtmlUrl string `json:"html_url"`
}

func newGithubClient(cfg *GithubConfig) *githubClient {
	return &githubClient{
		cfg: cfg,
		c:   http.Client{},
	}
}

func (c *githubClient) uploadGist(content string) (string, error) {
	fileName := fmt.Sprintf("%s_retro.md", time.Now().Format("2006_01_02"))

	reqBody := gistCreateReq{
		Public: false,
		Files: map[string]gistCreateReqFile{
			fileName: {
				Content: content,
			},
		},
	}

	rawBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		githubUrlBase+gistCreatePath,
		bytes.NewBuffer(rawBody),
	)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(c.cfg.Username, c.cfg.AccessToken)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := c.c.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf(
			"Github API responded with %s",
			resp.Status,
		)
	}

	defer resp.Body.Close()

	respBody := new(gistCreateResp)
	if err = json.NewDecoder(resp.Body).Decode(respBody); err != nil {
		return "", err
	}

	return respBody.HtmlUrl, nil
}
