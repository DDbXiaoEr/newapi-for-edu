package model

import (
	"slices"
	"sort"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/setting/group_channel_setting"
	"github.com/QuantumNous/new-api/setting/ratio_setting"
)

// groupBoundChannelIDs returns the explicit channel pin for group+model.
// Exact name wins, then the routing-normalized name, then a legacy "*" pin.
func groupBoundChannelIDs(group, model string) ([]int, bool) {
	if ids, ok := group_channel_setting.GetExactBoundChannelIDs(group, model); ok {
		return ids, true
	}
	normalized := ratio_setting.RoutingMatchModelName(model)
	if normalized != "" && normalized != model {
		if ids, ok := group_channel_setting.GetExactBoundChannelIDs(group, normalized); ok {
			return ids, true
		}
	}
	return group_channel_setting.GetWildcardChannelIDs(group)
}

// GroupHasModelBinding reports whether group+model has an explicit channel pin.
func GroupHasModelBinding(group, model string) bool {
	_, ok := groupBoundChannelIDs(group, model)
	return ok
}

// ChannelAllowedForGroup reports whether channelID may serve modelName for the
// given group. A pinned model uses its explicit channel list (ignoring the
// channel's own Group field); an unbound model keeps the legacy abilities-based
// membership.
func ChannelAllowedForGroup(group, modelName string, channelID int) bool {
	if group == "" || modelName == "" || channelID <= 0 {
		return false
	}
	if ids, ok := groupBoundChannelIDs(group, modelName); ok {
		if !slices.Contains(ids, channelID) {
			return false
		}
		return channelSupportsModelForGroup(channelID, modelName)
	}
	return IsChannelEnabledForGroupModel(group, modelName, channelID)
}

// channelModelsContain reports whether a channel model list covers the
// requested model, either exactly or through the routing-normalized name.
func channelModelsContain(models []string, model, normalized string) bool {
	for _, m := range models {
		if m == model {
			return true
		}
	}
	if normalized == "" || normalized == model {
		return false
	}
	for _, m := range models {
		if m == normalized {
			return true
		}
	}
	return false
}

func channelModelMatches(ch *Channel, model string) bool {
	if ch == nil {
		return false
	}
	return channelModelsContain(ch.GetModels(), model, ratio_setting.RoutingMatchModelName(model))
}

// channelSupportsModelForGroup reports whether the channel exists, is enabled
// and supports the model in the current storage mode.
func channelSupportsModelForGroup(channelID int, model string) bool {
	if !common.MemoryCacheEnabled {
		ch, err := GetChannelById(channelID, false)
		if err != nil || ch == nil || ch.Status != common.ChannelStatusEnabled {
			return false
		}
		return channelModelMatches(ch, model)
	}
	channelSyncLock.RLock()
	defer channelSyncLock.RUnlock()
	ch, ok := channelsIDM[channelID]
	if !ok || ch.Status != common.ChannelStatusEnabled {
		return false
	}
	return channelModelMatches(ch, model)
}

// boundCandidateIDsForModelLocked returns the enabled bound channels serving
// model (exact or routing-normalized). Caller must hold channelSyncLock (read).
func boundCandidateIDsForModelLocked(ids []int, model string) []int {
	normalized := ratio_setting.RoutingMatchModelName(model)
	candidates := make([]int, 0, len(ids))
	for _, id := range ids {
		ch, ok := channelsIDM[id]
		if !ok || ch.Status != common.ChannelStatusEnabled {
			continue
		}
		if channelModelsContain(ch.GetModels(), model, normalized) {
			candidates = append(candidates, id)
		}
	}
	return candidates
}

// boundChannelsByIDs loads channels by id. Used by the non-cache path.
func boundChannelsByIDs(ids []int) []*Channel {
	if len(ids) == 0 {
		return nil
	}
	var channels []*Channel
	if err := DB.Where("id IN ?", ids).Find(&channels).Error; err != nil {
		return nil
	}
	return channels
}

