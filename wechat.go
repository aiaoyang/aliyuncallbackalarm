package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

// token cache
type tokenCache struct {
	token     string
	timeout   time.Duration
	timestamp time.Time
}

//Content 文本消息内容
type Content struct {
	Content string `json:"content"`
}

//Message 消息主体参数
type Message struct {
	ToUser  string  `json:"touser"`
	ToParty string  `json:"toparty"`
	ToTag   string  `json:"totag"`
	MsgType string  `json:"msgtype"`
	AgentID int     `json:"agentid"`
	Text    Content `json:"text"`
}

// wechat response
type wechatResp struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIN   int64  `json:"expires_in"`
}

// wechat send msg response
type wechatMsgResp struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Invaliduser  string `json:"invaliduser"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
}

func (t *tokenCache) Get() string {
	if t.token == "" || time.Now().Add(-t.timeout).After(t.timestamp) {
		token, err := genAccessToken()
		if err != nil {
			log.Println(err)
		}
		t.token = token
		t.timestamp = time.Now()
		log.Printf("gen token from tencent\n")
		return token
	}
	// log.Printf("gen token from cache %s\n", t.token)
	return t.token
}
func genAccessToken() (string, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	f, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Println(err)
	}
	v.ReadConfig(bytes.NewBuffer(f))
	// 判断是否正确取得配置
	corpID, ok := v.Get("wechat.corpID").(string)
	if !ok {
		return "", fmt.Errorf("corpID")
	}
	secret, ok := v.Get("wechat.secret").(string)
	if !ok {
		return "", fmt.Errorf("secret")
	}
	resp, err := http.Get(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s",
		corpID,
		secret,
	))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	wxres := &wechatResp{}
	err = json.Unmarshal(b, wxres)
	if err != nil {
		return "", err
	}
	if wxres.ErrCode != 0 {
		return "", fmt.Errorf("get wx token : %s", wxres.ErrMsg)
	}
	return wxres.AccessToken, nil
}

var cacheToken = &tokenCache{
	token:     "",
	timeout:   time.Second * 7200,
	timestamp: time.Now(),
}

func sendWechatMSG(user string, msgInfo string) {
	token := cacheToken.Get()
	u := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token)
	wxmsg := &Message{
		ToUser:  user,
		MsgType: "text",
		AgentID: 1000002,
		Text: Content{
			Content: msgInfo,
		},
	}
	b, err := json.Marshal(wxmsg)
	if err != nil {
		log.Println(err)
	}
	resp, err := http.Post(u, "application/json;charset=utf-8", bytes.NewReader(b))
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	resStruc := &wechatMsgResp{}
	err = json.Unmarshal(res, resStruc)
	if err != nil {
		log.Println(err)
	}
	if resStruc.Errcode != 0 {
		log.Println(resStruc.Errmsg)
	}
}
