package dingtalk

import (
	"io"
	"os"
	"testing"
)

func TestSendText(t *testing.T) {
	res, err := SendText("https://oapi.dingtalk.com/robot/send?access_token=", "this is test")
	if err != nil {
		t.Error(err)
	}
	io.Copy(os.Stdout, res)
}
