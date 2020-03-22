package bbb

import (
	"encoding/xml"
	"log"

	"github.com/baltuonis/prometheus-bigbluebutton-exporter/command"
	"github.com/baltuonis/prometheus-bigbluebutton-exporter/models"
)

type BBBClient struct {
	BaseURL string
	Secret  string
	Debug   bool
}

// GetMeetings gets all meetings data
func (bbbClient *BBBClient) GetMeetings() *models.GetMeetingsResponse {
	checksum := command.GetChecksum("getMeetings", "", bbbClient.Secret)

	bbbResponse := command.HttpGet(bbbClient.BaseURL + "getMeetings?checksum=" + checksum)

	if "ERROR" == bbbResponse {
		log.Println("GetMeetings: HTTP ERROR.")
		return nil
	}

	var response models.GetMeetingsResponse
	err := xml.Unmarshal([]byte(bbbResponse), &response)

	if nil != err {
		log.Println("GetMeetings: XML PARSE ERROR: " + err.Error())
		log.Println("GetMeetings: BBB server response: " + bbbResponse)
		return nil
	}

	if "SUCCESS" == response.ReturnCode {
		if bbbClient.Debug {
			log.Println("GetMeetings: REQUEST SUCCESS.")
		}
		return &response
	}

	return nil
}
