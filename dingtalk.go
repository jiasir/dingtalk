package dingtalk

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Text struct {
	Content string `json:"content"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type MsgText struct {
	MsgType string `json:"msgtype"`
	Text    Text   `json:"text"`
	At      At     `json:"at"`
}

func SendText(accessToken string, content string) (io.ReadCloser, error) {

	var msgText MsgText
	msgText.MsgType = "text"
	msgText.Text.Content = content

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(msgText)

	res, err := http.Post(accessToken, "application/json; charset=utf-8", b)
	return res.Body, err

}
