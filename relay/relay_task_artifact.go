package relay

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/model"
	relaychannel "github.com/QuantumNous/new-api/relay/channel"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/service"
	"github.com/QuantumNous/new-api/types"
)

var ErrTaskArtifactPluginUnavailable = errors.New("task artifact plugin unavailable")

// InitTaskArtifactAdaptor builds a task adaptor initialized with the channel
// metadata needed to list artifacts and construct upstream content requests.
func InitTaskArtifactAdaptor(task *model.Task) (relaychannel.TaskAdaptor, error) {
	if task == nil || !TaskHasPluginExecution(task) {
		return nil, ErrTaskArtifactPluginUnavailable
	}
	channelModel, err := model.CacheGetChannel(task.ChannelId)
	if err != nil {
		return nil, fmt.Errorf("%w: channel unavailable", ErrTaskArtifactPluginUnavailable)
	}
	adaptor := GetTaskAdaptor(task.Platform)
	if adaptor == nil {
		return nil, ErrTaskArtifactPluginUnavailable
	}
	pluginKey := task.PrivateData.Key
	if pluginKey == "" {
		pluginKey = channelModel.Key
	}
	baseURL := channelModel.GetBaseURL()
	if baseURL == "" {
		baseURL = constant.GetChannelBaseURL(channelModel.Type)
	}
	adaptor.Init(&relaycommon.RelayInfo{
		ChannelMeta: &relaycommon.ChannelMeta{
			ChannelType:    channelModel.Type,
			ChannelBaseUrl: baseURL,
			ApiKey:         pluginKey,
			ChannelSetting: channelModel.GetSetting(),
		},
	})
	return adaptor, nil
}

func TaskHasPluginExecution(task *model.Task) bool {
	return task != nil &&
		task.PrivateData.Execution != nil &&
		task.PrivateData.Execution.TaskPlugin != nil &&
		strings.TrimSpace(task.PrivateData.Execution.TaskPlugin.Key) != ""
}

// ListTaskArtifacts returns the artifacts exposed by a task plugin.
func ListTaskArtifacts(task *model.Task) ([]types.TaskArtifact, error) {
	adaptor, err := InitTaskArtifactAdaptor(task)
	if err != nil {
		return nil, err
	}
	provider, ok := adaptor.(relaychannel.TaskArtifactProvider)
	if !ok {
		return nil, nil
	}
	artifacts, err := provider.ListArtifacts(task)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrTaskArtifactPluginUnavailable, err)
	}
	return artifacts, nil
}

type taskArtifactSource struct {
	task            *model.Task
	provider        relaychannel.TaskArtifactProvider
	contentProvider relaychannel.TaskContentRequestProvider
}

func NewTaskArtifactSource(task *model.Task) (service.TaskArtifactSource, error) {
	if task == nil || !TaskHasPluginExecution(task) {
		return nil, nil
	}
	adaptor, err := InitTaskArtifactAdaptor(task)
	if err != nil {
		return nil, err
	}
	provider, ok := adaptor.(relaychannel.TaskArtifactProvider)
	if !ok {
		return nil, nil
	}
	contentProvider, ok := adaptor.(relaychannel.TaskContentRequestProvider)
	if !ok {
		return nil, nil
	}
	return &taskArtifactSource{
		task:            task,
		provider:        provider,
		contentProvider: contentProvider,
	}, nil
}

func (s *taskArtifactSource) ListArtifacts() ([]types.TaskArtifact, error) {
	artifacts, err := s.provider.ListArtifacts(s.task)
	if err != nil {
		return nil, err
	}
	return artifacts, nil
}

func (s *taskArtifactSource) BuildContentRequest(artifactKey string) (*service.TaskArtifactFetchRequest, error) {
	descriptor, err := s.contentProvider.BuildContentRequest(s.task, artifactKey, relaychannel.TaskArtifactClientRequest{
		Method: http.MethodGet,
	})
	if err != nil {
		return nil, err
	}
	if descriptor == nil {
		return nil, errors.New("artifact content plugin returned no request")
	}
	return &service.TaskArtifactFetchRequest{
		URL:            descriptor.URL,
		Method:         descriptor.Method,
		Headers:        descriptor.Headers,
		Body:           descriptor.Body,
		Credentialless: descriptor.Credentialless,
	}, nil
}
