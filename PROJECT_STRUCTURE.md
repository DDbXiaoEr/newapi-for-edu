# PROJECT_STRUCTURE.md — Project File Structure

> 此文件记录项目完整文件结构。每次对项目进行文件增删或结构调整后，需同步更新此文件。
> AGENTS 读取此文件即可了解项目结构，无需每次重新探索整个项目。

```
newapi_2_-edu/
├── AGENTS.md                           # Project conventions & guidelines
├── CLAUDE.md                           # Claude-specific instructions
├── PROJECT_STRUCTURE.md                # This file - project file structure
├── NEWAPIREADME.md                     # Main README
├── README.md                           # README (en)
├── README.fr.md                        # README (French)
├── README.ja.md                        # README (Japanese)
├── README.zh_CN.md                     # README (Simplified Chinese)
├── README.zh_TW.md                     # README (Traditional Chinese)
├── LICENSE                             # License file
├── VERSION                             # Version file
├── Makefile                            # Build automation
├── docker-compose.yml                  # Docker Compose config
├── Dockerfile.allinone                 # All-in-one Docker image
├── Dockerfile.backend                  # Backend-only Docker image
├── go.mod                              # Go module definition
├── go.sum                              # Go dependency checksums
├── main.go                             # Application entry point
├── main-backend.go                     # Backend-only entry point
├── new-api.service                     # systemd service file
├── start.sh                            # Startup script
├── sources.list                        # APT sources list
│
├── bin/                                # Migration & test scripts
│   ├── migration_v0.2-v0.3.sql
│   ├── migration_v0.3-v0.4.sql
│   └── time_test.sh
│
├── common/                             # Shared utilities (Rule 1: JSON marshaling)
│   ├── json.go                         # JSON marshal/unmarshal wrappers
│   ├── api_type.go
│   ├── audio.go
│   ├── body_storage.go
│   ├── constants.go
│   ├── copy.go
│   ├── crypto.go
│   ├── custom-event.go
│   ├── database.go
│   ├── disk_cache.go
│   ├── disk_cache_config.go
│   ├── email.go
│   ├── email-outlook-auth.go
│   ├── embed-file-system.go
│   ├── endpoint_defaults.go
│   ├── endpoint_type.go
│   ├── env.go
│   ├── gin.go
│   ├── go-channel.go
│   ├── gopool.go
│   ├── hash.go
│   ├── init.go
│   ├── ip.go
│   ├── jwt.go
│   ├── model.go
│   ├── page_info.go
│   ├── performance_config.go
│   ├── pprof.go
│   ├── pyro.go
│   ├── quota.go
│   ├── rate-limit.go
│   ├── redis.go
│   ├── ssrf_protection.go
│   ├── str.go
│   ├── sys_log.go
│   ├── system_monitor.go
│   ├── system_monitor_unix.go
│   ├── system_monitor_windows.go
│   ├── topup-ratio.go
│   ├── totp.go
│   ├── url_validator.go
│   ├── url_validator_test.go
│   ├── utils.go
│   ├── validate.go
│   ├── verification.go
│   └── limiter/
│       ├── limiter.go
│       └── lua/
│           └── rate_limit.lua
│
├── constant/                           # Constants & enums
│   ├── api_type.go
│   ├── azure.go
│   ├── cache_key.go
│   ├── channel.go
│   ├── context_key.go
│   ├── endpoint_type.go
│   ├── env.go
│   ├── finish_reason.go
│   ├── midjourney.go
│   ├── multi_key_mode.go
│   ├── README.md
│   ├── setup.go
│   ├── task.go
│   └── waffo_pay_method.go
│
├── controller/                         # HTTP request handlers (Gin)
│   ├── billing.go
│   ├── channel.go
│   ├── channel_upstream_update.go
│   ├── channel_upstream_update_test.go
│   ├── channel_affinity_cache.go
│   ├── channel-billing.go
│   ├── channel-test.go
│   ├── checkin.go
│   ├── codex_oauth.go
│   ├── codex_usage.go
│   ├── console_migrate.go
│   ├── custom_oauth.go
│   ├── deployment.go
│   ├── group.go
│   ├── image.go
│   ├── log.go
│   ├── login_ldap_test.go
│   ├── midjourney.go
│   ├── misc.go
│   ├── missing_models.go
│   ├── model.go
│   ├── model_meta.go
│   ├── model_sync.go
│   ├── oauth.go
│   ├── option.go
│   ├── passkey.go
│   ├── payment_webhook_availability.go
│   ├── payment_webhook_availability_test.go
│   ├── performance.go
│   ├── playground.go
│   ├── prefill_group.go
│   ├── pricing.go
│   ├── rankings.go
│   ├── ratio_config.go
│   ├── ratio_sync.go
│   ├── redemption.go
│   ├── relay.go
│   ├── secure_verification.go
│   ├── setup.go
│   ├── subscription.go
│   ├── subscription_payment_creem.go
│   ├── subscription_payment_epay.go
│   ├── subscription_payment_stripe.go
│   ├── swag_video.go
│   ├── task.go
│   ├── telegram.go
│   ├── token.go
│   ├── token_test.go
│   ├── topup.go
│   ├── topup_creem.go
│   ├── topup_epay_guard_test.go
│   ├── topup_stripe.go
│   ├── topup_waffo.go
│   ├── topup_waffo_pancake.go
│   ├── topup_waffo_pancake_test.go
│   ├── twofa.go
│   ├── uptime_kuma.go
│   ├── usedata.go
│   ├── user.go
│   ├── vendor_meta.go
│   ├── video_proxy.go
│   ├── video_proxy_gemini.go
│   └── wechat.go
│
├── dto/                                # Data Transfer Objects (request/response)
│   ├── audio.go
│   ├── channel_settings.go
│   ├── claude.go
│   ├── embedding.go
│   ├── error.go
│   ├── gemini.go
│   ├── gemini_generation_config_test.go
│   ├── gemini_isstream_test.go
│   ├── midjourney.go
│   ├── notify.go
│   ├── openai_compaction.go
│   ├── openai_image.go
│   ├── openai_request.go
│   ├── openai_request_zero_value_test.go
│   ├── openai_response.go
│   ├── openai_responses_compaction_request.go
│   ├── openai_video.go
│   ├── playground.go
│   ├── pricing.go
│   ├── ratio_sync.go
│   ├── realtime.go
│   ├── request_common.go
│   ├── rerank.go
│   ├── sensitive.go
│   ├── suno.go
│   ├── task.go
│   ├── user_settings.go
│   ├── values.go
│   └── video.go
│
├── i18n/                               # Backend internationalization
│   ├── i18n.go
│   ├── keys.go
│   └── locales/
│       ├── en.yaml
│       ├── zh-CN.yaml
│       └── zh-TW.yaml
│
├── logger/                             # Logging
│   ├── logger.go
│   ├── syslog_unix.go
│   └── syslog_windows.go
│
├── middleware/                         # Gin middleware
│   ├── auth.go
│   ├── body_cleanup.go
│   ├── cache.go
│   ├── cors.go
│   ├── disable-cache.go
│   ├── distributor.go
│   ├── email-verification-rate-limit.go
│   ├── gzip.go
│   ├── header_nav.go
│   ├── i18n.go
│   ├── jimeng_adapter.go
│   ├── kling_adapter.go
│   ├── logger.go
│   ├── model-rate-limit.go
│   ├── performance.go
│   ├── rate-limit.go
│   ├── recover.go
│   ├── request-id.go
│   ├── secure_verification.go
│   ├── stats.go
│   ├── turnstile-check.go
│   └── utils.go
│
├── model/                              # Data models & DB access (GORM)
│   ├── main.go                         # DB init, cross-DB compat vars
│   ├── ability.go
│   ├── channel.go
│   ├── channel_cache.go
│   ├── channel_satisfy.go
│   ├── checkin.go
│   ├── custom_oauth_provider.go
│   ├── db_time.go
│   ├── errors.go
│   ├── log.go
│   ├── midjourney.go
│   ├── missing_models.go
│   ├── model_extra.go
│   ├── model_meta.go
│   ├── option.go
│   ├── passkey.go
│   ├── payment_method_guard_test.go
│   ├── prefill_group.go
│   ├── pricing.go
│   ├── pricing_default.go
│   ├── pricing_refresh.go
│   ├── redemption.go
│   ├── setup.go
│   ├── subscription.go
│   ├── task.go
│   ├── task_cas_test.go
│   ├── token.go
│   ├── token_cache.go
│   ├── topup.go
│   ├── twofa.go
│   ├── usedata.go
│   ├── usedata_rankings.go
│   ├── user.go
│   ├── user_cache.go
│   ├── user_oauth_binding.go
│   ├── utils.go
│   └── vendor_meta.go
│
├── oauth/                              # OAuth provider implementations
│   ├── provider.go
│   ├── registry.go
│   ├── types.go
│   ├── cas.go
│   ├── discord.go
│   ├── generic.go
│   ├── github.go
│   ├── linuxdo.go
│   └── oidc.go
│
├── pkg/                                # Internal packages
│   ├── cachex/
│   │   ├── codec.go
│   │   ├── hybrid_cache.go
│   │   └── namespace.go
│   └── ionet/
│       ├── client.go
│       ├── container.go
│       ├── deployment.go
│       ├── hardware.go
│       ├── jsonutil.go
│       └── types.go
│
├── relay/                              # AI API relay/proxy layer
│   ├── relay_adaptor.go                # Channel adapter interface
│   ├── relay_task.go                   # Async task relay
│   ├── audio_handler.go
│   ├── chat_completions_via_responses.go
│   ├── claude_handler.go
│   ├── compatible_handler.go
│   ├── embedding_handler.go
│   ├── gemini_handler.go
│   ├── image_handler.go
│   ├── mjproxy_handler.go
│   ├── param_override_error.go
│   ├── rerank_handler.go
│   ├── responses_handler.go
│   ├── websocket.go
│   ├── channel/
│   │   ├── adapter.go
│   │   ├── api_request.go
│   │   ├── api_request_test.go
│   │   ├── ai360/
│   │   │   └── constants.go
│   │   ├── ali/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   ├── image.go
│   │   │   ├── image_wan.go
│   │   │   ├── rerank.go
│   │   │   └── text.go
│   │   ├── aws/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   ├── relay-aws.go
│   │   │   └── relay_aws_test.go
│   │   ├── baidu/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   └── relay-baidu.go
│   │   ├── baidu_v2/
│   │   │   ├── adaptor.go
│   │   │   └── constants.go
│   │   ├── claude/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   ├── relay-claude.go
│   │   │   ├── relay_claude_test.go
│   │   │   └── message_delta_usage_patch_test.go
│   │   ├── cloudflare/
│   │   │   ├── adaptor.go
│   │   │   ├── constant.go
│   │   │   ├── dto.go
│   │   │   └── relay_cloudflare.go
│   │   ├── codex/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   └── oauth_key.go
│   │   ├── cohere/
│   │   │   ├── adaptor.go
│   │   │   ├── constant.go
│   │   │   ├── dto.go
│   │   │   └── relay-cohere.go
│   │   ├── coze/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   └── relay-coze.go
│   │   ├── deepseek/
│   │   │   ├── adaptor.go
│   │   │   └── constants.go
│   │   ├── dify/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   └── relay-dify.go
│   │   ├── gemini/
│   │   │   ├── adaptor.go
│   │   │   ├── constant.go
│   │   │   ├── relay-gemini.go
│   │   │   ├── relay-gemini-native.go
│   │   │   └── relay_gemini_usage_test.go
│   │   ├── jimeng/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── image.go
│   │   │   └── sign.go
│   │   ├── jina/
│   │   │   ├── adaptor.go
│   │   │   ├── constant.go
│   │   │   └── relay-jina.go
│   │   ├── lingyiwanwu/
│   │   │   └── constrants.go
│   │   ├── minimax/
│   │   │   ├── adaptor.go
│   │   │   ├── adaptor_test.go
│   │   │   ├── constants.go
│   │   │   ├── image.go
│   │   │   ├── relay-minimax.go
│   │   │   └── tts.go
│   │   ├── mistral/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   └── text.go
│   │   ├── mokaai/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   └── relay-mokaai.go
│   │   ├── moonshot/
│   │   │   ├── adaptor.go
│   │   │   └── constants.go
│   │   ├── ollama/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   ├── relay-ollama.go
│   │   │   └── stream.go
│   │   ├── openai/
│   │   │   ├── adaptor.go
│   │   │   ├── audio.go
│   │   │   ├── chat_via_responses.go
│   │   │   ├── constant.go
│   │   │   ├── helper.go
│   │   │   ├── relay-openai.go
│   │   │   ├── relay_responses.go
│   │   │   └── relay_responses_compact.go
│   │   ├── openrouter/
│   │   │   ├── constant.go
│   │   │   └── dto.go
│   │   ├── palm/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   └── relay-palm.go
│   │   ├── perplexity/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   └── relay-perplexity.go
│   │   ├── replicate/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   └── dto.go
│   │   ├── siliconflow/
│   │   │   ├── adaptor.go
│   │   │   ├── constant.go
│   │   │   ├── dto.go
│   │   │   └── relay-siliconflow.go
│   │   ├── submodel/
│   │   │   ├── adaptor.go
│   │   │   └── constants.go
│   │   ├── tencent/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   └── relay-tencent.go
│   │   ├── vertex/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   ├── relay-vertex.go
│   │   │   └── service_account.go
│   │   ├── volcengine/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── protocols.go
│   │   │   └── tts.go
│   │   ├── xai/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   └── text.go
│   │   ├── xinference/
│   │   │   ├── constant.go
│   │   │   └── dto.go
│   │   ├── xunfei/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   └── relay-xunfei.go
│   │   ├── zhipu/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   └── relay-zhipu.go
│   │   ├── zhipu_4v/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   ├── image.go
│   │   │   └── relay-zhipu_v4.go
│   │   └── task/
│   │       ├── taskcommon/
│   │       │   └── helpers.go
│   │       ├── ali/
│   │       │   ├── adaptor.go
│   │       │   └── constants.go
│   │       ├── doubao/
│   │       │   ├── adaptor.go
│   │       │   └── constants.go
│   │       ├── gemini/
│   │       │   ├── adaptor.go
│   │       │   ├── billing.go
│   │       │   ├── dto.go
│   │       │   └── image.go
│   │       ├── hailuo/
│   │       │   ├── adaptor.go
│   │       │   ├── constants.go
│   │       │   └── models.go
│   │       ├── jimeng/
│   │       │   └── adaptor.go
│   │       ├── kling/
│   │       │   └── adaptor.go
│   │       ├── sora/
│   │       │   ├── adaptor.go
│   │       │   └── constants.go
│   │       ├── suno/
│   │       │   ├── adaptor.go
│   │       │   └── models.go
│   │       ├── vertex/
│   │       │   └── adaptor.go
│   │       └── vidu/
│   │           └── adaptor.go
│   ├── common/
│   │   ├── billing.go
│   │   ├── override.go
│   │   ├── override_test.go
│   │   ├── relay_info.go
│   │   ├── relay_info_test.go
│   │   ├── relay_utils.go
│   │   ├── request_conversion.go
│   │   ├── stream_status.go
│   │   └── stream_status_test.go
│   ├── common_handler/
│   │   └── rerank.go
│   ├── constant/
│   │   └── relay_mode.go
│   ├── helper/
│   │   ├── common.go
│   │   ├── model_mapped.go
│   │   ├── price.go
│   │   ├── stream_result.go
│   │   ├── stream_scanner.go
│   │   ├── stream_scanner_test.go
│   │   └── valid_request.go
│   └── reasonmap/
│       └── reasonmap.go
│
├── router/                             # HTTP routing (Gin)
│   ├── main.go                         # Router setup
│   ├── api-router.go                   # API routes
│   ├── dashboard.go                    # Dashboard routes
│   ├── relay-router.go                 # Relay/proxy routes
│   ├── video-router.go                 # Video proxy routes
│   └── web-router.go                   # Web/frontend routes
│
├── service/                            # Business logic layer
│   ├── audio.go
│   ├── billing.go
│   ├── billing_session.go
│   ├── channel.go
│   ├── channel_select.go
│   ├── channel_affinity.go
│   ├── channel_affinity_template_test.go
│   ├── channel_affinity_usage_cache_test.go
│   ├── codex_credential_refresh.go
│   ├── codex_credential_refresh_task.go
│   ├── codex_oauth.go
│   ├── codex_wham_usage.go
│   ├── convert.go
│   ├── download.go
│   ├── epay.go
│   ├── error.go
│   ├── error_test.go
│   ├── file_decoder.go
│   ├── file_service.go
│   ├── funding_source.go
│   ├── group.go
│   ├── http.go
│   ├── http_client.go
│   ├── image.go
│   ├── ldap.go
│   ├── ldap_test.go
│   ├── log_info_generate.go
│   ├── midjourney.go
│   ├── notify-limit.go
│   ├── openai_chat_responses_compat.go
│   ├── openai_chat_responses_mode.go
│   ├── quota.go
│   ├── rankings.go
│   ├── sensitive.go
│   ├── str.go
│   ├── subscription_reset_task.go
│   ├── task.go
│   ├── task_billing.go
│   ├── task_billing_test.go
│   ├── task_polling.go
│   ├── text_quota.go
│   ├── text_quota_test.go
│   ├── token_counter.go
│   ├── token_estimator.go
│   ├── tokenizer.go
│   ├── usage_helpr.go
│   ├── user_notify.go
│   ├── violation_fee.go
│   ├── waffo_pancake.go
│   ├── waffo_pancake_test.go
│   ├── webhook.go
│   ├── openaicompat/
│   │   ├── chat_to_responses.go
│   │   ├── policy.go
│   │   ├── regex.go
│   │   └── responses_to_chat.go
│   └── passkey/
│       ├── service.go
│       ├── session.go
│       └── user.go
│
├── setting/                            # Configuration management
│   ├── auto_group.go
│   ├── chat.go
│   ├── midjourney.go
│   ├── payment_creem.go
│   ├── payment_stripe.go
│   ├── payment_waffo.go
│   ├── payment_waffo_pancake.go
│   ├── rate_limit.go
│   ├── sensitive.go
│   ├── user_usable_group.go
│   ├── config/
│   │   └── config.go
│   ├── console_setting/
│   │   ├── config.go
│   │   └── validation.go
│   ├── model_setting/
│   │   ├── claude.go
│   │   ├── gemini.go
│   │   ├── global.go
│   │   ├── grok.go
│   │   └── qwen.go
│   ├── operation_setting/
│   │   ├── channel_affinity_setting.go
│   │   ├── checkin_setting.go
│   │   ├── general_setting.go
│   │   ├── monitor_setting.go
│   │   ├── operation_setting.go
│   │   ├── payment_setting.go
│   │   ├── payment_setting_old.go
│   │   ├── quota_setting.go
│   │   ├── status_code_ranges.go
│   │   ├── status_code_ranges_test.go
│   │   ├── token_setting.go
│   │   └── tools.go
│   ├── performance_setting/
│   │   └── config.go
│   ├── ratio_setting/
│   │   ├── cache_ratio.go
│   │   ├── compact_suffix.go
│   │   ├── expose_ratio.go
│   │   ├── exposed_cache.go
│   │   ├── group_ratio.go
│   │   └── model_ratio.go
│   ├── reasoning/
│   │   └── suffix.go
│   └── system_setting/
│       ├── cas.go
│       ├── discord.go
│       ├── fetch_setting.go
│       ├── ldap.go
│       ├── legal.go
│       ├── oidc.go
│       ├── passkey.go
│       └── system_setting_old.go
│
├── types/                              # Type definitions
│   ├── channel_error.go
│   ├── error.go
│   ├── file_data.go
│   ├── file_source.go
│   ├── price_data.go
│   ├── relay_format.go
│   ├── request_meta.go
│   ├── rw_map.go
│   └── set.go
│
├── deploy/                             # Deployment configs
│   ├── docker-compose.yml
│   ├── generate-certs.sh
│   └── nginx.conf
│
├── docs/                               # Documentation
│   ├── ionet-client.md
│   ├── translation-glossary.md
│   ├── translation-glossary.fr.md
│   ├── translation-glossary.ru.md
│   ├── channel/
│   │   └── other_setting.md
│   ├── images/
│   │   ├── aionui.png
│   │   ├── aliyun.png
│   │   ├── cherry-studio.png
│   │   ├── io-net.png
│   │   ├── pku.png
│   │   └── ucloud.png
│   ├── installation/
│   │   └── BT.md
│   └── openapi/
│       ├── api.json
│       └── relay.json
│
├── electron/                           # Electron desktop app
│   ├── main.js
│   ├── preload.js
│   ├── build.sh
│   ├── create-tray-icon.js
│   ├── entitlements.mac.plist
│   ├── package.json
│   ├── package-lock.json
│   ├── README.md
│   ├── icon.png
│   ├── tray-icon-windows.png
│   ├── tray-iconTemplate.png
│   └── tray-iconTemplate@2x.png
│
└── web/                                # Frontend (React)
    ├── classic/                        # Classic React frontend (Vite + Semi UI)
    │   ├── package.json
    │   ├── bun.lock
    │   ├── Makefile
    │   ├── index.html
    │   ├── vite.config.js
    │   ├── tailwind.config.js
    │   ├── postcss.config.js
    │   ├── jsconfig.json
    │   ├── vercel.json
    │   ├── i18next.config.js
    │   ├── public/
    │   │   ├── favicon.ico
    │   │   ├── logo.png
    │   │   ├── robots.txt
    │   │   ├── ratio.png
    │   │   ├── cover-4.webp
    │   │   ├── azure_model_name.png
    │   │   ├── pay-apple.png
    │   │   ├── pay-card.png
    │   │   └── pay-google.png
    │   └── src/
    │       ├── App.jsx
    │       ├── index.jsx
    │       ├── index.css
    │       ├── components/
    │       │   ├── auth/
    │       │   ├── common/
    │       │   ├── dashboard/
    │       │   ├── layout/
    │       │   ├── model-deployments/
    │       │   ├── playground/
    │       │   ├── settings/
    │       │   ├── setup/
    │       │   ├── table/
    │       │   └── topup/
    │       ├── constants/
    │       │   ├── channel-affinity-template.constants.js
    │       │   ├── channel.constants.js
    │       │   ├── common.constant.js
    │       │   ├── console.constants.js
    │       │   ├── dashboard.constants.js
    │       │   ├── playground.constants.js
    │       │   ├── redemption.constants.js
    │       │   ├── toast.constants.js
    │       │   ├── user.constants.js
    │       │   └── index.js
    │       ├── contexts/
    │       │   └── PlaygroundContext.jsx
    │       ├── helpers/
    │       │   ├── api.js
    │       │   ├── auth.jsx
    │       │   ├── base64.js
    │       │   ├── boolean.js
    │       │   ├── dashboard.jsx
    │       │   ├── data.js
    │       │   ├── history.js
    │       │   ├── log.js
    │       │   ├── passkey.js
    │       │   ├── quota.js
    │       │   ├── render.jsx
    │       │   ├── secureApiCall.js
    │       │   ├── statusCodeRules.js
    │       │   ├── subscriptionFormat.js
    │       │   ├── token.js
    │       │   ├── utils.jsx
    │       │   └── index.js
    │       ├── i18n/
    │       │   ├── i18n.js
    │       │   └── language.js
    │       ├── locales/
    │       │   ├── en.json
    │       │   ├── fr.json
    │       │   ├── ja.json
    │       │   ├── ru.json
    │       │   ├── vi.json
    │       │   └── zh.json
    │       ├── pages/
    │       │   ├── About/
    │       │   ├── Channel/
    │       │   ├── Chat/
    │       │   ├── Chat2Link/
    │       │   ├── Dashboard/
    │       │   ├── Forbidden/
    │       │   ├── Home/
    │       │   ├── Log/
    │       │   ├── Midjourney/
    │       │   ├── Model/
    │       │   ├── ModelDeployment/
    │       │   ├── NotFound/
    │       │   ├── Playground/
    │       │   ├── Pricing/
    │       │   ├── PrivacyPolicy/
    │       │   ├── Redemption/
    │       │   ├── Setting/
    │       │   ├── Setup/
    │       │   ├── Subscription/
    │       │   ├── Task/
    │       │   ├── Token/
    │       │   ├── TopUp/
    │       │   ├── User/
    │       │   └── UserAgreement/
    │       └── services/
    │           └── secureVerification.js
    │
    └── default/                        # Default/new React frontend (Rsbuild + TanStack)
        ├── AGENTS.md
        ├── package.json
        ├── bun.lock
        ├── Makefile
        ├── index.html
        ├── rsbuild.config.ts
        ├── postcss.config.mjs
        ├── tsconfig.json
        ├── tsconfig.app.json
        ├── tsconfig.node.json
        ├── components.json
        ├── cz.yaml
        ├── knip.config.ts
        ├── netlify.toml
        ├── public/
        │   ├── favicon.ico
        │   ├── logo.png
        │   ├── pay-apple.png
        │   ├── pay-card.png
        │   ├── pay-google.png
        │   ├── waffo-logo-dark.svg
        │   └── waffo-logo-light.svg
        ├── scripts/
        │   ├── add-copyright.mjs
        │   ├── format-with-protected-headers.mjs
        │   └── sync-i18n.mjs
        └── src/
            ├── main.tsx
            ├── env.d.ts
            ├── routeTree.gen.ts
            ├── tanstack-table.d.ts
            ├── assets/
            ├── components/
            │   ├── ai-elements/
            │   ├── data-table/
            │   ├── layout/
            │   ├── ui/
            │   ├── animate-in-view.tsx
            │   ├── auto-skeleton.tsx
            │   ├── coming-soon.tsx
            │   ├── command-menu.tsx
            │   ├── config-drawer.tsx
            │   ├── confirm-dialog.tsx
            │   ├── copy-button.tsx
            │   ├── date-picker.tsx
            │   ├── datetime-picker.tsx
            │   ├── dialog.tsx
            │   ├── drawer-layout.ts
            │   ├── empty-state.tsx
            │   ├── error-state.tsx
            │   ├── group-badge.tsx
            │   ├── json-code-editor.tsx
            │   ├── json-editor.tsx
            │   ├── language-switcher.tsx
            │   ├── learn-more.tsx
            │   ├── loading-state.tsx
            │   ├── long-text.tsx
            │   ├── masked-value-display.tsx
            │   ├── model-group-selector.tsx
            │   ├── multi-select.tsx
            │   ├── navigation-progress.tsx
            │   ├── notification-popover.tsx
            │   ├── page-transition.tsx
            │   ├── password-input.tsx
            │   ├── profile-dropdown.tsx
            │   ├── provider-badge.tsx
            │   ├── react-icon-by-name.tsx
            │   ├── risk-acknowledgement-dialog.tsx
            │   ├── search.tsx
            │   ├── sign-out-dialog.tsx
            │   ├── skip-to-main.tsx
            │   ├── status-badge.tsx
            │   ├── table-id.tsx
            │   ├── tag-input.tsx
            │   ├── theme-quick-switcher.tsx
            │   ├── theme-switch.tsx
            │   ├── truncated-text.tsx
            │   └── turnstile.tsx
            ├── config/
            ├── context/
            ├── features/
            │   ├── about/
            │   ├── auth/
            │   ├── channels/
            │   ├── chat/
            │   ├── dashboard/
            │   ├── errors/
            │   ├── home/
            │   ├── keys/
            │   ├── legal/
            │   ├── models/
            │   ├── performance-metrics/
            │   ├── playground/
            │   ├── pricing/
            │   ├── profile/
            │   ├── rankings/
            │   ├── redemption-codes/
            │   ├── setup/
            │   ├── subscriptions/
            │   ├── system-info/
            │   ├── system-settings/
            │   ├── usage-logs/
            │   ├── users/
            │   └── wallet/
            ├── hooks/
            ├── i18n/
            │   ├── config.ts
            │   ├── languages.ts
            │   ├── static-keys.ts
            │   └── locales/
            │       ├── en.json
            │       ├── fr.json
            │       ├── ja.json
            │       ├── ru.json
            │       ├── vi.json
            │       └── zh.json
            ├── lib/
            │   ├── api.ts
            │   ├── avatar.ts
            │   ├── build-metadata.ts
            │   ├── colors.ts
            │   ├── constants.ts
            │   ├── cookies.ts
            │   ├── copy-to-clipboard.ts
            │   ├── currency.ts
            │   ├── dayjs.ts
            │   ├── dom-utils.ts
            │   ├── format.ts
            │   ├── frontend-cache.ts
            │   ├── handle-server-error.ts
            │   ├── http-status-code-rules.ts
            │   ├── lobe-icon.tsx
            │   ├── motion.ts
            │   ├── nav-modules.ts
            │   ├── oauth.ts
            │   ├── passkey.ts
            │   ├── roles.ts
            │   ├── secure-verification.ts
            │   ├── session-flag.ts
            │   ├── show-submitted-data.tsx
            │   ├── theme-customization.ts
            │   ├── theme-radius.ts
            │   ├── time.ts
            │   ├── use-chart-theme.ts
            │   ├── use-controllable-state.ts
            │   ├── utils.ts
            │   └── vchart.ts
            ├── routes/
            │   ├── __root.tsx
            │   ├── index.tsx
            │   ├── privacy-policy.tsx
            │   ├── user-agreement.tsx
            │   ├── (auth)/
            │   ├── (errors)/
            │   ├── _authenticated/
            │   ├── about/
            │   ├── console/
            │   ├── oauth/
            │   ├── pricing/
            │   ├── rankings/
            │   └── setup/
            ├── stores/
            │   ├── auth-store.ts
            │   ├── notification-store.ts
            │   └── system-config-store.ts
            └── styles/
```

