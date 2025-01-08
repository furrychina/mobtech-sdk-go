package mobpush

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	apiCreatePush = "http://api.push.mob.com/v3/push/createPush"
)

func NewMessage(appKey string, target *PushTarget, notify *PushNotify) *PushObject {
	return &PushObject{
		Source:     "webapi",
		AppKey:     appKey,
		PushTarget: target,
		PushNotify: notify,
	}
}

func NewNotify(title, content string, extrasMapList []ExtrasMap) *PushNotify {
	return &PushNotify{
		Title:         title,
		Content:       content,
		Type:          TypeNotify,
		Policy:        PolicyTCPFirst,
		Plats:         []int{IOS, Android},
		ExtrasMapList: extrasMapList,
		AndroidNotify: &AndroidNotify{
			NativeCategory: "msg",
		},
	}
}

func SendPush(appSecret string, pushObject *PushObject) (*Response, error) {
	// 检查推送设备列表是否为空，空则跳过推送
	if len(pushObject.PushTarget.Rids) == 0 {
		return nil, nil
	}
	// 构造推送消息
	requestBody, _ := json.Marshal(pushObject)
	// 将请求体和密钥拼接，生成签名
	sign := md5.Sum(append(requestBody, []byte(appSecret)...))
	// 发送请求
	req, err := http.NewRequest("POST", apiCreatePush, bytes.NewBuffer(requestBody))
	if err != nil {
		return &Response{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("key", pushObject.AppKey)
	req.Header.Set("sign", fmt.Sprintf("%x", sign))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &Response{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &Response{}, err
	}

	// 打印响应体内容，便于调试
	//fmt.Print("Response Body:", string(body))

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	// 将响应体转换为结构体
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return &Response{}, err
	}
	return &result, nil
}
