package common_test

import (
	"testing"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/constant"
	"github.com/stretchr/testify/assert"
)

func TestWanEndpointsDistinguishImagesFromVideos(t *testing.T) {
	for _, name := range []string{
		"wan2.7-image-pro", "wan2.7-image", "wan2.6-image", "wan2.6-t2i",
		"wan2.5-t2i-preview", "wan2.2-t2i-flash", "wan2.2-t2i-plus",
		"wanx2.1-t2i-turbo", "wanx2.1-t2i-plus", "wanx2.0-t2i-turbo",
	} {
		t.Run(name, func(t *testing.T) {
			assert.Contains(t, common.GetEndpointTypesByChannelType(constant.ChannelTypeAli, name), constant.EndpointTypeImageGeneration)
		})
	}
	for _, name := range []string{
		"wanx2.1-t2v-plus", "wanx2.1-t2v-turbo", "wanx2.1-i2v-plus", "wanx2.1-i2v-turbo",
	} {
		t.Run(name, func(t *testing.T) {
			assert.NotContains(t, common.GetEndpointTypesByChannelType(constant.ChannelTypeAli, name), constant.EndpointTypeImageGeneration)
		})
	}
}

func TestSeedreamIsImageGenerationModel(t *testing.T) {
	for _, name := range []string{
		"doubao-seedream-4-0-250828",
		"seedream-4-0-250828",
		"doubao-seedream-5-0-pro-260628",
	} {
		t.Run(name, func(t *testing.T) {
			assert.True(t, common.IsImageGenerationModel(name))
			assert.Equal(t, constant.EndpointTypeImageGeneration, common.GetEndpointTypesByChannelType(constant.ChannelTypeVolcEngine, name)[0])
		})
	}
	assert.False(t, common.IsImageGenerationModel("doubao-seedance-2-5-260628"))
}
