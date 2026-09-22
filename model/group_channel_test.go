package model

import (
	"testing"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/setting/group_channel_setting"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupGroupChannelTest(t *testing.T) {
	t.Helper()
	truncateTables(t)
	require.NoError(t, DB.Exec("DELETE FROM abilities").Error)
	require.NoError(t, DB.Exec("DELETE FROM channels").Error)

	previous := group_channel_setting.GetGroupChannelsCopy()
	t.Cleanup(func() { group_channel_setting.SetGroupChannels(previous) })
	group_channel_setting.SetGroupChannels(nil)
}

func createGroupChannel(t *testing.T, name string, group string, models string) *Channel {
	t.Helper()
	channel := &Channel{
		Name:   name,
		Key:    "test-key",
		Status: common.ChannelStatusEnabled,
		Models: models,
		Group:  group,
	}
	require.NoError(t, DB.Create(channel).Error)
	require.NoError(t, channel.UpdateAbilities(nil))
	return channel
}

func TestGroupEnabledModelsUsesBinding(t *testing.T) {
	setupGroupChannelTest(t)
	group := "bound-model-group"
	bound := createGroupChannel(t, "bound", group, "model-a,model-shared")
	other := createGroupChannel(t, "other", group, "model-b,model-shared")

	group_channel_setting.SetGroupChannels(map[string][]int{group: {bound.Id}})

	assert.ElementsMatch(t, []string{"model-a", "model-shared"}, GetGroupEnabledModels(group))

	// Empty list means unbound: legacy union of both channels.
	group_channel_setting.SetGroupChannels(map[string][]int{group: {}})
	assert.False(t, GroupHasBinding(group))
	assert.ElementsMatch(t, []string{"model-a", "model-b", "model-shared"}, GetGroupEnabledModels(group))

	// Sanity: abilities were built for both channels.
	var count int64
	require.NoError(t, DB.Model(&Ability{}).Where("channel_id IN ?", []int{bound.Id, other.Id}).Count(&count).Error)
	assert.EqualValues(t, 4, count)
}

func TestChannelAllowedForGroupWithBinding(t *testing.T) {
	setupGroupChannelTest(t)
	group := "bound-allow-group"
	bound := createGroupChannel(t, "bound", group, "model-a")
	other := createGroupChannel(t, "other", group, "model-a")

	previousCache := common.MemoryCacheEnabled
	common.MemoryCacheEnabled = false
	t.Cleanup(func() { common.MemoryCacheEnabled = previousCache })

	group_channel_setting.SetGroupChannels(map[string][]int{group: {bound.Id}})

	assert.True(t, ChannelAllowedForGroup(group, "model-a", bound.Id))
	assert.False(t, ChannelAllowedForGroup(group, "model-a", other.Id))
	assert.False(t, ChannelAllowedForGroup(group, "model-missing", bound.Id))
	// Unbound group falls back to abilities membership.
	group_channel_setting.SetGroupChannels(nil)
	assert.True(t, ChannelAllowedForGroup(group, "model-a", other.Id))
}

func TestGetChannelHonorsBindingDBPath(t *testing.T) {
	setupGroupChannelTest(t)
	group := "bound-db-group"
	legacy := createGroupChannel(t, "legacy", group, "shared-model")
	bound := createGroupChannel(t, "bound", group, "shared-model")

	previousCache := common.MemoryCacheEnabled
	common.MemoryCacheEnabled = false
	t.Cleanup(func() { common.MemoryCacheEnabled = previousCache })

	group_channel_setting.SetGroupChannels(map[string][]int{group: {bound.Id}})

	for i := 0; i < 10; i++ {
		channel, err := GetChannel(group, "shared-model", 0, nil)
		require.NoError(t, err)
		require.NotNil(t, channel)
		assert.Equal(t, bound.Id, channel.Id, "bound group must not select unbound channel %d", legacy.Id)
	}

	// A model only served by the unbound channel must disappear.
	onlyLegacy := createGroupChannel(t, "legacy-only", group, "legacy-only-model")
	_ = onlyLegacy
	channel, err := GetChannel(group, "legacy-only-model", 0, nil)
	require.NoError(t, err)
	assert.Nil(t, channel)
}

func TestGetRandomSatisfiedChannelHonorsBindingCachePath(t *testing.T) {
	setupGroupChannelTest(t)
	group := "bound-cache-group"
	legacy := createGroupChannel(t, "legacy", group, "cache-model")
	bound := createGroupChannel(t, "bound", group, "cache-model")

	previousCache := common.MemoryCacheEnabled
	common.MemoryCacheEnabled = true
	t.Cleanup(func() { common.MemoryCacheEnabled = previousCache })

	group_channel_setting.SetGroupChannels(map[string][]int{group: {bound.Id}})
	InitChannelCache()

	for i := 0; i < 20; i++ {
		channel, err := GetRandomSatisfiedChannel(group, "cache-model", 0, nil)
		require.NoError(t, err)
		require.NotNil(t, channel)
		assert.Equal(t, bound.Id, channel.Id, "bound group must not select unbound channel %d", legacy.Id)
	}
}
