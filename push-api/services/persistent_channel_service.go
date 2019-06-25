package services

import (
	"go.uber.org/zap"

	"github.com/rafaeleyng/push-api/push-api/models"
)

type (
	PersistentChannelService interface {
		TriggerRevivePersistentChannels()
	}

	persistentChannelService struct {
		logger *zap.Logger
		publicationService PublicationService
		channelService ChannelService
	}
)

func (s *persistentChannelService) TriggerRevivePersistentChannels() {
	channels, channelResult := s.channelService.GetAll()
	if channelResult != ChannelRetrievalSuccess {
		s.logger.Error("failed to retrieve persistent channels to revive")
		return
	} else if len(channels) == 0 {
		s.logger.Debug("no channels to revive")
		return
	}

	channelsIds := make([]string, len(channels))

	for i, channel := range channels {
		channelsIds[i] = channel.Id
	}

	messageResult := s.publicationService.PublishMessage(&models.Message{
		Channels: channelsIds,
		Content:  "ping",
	})

	if messageResult == PublishingFailure {
		s.logger.Error("failed to revive persistent channels")
		return
	}

	s.logger.Debug("did revive persistent channels")
}

func NewPersistentChannelService(logger *zap.Logger, publicationService PublicationService, channelService ChannelService) PersistentChannelService {
	return &persistentChannelService{
		logger: logger.Named("persistentChannelService"),
		publicationService: publicationService,
		channelService: channelService,
	}
}
