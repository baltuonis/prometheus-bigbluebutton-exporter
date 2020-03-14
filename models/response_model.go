package models

// GetMeetingsResponse is response of /getMeetings
type GetMeetingsResponse struct {
	ReturnCode string   `xml:"returncode"` // 是否成功
	Meetings   meetings `xml:"meetings"`
}

//-----------------------------------------------------------------------------
// 建立会议室返回XML的数据结构, 即create接口调用的返回值实例
type CreateMeetingResponse struct {
	Returncode           string `xml:"returncode"`           // 是否成功
	MeetingID            string `xml:"meetingID"`            // 会议ID
	CreateTime           string `xml:"createTime"`           // 会议创建时间
	AttendeePW           string `xml:"attendeePW"`           // 与会者密码
	ModeratorPW          string `xml:"moderatorPW"`          // 会议管理员密码
	HasBeenForciblyEnded string `xml:"hasBeenForciblyEnded"` // 是否可以被强制结束
	MessageKey           string `xml:"messageKey"`           // 返回消息Key
	Message              string `xml:"message"`              // 返回消息
}

//-----------------------------------------------------------------------------
// 检查会议室是否在运行返回XML的数据结构, 即isMeetingRunning接口调用的返回值实例
type IsMeetingRunningResponse struct {
	ReturnCode string `xml:"returncode"` // 是否成功
	Running    bool   `xml:"running"`    // 会议室状态
}

//-----------------------------------------------------------------------------
// 关闭会议室返回XML的数据结构, 即end接口调用的返回值实例
type EndResponse struct {
	ReturnCode string `xml:"returncode"` // 是否成功
	MessageKey string `xml:"messageKey"` // 返回消息Key
	Message    string `xml:"message"`    // 返回消息
}

//-----------------------------------------------------------------------------
// 获取会议信息返回XML的数据结构, 即getMeetingInfo接口调用的返回值实例
type GetMeetingInfoResponse struct {
	ReturnCode            string    `xml:"returncode"`            // 是否成功
	MeetingName           string    `xml:"meetingName"`           // 会议名称
	MeetingID             string    `xml:"meetingID"`             // 会议ID
	InternalMeetingID     string    `xml:"internalMeetingID"`     // 内部会议ID, 由系统随机分发
	CreateTime            string    `xml:"createTime"`            // 会议室创建时间
	CreateDate            string    `xml:"createDate"`            // 会议室创建日期
	VoiceBridge           string    `xml:"voiceBridge"`           // 通过电话拨入语音会议时需要输入的PIN码
	DialNumber            string    `xml:"dialNumber"`            // 可通过电话直接拨入语音会议的号码
	AttendeePW            string    `xml:"attendeePW"`            // 与会者密码
	ModeratorPW           string    `xml:"moderatorPW"`           // 管理员密码
	Running               bool      `xml:"running"`               // 是否正在运行
	Duration              int       `xml:"duration"`              // 会议室有效时长
	HasUserJoined         bool      `xml:"hasUserJoined"`         // 是否有用户加入
	Recording             bool      `xml:"recording"`             // 是否可以录制会议
	HasBeenForciblyEnded  bool      `xml:"hasBeenForciblyEnded"`  // 是否已经被强制结束
	StartTime             string    `xml:"startTime"`             // 会议开始时间
	EndTime               string    `xml:"endTime"`               // 会议结束时间
	ParticipantCount      int       `xml:"participantCount"`      // 会议室内参与人数
	ListenerCount         int       `xml:"listenerCount"`         // 聆听着数量
	VoiceParticipantCount int       `xml:"voiceParticipantCount"` // 语音参与者数量
	VideoCount            int       `xml:"videoCount"`            // 视频参与者数量
	MaxUsers              int       `xml:"maxUsers"`              // 最大可进入人数
	ModeratorCount        int       `xml:"moderatorCount"`        // 会议管理员数量
	Attendees             attendees `xml:"attendees"`             // 与会者信息
	Metadata              string    `xml:"metadata"`              // 元数据
	MessageKey            string    `xml:"messageKey"`            // 返回的消息Key
	Message               string    `xml:"message"`               // 返回的消息
}

type attendees struct {
	Attendees []attendee `xml:"attendee"`
}

type meetings struct {
	Meetings []GetMeetingInfoResponse `xml:"meeting"`
}

type attendee struct {
	UserID          string `xml:"userID"`          // 参与者ID, 加入会议室时设定的
	FullName        string `xml:"fullName"`        // 参与者名称
	Role            string `xml:"role"`            // 角色, 即是管理员还是与会者
	IsPresenter     bool   `xml:"isPresenter"`     // 是否会议主持人
	IsListeningOnly bool   `xml:"isListeningOnly"` // 是否只能听讲
	HasJoinedVoice  bool   `xml:"hasJoinedVoice"`  // 是否有声音, 即是否有麦克风发音
	HasVideo        bool   `xml:"hasVideo"`        // 是否有视频
	Customdata      string `xml:"customdata"`      // 自定义数据
}
