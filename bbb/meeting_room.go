package bbb

import (
	"encoding/xml"
	"log"
	"net/url"

	"github.com/baltuonis/prometheus-bigbluebutton-exporter/command"
	"github.com/baltuonis/prometheus-bigbluebutton-exporter/models"
)

const (
	BASE_URL = ""
	SALT     = "" // 公钥
)

// MeetingRoom BBB response object
type MeetingRoom struct {
	Name_                   string
	MeetingID_              string
	AttendeePW_             string
	ModeratorPW_            string
	Welcome                 string
	DialNumber              string
	VoiceBridge             string
	WebVoice                string
	LogoutURL               string
	Record                  string
	Duration                int
	Meta                    string
	ModeratorOnlyMessage    string
	AutoStartRecording      bool
	AllowStartStopRecording bool

	MeetingInfo    models.GetMeetingInfoResponse
	Participantses []Participants
}

// IsMeeetingRunning - checks if meeting is running
func (meetingRoom *MeetingRoom) IsMeetingRunning() bool {
	if "" == meetingRoom.MeetingID_ {
		log.Println("ERROR: PARAM ERROR.")
		return false
	}

	createParam := "meetingID=" + url.QueryEscape(meetingRoom.MeetingID_)
	checksum := command.GetChecksum("isMeetingRunning", createParam, SALT)

	createResponse := command.HttpGet(BASE_URL + "isMeetingRunning?" + createParam +
		"&checksum=" + checksum)

	if "ERROR" == createResponse {
		log.Println("ERROR: HTTP ERROR.")
		return false
	}

	responseXML := models.IsMeetingRunningResponse{}
	err := xml.Unmarshal([]byte(createResponse), responseXML)

	if nil != err {
		log.Println("XML PARSE ERROR: " + err.Error())
		return false
	}

	if "SUCCESS" == responseXML.ReturnCode {
		log.Println("CALLED SUCCESS.")
		return responseXML.Running
	}

	return false
}

// GetMeetingInfo - gets meeting info
func (meetingRoom *MeetingRoom) GetMeetingInfo() *models.GetMeetingInfoResponse {
	if "" == meetingRoom.MeetingID_ || "" == meetingRoom.ModeratorPW_ {
		log.Println("ERROR: PARAM ERROR.")
		return nil
	}

	createParam := "meetingID=" + url.QueryEscape(meetingRoom.MeetingID_) +
		"&password=" + url.QueryEscape(meetingRoom.ModeratorPW_)
	checksum := command.GetChecksum("getMeetingInfo", createParam, SALT)

	createResponse := command.HttpGet(BASE_URL + "getMeetingInfo?" + createParam +
		"&checksum=" + checksum)

	if "ERROR" == createResponse {
		log.Println("ERROR: HTTP ERROR.")
		return nil
	}

	err := xml.Unmarshal([]byte(createResponse), &meetingRoom.MeetingInfo)

	if nil != err {
		log.Println("XML PARSE ERROR: " + err.Error())
		return nil
	}

	if "SUCCESS" == meetingRoom.MeetingInfo.ReturnCode {
		log.Println("GET MEETING INFO SUCCESS.")
		return &meetingRoom.MeetingInfo
	}

	return nil
}
