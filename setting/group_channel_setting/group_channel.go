package group_channel_setting

import (
	"bytes"
	"sort"
	"sync"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/setting/config"
)

// OptionKeyGroupChannels is the flat option key used to persist the
// group -> model -> channel binding map.
const OptionKeyGroupChannels = "group_channel_setting.group_channels"

// WildcardModel is the legacy group-wide binding key. Old payloads stored
// group -> []channelID; they are migrated to group -> {"*": ids} so existing
// bindings keep covering every model until an admin saves a per-model pin.
const WildcardModel = "*"

// GroupChannelSetting stores per-model channel pins for user groups.
// Semantics:
//   - group+model with a non-empty channel list is authoritative for that model.
//   - group+"*" (legacy) is authoritative for models without a specific pin.
//   - otherwise routing keeps the legacy channel.Group / abilities behavior.
type GroupChannelSetting struct {
	GroupChannels *groupChannelsStore `json:"group_channels"`
}

type groupChannelsStore struct {
	mu   sync.RWMutex
	data map[string]map[string][]int
}

func newGroupChannelsStore() *groupChannelsStore {
	return &groupChannelsStore{data: make(map[string]map[string][]int)}
}

func (s *groupChannelsStore) UnmarshalJSON(b []byte) error {
	parsed, err := parseBindings(b)
	if err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = parsed
	return nil
}

func (s *groupChannelsStore) MarshalJSON() ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return common.Marshal(s.data)
}

var groupChannelSetting = GroupChannelSetting{
	GroupChannels: newGroupChannelsStore(),
}

func init() {
	config.GlobalConfig.Register("group_channel_setting", &groupChannelSetting)
}

func parseBindings(data []byte) (map[string]map[string][]int, error) {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 || string(trimmed) == "null" {
		return map[string]map[string][]int{}, nil
	}
	var object map[string]common.RawMessage
	if err := common.Unmarshal(trimmed, &object); err != nil {
		return nil, err
	}
	out := make(map[string]map[string][]int, len(object))
	for group, raw := range object {
		switch common.GetJsonType(raw) {
		case "array":
			var ids []int
			if err := common.Unmarshal(raw, &ids); err != nil {
				return nil, err
			}
			normalized := normalizeIDs(ids)
			if len(normalized) == 0 {
				continue
			}
			out[group] = map[string][]int{WildcardModel: normalized}
		case "object":
			var models map[string][]int
			if err := common.Unmarshal(raw, &models); err != nil {
				return nil, err
			}
			normalizedModels := normalizeModelIDs(models)
			if len(normalizedModels) == 0 {
				continue
			}
			out[group] = normalizedModels
		}
	}
	return out, nil
}

// GetExactBoundChannelIDs returns the pin stored under group+model with no
// wildcard fallback. The bool is false when that exact key is absent.
func GetExactBoundChannelIDs(group, model string) ([]int, bool) {
	if group == "" || model == "" {
		return nil, false
	}
	store := groupChannelSetting.GroupChannels
	store.mu.RLock()
	defer store.mu.RUnlock()
	models, ok := store.data[group]
	if !ok {
		return nil, false
	}
	ids, ok := models[model]
	if !ok || len(ids) == 0 {
		return nil, false
	}
	return append([]int(nil), ids...), true
}

// GetBoundChannelIDs returns the pinned channel ids for group+model.
// A model-specific pin wins over a legacy "*" pin. The bool is false when
// that model is unbound and callers must fall back to legacy routing.
func GetBoundChannelIDs(group, model string) ([]int, bool) {
	if ids, ok := GetExactBoundChannelIDs(group, model); ok {
		return ids, true
	}
	if model != WildcardModel {
		return GetExactBoundChannelIDs(group, WildcardModel)
	}
	return nil, false
}

// GetWildcardChannelIDs returns the legacy group-wide pin, if any.
func GetWildcardChannelIDs(group string) ([]int, bool) {
	return GetExactBoundChannelIDs(group, WildcardModel)
}

// GetGroupModelsCopy returns a copy of the per-model pins for a group.
func GetGroupModelsCopy(group string) map[string][]int {
	store := groupChannelSetting.GroupChannels
	store.mu.RLock()
	defer store.mu.RUnlock()
	return cloneModelIDs(store.data[group])
}

// GetGroupChannelsCopy returns a deep copy of the whole binding map.
func GetGroupChannelsCopy() map[string]map[string][]int {
	store := groupChannelSetting.GroupChannels
	store.mu.RLock()
	defer store.mu.RUnlock()
	out := make(map[string]map[string][]int, len(store.data))
	for group, models := range store.data {
		out[group] = cloneModelIDs(models)
	}
	return out
}

// SetGroupChannels replaces the whole binding map. Empty lists and empty
// groups are dropped so that "unbind" is represented by removing the key.
func SetGroupChannels(bindings map[string]map[string][]int) {
	normalized := make(map[string]map[string][]int, len(bindings))
	for group, models := range bindings {
		nm := normalizeModelIDs(models)
		if len(nm) == 0 {
			continue
		}
		normalized[group] = nm
	}
	store := groupChannelSetting.GroupChannels
	store.mu.Lock()
	defer store.mu.Unlock()
	store.data = normalized
}

// GroupChannels2JSONString serializes the binding map for persistence.
func GroupChannels2JSONString() string {
	bytes, err := groupChannelSetting.GroupChannels.MarshalJSON()
	if err != nil {
		return "{}"
	}
	return string(bytes)
}

func cloneModelIDs(models map[string][]int) map[string][]int {
	if len(models) == 0 {
		return map[string][]int{}
	}
	out := make(map[string][]int, len(models))
	for model, ids := range models {
		out[model] = append([]int(nil), ids...)
	}
	return out
}

func normalizeModelIDs(models map[string][]int) map[string][]int {
	if len(models) == 0 {
		return nil
	}
	out := make(map[string][]int, len(models))
	for model, ids := range models {
		normalized := normalizeIDs(ids)
		if len(normalized) == 0 {
			continue
		}
		out[model] = normalized
	}
	if len(out) == 0 {
		return nil
	}
	return out
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
