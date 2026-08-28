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
├── go.mod                              # Go module definition (main module)
├── go.sum                              # Go dependency checksums
├── main.go                             # Application entry point (embeds web/default/dist)
├── main-backend.go                     # Backend-only entry point
├── new-api.service                     # systemd service file
├── start.sh                            # Startup script
├── sources.list                        # APT sources list
│
├── bin
├── common
│   ├── limiter
│   │   ├── lua
│   │   │   └── rate_limit.lua
│   │   └── limiter.go
│   ├── api_type.go
│   ├── audio.go
│   ├── body_storage_test.go
│   ├── body_storage.go
│   ├── constants.go
│   ├── copy.go
│   ├── crypto.go
│   ├── custom-event.go
│   ├── database.go
│   ├── disk_cache_config.go
│   ├── disk_cache.go
│   ├── email_ntlm_auth.go
│   ├── email_test.go
│   ├── email-outlook-auth.go
│   ├── email.go
│   ├── embed-file-system.go
│   ├── endpoint_defaults.go
│   ├── endpoint_type.go
│   ├── env.go
│   ├── gin.go
│   ├── go-channel.go
│   ├── gopool.go
│   ├── hash.go
│   ├── init.go
│   ├── ip.go
│   ├── json_test.go
│   ├── json.go
│   ├── jwt.go
│   ├── model.go
│   ├── node_identity.go
│   ├── page_info.go
│   ├── performance_config.go
│   ├── pprof.go
│   ├── proxy_url.go
│   ├── pyro.go
│   ├── quota_math_test.go
│   ├── quota_math.go
│   ├── quota.go
│   ├── rate-limit.go
│   ├── redis.go
│   ├── request_body_limit.go
│   ├── session_cookie.go
│   ├── ssrf_protection_test.go
│   ├── ssrf_protection.go
│   ├── str.go
│   ├── sys_log.go
│   ├── system_monitor_unix.go
│   ├── system_monitor_windows.go
│   ├── system_monitor.go
│   ├── topup-ratio.go
│   ├── totp.go
│   ├── url_validator_test.go
│   ├── url_validator.go
│   ├── user_session_test.go
│   ├── utils.go
│   ├── validate.go
│   └── verification.go
├── constant
│   ├── api_type.go
│   ├── azure.go
│   ├── cache_key.go
│   ├── channel.go
│   ├── context_key.go
│   ├── endpoint_type.go
│   ├── env.go
│   ├── finish_reason.go
│   ├── midjourney.go
│   ├── multi_key_mode.go
│   ├── setup.go
│   ├── task.go
│   └── waffo_pay_method.go
├── controller
│   ├── audit.go
│   ├── auth_flow_test.go
│   ├── auth_session_test.go
│   ├── auth_session.go
│   ├── authz.go
│   ├── billing.go
│   ├── channel_affinity_cache.go
│   ├── channel_authz_test.go
│   ├── channel_authz.go
│   ├── channel_test_internal_test.go
│   ├── channel_upstream_update_test.go
│   ├── channel_upstream_update.go
│   ├── channel-billing.go
│   ├── channel-test.go
│   ├── channel.go
│   ├── checkin.go
│   ├── codex_oauth.go
│   ├── codex_usage.go
│   ├── console_migrate.go
│   ├── custom_oauth.go
│   ├── deployment.go
│   ├── group.go
│   ├── image.go
│   ├── log.go
│   ├── login_ldap_test.go
│   ├── midjourney.go
│   ├── misc_oidc_test.go
│   ├── misc.go
│   ├── missing_models.go
│   ├── model_list_test.go
│   ├── model_meta.go
│   ├── model_owned_by_test.go
│   ├── model_sync.go
│   ├── model.go
│   ├── oauth.go
│   ├── option_claude_test.go
│   ├── option_gemini_test.go
│   ├── option.go
│   ├── passkey_test.go
│   ├── passkey.go
│   ├── payment_compliance.go
│   ├── payment_webhook_availability_test.go
│   ├── payment_webhook_availability.go
│   ├── perf_metrics.go
│   ├── performance.go
│   ├── playground.go
│   ├── prefill_group.go
│   ├── pricing.go
│   ├── rankings.go
│   ├── ratio_config.go
│   ├── ratio_sync.go
│   ├── redemption.go
│   ├── relay.go
│   ├── return_path_test.go
│   ├── return_path.go
│   ├── secure_verification.go
│   ├── setup.go
│   ├── subscription_payment_creem.go
│   ├── subscription_payment_epay.go
│   ├── subscription_payment_stripe.go
│   ├── subscription_payment_waffo_pancake.go
│   ├── subscription.go
│   ├── swag_video.go
│   ├── system_info.go
│   ├── system_task_handlers.go
│   ├── system_task.go
│   ├── task.go
│   ├── telegram_test.go
│   ├── telegram.go
│   ├── theme_compat_test.go
│   ├── token_auto_groups_test.go
│   ├── token_test.go
│   ├── token.go
│   ├── topup_creem.go
│   ├── topup_epay_guard_test.go
│   ├── topup_quota_limit_test.go
│   ├── topup_stripe.go
│   ├── topup_waffo_pancake_test.go
│   ├── topup_waffo_pancake.go
│   ├── topup_waffo.go
│   ├── topup.go
│   ├── twofa.go
│   ├── uptime_kuma.go
│   ├── usedata_flow_test.go
│   ├── usedata.go
│   ├── user_manage_test.go
│   ├── user.go
│   ├── vendor_meta.go
│   ├── video_proxy_gemini.go
│   ├── video_proxy.go
│   └── wechat.go
├── dto
│   ├── midjourney.go
│   ├── suno.go
│   ├── task.go
│   └── video.go
├── i18n
│   ├── locales
│   │   ├── en.yaml
│   │   ├── zh-CN.yaml
│   │   └── zh-TW.yaml
│   ├── i18n.go
│   └── keys.go
├── logger
│   ├── logger.go
│   ├── syslog_unix.go
│   └── syslog_windows.go
├── logs
│   ├── oneapi-20260422151108.log
│   ├── oneapi-20260608144107.log
│   ├── oneapi-20260608153512.log
│   ├── oneapi-20260613141018.log
│   ├── oneapi-20260614145603.log
│   ├── oneapi-20260615174323.log
│   ├── oneapi-20260616145826.log
│   ├── oneapi-20260617144426.log
│   ├── oneapi-20260617145839.log
│   ├── oneapi-20260618145234.log
│   ├── oneapi-20260618161124.log
│   ├── oneapi-20260618164410.log
│   ├── oneapi-20260618165351.log
│   ├── oneapi-20260620144301.log
│   ├── oneapi-20260621151632.log
│   ├── oneapi-20260621152310.log
│   ├── oneapi-20260621155118.log
│   ├── oneapi-20260623150004.log
│   ├── oneapi-20260623163412.log
│   ├── oneapi-20260626161235.log
│   ├── oneapi-20260626163554.log
│   ├── oneapi-20260626165035.log
│   ├── oneapi-20260626165642.log
│   ├── oneapi-20260626172043.log
│   ├── oneapi-20260626172412.log
│   ├── oneapi-20260626173111.log
│   ├── oneapi-20260627145608.log
│   ├── oneapi-20260627150749.log
│   ├── oneapi-20260627151941.log
│   ├── oneapi-20260627160011.log
│   ├── oneapi-20260627160134.log
│   ├── oneapi-20260627160203.log
│   ├── oneapi-20260627160248.log
│   ├── oneapi-20260627160441.log
│   ├── oneapi-20260627160525.log
│   ├── oneapi-20260627160536.log
│   ├── oneapi-20260627160608.log
│   ├── oneapi-20260629175842.log
│   ├── oneapi-20260629175958.log
│   ├── oneapi-20260629181005.log
│   ├── oneapi-20260701150743.log
│   ├── oneapi-20260701171911.log
│   ├── oneapi-20260701173822.log
│   ├── oneapi-20260703153508.log
│   ├── oneapi-20260703154918.log
│   ├── oneapi-20260703162842.log
│   ├── oneapi-20260703164450.log
│   ├── oneapi-20260703170738.log
│   ├── oneapi-20260705144250.log
│   ├── oneapi-20260705150254.log
│   ├── oneapi-20260706154323.log
│   ├── oneapi-20260706160220.log
│   ├── oneapi-20260718144819.log
│   ├── oneapi-20260718145642.log
│   ├── oneapi-20260816162632.log
│   ├── oneapi-20260816164030.log
│   ├── oneapi-20260816170012.log
│   ├── oneapi-20260828152511.log
│   └── oneapi-20260828152544.log
├── middleware
│   ├── audit.go
│   ├── auth_origin_test.go
│   ├── auth_origin.go
│   ├── auth_test.go
│   ├── auth.go
│   ├── body_cleanup.go
│   ├── cache.go
│   ├── cors.go
│   ├── disable-cache.go
│   ├── distributor.go
│   ├── email-verification-rate-limit.go
│   ├── gzip.go
│   ├── header_nav_test.go
│   ├── header_nav.go
│   ├── i18n.go
│   ├── jimeng_adapter.go
│   ├── kling_adapter.go
│   ├── logger.go
│   ├── model_rate_limit_test.go
│   ├── model-rate-limit.go
│   ├── performance.go
│   ├── rate_limit_test.go
│   ├── rate-limit.go
│   ├── recover.go
│   ├── request_body_limit.go
│   ├── request-id.go
│   ├── secure_verification.go
│   ├── stats.go
│   ├── token_auto_groups_context_test.go
│   ├── trusted_proxies_test.go
│   ├── trusted_proxies.go
│   ├── turnstile-check.go
│   └── utils.go
├── model
│   ├── ability.go
│   ├── auth_flow_test.go
│   ├── auth_flow.go
│   ├── authz_role.go
│   ├── casbin_rule.go
│   ├── channel_cache.go
│   ├── channel_satisfy.go
│   ├── channel_settings_test.go
│   ├── channel_status_test.go
│   ├── channel.go
│   ├── checkin.go
│   ├── clickhouse_log_test.go
│   ├── custom_oauth_provider.go
│   ├── db_time.go
│   ├── errors.go
│   ├── external_identity_claim_test.go
│   ├── external_identity_claim.go
│   ├── frontend_option_migration_test.go
│   ├── frontend_option_migration.go
│   ├── gorm_logger_test.go
│   ├── gorm_logger.go
│   ├── locking_test.go
│   ├── locking.go
│   ├── log_format_test.go
│   ├── log.go
│   ├── main.go
│   ├── midjourney.go
│   ├── missing_models.go
│   ├── model_extra.go
│   ├── model_meta.go
│   ├── model_owner_test.go
│   ├── option_auto_group_test.go
│   ├── option.go
│   ├── passkey.go
│   ├── payment_method_guard_test.go
│   ├── perf_metric.go
│   ├── prefill_group.go
│   ├── pricing_default.go
│   ├── pricing_endpoint_test.go
│   ├── pricing_refresh.go
│   ├── pricing.go
│   ├── quota_reserve_test.go
│   ├── quota_reserve.go
│   ├── redemption_test.go
│   ├── redemption.go
│   ├── setup.go
│   ├── subscription_auth_test.go
│   ├── subscription_reset_test.go
│   ├── subscription.go
│   ├── system_instance.go
│   ├── system_task_test.go
│   ├── system_task.go
│   ├── task_cas_test.go
│   ├── task.go
│   ├── token_auto_groups_cache_test.go
│   ├── token_cache.go
│   ├── token.go
│   ├── topup.go
│   ├── twofa.go
│   ├── usedata_flow_test.go
│   ├── usedata_flow.go
│   ├── usedata_rankings.go
│   ├── usedata.go
│   ├── user_auth_cache.go
│   ├── user_authentication_test.go
│   ├── user_cache_auth_version_test.go
│   ├── user_cache.go
│   ├── user_oauth_binding.go
│   ├── user_pagination_test.go
│   ├── user_session_migration_test.go
│   ├── user_session_test.go
│   ├── user_session.go
│   ├── user_update_test.go
│   ├── user.go
│   ├── utils.go
│   └── vendor_meta.go
├── oauth
│   ├── cas.go
│   ├── discord.go
│   ├── generic.go
│   ├── github.go
│   ├── linuxdo.go
│   ├── oidc_test.go
│   ├── oidc.go
│   ├── provider.go
│   ├── registry.go
│   └── types.go
├── pkg
│   ├── billingexpr
│   │   ├── billingexpr_test.go
│   │   ├── compile.go
│   │   ├── round.go
│   │   ├── run.go
│   │   ├── settle_clamp_test.go
│   │   ├── settle.go
│   │   └── types.go
│   ├── cachex
│   │   ├── codec.go
│   │   ├── hybrid_cache.go
│   │   └── namespace.go
│   ├── ionet
│   │   ├── client.go
│   │   ├── container.go
│   │   ├── deployment.go
│   │   ├── hardware.go
│   │   ├── jsonutil.go
│   │   └── types.go
│   └── perf_metrics
│       ├── flush.go
│       ├── metrics.go
│       └── types.go
├── relay
│   ├── channel
│   │   ├── advancedcustom
│   │   │   ├── adaptor_test.go
│   │   │   └── adaptor.go
│   │   ├── ai360
│   │   │   └── constants.go
│   │   ├── ali
│   │   │   ├── adaptor_test.go
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   ├── image_wan.go
│   │   │   ├── image.go
│   │   │   ├── rerank.go
│   │   │   ├── text_test.go
│   │   │   └── text.go
│   │   ├── aws
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   ├── relay_aws_test.go
│   │   │   └── relay-aws.go
│   │   ├── baidu
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   └── relay-baidu.go
│   │   ├── baidu_v2
│   │   │   ├── adaptor.go
│   │   │   └── constants.go
│   │   ├── claude
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   ├── message_delta_usage_patch_test.go
│   │   │   ├── relay_claude_test.go
│   │   │   ├── relay-claude.go
│   │   │   └── tool_billing_test.go
│   │   ├── cloudflare
│   │   │   ├── adaptor.go
│   │   │   ├── constant.go
│   │   │   ├── dto.go
│   │   │   └── relay_cloudflare.go
│   │   ├── codex
│   │   │   ├── adaptor_test.go
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   └── oauth_key.go
│   │   ├── cohere
│   │   │   ├── adaptor.go
│   │   │   ├── constant.go
│   │   │   ├── dto.go
│   │   │   └── relay-cohere.go
│   │   ├── coze
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   └── relay-coze.go
│   │   ├── deepseek
│   │   │   ├── adaptor.go
│   │   │   └── constants.go
│   │   ├── dify
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   └── relay-dify.go
│   │   ├── gemini
│   │   │   ├── adaptor_responses_test.go
│   │   │   ├── adaptor.go
│   │   │   ├── constant.go
│   │   │   ├── relay_gemini_usage_test.go
│   │   │   ├── relay_responses_test.go
│   │   │   ├── relay_responses.go
│   │   │   ├── relay-gemini-native.go
│   │   │   └── relay-gemini.go
│   │   ├── jimeng
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── image.go
│   │   │   └── sign.go
│   │   ├── jina
│   │   │   ├── adaptor.go
│   │   │   ├── constant.go
│   │   │   └── relay-jina.go
│   │   ├── lingyiwanwu
│   │   │   └── constrants.go
│   │   ├── minimax
│   │   │   ├── adaptor_test.go
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── image.go
│   │   │   ├── relay-minimax.go
│   │   │   └── tts.go
│   │   ├── mistral
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   └── text.go
│   │   ├── mokaai
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   └── relay-mokaai.go
│   │   ├── moonshot
│   │   │   ├── adaptor_test.go
│   │   │   ├── adaptor.go
│   │   │   └── constants.go
│   │   ├── newapi
│   │   │   ├── adaptor.go
│   │   │   └── constants.go
│   │   ├── ollama
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   ├── relay-ollama.go
│   │   │   ├── stream_test.go
│   │   │   └── stream.go
│   │   ├── openai
│   │   │   ├── adaptor.go
│   │   │   ├── audio.go
│   │   │   ├── chat_via_responses_test.go
│   │   │   ├── chat_via_responses.go
│   │   │   ├── constant.go
│   │   │   ├── helper.go
│   │   │   ├── image_edit_test.go
│   │   │   ├── image_stream_test.go
│   │   │   ├── relay_image.go
│   │   │   ├── relay_realtime.go
│   │   │   ├── relay_responses_billing_test.go
│   │   │   ├── relay_responses_compact.go
│   │   │   ├── relay_responses.go
│   │   │   ├── relay-openai.go
│   │   │   ├── responses_via_chat.go
│   │   │   ├── stream_tool_billing_test.go
│   │   │   └── usage.go
│   │   ├── openrouter
│   │   │   ├── constant.go
│   │   │   └── dto.go
│   │   ├── palm
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   └── relay-palm.go
│   │   ├── perplexity
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   └── relay-perplexity.go
│   │   ├── replicate
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   └── dto.go
│   │   ├── siliconflow
│   │   │   ├── adaptor.go
│   │   │   ├── constant.go
│   │   │   ├── dto.go
│   │   │   └── relay-siliconflow.go
│   │   ├── sub2api
│   │   │   ├── adaptor_test.go
│   │   │   ├── adaptor.go
│   │   │   └── constants.go
│   │   ├── submodel
│   │   │   ├── adaptor.go
│   │   │   └── constants.go
│   │   ├── task
│   │   │   ├── ali
│   │   │   ├── doubao
│   │   │   ├── gemini
│   │   │   ├── hailuo
│   │   │   ├── jimeng
│   │   │   ├── kling
│   │   │   ├── sora
│   │   │   ├── suno
│   │   │   ├── taskcommon
│   │   │   ├── vertex
│   │   │   └── vidu
│   │   ├── tencent
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dispatch_test.go
│   │   │   ├── dispatch.go
│   │   │   ├── dto.go
│   │   │   └── relay-tencent.go
│   │   ├── vertex
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   ├── relay-vertex.go
│   │   │   ├── service_account.go
│   │   │   └── url_builder.go
│   │   ├── volcengine
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── protocols.go
│   │   │   └── tts.go
│   │   ├── xai
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   └── text.go
│   │   ├── xinference
│   │   │   ├── constant.go
│   │   │   └── dto.go
│   │   ├── xunfei
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   └── relay-xunfei.go
│   │   ├── zhipu
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   └── relay-zhipu.go
│   │   ├── zhipu_4v
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   ├── image.go
│   │   │   └── relay-zhipu_v4.go
│   │   ├── adapter.go
│   │   ├── api_request_getbody_test.go
│   │   ├── api_request_redirect_test.go
│   │   ├── api_request_test.go
│   │   └── api_request.go
│   ├── common
│   │   ├── billing.go
│   │   ├── outbound_body_test.go
│   │   ├── outbound_body.go
│   │   ├── override_test.go
│   │   ├── override.go
│   │   ├── relay_info_gemini_settings_test.go
│   │   ├── relay_info_test.go
│   │   ├── relay_info.go
│   │   ├── relay_utils_test.go
│   │   ├── relay_utils.go
│   │   ├── request_conversion.go
│   │   ├── stream_status_test.go
│   │   ├── stream_status.go
│   │   ├── tool_usage_test.go
│   │   └── tool_usage.go
│   ├── common_handler
│   │   └── rerank.go
│   ├── constant
│   │   ├── relay_mode_test.go
│   │   └── relay_mode.go
│   ├── helper
│   │   ├── billing_expr_request_test.go
│   │   ├── billing_expr_request.go
│   │   ├── common.go
│   │   ├── max_tokens_bounds_test.go
│   │   ├── model_mapped.go
│   │   ├── openai_image_request_test.go
│   │   ├── price_test.go
│   │   ├── price.go
│   │   ├── stream_result.go
│   │   ├── stream_scanner_test.go
│   │   ├── stream_scanner.go
│   │   └── valid_request.go
│   ├── alpha_search_handler_test.go
│   ├── alpha_search_handler.go
│   ├── audio_handler.go
│   ├── chat_completions_via_responses_test.go
│   ├── chat_completions_via_responses.go
│   ├── claude_handler.go
│   ├── compatible_handler.go
│   ├── embedding_handler.go
│   ├── gemini_handler.go
│   ├── image_handler.go
│   ├── mjproxy_handler.go
│   ├── param_override_error.go
│   ├── relay_adaptor.go
│   ├── relay_task.go
│   ├── rerank_handler.go
│   ├── responses_handler.go
│   └── websocket.go
├── relaykit
│   ├── dto
│   │   ├── alpha_search_request.go
│   │   ├── audio.go
│   │   ├── billing_usage_test.go
│   │   ├── billing_usage.go
│   │   ├── channel_settings_test.go
│   │   ├── channel_settings.go
│   │   ├── claude.go
│   │   ├── embedding.go
│   │   ├── error.go
│   │   ├── gemini_generation_config_test.go
│   │   ├── gemini_isstream_test.go
│   │   ├── gemini_response_test.go
│   │   ├── gemini.go
│   │   ├── notify.go
│   │   ├── openai_compaction.go
│   │   ├── openai_image.go
│   │   ├── openai_request_zero_value_test.go
│   │   ├── openai_request.go
│   │   ├── openai_response.go
│   │   ├── openai_responses_compaction_request.go
│   │   ├── openai_video.go
│   │   ├── playground.go
│   │   ├── pricing.go
│   │   ├── ratio_sync.go
│   │   ├── realtime.go
│   │   ├── request_common.go
│   │   ├── rerank.go
│   │   ├── sensitive.go
│   │   ├── user_settings.go
│   │   └── values.go
│   ├── reasonmap
│   │   └── reasonmap.go
│   ├── relayconvert
│   │   ├── convmeta
│   │   │   ├── format.go
│   │   │   ├── meta_test.go
│   │   │   ├── meta.go
│   │   │   └── options.go
│   │   ├── internal
│   │   │   ├── claude_messages
│   │   │   ├── gemini_chat
│   │   │   ├── jsonutil
│   │   │   ├── media
│   │   │   ├── oai_chat
│   │   │   ├── oai_responses
│   │   │   └── shared
│   │   ├── kitutil
│   │   │   ├── json.go
│   │   │   ├── log_test.go
│   │   │   ├── log.go
│   │   │   ├── mask.go
│   │   │   └── value.go
│   │   ├── reasoning
│   │   │   └── suffix.go
│   │   ├── testdata
│   │   │   └── golden
│   │   ├── boundary_test.go
│   │   ├── claude_default_max_tokens_test.go
│   │   ├── golden_test.go
│   │   ├── media.go
│   │   ├── request_compat.go
│   │   ├── request_registry_test.go
│   │   ├── request_registry.go
│   │   ├── response_compat.go
│   │   ├── response_registry_test.go
│   │   ├── response_registry.go
│   │   ├── safety_settings_test.go
│   │   ├── terminal_stream_test.go
│   │   ├── text_converter_registry_test.go
│   │   └── text_converter_registry.go
│   ├── types
│   │   ├── channel_error.go
│   │   ├── endpoint_type.go
│   │   ├── error.go
│   │   ├── file_data.go
│   │   ├── file_source.go
│   │   ├── relay_format.go
│   │   └── request_meta.go
│   ├── go.mod
│   └── go.sum
├── router
│   ├── api-router.go
│   ├── authz-router.go
│   ├── channel_router_test.go
│   ├── channel-router.go
│   ├── dashboard.go
│   ├── main.go
│   ├── relay_router_test.go
│   ├── relay-router.go
│   ├── retired_frontend_routes_test.go
│   ├── video-router.go
│   └── web-router.go
├── service
│   ├── authz
│   │   ├── adapter.go
│   │   ├── assignment.go
│   │   ├── authz_test.go
│   │   ├── enforcer.go
│   │   ├── override.go
│   │   ├── permission.go
│   │   ├── registry.go
│   │   ├── resolver.go
│   │   ├── resources_channel.go
│   │   ├── role.go
│   │   └── seed.go
│   ├── passkey
│   │   ├── service.go
│   │   ├── session.go
│   │   └── user.go
│   ├── audio.go
│   ├── auth_cleanup.go
│   ├── auth_session_test.go
│   ├── auth_session.go
│   ├── auth_token_test.go
│   ├── auth_token.go
│   ├── billing_session.go
│   ├── billing_usage.go
│   ├── billing.go
│   ├── channel_affinity_template_test.go
│   ├── channel_affinity_usage_cache_test.go
│   ├── channel_affinity.go
│   ├── channel_select_auto_groups_test.go
│   ├── channel_select.go
│   ├── channel.go
│   ├── codex_channel_models.go
│   ├── codex_credential_refresh_task.go
│   ├── codex_credential_refresh.go
│   ├── codex_models.go
│   ├── codex_oauth.go
│   ├── codex_wham_usage.go
│   ├── convert_test.go
│   ├── convert.go
│   ├── download.go
│   ├── epay.go
│   ├── error_test.go
│   ├── error.go
│   ├── file_decoder.go
│   ├── file_service.go
│   ├── funding_source.go
│   ├── group_auto_groups_test.go
│   ├── group.go
│   ├── http_client_transport_test.go
│   ├── http_client.go
│   ├── http_transport_policy.go
│   ├── http_transport_sharded.go
│   ├── http.go
│   ├── image.go
│   ├── ldap_test.go
│   ├── ldap.go
│   ├── log_info_generate.go
│   ├── midjourney.go
│   ├── notify-limit.go
│   ├── openai_chat_responses_compat.go
│   ├── openai_chat_responses_mode.go
│   ├── protected_fetch_client_test.go
│   ├── protected_fetch_client.go
│   ├── quota_saturation_test.go
│   ├── quota.go
│   ├── rankings.go
│   ├── request_converter.go
│   ├── return_path_test.go
│   ├── return_path.go
│   ├── sensitive.go
│   ├── str.go
│   ├── subscription_reset_task.go
│   ├── system_instance.go
│   ├── system_task_test.go
│   ├── system_task.go
│   ├── task_billing_test.go
│   ├── task_billing.go
│   ├── task_polling_test.go
│   ├── task_polling.go
│   ├── task.go
│   ├── text_quota_test.go
│   ├── text_quota.go
│   ├── tiered_settle_test.go
│   ├── tiered_settle.go
│   ├── token_counter.go
│   ├── token_estimator.go
│   ├── tokenizer.go
│   ├── usage_helpr.go
│   ├── user_notify.go
│   ├── violation_fee.go
│   ├── waffo_pancake.go
│   └── webhook.go
├── setting
│   ├── billing_setting
│   │   └── tiered_billing.go
│   ├── config
│   │   ├── config_test.go
│   │   └── config.go
│   ├── console_setting
│   │   ├── config.go
│   │   └── validation.go
│   ├── model_setting
│   │   ├── claude_test.go
│   │   ├── claude.go
│   │   ├── gemini_test.go
│   │   ├── gemini.go
│   │   ├── global.go
│   │   ├── grok.go
│   │   └── qwen.go
│   ├── operation_setting
│   │   ├── channel_affinity_setting.go
│   │   ├── checkin_setting.go
│   │   ├── general_setting.go
│   │   ├── monitor_setting_test.go
│   │   ├── monitor_setting.go
│   │   ├── operation_setting.go
│   │   ├── payment_setting_old.go
│   │   ├── payment_setting.go
│   │   ├── quota_setting.go
│   │   ├── status_code_ranges_test.go
│   │   ├── status_code_ranges.go
│   │   ├── token_setting.go
│   │   ├── tools_price_test.go
│   │   └── tools.go
│   ├── perf_metrics_setting
│   │   └── config.go
│   ├── performance_setting
│   │   └── config.go
│   ├── ratio_setting
│   │   ├── cache_ratio.go
│   │   ├── compact_suffix.go
│   │   ├── expose_ratio.go
│   │   ├── exposed_cache.go
│   │   ├── group_ratio.go
│   │   └── model_ratio.go
│   ├── reasoning
│   │   └── suffix.go
│   ├── system_setting
│   │   ├── cas.go
│   │   ├── discord.go
│   │   ├── fetch_setting.go
│   │   ├── ldap.go
│   │   ├── legal.go
│   │   ├── oidc_test.go
│   │   ├── oidc.go
│   │   ├── passkey.go
│   │   └── system_setting_old.go
│   ├── auto_group_test.go
│   ├── auto_group.go
│   ├── chat.go
│   ├── midjourney.go
│   ├── payment_creem.go
│   ├── payment_stripe.go
│   ├── payment_waffo_pancake.go
│   ├── payment_waffo.go
│   ├── rate_limit.go
│   ├── sensitive.go
│   └── user_usable_group.go
├── types
│   ├── price_data.go
│   ├── rw_map.go
│   └── set.go
├── Dockerfile.allinone
├── Dockerfile.backend
├── exec.log
├── go.mod
├── go.sum
├── main-backend.go
└── main.go
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
| **控制器** | `controller/` | 请求处理 |
| **服务** | `service/` | 业务逻辑 — 认证、计费、配额、分词、支付、中继 |
| **模型** | `model/` | GORM 数据模型 — 用户、令牌、渠道、日志、定价、订阅、会话、AuthFlow |
| **中继/代理** | `relay/` | AI 提供商适配器 — 40+ 个提供商，每个在 `relay/channel/<provider>/` |
| **中继核心库** | `relaykit/` | 独立 Go module — DTO、协议转换层 (`relayconvert`)、中继类型、reasonmap |
| **中间件** | `middleware/` | 横切关注点 — 认证、限流、CORS、日志、分发 |
| **配置** | `setting/` | 分域配置（模型、运营、系统、比率、性能） |
| **共享** | `common/` | 工具函数 — JSON、加密、Redis、限流、JWT、校验 |
| **常量** | `constant/` | 枚举和常量 |
| **DTO** | `dto/` | 异步任务相关 DTO（其余 DTO 在 `relaykit/dto/`） |
| **类型** | `types/` | 基础类型（其余在 `relaykit/types/`） |
| **国际化** | `i18n/` (后端) + `web/.../src/i18n/` (前端) | 双层国际化 |
| **前端** | `web/classic/` (Vite+JS 旧版, 不嵌入), `web/default/` (Rsbuild+TS+TanStack, 由 main.go 嵌入) | 两套独立 React 前端 |

