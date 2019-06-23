package services

import (
	"encoding/json"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/rafaeleyng/push-api/push-api/models"
)

type (
	PublishingResult int

	PublicationService interface {
		PublishMessage(*models.Message) PublishingResult
	}

	publicationService struct{
		logger *zap.Logger
		machineryServer *machinery.Server
		taskName string
	}
)

const (
	PublishingSuccess PublishingResult = iota
	PublishingInvalid
	PublishingFailure
)

func (s *publicationService) buildTaskSignature(messageJson *string) *tasks.Signature {
	return &tasks.Signature{
		Name: s.taskName,
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: *messageJson,
			},
		},
	}
}

func (s *publicationService) PublishMessage(message *models.Message) PublishingResult {
	bytes, err := json.Marshal(message)
	if err != nil {
		s.logger.Error("error marshaling message", zap.Any("message", message), zap.Error(err))
		return PublishingInvalid
	}
	messageJson := string(bytes)
	signature := s.buildTaskSignature(&messageJson)

	_, err = s.machineryServer.SendTask(signature)
	if err != nil {
		s.logger.Error("error publishing message", zap.Any("message", message), zap.Error(err))
		return PublishingFailure
	}

	s.logger.Debug("message published", zap.String("message", messageJson))
	return PublishingSuccess
}

func NewPublicationService(config *viper.Viper, logger *zap.Logger, machineryServer *machinery.Server) PublicationService {
	return &publicationService{
		logger: logger.Named("publicationService"),
		machineryServer: machineryServer,
		taskName: config.GetString("redis.pubsub.publish_task"),
	}
}