## 各层职责

| 层级 | 目录 | 职责 |
|------|------|------|
| **入口** | `main.go`, `main-backend.go` | 应用启动 |
| **路由** | `router/` | HTTP 路由定义 (API, relay, dashboard, web, video) |
| **控制器** | `controller/` | 请求处理（约60个文件） |
| **服务** | `service/` | 业务逻辑（~55文件）— 认证、计费、配额、分词、支付 |
| **模型** | `model/` | GORM 数据模型（~30文件）— 用户、令牌、渠道、日志、定价、订阅 |
| **中继/代理** | `relay/` | AI 提供商适配器 — 40+ 个提供商，每个在 `relay/channel/<provider>/` |
| **中间件** | `middleware/` | 横切关注点（~16文件）— 认证、限流、CORS、日志、分发 |
| **配置** | `setting/` | 分域配置（模型、运营、系统、比率、性能） |
| **共享** | `common/` | 工具函数（~35文件）— JSON、加密、Redis、限流、JWT、校验 |
| **常量** | `constant/` | 枚举和常量（~14文件） |
| **DTO** | `dto/` | 请求/响应结构体（~25文件） |
| **类型** | `types/` | 核心类型定义（~8文件） |
| **国际化** | `i18n/` (后端) + `web/.../src/i18n/` (前端) | 双层国际化 |
| **前端** | `web/classic/` (Vite+JS), `web/default/` (Rsbuild+TS+TanStack) | 两套独立 React 前端 |

## 40+ AI Provider Adapters

`relay/channel/` 下的提供商适配器:

`ai360/` `ali/` `aws/` `baidu/` `baidu_v2/` `claude/` `cloudflare/` `codex/` `cohere/` `coze/` `deepseek/` `dify/` `gemini/` `jimeng/` `jina/` `lingyiwanwu/` `minimax/` `mistral/` `mokaai/` `moonshot/` `ollama/` `openai/` `openrouter/` `palm/` `perplexity/` `replicate/` `siliconflow/` `submodel/` `tencent/` `vertex/` `volcengine/` `xai/` `xinference/` `xunfei/` `zhipu/` `zhipu_4v/`

异步任务类渠道 `relay/channel/task/`: `ali/` `doubao/` `gemini/` `hailuo/` `jimeng/` `kling/` `sora/` `suno/` `vertex/` `vidu/`
