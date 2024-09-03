package mobpush

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func NewMessage(target *PushTarget, notify *PushNotify) *PushObject {
	return &PushObject{
		Source:     "webapi",
		AppKey:     os.Getenv("MOB_PUSH_APP_KEY"),
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

func (o *PushObject) PushRid(pushObject *PushObject) (*Response, error) {
	// 检查KEY和SECRET是否为空，为空则提示错误
	if os.Getenv("MOB_PUSH_APP_KEY") == "" || os.Getenv("MOB_PUSH_APP_SECRET") == "" {
		return nil, errors.New("MOB_PUSH_APP_KEY 或 MOB_PUSH_APP_SECRET 没有配置，跳过推送，请检查配置")
	}
	// 检查推送设备列表是否为空，空则跳过推送
	if len(pushObject.PushTarget.Rids) == 0 {
		return nil, nil
	}
	resp, err := sendPush(pushObject)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func sendPush(pushObject *PushObject) (Response, error) {
	// 构造推送消息
	requestBody, _ := json.Marshal(pushObject)
	// 将请求体和密钥拼接，生成签名
	sign := md5.Sum(append(requestBody, []byte(os.Getenv("MOB_PUSH_APP_SECRET"))...))
	// 发送请求
	req, err := http.NewRequest("POST", apiCreatePush, bytes.NewBuffer(requestBody))
	if err != nil {
		return Response{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("key", os.Getenv("MOB_PUSH_APP_KEY"))
	req.Header.Set("sign", fmt.Sprintf("%x", sign))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	// 打印响应体内容，便于调试
	//fmt.Print("Response Body:", string(body))

	defer resp.Body.Close()

	// 将响应体转换为结构体
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return Response{}, err
	}
	return result, nil
}
