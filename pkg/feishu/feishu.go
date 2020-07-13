package feishu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Bowser1704/Alertmanager-feishu-webhook/log"
	"github.com/Bowser1704/Alertmanager-feishu-webhook/pkg/model"
	"github.com/spf13/viper"
)

// FS is feishu message struct
type FS struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Token string `json:"token"`
}

type fsmsg struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func (f *FS) renderTitle(data interface{}) (string, error) {
	return f.Title, nil
}

func (f *FS) renderText(data interface{}) (string, error) {
	return f.Text, nil
}

// Firing fire the message to the target
func (f *FS) Firing() {
	fsurl := viper.GetString("fs")
	if fsurl == "" {
		fsurl = "https://open.feishu.cn/open-apis/bot/hook/"
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(fsmsg{
		Title: f.Title,
		Text:  f.Text,
	})
	client := &http.Client{}
	res, err := client.Post(fsurl+f.Token, "application/json", b)
	if err != nil {
		log.Error("feishu post failed")
	}
	defer res.Body.Close()

	// 需不需要输出 response body?
	// result, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Error("feishu response error")
	// }
	// _, err = base64.StdEncoding.Decode(result, result)
	// if err != nil {
	// 	log.Error("base64 decode feishu response failed")
	// }

	// log.Info("Feishu response: " +
}

// Init transfer alertmanager message to feishu message
func (f *FS) Init(webMsg *model.WebhookMessage) {
	f.Title = webMsg.Status + " : " + webMsg.CommonAnnotations["message"]
	f.Text = GenerateText(webMsg)
}

// GenerateText generate feishu text from alertmanager message
func GenerateText(webMsg *model.WebhookMessage) string {
	// transfer time zone from UTC to local time zone.
	startsAt := webMsg.Alerts[0].StartsAt.Local().String()
	endsAt := webMsg.Alerts[0].EndsAt.Local().String()
	text := fmt.Sprintf("StartAt: %s \nEndsAt: %s \n", strings.Split(startsAt, ".")[0], strings.Split(endsAt, ".")[0])

	return text
}
