package bbb

import (
	"net/url"

	"github.com/baltuonis/prometheus-bigbluebutton-exporter/command"
)

/*******************************************************************************
* 参与者类, 抽象会议的参与者, 可以是管理员也可以是与会者
*******************************************************************************/
type Participants struct {
	IsAdmin_     int    // 标识参与者是否管理员, 1代表管理员, 0代表非管理员
	FullName_    string // 必填, 用户名
	MeetingID_   string // 必填, 试图加入的会议ID
	Password_    string // 必填, 会议室密码, 这里特指与会者密码, 如果传入管理员密码则以管理员身份进入
	CreateTime   string // 可选, 会议室创建时间, 用来匹配MeetingID, 避免同一个参会者多次进入
	UserID       string // 可选, 标识用户身份的ID, 在调用GetMeetingInfo时将被返回
	WebVoiceConf string // 可选, VOIP协议扩展
	ConfigToken  string // 可选, 由SetConfigXML调用返回的Token
	AvatarURL    string // 可选, 用户头像URL, 当config.xml中displayAvatar为true时提供
	Redirect     string // 可选, 实验, 当HTML5不可用时, 用来重定向到Flash客户端
	ClientURL    string // 可选, 试验, 用来显示自动以的客户端名称
	JoinURL      string // 该参与者加入会议的URL

}

/*******************************************************************************
* 根据参与者类的配置获得要加入的会议室的地址, 获取的地址可以直接进入到会议室当中
* 返回: 加入指定会议室的URL
*******************************************************************************/
func (participants *Participants) GetJoinURL() string {
	if "" == participants.FullName_ || "" == participants.MeetingID_ ||
		"" == participants.Password_ {
		return "ERROR: PARAM ERROR."
	}

	// 构造必填参数
	fullName := "fullName=" + url.QueryEscape(participants.FullName_)     // 用户名
	meetingID := "&meetingID=" + url.QueryEscape(participants.MeetingID_) // 试图加入的会议ID
	password := "&password=" + url.QueryEscape(participants.Password_)    // 会议室密码, 这里特指与会者密码, 如果传入管理员密码则以管理员身份进入

	var createTime string  // 会议室创建时间, 用来匹配MeetingID, 避免同一个参会者多次进入
	var userID string      // 标识用户身份的ID, 在调用GetMeetingInfo时将返回
	var configToken string // 有SetConfigXML调用返回的Token
	var avatarURL string   // 用户头像的URL, 当config.xml中displayAvatar为true时提供
	var redirect string    // 当HTML5不可用时, 是否重定向到Flash客户端
	var clientURL string   // 重定向URL

	if "" != participants.CreateTime {
		createTime = "&createTime=" + url.QueryEscape(participants.CreateTime)
	}

	if "" != participants.UserID {
		userID = "&userID=" + url.QueryEscape(participants.UserID)
	}

	if "" != participants.ConfigToken {
		configToken = "&configToken=" + url.QueryEscape(participants.ConfigToken)
	}

	if "" != participants.AvatarURL {
		avatarURL = "&avatarURL=" + url.QueryEscape(participants.AvatarURL)
	}

	if "" != participants.ClientURL {
		redirect = "&redirect=true"
		clientURL = "&clientURL=" + url.QueryEscape(participants.ClientURL)
	}

	// 合成请求参数
	joinParam := fullName + meetingID + password + createTime + userID +
		configToken + avatarURL + redirect + clientURL

	checksum := command.GetChecksum("join", joinParam, SALT)
	joinUrl := BASE_URL + "join?" + joinParam + "&checksum=" + checksum
	participants.JoinURL = joinUrl

	return joinUrl
}