## 40+ AI Provider Adapters

`relay/channel/` 下的提供商适配器:

`advancedcustom/` `ai360/` `ali/` `aws/` `baidu/` `baidu_v2/` `claude/` `cloudflare/` `codex/` `cohere/` `coze/` `deepseek/` `dify/` `gemini/` `jimeng/` `jina/` `lingyiwanwu/` `minimax/` `mistral/` `mokaai/` `moonshot/` `newapi/` `ollama/` `openai/` `openrouter/` `palm/` `perplexity/` `replicate/` `siliconflow/` `sub2api/` `submodel/` `tencent/` `vertex/` `volcengine/` `xai/` `xinference/` `xunfei/` `zhipu/` `zhipu_4v/`

异步任务类渠道 `relay/channel/task/`: `ali/` `doubao/` `gemini/` `hailuo/` `jimeng/` `kling/` `sora/` `suno/` `vertex/` `vidu/`

## 协议转换层 (relaykit/relayconvert)

- `oai_chat/` — Chat Completions <-> Responses / Claude / Gemini 双向转换
- `oai_responses/` — Responses -> Chat / Claude / Gemini 请求与响应转换
- `claude_messages/`, `gemini_chat/` — Claude / Gemini 通道适配转换
- `kitutil/`, `reasoning/`, `shared/` — 日志、推理内容、公共工具
- `convmeta/` — 转换格式元数据
