package group_channel_setting

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseBindingsMigratesLegacyGroupList(t *testing.T) {
	parsed, err := parseBindings([]byte(`{"default":[2,1,1],"vip":{"gpt-4":[9],"claude":[]}}`))
	require.NoError(t, err)
	assert.Equal(t, map[string]map[string][]int{
		"default": {WildcardModel: {1, 2}},
		"vip":     {"gpt-4": {9}},
	}, parsed)
}

func TestGetBoundChannelIDsPrefersModelPin(t *testing.T) {
	previous := GetGroupChannelsCopy()
	t.Cleanup(func() { SetGroupChannels(previous) })

	SetGroupChannels(map[string]map[string][]int{
		"default": {
			WildcardModel: {1, 2},
			"gpt-4":       {3},
		},
	})

	ids, ok := GetBoundChannelIDs("default", "gpt-4")
	require.True(t, ok)
	assert.Equal(t, []int{3}, ids)

	_, exact := GetExactBoundChannelIDs("default", "claude")
	assert.False(t, exact)

	ids, ok = GetBoundChannelIDs("default", "claude")
	require.True(t, ok)
	assert.Equal(t, []int{1, 2}, ids)

	_, ok = GetBoundChannelIDs("vip", "gpt-4")
	assert.False(t, ok)
}

func TestSetGroupChannelsDropsEmptyPins(t *testing.T) {
	previous := GetGroupChannelsCopy()
	t.Cleanup(func() { SetGroupChannels(previous) })

	SetGroupChannels(map[string]map[string][]int{
		"default": {"gpt-4": {}, "claude": {5}},
		"vip":     {},
	})

	assert.Equal(t, map[string]map[string][]int{
		"default": {"claude": {5}},
	}, GetGroupChannelsCopy())
}
