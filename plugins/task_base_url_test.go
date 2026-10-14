package plugins_test

import (
	"testing"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/pkg/jsplugin"
	builtinplugins "github.com/QuantumNous/new-api/plugins"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// A channel may carry the vendor's OpenAI-compatible base URL for its chat
// models. Native generation APIs are rooted at the service host, so the
// plugins must strip the compatible/versioned suffix before appending their
// own path. This mirrors the production Alibaba 404 regression.
func TestBuiltinVideoPluginsNormalizeCompatibleBaseURL(t *testing.T) {
	for _, tc := range []struct {
		name    string
		key     string
		base    string
		request map[string]any
		ctx     map[string]any
		want    string
	}{
		{
			name:    "doubao strips ark api base",
			key:     "doubao",
			base:    "https://ark.cn-beijing.volces.com/api/v3",
			request: map[string]any{"model": "doubao-seedance-2-0-260128", "prompt": "a cat"},
			want:    "https://ark.cn-beijing.volces.com/api/v3/contents/generations/tasks",
		},
		{
			name:    "google strips openai compatible base",
			key:     "google",
			base:    "https://generativelanguage.googleapis.com/v1beta/openai",
			request: map[string]any{"prompt": "a cat"},
			ctx:     map[string]any{"upstreamModel": "veo-3.0-generate-001"},
			want:    "https://generativelanguage.googleapis.com/v1beta/models/veo-3.0-generate-001:predictLongRunning",
		},
		{
			name:    "hailuo strips minimax v1 base",
			key:     "hailuo",
			base:    "https://api.minimax.chat/v1",
			request: map[string]any{"prompt": "a cat"},
			ctx:     map[string]any{"upstreamModel": "MiniMax-Hailuo-02"},
			want:    "https://api.minimax.chat/v1/video_generation",
		},
		{
			name:    "sora strips openai v1 base",
			key:     "sora",
			base:    "https://api.openai.com/v1",
			request: map[string]any{"prompt": "a cat"},
			ctx:     map[string]any{"upstreamModel": "sora-2"},
			want:    "https://api.openai.com/v1/videos",
		},
		{
			name:    "kling strips relay prefix and version",
			key:     "kling",
			base:    "https://relay.example.com/kling",
			request: map[string]any{"prompt": "a cat"},
			ctx:     map[string]any{"upstreamModel": "kling-v1", "apiKey": "sk-relay"},
			want:    "https://relay.example.com/kling/v1/videos/text2video",
		},
		{
			name:    "vidu strips ent v2 base",
			key:     "vidu",
			base:    "https://api.vidu.cn/ent/v2",
			request: map[string]any{"prompt": "a cat"},
			ctx:     map[string]any{"upstreamModel": "viduq1"},
			want:    "https://api.vidu.cn/ent/v2/text2video",
		},
		{
			name:    "jimeng strips relay prefix",
			key:     "jimeng",
			base:    "https://relay.example.com/jimeng",
			request: map[string]any{"prompt": "a cat"},
			ctx:     map[string]any{"upstreamModel": "jimeng_high_aes_general_v21_L", "apiKey": "sk-relay"},
			want:    "https://relay.example.com/jimeng/?Action=CVSync2AsyncSubmitTask&Version=2022-08-31",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			source, err := builtinplugins.Source(tc.key)
			require.NoError(t, err)
			plugin, err := jsplugin.NewRegistry().RegisterFactory(source, jsplugin.Options{Key: tc.key})
			require.NoError(t, err)

			ctx := map[string]any{
				"requestBody": tc.request,
				"baseUrl":     tc.base,
				"apiKey":      "sk-test",
			}
			for key, value := range tc.ctx {
				ctx[key] = value
			}
			value, callErr := plugin.Engine.Call(t.Context(), "buildSubmitRequest", ctx)
			require.NoError(t, callErr)
			encoded, err := common.Marshal(value)
			require.NoError(t, err)
			var result map[string]any
			require.NoError(t, common.Unmarshal(encoded, &result))
			assert.Equal(t, tc.want, result["url"])
		})
	}
}
