package controller

import (
	"sort"
	"strings"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/setting/group_channel_setting"
	"github.com/QuantumNous/new-api/setting/ratio_setting"

	"github.com/gin-gonic/gin"
)

type groupChannelBindingRequest struct {
	Group      string `json:"group"`
	Model      string `json:"model"`
	ChannelIDs []int  `json:"channel_ids"`
}

// GetGroupChannelBindings returns the group -> model -> channels mapping
// together with the channel list used by the binding editor.
func GetGroupChannelBindings(c *gin.Context) {
	bindings := group_channel_setting.GetGroupChannelsCopy()
	summaries := model.GetGroupChannelSummaries()

	existing := make(map[int]struct{}, len(summaries))
	for _, summary := range summaries {
		existing[summary.Id] = struct{}{}
	}
	staleChannels := make(map[string]map[string][]int)
	for group, models := range bindings {
		for modelName, ids := range models {
			for _, id := range ids {
				if _, ok := existing[id]; !ok {
					if staleChannels[group] == nil {
						staleChannels[group] = make(map[string][]int)
					}
					staleChannels[group][modelName] = append(staleChannels[group][modelName], id)
				}
			}
		}
	}

	groups := make([]string, 0)
	groupModels := make(map[string][]string)
	for name := range ratio_setting.GetGroupRatioCopy() {
		groups = append(groups, name)
		groupModels[name] = model.GetGroupEnabledModels(name)
	}
	sort.Strings(groups)

	common.ApiSuccess(c, gin.H{
		"group_channels": bindings,
		"channels":       summaries,
		"stale_channels": staleChannels,
		"groups":         groups,
		"group_models":   groupModels,
	})
}

// UpdateGroupChannelBinding sets (or clears, when channel_ids is empty) the
// explicit channel pin for a single group+model. It persists through the
// generic option path so the change hot-reloads.
func UpdateGroupChannelBinding(c *gin.Context) {
	var req groupChannelBindingRequest
	if err := common.DecodeJson(c.Request.Body, &req); err != nil {
		common.ApiErrorMsg(c, "invalid request body")
		return
	}

	group := strings.TrimSpace(req.Group)
	if group == "" {
		common.ApiErrorMsg(c, "group is required")
		return
	}
	if !ratio_setting.ContainsGroupRatio(group) {
		common.ApiErrorMsg(c, "group is not defined in group ratio")
		return
	}
	modelName := strings.TrimSpace(req.Model)
	if modelName == "" {
		common.ApiErrorMsg(c, "model is required")
		return
	}

	existing := make(map[int]struct{})
	for _, summary := range model.GetGroupChannelSummaries() {
		existing[summary.Id] = struct{}{}
	}

	seen := make(map[int]struct{}, len(req.ChannelIDs))
	ids := make([]int, 0, len(req.ChannelIDs))
	for _, id := range req.ChannelIDs {
		if id <= 0 {
			continue
		}
		if _, ok := existing[id]; !ok {
			common.ApiErrorMsg(c, "channel does not exist")
			return
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		ids = append(ids, id)
	}
	sort.Ints(ids)

	bindings := group_channel_setting.GetGroupChannelsCopy()
	models := bindings[group]
	if models == nil {
		models = make(map[string][]int)
	}
	if len(ids) == 0 {
		delete(models, modelName)
	} else {
		models[modelName] = ids
	}
	if len(models) == 0 {
		delete(bindings, group)
	} else {
		bindings[group] = models
	}

	payload, err := common.Marshal(bindings)
	if err != nil {
		common.ApiError(c, err)
		return
	}
	if err := model.UpdateOption(group_channel_setting.OptionKeyGroupChannels, string(payload)); err != nil {
		common.ApiError(c, err)
		return
	}

	common.ApiSuccess(c, gin.H{
		"group_channels": group_channel_setting.GetGroupChannelsCopy(),
	})
}
