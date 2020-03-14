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

	createResponse := command.HttpGet(bbbClient.BaseURL + "getMeetings?checksum=" + checksum)

	if "ERROR" == createResponse {
		log.Println("ERROR: HTTP ERROR.")
		return nil
	}

	var response models.GetMeetingsResponse
	err := xml.Unmarshal([]byte(createResponse), &response)

	if nil != err {
		log.Println("XML PARSE ERROR: " + err.Error())
		return nil
	}

	if "SUCCESS" == response.ReturnCode {
		if bbbClient.Debug {
			log.Println("GET MEETING INFO SUCCESS.")
		}
		return &response
	}

	return nil
}
