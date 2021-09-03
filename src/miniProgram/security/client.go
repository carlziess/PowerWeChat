package security

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/security/response"
)

type Client struct {
	*kernel.BaseClient
}

// 向插件开发者发起使用插件的申请
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.imgSecCheck.html
func (comp *Client) imgSecCheck(media *power.HashMap) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"media": media,
	}

	_, err := comp.HttpPostJson("wxa/img_sec_check", data, nil, nil, result)

	return result, err
}

// 异步校验图片/音频是否含有违法违规内容
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.mediaCheckAsync.html
func (comp *Client) MediaCheckAsync(mediaURL string, mediaType int, version int, openID string, scene int) (*response.ResponseSecurityMediaCheckASync, error) {

	result := &response.ResponseSecurityMediaCheckASync{}

	data := &object.HashMap{
		"media_url":  mediaURL,
		"media_type": mediaType,
		"version":    version,
		"openid":     openID,
		"scene":      scene,
	}

	_, err := comp.HttpPostJson("wxa/media_check_async", data, nil, nil, result)

	return result, err
}

// 检查一段文本是否含有违法违规内容
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.msgSecCheck.html
func (comp *Client) MsgSecCheck(
	openID string, scene int, version int, content string,
	nickname int, title int, signature string) (*response.ResponseSecurityMediaCheckASync, error) {

	result := &response.ResponseSecurityMediaCheckASync{}

	data := &object.HashMap{
		"openid":    openID,
		"scene":     scene,
		"version":   version,
		"content":   content,
		"nickname":  nickname,
		"title":     title,
		"signature": signature,
	}

	_, err := comp.HttpPostJson("wxa/msg_sec_check", data, nil, nil, result)

	return result, err
}
