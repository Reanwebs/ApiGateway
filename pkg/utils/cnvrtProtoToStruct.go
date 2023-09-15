package utils

import (
	"errors"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb/conference"

	"github.com/golang/protobuf/ptypes"
)

func ConvertProtoToStruct(protoResponse *conference.ScheduledConferenceResponse) (*models.ScheduledConferenceResponse, error) {
	if protoResponse == nil {
		return nil, errors.New("protobuf response is nil")
	}

	result := &models.ScheduledConferenceResponse{}

	result.Result = protoResponse.Result

	for _, protoConf := range protoResponse.Data {
		timeVal, err := ptypes.Timestamp(protoConf.Time)
		if err != nil {
			return nil, err
		}

		scheduledConf := models.ScheduledConference{
			UserID:           protoConf.UserID,
			ScheduleID:       protoConf.ScheduleID,
			Title:            protoConf.Title,
			Description:      protoConf.Description,
			Interest:         protoConf.Interest,
			Chat:             protoConf.Chat,
			Participantlimit: protoConf.Participantlimit,
			Time:             timeVal,
			Status:           protoConf.Status,
			Durations:        protoConf.Durations,
		}
		result.ScheduledConference = append(result.ScheduledConference, scheduledConf)
	}

	return result, nil
}
