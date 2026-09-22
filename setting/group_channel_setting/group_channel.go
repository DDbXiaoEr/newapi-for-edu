package group_channel_setting

import (
	"sort"

	"github.com/QuantumNous/new-api/setting/config"
	"github.com/QuantumNous/new-api/types"
)

// OptionKeyGroupChannels is the flat option key used to persist the
// group -> channel binding map. It follows the "module.field" convention so
// the generic option update path hot-reloads the registered config below.
const OptionKeyGroupChannels = "group_channel_setting.group_channels"

// GroupChannelSetting stores the explicit channel binding for user groups.
// Semantics:
//   - A group present with a non-empty channel list is authoritative: requests
//     using that group only route through those channels, ignoring each
//     channel's own Group field.
//   - A group absent (or present with an empty list) is unbound: routing keeps
//     the legacy channel.Group / abilities behavior.
type GroupChannelSetting struct {
	GroupChannels *types.RWMap[string, []int] `json:"group_channels"`
}

var groupChannelSetting = GroupChannelSetting{
	GroupChannels: types.NewRWMap[string, []int](),
}

func init() {
	config.GlobalConfig.Register("group_channel_setting", &groupChannelSetting)
}

// GetBoundChannelIDs returns the bound channel ids for a group. The bool is
// false when the group is unbound (absent or empty), in which case callers must
// fall back to the legacy behavior.
func GetBoundChannelIDs(group string) ([]int, bool) {
	ids, ok := groupChannelSetting.GroupChannels.Get(group)
	if !ok || len(ids) == 0 {
		return nil, false
	}
	return ids, true
}

// GetGroupChannelsCopy returns a copy of the whole binding map.
func GetGroupChannelsCopy() map[string][]int {
	return groupChannelSetting.GroupChannels.ReadAll()
}

// SetGroupChannels replaces the whole binding map. Empty lists are dropped so
// that "unbind" is represented by removing the key.
func SetGroupChannels(bindings map[string][]int) {
	normalized := make(map[string][]int, len(bindings))
	for group, ids := range bindings {
		if len(ids) == 0 {
			continue
		}
		normalized[group] = normalizeIDs(ids)
	}
	groupChannelSetting.GroupChannels.Clear()
	groupChannelSetting.GroupChannels.AddAll(normalized)
}

// GroupChannels2JSONString serializes the binding map for persistence.
func GroupChannels2JSONString() string {
	return groupChannelSetting.GroupChannels.MarshalJSONString()
}

func normalizeIDs(ids []int) []int {
	if len(ids) == 0 {
		return nil
	}
	seen := make(map[int]struct{}, len(ids))
	unique := make([]int, 0, len(ids))
	for _, id := range ids {
		if id <= 0 {
			continue
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		unique = append(unique, id)
	}
	sort.Ints(unique)
	return unique
}