// boundAbilitiesForModel builds abilities-equivalent rows for the bound
// channels so the existing priority/weight selection can be reused.
func boundAbilitiesForModel(ids []int, model string) []Ability {
	channels := boundChannelsByIDs(ids)
	normalized := ratio_setting.RoutingMatchModelName(model)
	abilities := make([]Ability, 0, len(channels))
	for _, ch := range channels {
		if ch.Status != common.ChannelStatusEnabled {
			continue
		}
		if !channelModelsContain(ch.GetModels(), model, normalized) {
			continue
		}
		abilities = append(abilities, Ability{
			Group:     "(bound)",
			Model:     model,
			ChannelId: ch.Id,
			Enabled:   true,
			Priority:  ch.Priority,
			Weight:    uint(ch.GetWeight()),
			Tag:       ch.Tag,
		})
	}
	return abilities
}

// GroupChannelSummary is the admin-facing view of a channel used by the
// group binding editor (id, name, status and supported models).
type GroupChannelSummary struct {
	Id     int      `json:"id"`
	Name   string   `json:"name"`
	Status int      `json:"status"`
	Models []string `json:"models"`
}

// GetGroupChannelSummaries lists every channel for the binding editor.
func GetGroupChannelSummaries() []GroupChannelSummary {
	var channels []*Channel
	if err := DB.Find(&channels).Error; err != nil {
		return nil
	}
	summaries := make([]GroupChannelSummary, 0, len(channels))
	for _, ch := range channels {
		summaries = append(summaries, GroupChannelSummary{
			Id:     ch.Id,
			Name:   ch.Name,
			Status: ch.Status,
			Models: ch.GetModels(),
		})
	}
	sort.Slice(summaries, func(i, j int) bool { return summaries[i].Id < summaries[j].Id })
	return summaries
}

func anyBoundChannelSupportsModel(ids []int, model string) bool {
	for _, id := range ids {
		if channelSupportsModelForGroup(id, model) {
			return true
		}
	}
	return false
}

// boundGroupEnabledModels returns the union of models served by the bound
// enabled channels.
func boundGroupEnabledModels(ids []int) []string {
	modelSet := make(map[string]struct{})
	if common.MemoryCacheEnabled {
		channelSyncLock.RLock()
		for _, id := range ids {
			ch, ok := channelsIDM[id]
			if !ok || ch.Status != common.ChannelStatusEnabled {
				continue
			}
			for _, m := range ch.GetModels() {
				modelSet[m] = struct{}{}
			}
		}
		channelSyncLock.RUnlock()
	} else {
		for _, ch := range boundChannelsByIDs(ids) {
			if ch.Status != common.ChannelStatusEnabled {
				continue
			}
			for _, m := range ch.GetModels() {
				modelSet[m] = struct{}{}
			}
		}
	}
	models := make([]string, 0, len(modelSet))
	for m := range modelSet {
		models = append(models, m)
	}
	sort.Strings(models)
	return models
}

func collectGroupEnabledModels(group string) []string {
	modelSet := make(map[string]struct{})
	if ids, ok := group_channel_setting.GetWildcardChannelIDs(group); ok {
		for _, m := range boundGroupEnabledModels(ids) {
			modelSet[m] = struct{}{}
		}
	} else {
		var models []string
		DB.Table("abilities").Where(commonGroupCol+" = ? and enabled = ?", group, true).Distinct("model").Pluck("model", &models)
		for _, m := range models {
			modelSet[m] = struct{}{}
		}
	}
	for model, ids := range group_channel_setting.GetGroupModelsCopy(group) {
		if model == group_channel_setting.WildcardModel {
			continue
		}
		if anyBoundChannelSupportsModel(ids, model) {
			modelSet[model] = struct{}{}
		} else {
			delete(modelSet, model)
		}
	}
	out := make([]string, 0, len(modelSet))
	for m := range modelSet {
		out = append(out, m)
	}
	sort.Strings(out)
	return out
}
