package bapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-programming-tour-book/tag-service/pkg/errcode"
	"io"
	"mime/multipart"
	"net/http"
)

const (
	APP_KEY    = "eddycjy"
	APP_SECRET = "go-programming-tour-book"
)

type AccessToken struct {
	Token string `json:"token"`
}

type API struct {
	URL string
}

func NewAPI(url string) *API {
	return &API{URL: url}
}

func (a *API) getAccessToken(ctx context.Context) (string, error) {
	token, err := a.httpPost(ctx, fmt.Sprintf("%s", "http://localhost:8000/auth"))
	if err != nil {
		return "", err
	}
	var accessToken AccessToken
	accessToken.Token = token
	return accessToken.Token, nil
}

func (a *API) GetTagList(ctx context.Context, name string) ([]byte, error) {
	token, err := a.getAccessToken(ctx)
	if err != nil {
		fmt.Println("err happen")
		return nil, err
	}
	body, err := a.httpGet(ctx, fmt.Sprintf("%s?name=%s&state=1", "http://localhost:8000/api/v1/tags", name), token)
	if err != nil {
		return nil, errcode.ToRPCError(errcode.ErrorGetTagListFail)
	}
	return body, nil
}

func (a *API) httpGet(ctx context.Context, path, token string) ([]byte, error) {
	fmt.Println(path)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败：", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error occur when close the stream: %v", err)
		}
	}(resp.Body)
	bodyByte, _ := io.ReadAll(resp.Body)
	return bodyByte, nil
}

func (a *API) httpPost(ctx context.Context, path string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("app_key", APP_KEY)
	_ = writer.WriteField("app_secret", APP_SECRET)

	err := writer.Close()
	if err != nil {
		return "", err
	}
	resp, err := http.Post(path, writer.FormDataContentType(), body)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error occur when close the stream: %v", err)
		}
	}(resp.Body)
	bodyByte, _ := io.ReadAll(resp.Body)
	var accessToken AccessToken
	err = json.Unmarshal(bodyByte, &accessToken)
	if err != nil {
		return "", err
	}
	return accessToken.Token, nil
}
