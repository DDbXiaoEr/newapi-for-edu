# PROJECT_STRUCTURE.md — Project File Structure

> 此文件记录项目完整文件结构。每次对项目进行文件增删或结构调整后，需同步更新此文件。
> AGENTS 读取此文件即可了解项目结构，无需每次重新探索整个项目。
> 本文件由 `git ls-files` 自动生成（不含被 .gitignore 忽略的构建产物/依赖）。

```
newapi_2_-edu/
├── .dockerignore
├── .gitignore
├── AGENTS.md
├── CLAUDE.md
├── LICENSE
├── Makefile
├── NEWAPIREADME.md
├── PROJECT_NAVIGATION.md
├── PROJECT_STRUCTURE.md
├── README.fr.md
├── README.ja.md
├── README.md
├── README.zh_CN.md
├── README.zh_TW.md
├── VERSION
├── bin/
│   ├── migration_v0.2-v0.3.sql
│   ├── migration_v0.3-v0.4.sql
│   └── time_test.sh
├── common/
│   ├── account_password.go
│   ├── advanced_custom_presets.go
│   ├── api_type.go
│   ├── api_type_task_plugin_test.go
│   ├── audio.go
│   ├── audio_case_test.go
│   ├── body_storage.go
│   ├── body_storage_test.go
│   ├── constants.go
│   ├── copy.go
│   ├── crypto.go
│   ├── custom-event.go
│   ├── database.go
│   ├── disk_cache.go
│   ├── disk_cache_config.go
│   ├── email-outlook-auth.go
│   ├── email.go
│   ├── email_ntlm_auth.go
│   ├── email_test.go
│   ├── embed-file-system.go
│   ├── endpoint_defaults.go
│   ├── endpoint_type.go
│   ├── env.go
│   ├── env_test.go
│   ├── etag.go
│   ├── gin.go
│   ├── go-channel.go
│   ├── gopool.go
│   ├── hash.go
│   ├── init.go
│   ├── ip.go
│   ├── json.go
│   ├── json_test.go
│   ├── jwt.go
│   ├── limiter/
│   │   ├── limiter.go
│   │   └── lua/
│   │       └── rate_limit.lua
│   ├── model.go
│   ├── model_test.go
│   ├── node_identity.go
│   ├── page_info.go
│   ├── password_crypto.go
│   ├── performance_config.go
│   ├── pprof.go
│   ├── proxy_url.go
│   ├── pyro.go
│   ├── quota.go
│   ├── quota_math.go
│   ├── quota_math_test.go
│   ├── rate-limit.go
│   ├── rate-limit_test.go
│   ├── redis.go
│   ├── request_body_limit.go
│   ├── ssrf_protection.go
│   ├── ssrf_protection_test.go
│   ├── str.go
│   ├── sys_log.go
│   ├── system_monitor.go
│   ├── system_monitor_unix.go
│   ├── system_monitor_windows.go
│   ├── topup-ratio.go
│   ├── totp.go
│   ├── trusted_proxies.go
│   ├── url_validator.go
│   ├── url_validator_test.go
│   ├── user_session_test.go
│   ├── utils.go
│   ├── validate.go
│   └── verification.go
├── constant/
│   ├── README.md
│   ├── api_type.go
│   ├── azure.go
│   ├── cache_key.go
│   ├── channel.go
│   ├── channel_test.go
│   ├── context_key.go
│   ├── endpoint_type.go
│   ├── env.go
│   ├── finish_reason.go
│   ├── midjourney.go
│   ├── multi_key_mode.go
│   ├── setup.go
│   ├── task.go
│   ├── task_test.go
│   └── waffo_pay_method.go
├── controller/
│   ├── access_token.go
│   ├── access_token_audit_test.go
│   ├── audit.go
│   ├── auth_flow_test.go
│   ├── auth_session.go
│   ├── auth_session_test.go
│   ├── authz.go
│   ├── billing.go
│   ├── billing_option_test.go
│   ├── channel-billing.go
│   ├── channel-test.go
│   ├── channel.go
│   ├── channel_affinity_cache.go
│   ├── channel_authz.go
│   ├── channel_authz_test.go
│   ├── channel_billing_test.go
│   ├── channel_inference.go
│   ├── channel_inference_test.go
│   ├── channel_multi_key_test.go
│   ├── channel_pin_retry_test.go
│   ├── channel_task_plugin_bind_test.go
│   ├── channel_task_plugin_validation_test.go
│   ├── channel_test_internal_test.go
│   ├── channel_test_request_test.go
│   ├── channel_upstream_update.go
│   ├── channel_upstream_update_test.go
│   ├── checkin.go
│   ├── codex_oauth.go
│   ├── codex_oauth_test.go
│   ├── codex_usage.go
│   ├── console_migrate.go
│   ├── custom_oauth.go
│   ├── deployment.go
│   ├── email_binding.go
│   ├── group.go
│   ├── group_channel.go
│   ├── group_channel_test.go
│   ├── image.go
│   ├── log.go
│   ├── login_ldap_test.go
│   ├── login_verification.go
│   ├── midjourney.go
│   ├── misc.go
│   ├── misc_oidc_test.go
│   ├── missing_models.go
│   ├── model.go
│   ├── model_list_test.go
│   ├── model_management_test.go
│   ├── model_meta.go
│   ├── model_owned_by_test.go
│   ├── model_pricing_config.go
│   ├── model_sync.go
│   ├── oauth.go
│   ├── option.go
│   ├── option_claude_test.go
│   ├── option_gemini_test.go
│   ├── passkey.go
│   ├── passkey_test.go
│   ├── payment_compliance.go
│   ├── payment_webhook_availability.go
│   ├── payment_webhook_availability_test.go
│   ├── perf_metrics.go
│   ├── performance.go
│   ├── playground.go
│   ├── plugin_endpoint_test.go
│   ├── plugin_native_e2e_test.go
│   ├── plugin_protocol.go
│   ├── plugin_protocol_limiter.go
│   ├── plugin_protocol_limiter_test.go
│   ├── plugin_protocol_sdk_test.go
│   ├── plugin_protocol_test.go
│   ├── prefill_group.go
│   ├── pricing.go
│   ├── rankings.go
│   ├── ratio_config.go
│   ├── ratio_sync.go
│   ├── ratio_sync_test.go
│   ├── redemption.go
│   ├── redemption_batch_test.go
│   ├── relay.go
│   ├── relay_count_tokens_test.go
│   ├── relay_error_log_test.go
│   ├── relay_task_plugin_test.go
│   ├── request_policy.go
│   ├── responses_websocket.go
│   ├── responses_websocket_test.go
│   ├── return_path.go
│   ├── return_path_test.go
│   ├── revalidated_response.go
│   ├── secure_verification.go
│   ├── security_account_test.go
│   ├── security_enrollment_test.go
│   ├── setup.go
│   ├── subscription.go
│   ├── subscription_payment_creem.go
│   ├── subscription_payment_epay.go
│   ├── subscription_payment_stripe.go
│   ├── subscription_payment_waffo_pancake.go
│   ├── swag_video.go
│   ├── system_info.go
│   ├── system_task.go
│   ├── system_task_handlers.go
│   ├── system_task_test.go
│   ├── task.go
│   ├── task_generic_test.go
│   ├── task_log_view_test.go
│   ├── task_plugin.go
│   ├── task_plugin_debug.go
│   ├── task_plugin_debug_test.go
│   ├── task_plugin_test.go
│   ├── telegram.go
│   ├── telegram_test.go
│   ├── theme_compat_test.go
│   ├── token.go
│   ├── token_auto_groups_test.go
│   ├── token_test.go
│   ├── topup.go
│   ├── topup_creem.go
│   ├── topup_epay_guard_test.go
│   ├── topup_quota_limit_test.go
│   ├── topup_stripe.go
│   ├── topup_waffo.go
│   ├── topup_waffo_pancake.go
│   ├── topup_waffo_pancake_test.go
│   ├── twofa.go
│   ├── uptime_kuma.go
│   ├── usedata.go
│   ├── usedata_flow_test.go
│   ├── user.go
│   ├── user_manage_test.go
│   ├── user_quota.go
│   ├── vendor_meta.go
│   ├── video_proxy.go
│   └── wechat.go
├── deploy/
│   ├── docker-compose.yml
│   ├── generate-certs.sh
│   └── nginx.conf
├── docker-compose.yml
├── docker-compose/
│   ├── db.yml
│   ├── docker-compose.yml
│   └── ldap/
│       ├── entrypoint.sh
│       └── init-data.ldif
├── docker/
│   ├── Dockerfile.allinone
│   └── Dockerfile.backend
├── docs/
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
│   ├── ionet-client.md
│   ├── openapi/
│   │   ├── api.json
│   │   └── relay.json
│   ├── plugin-api/
│   │   ├── README.md
│   │   ├── v1.d.ts
│   │   ├── v1.md
│   │   └── v1.schema.json
│   ├── translation-glossary.fr.md
│   ├── translation-glossary.md
│   └── translation-glossary.ru.md
├── dto/
│   ├── channel_constraints.go
│   ├── channel_constraints_test.go
│   ├── midjourney.go
│   ├── plugin_protocol.go
│   ├── suno.go
│   ├── task.go
│   ├── task_plugin.go
│   └── video.go
├── e2e/
│   └── doc_parse_test.go
├── electron/
│   ├── README.md
│   ├── build.sh
│   ├── create-tray-icon.js
│   ├── entitlements.mac.plist
│   ├── icon.png
│   ├── main.js
│   ├── package-lock.json
│   ├── package.json
│   ├── preload.js
│   ├── tray-icon-windows.png
│   ├── tray-iconTemplate.png
│   └── tray-iconTemplate@2x.png
├── exec.log
├── go.mod
├── go.sum
├── i18n/
│   ├── i18n.go
│   ├── keys.go
│   └── locales/
│       ├── en.yaml
│       ├── zh-CN.yaml
│       └── zh-TW.yaml
├── kubernetes/
│   ├── configmap.yaml
│   ├── deployment.yaml
│   ├── hpa.yaml
│   ├── external-clickhouse.yaml
│   ├── external-ldap.yaml
│   ├── external-postgres.yaml
│   ├── external-redis.yaml
│   ├── kustomization.yaml
│   ├── namespace.yaml
│   ├── secret.yaml
│   └── service.yaml
├── logger/
│   ├── logger.go
│   ├── syslog_unix.go
│   └── syslog_windows.go
├── main-backend.go
├── main.go
├── middleware/
│   ├── audit.go
│   ├── auth.go
│   ├── auth_test.go
│   ├── body_cleanup.go
│   ├── cache.go
│   ├── cors.go
│   ├── disable-cache.go
│   ├── distributor.go
│   ├── distributor_test.go
│   ├── email-verification-rate-limit.go
│   ├── gzip.go
│   ├── header_nav.go
│   ├── header_nav_test.go
│   ├── i18n.go
│   ├── logger.go
│   ├── model-rate-limit.go
│   ├── model_rate_limit_test.go
│   ├── performance.go
│   ├── rate-limit.go
│   ├── rate_limit_test.go
│   ├── recover.go
│   ├── request-id.go
│   ├── request_body_limit.go
│   ├── secure_verification.go
│   ├── stats.go
│   ├── task_artifact_access.go
│   ├── task_artifact_access_test.go
│   ├── task_plugin.go
│   ├── task_plugin_model.go
│   ├── task_plugin_origin_task_test.go
│   ├── task_plugin_test.go
│   ├── token_auto_groups_context_test.go
│   ├── trusted_proxies.go
│   ├── trusted_proxies_test.go
│   ├── turnstile-check.go
│   └── utils.go
├── model/
│   ├── ability.go
│   ├── account_security.go
│   ├── audit_log.go
│   ├── audit_other.go
│   ├── auth_flow.go
│   ├── auth_flow_test.go
│   ├── authz_role.go
│   ├── casbin_rule.go
│   ├── channel.go
│   ├── channel_cache.go
│   ├── channel_constraint.go
│   ├── channel_constraint_test.go
│   ├── channel_satisfy.go
│   ├── channel_settings_test.go
│   ├── channel_status_test.go
│   ├── checkin.go
│   ├── clickhouse_log_test.go
│   ├── custom_oauth_provider.go
│   ├── db_time.go
│   ├── email_binding.go
│   ├── errors.go
│   ├── external_identity_claim.go
│   ├── external_identity_claim_test.go
│   ├── frontend_option_migration.go
│   ├── frontend_option_migration_test.go
│   ├── gorm_logger.go
│   ├── gorm_logger_test.go
│   ├── group_channel.go
│   ├── group_channel_test.go
│   ├── json_column_test.go
│   ├── legacy_dalle_pricing.go
│   ├── locking.go
│   ├── locking_test.go
│   ├── log.go
│   ├── log_format_test.go
│   ├── log_other.go
│   ├── log_other_test.go
│   ├── login_verification.go
│   ├── main.go
│   ├── midjourney.go
│   ├── migration_dialector.go
│   ├── migration_dialector_test.go
│   ├── missing_models.go
│   ├── model_extra.go
│   ├── model_meta.go
│   ├── model_metadata_sync.go
│   ├── model_owner_test.go
│   ├── model_pricing_config.go
│   ├── model_pricing_conversion.go
│   ├── option.go
│   ├── option_auto_group_test.go
│   ├── option_primary_key_migration.go
│   ├── option_task_plugin_test.go
│   ├── passkey.go
│   ├── passkey_option.go
│   ├── password_crypto.go
│   ├── payment_method_guard_test.go
│   ├── perf_metric.go
│   ├── prefill_group.go
│   ├── prefill_group_migration.go
│   ├── prefill_group_migration_test.go
│   ├── pricing.go
│   ├── pricing_default.go
│   ├── pricing_endpoint_test.go
│   ├── pricing_refresh.go
│   ├── pricing_usage_schema_test.go
│   ├── quota_reserve.go
│   ├── quota_reserve_test.go
│   ├── redemption.go
│   ├── redemption_test.go
│   ├── request_policy.go
│   ├── request_policy_test.go
│   ├── setup.go
│   ├── subscription.go
│   ├── subscription_auth_test.go
│   ├── subscription_reset_test.go
│   ├── system_instance.go
│   ├── system_task.go
│   ├── system_task_test.go
│   ├── task.go
│   ├── task_cas_test.go
│   ├── task_model_alias.go
│   ├── task_openai_video_test.go
│   ├── task_plugin.go
│   ├── task_plugin_channel_select_test.go
│   ├── task_plugin_test.go
│   ├── token.go
│   ├── token_auto_groups_cache_test.go
│   ├── token_cache.go
│   ├── token_migration.go
│   ├── token_migration_test.go
│   ├── topup.go
│   ├── twofa.go
│   ├── twofa_enrollment.go
│   ├── usedata.go
│   ├── usedata_flow.go
│   ├── usedata_flow_test.go
│   ├── usedata_rankings.go
│   ├── user.go
│   ├── user_auth_cache.go
│   ├── user_authentication_test.go
│   ├── user_cache.go
│   ├── user_cache_auth_version_test.go
│   ├── user_oauth_binding.go
│   ├── user_pagination_test.go
│   ├── user_quota_adjustment.go
│   ├── user_session.go
│   ├── user_session_migration_test.go
│   ├── user_session_test.go
│   ├── user_update_test.go
│   ├── utils.go
│   ├── vendor_management.go
│   └── vendor_meta.go
├── new-api.service
├── oauth/
│   ├── cas.go
│   ├── discord.go
│   ├── generic.go
│   ├── github.go
│   ├── linuxdo.go
│   ├── oidc.go
│   ├── oidc_test.go
│   ├── provider.go
│   ├── registry.go
│   ├── telegram.go
│   └── types.go
├── pkg/
│   ├── billingexpr/
│   │   ├── billingexpr_test.go
│   │   ├── compile.go
│   │   ├── compile_usage_test.go
│   │   ├── fixed.go
│   │   ├── round.go
│   │   ├── run.go
│   │   ├── settle.go
│   │   ├── settle_clamp_test.go
│   │   ├── task_usage_test.go
│   │   ├── testdata/
│   │   │   └── frontend_simulation.json
│   │   └── types.go
│   ├── cachex/
│   │   ├── codec.go
│   │   ├── hybrid_cache.go
│   │   └── namespace.go
│   ├── ionet/
│   │   ├── client.go
│   │   ├── container.go
│   │   ├── deployment.go
│   │   ├── hardware.go
│   │   ├── jsonutil.go
│   │   └── types.go
│   ├── jsplugin/
│   │   ├── cli.go
│   │   ├── cli_test.go
│   │   ├── engine.go
│   │   ├── engine_test.go
│   │   ├── fixture.go
│   │   ├── fixture_test.go
│   │   ├── icon.go
│   │   ├── json_state.go
│   │   ├── model_fold.go
│   │   ├── protocol_supports_test.go
│   │   ├── registry.go
│   │   ├── registry_disabled_factory_test.go
│   │   ├── registry_master_enabled_test.go
│   │   ├── registry_test.go
│   │   ├── request.go
│   │   ├── routing.go
│   │   ├── routing_test.go
│   │   └── utils.go
│   ├── perf_metrics/
│   │   ├── flush.go
│   │   ├── metrics.go
│   │   ├── metrics_test.go
│   │   ├── outcome.go
│   │   └── types.go
│   ├── testdb/
│   │   └── testdb.go
│   └── wsmanager/
│       ├── wsmanager.go
│       └── wsmanager_test.go
├── plugins/
│   ├── .oxfmtrc.json
│   ├── .oxlintrc.json
│   ├── alibaba_responses_test.go
│   ├── alibaba_wan3_test.go
│   ├── builtin_plugins_test.go
│   ├── doubao_responses_test.go
│   ├── embed.go
│   ├── google_responses_test.go
│   ├── hailuo_responses_test.go
│   ├── jimeng_responses_test.go
│   ├── kling_responses_test.go
│   ├── sora_responses_test.go
│   ├── sunoapi_responses_test.go
│   ├── tasks/
│   │   ├── alibaba/
│   │   │   └── plugin.js
│   │   ├── doubao/
│   │   │   └── plugin.js
│   │   ├── google/
│   │   │   └── plugin.js
│   │   ├── hailuo/
│   │   │   └── plugin.js
│   │   ├── jimeng/
│   │   │   └── plugin.js
│   │   ├── kling/
│   │   │   └── plugin.js
│   │   ├── sora/
│   │   │   └── plugin.js
│   │   ├── sunoapi/
│   │   │   └── plugin.js
│   │   ├── vertex-ai/
│   │   │   └── plugin.js
│   │   └── vidu/
│   │       └── plugin.js
│   ├── veo_poll_test.go
│   ├── vertex_ai_responses_test.go
│   ├── video_responses_test_helpers_test.go
│   └── vidu_responses_test.go
├── relay/
│   ├── alpha_search_handler.go
│   ├── alpha_search_handler_test.go
│   ├── audio_handler.go
│   ├── channel/
│   │   ├── adapter.go
│   │   ├── advancedcustom/
│   │   │   ├── adaptor.go
│   │   │   ├── adaptor_test.go
│   │   │   ├── presets_test.go
│   │   │   └── rerank.go
│   │   ├── ai360/
│   │   │   └── constants.go
│   │   ├── ali/
│   │   │   ├── adaptor.go
│   │   │   ├── adaptor_test.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   ├── image.go
│   │   │   ├── image_wan.go
│   │   │   ├── rerank.go
│   │   │   ├── text.go
│   │   │   └── text_test.go
│   │   ├── api_request.go
│   │   ├── api_request_getbody_test.go
│   │   ├── api_request_redirect_test.go
│   │   ├── api_request_test.go
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
│   │   │   ├── adaptor_test.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   ├── message_delta_usage_patch_test.go
│   │   │   ├── relay-claude.go
│   │   │   ├── relay_claude_test.go
│   │   │   ├── relay_responses.go
│   │   │   └── tool_billing_test.go
│   │   ├── cloudflare/
│   │   │   ├── adaptor.go
│   │   │   ├── constant.go
│   │   │   ├── dto.go
│   │   │   └── relay_cloudflare.go
│   │   ├── codex/
│   │   │   ├── adaptor.go
│   │   │   ├── adaptor_test.go
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
│   │   │   ├── adaptor_responses_test.go
│   │   │   ├── constant.go
│   │   │   ├── relay-gemini-native.go
│   │   │   ├── relay-gemini.go
│   │   │   ├── relay_gemini_usage_test.go
│   │   │   ├── relay_responses.go
│   │   │   └── relay_responses_test.go
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
│   │   │   ├── adaptor_test.go
│   │   │   └── constants.go
│   │   ├── newapi/
│   │   │   ├── adaptor.go
│   │   │   └── constants.go
│   │   ├── ollama/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   ├── relay-ollama.go
│   │   │   ├── stream.go
│   │   │   └── stream_test.go
│   │   ├── openai/
│   │   │   ├── adaptor.go
│   │   │   ├── audio.go
│   │   │   ├── chat_via_responses.go
│   │   │   ├── chat_via_responses_test.go
│   │   │   ├── constant.go
│   │   │   ├── helper.go
│   │   │   ├── image_edit_test.go
│   │   │   ├── image_stream_test.go
│   │   │   ├── relay-openai.go
│   │   │   ├── relay_image.go
│   │   │   ├── relay_realtime.go
│   │   │   ├── relay_responses.go
│   │   │   ├── relay_responses_billing_test.go
│   │   │   ├── relay_responses_compact.go
│   │   │   ├── responses_via_chat.go
│   │   │   ├── stream_tool_billing_test.go
│   │   │   └── usage.go
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
│   │   ├── sub2api/
│   │   │   ├── adaptor.go
│   │   │   ├── adaptor_test.go
│   │   │   └── constants.go
│   │   ├── submodel/
│   │   │   ├── adaptor.go
│   │   │   └── constants.go
│   │   ├── task/
│   │   │   ├── jsplugin/
│   │   │   │   ├── adaptor.go
│   │   │   │   ├── adaptor_test.go
│   │   │   │   ├── auth.go
│   │   │   │   ├── auth_test.go
│   │   │   │   ├── performance_test.go
│   │   │   │   └── submit_stream.go
│   │   │   └── taskcommon/
│   │   │       └── helpers.go
│   │   ├── tencent/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dispatch.go
│   │   │   ├── dispatch_test.go
│   │   │   ├── dto.go
│   │   │   └── relay-tencent.go
│   │   ├── vertex/
│   │   │   ├── adaptor.go
│   │   │   ├── constants.go
│   │   │   ├── dto.go
│   │   │   ├── relay-vertex.go
│   │   │   ├── service_account.go
│   │   │   └── url_builder.go
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
│   │   └── zhipu_4v/
│   │       ├── adaptor.go
│   │       ├── constants.go
│   │       ├── dto.go
│   │       ├── image.go
│   │       └── relay-zhipu_v4.go
│   ├── chat_completions_via_responses.go
│   ├── chat_completions_via_responses_test.go
│   ├── claude_handler.go
│   ├── common/
│   │   ├── billing.go
│   │   ├── conversion_diagnostics.go
│   │   ├── outbound_body.go
│   │   ├── outbound_body_test.go
│   │   ├── override.go
│   │   ├── override_test.go
│   │   ├── relay_info.go
│   │   ├── relay_info_gemini_settings_test.go
│   │   ├── relay_info_test.go
│   │   ├── relay_utils.go
│   │   ├── relay_utils_test.go
│   │   ├── request_conversion.go
│   │   ├── response_model.go
│   │   ├── stream_status.go
│   │   ├── stream_status_test.go
│   │   ├── tool_usage.go
│   │   └── tool_usage_test.go
│   ├── common_handler/
│   │   └── rerank.go
│   ├── compatible_handler.go
│   ├── constant/
│   │   ├── relay_mode.go
│   │   └── relay_mode_test.go
│   ├── convert_request_error.go
│   ├── convert_request_error_test.go
│   ├── embedding_handler.go
│   ├── gemini_handler.go
│   ├── helper/
│   │   ├── billing_expr_request.go
│   │   ├── billing_expr_request_test.go
│   │   ├── common.go
│   │   ├── max_tokens_bounds_test.go
│   │   ├── model_mapped.go
│   │   ├── model_modifier.go
│   │   ├── openai_image_request_test.go
│   │   ├── price.go
│   │   ├── price_test.go
│   │   ├── reasoning_suffix.go
│   │   ├── reasoning_suffix_test.go
│   │   ├── stream_result.go
│   │   ├── stream_scanner.go
│   │   ├── stream_scanner_test.go
│   │   └── valid_request.go
│   ├── image_handler.go
│   ├── mjproxy_handler.go
│   ├── param_override_error.go
│   ├── plugin_protocol.go
│   ├── plugin_protocol_test.go
│   ├── relay_adaptor.go
│   ├── relay_adaptor_jsplugin_test.go
│   ├── relay_task.go
│   ├── relay_task_test.go
│   ├── request_billing.go
│   ├── request_clone_test.go
│   ├── rerank_handler.go
│   ├── response_model_test.go
│   ├── responses_handler.go
│   ├── responses_request.go
│   ├── responses_websocket.go
│   ├── responses_websocket_test.go
│   ├── task_platform_error_test.go
│   ├── task_platform_test.go
│   └── websocket.go
├── relaykit/
│   ├── README.md
│   ├── dto/
│   │   ├── alpha_search_request.go
│   │   ├── audio.go
│   │   ├── billing_usage.go
│   │   ├── billing_usage_test.go
│   │   ├── channel_settings.go
│   │   ├── channel_settings_test.go
│   │   ├── claude.go
│   │   ├── embedding.go
│   │   ├── error.go
│   │   ├── gemini.go
│   │   ├── gemini_generation_config_test.go
│   │   ├── gemini_isstream_test.go
│   │   ├── gemini_response_test.go
│   │   ├── legacy_dalle_image.go
│   │   ├── notify.go
│   │   ├── openai_compaction.go
│   │   ├── openai_image.go
│   │   ├── openai_request.go
│   │   ├── openai_request_zero_value_test.go
│   │   ├── openai_response.go
│   │   ├── openai_responses_compaction_request.go
│   │   ├── openai_video.go
│   │   ├── playground.go
│   │   ├── pricing.go
│   │   ├── ratio_sync.go
│   │   ├── realtime.go
│   │   ├── reasoning_state.go
│   │   ├── request_common.go
│   │   ├── rerank.go
│   │   ├── sensitive.go
│   │   ├── usage_merge.go
│   │   ├── usage_merge_test.go
│   │   ├── user_settings.go
│   │   ├── values.go
│   │   └── values_test.go
│   ├── go.mod
│   ├── go.sum
│   ├── reasonmap/
│   │   └── reasonmap.go
│   ├── relayconvert/
│   │   ├── boundary_test.go
│   │   ├── claude_default_max_tokens_test.go
│   │   ├── convmeta/
│   │   │   ├── format.go
│   │   │   ├── meta.go
│   │   │   ├── meta_test.go
│   │   │   └── options.go
│   │   ├── golden_test.go
│   │   ├── internal/
│   │   │   ├── claude_messages/
│   │   │   │   ├── citations.go
│   │   │   │   ├── stream_billing_usage_test.go
│   │   │   │   ├── to_oai_chat_req.go
│   │   │   │   ├── to_oai_chat_resp.go
│   │   │   │   ├── to_oai_responses_hosted_stream.go
│   │   │   │   └── to_oai_responses_req.go
│   │   │   ├── convdiag/
│   │   │   │   └── collector.go
│   │   │   ├── gemini_chat/
│   │   │   │   ├── grounding.go
│   │   │   │   ├── to_oai_chat_req.go
│   │   │   │   ├── to_oai_chat_resp.go
│   │   │   │   └── to_oai_responses_hosted_stream.go
│   │   │   ├── jsonutil/
│   │   │   │   └── stringify.go
│   │   │   ├── media/
│   │   │   │   └── media.go
│   │   │   ├── oai_chat/
│   │   │   │   ├── citations.go
│   │   │   │   ├── to_claude_messages_req.go
│   │   │   │   ├── to_claude_messages_req_test.go
│   │   │   │   ├── to_claude_messages_resp.go
│   │   │   │   ├── to_claude_messages_resp_test.go
│   │   │   │   ├── to_gemini_chat_req.go
│   │   │   │   ├── to_gemini_chat_resp.go
│   │   │   │   ├── to_gemini_chat_resp_test.go
│   │   │   │   ├── to_oai_responses_req.go
│   │   │   │   ├── to_oai_responses_req_test.go
│   │   │   │   ├── to_oai_responses_resp.go
│   │   │   │   ├── to_oai_responses_resp_test.go
│   │   │   │   └── to_oai_responses_stream_resp.go
│   │   │   ├── oai_responses/
│   │   │   │   ├── req_helpers.go
│   │   │   │   ├── to_claude_messages_req.go
│   │   │   │   ├── to_claude_messages_resp.go
│   │   │   │   ├── to_claude_messages_stream_resp.go
│   │   │   │   ├── to_claude_messages_stream_resp_test.go
│   │   │   │   ├── to_gemini_chat_req.go
│   │   │   │   ├── to_gemini_chat_req_preprocess.go
│   │   │   │   ├── to_oai_chat_req.go
│   │   │   │   ├── to_oai_chat_req_test.go
│   │   │   │   ├── to_oai_chat_resp.go
│   │   │   │   ├── to_oai_chat_resp_test.go
│   │   │   │   └── to_oai_chat_stream_resp.go
│   │   │   ├── shared/
│   │   │   │   ├── claude/
│   │   │   │   │   ├── cache.go
│   │   │   │   │   ├── errors.go
│   │   │   │   │   ├── reasoning.go
│   │   │   │   │   ├── schema.go
│   │   │   │   │   ├── tool_choice.go
│   │   │   │   │   └── usage.go
│   │   │   │   └── gemini/
│   │   │   │       ├── request.go
│   │   │   │       └── schema.go
│   │   │   └── toolconv/
│   │   │       ├── decode.go
│   │   │       ├── encode.go
│   │   │       ├── hosted_values.go
│   │   │       ├── model.go
│   │   │       ├── policy_test.go
│   │   │       ├── response.go
│   │   │       └── response_artifacts.go
│   │   ├── kitutil/
│   │   │   ├── json.go
│   │   │   ├── json_test.go
│   │   │   ├── log.go
│   │   │   ├── log_test.go
│   │   │   ├── mask.go
│   │   │   └── value.go
│   │   ├── media.go
│   │   ├── reasoning/
│   │   │   ├── claude.go
│   │   │   ├── gemini.go
│   │   │   ├── gemini_test.go
│   │   │   ├── intent.go
│   │   │   ├── intent_test.go
│   │   │   ├── suffix.go
│   │   │   └── suffix_test.go
│   │   ├── request_compat.go
│   │   ├── request_registry.go
│   │   ├── request_registry_test.go
│   │   ├── response_compat.go
│   │   ├── response_registry.go
│   │   ├── response_registry_test.go
│   │   ├── safety_settings_test.go
│   │   ├── terminal_stream_test.go
│   │   ├── testdata/
│   │   │   └── golden/
│   │   │       ├── request/
│   │   │       │   ├── openai_responses_to_claude.golden.json
│   │   │       │   ├── openai_responses_to_openai.golden.json
│   │   │       │   ├── openai_to_claude.golden.json
│   │   │       │   └── openai_to_openai_responses.golden.json
│   │   │       ├── response/
│   │   │       │   ├── claude_to_openai.golden.json
│   │   │       │   ├── claude_to_openai_responses.golden.json
│   │   │       │   ├── gemini_to_claude.golden.json
│   │   │       │   ├── gemini_to_openai.golden.json
│   │   │       │   ├── gemini_to_openai_responses.golden.json
│   │   │       │   ├── openai_responses_to_claude.golden.json
│   │   │       │   ├── openai_responses_to_openai.golden.json
│   │   │       │   └── openai_to_openai_responses.golden.json
│   │   │       └── stream/
│   │   │           ├── claude_to_gemini.golden.json
│   │   │           ├── claude_to_openai.golden.json
│   │   │           ├── claude_to_openai_responses.golden.json
│   │   │           ├── gemini_to_claude.golden.json
│   │   │           ├── gemini_to_openai.golden.json
│   │   │           ├── gemini_to_openai_responses.golden.json
│   │   │           ├── openai_responses_to_claude.golden.json
│   │   │           ├── openai_responses_to_gemini.golden.json
│   │   │           ├── openai_responses_to_openai.golden.json
│   │   │           ├── openai_to_claude.golden.json
│   │   │           ├── openai_to_gemini.golden.json
│   │   │           └── openai_to_openai_responses.golden.json
│   │   ├── text_converter_registry.go
│   │   ├── text_converter_registry_test.go
│   │   └── tool_loss_policy_test.go
│   └── types/
│       ├── channel_error.go
│       ├── conversion.go
│       ├── endpoint_type.go
│       ├── error.go
│       ├── file_data.go
│       ├── file_source.go
│       ├── relay_format.go
│       └── request_meta.go
├── router/
│   ├── api-router.go
│   ├── authz-router.go
│   ├── channel-router.go
│   ├── channel_router_test.go
│   ├── dashboard.go
│   ├── main.go
│   ├── plugin-router.go
│   ├── plugin_router_test.go
│   ├── relay-router.go
│   ├── relay_router_test.go
│   ├── task-plugin-protocol-router.go
│   ├── task-router.go
│   ├── task_plugin_options_router_test.go
│   ├── task_plugin_protocol_router_test.go
│   ├── task_router_test.go
│   ├── video-router.go
│   ├── video_router_test.go
│   └── web-router.go
├── service/
│   ├── account_security.go
│   ├── audio.go
│   ├── auth_cleanup.go
│   ├── auth_session.go
│   ├── auth_session_test.go
│   ├── auth_token.go
│   ├── auth_token_test.go
│   ├── authz/
│   │   ├── adapter.go
│   │   ├── assignment.go
│   │   ├── authz_test.go
│   │   ├── enforcer.go
│   │   ├── override.go
│   │   ├── permission.go
│   │   ├── registry.go
│   │   ├── resolver.go
│   │   ├── resources_audit.go
│   │   ├── resources_channel.go
│   │   ├── resources_task_plugin.go
│   │   ├── role.go
│   │   └── seed.go
│   ├── billing.go
│   ├── billing_session.go
│   ├── billing_usage.go
│   ├── channel.go
│   ├── channel_affinity.go
│   ├── channel_affinity_template_test.go
│   ├── channel_affinity_usage_cache_test.go
│   ├── channel_select.go
│   ├── channel_select_auto_groups_test.go
│   ├── channel_select_binding_test.go
│   ├── channel_select_test.go
│   ├── codex_channel_models.go
│   ├── codex_credential_refresh.go
│   ├── codex_credential_refresh_task.go
│   ├── codex_models.go
│   ├── codex_oauth.go
│   ├── codex_wham_usage.go
│   ├── convert.go
│   ├── convert_test.go
│   ├── download.go
│   ├── email_binding.go
│   ├── epay.go
│   ├── error.go
│   ├── error_test.go
│   ├── file_decoder.go
│   ├── file_service.go
│   ├── funding_source.go
│   ├── group.go
│   ├── group_auto_groups_test.go
│   ├── http.go
│   ├── http_client.go
│   ├── http_client_transport_test.go
│   ├── http_transport_policy.go
│   ├── http_transport_sharded.go
│   ├── image.go
│   ├── image_billing.go
│   ├── group_assignment.go
│   ├── group_assignment_test.go
│   ├── ldap.go
│   ├── ldap_test.go
│   ├── log_info_generate.go
│   ├── login_verification.go
│   ├── midjourney.go
│   ├── notify-limit.go
│   ├── oauth_verification.go
│   ├── openai_chat_responses_compat.go
│   ├── openai_chat_responses_mode.go
│   ├── passkey/
│   │   ├── service.go
│   │   ├── session.go
│   │   └── user.go
│   ├── protected_fetch_client.go
│   ├── protected_fetch_client_test.go
│   ├── quota.go
│   ├── quota_saturation_test.go
│   ├── rankings.go
│   ├── relay_error.go
│   ├── relay_error_test.go
│   ├── request_converter.go
│   ├── request_policy.go
│   ├── response_converter.go
│   ├── responses_usage.go
│   ├── responses_usage_test.go
│   ├── return_path.go
│   ├── return_path_test.go
│   ├── security_verification.go
│   ├── sensitive.go
│   ├── str.go
│   ├── subscription_reset_task.go
│   ├── system_instance.go
│   ├── system_task.go
│   ├── system_task_test.go
│   ├── task.go
│   ├── task_artifact_access.go
│   ├── task_artifact_access_test.go
│   ├── task_artifact_store.go
│   ├── task_artifact_store_test.go
│   ├── task_billing.go
│   ├── task_billing_test.go
│   ├── task_plugin_audit.go
│   ├── task_plugin_view.go
│   ├── task_plugin_view_test.go
│   ├── task_polling.go
│   ├── task_polling_test.go
│   ├── text_quota.go
│   ├── text_quota_test.go
│   ├── tiered_settle.go
│   ├── tiered_settle_test.go
│   ├── token_counter.go
│   ├── token_estimator.go
│   ├── tokenizer.go
│   ├── twofa.go
│   ├── usage_helpr.go
│   ├── user_notify.go
│   ├── violation_fee.go
│   ├── waffo_pancake.go
│   ├── webhook.go
│   └── ws_close.go
├── setting/
│   ├── auto_group.go
│   ├── auto_group_test.go
│   ├── billing_setting/
│   │   ├── builtin_billing.go
│   │   ├── builtin_billing_test.go
│   │   ├── tiered_billing.go
│   │   └── tiered_billing_test.go
│   ├── chat.go
│   ├── config/
│   │   ├── config.go
│   │   └── config_test.go
│   ├── console_setting/
│   │   ├── config.go
│   │   └── validation.go
│   ├── group_channel_setting/
│   │   ├── group_channel.go
│   │   └── group_channel_test.go
│   ├── midjourney.go
│   ├── model_setting/
│   │   ├── claude.go
│   │   ├── claude_test.go
│   │   ├── gemini.go
│   │   ├── gemini_test.go
│   │   ├── global.go
│   │   ├── global_test.go
│   │   ├── grok.go
│   │   └── qwen.go
│   ├── operation_setting/
│   │   ├── channel_affinity_setting.go
│   │   ├── checkin_setting.go
│   │   ├── general_setting.go
│   │   ├── monitor_setting.go
│   │   ├── monitor_setting_test.go
│   │   ├── operation_setting.go
│   │   ├── payment_setting.go
│   │   ├── payment_setting_old.go
│   │   ├── quota_setting.go
│   │   ├── status_code_ranges.go
│   │   ├── status_code_ranges_test.go
│   │   ├── token_setting.go
│   │   ├── tools.go
│   │   └── tools_price_test.go
│   ├── payment_creem.go
│   ├── payment_stripe.go
│   ├── payment_waffo.go
│   ├── payment_waffo_pancake.go
│   ├── perf_metrics_setting/
│   │   └── config.go
│   ├── performance_setting/
│   │   └── config.go
│   ├── rate_limit.go
│   ├── ratio_setting/
│   │   ├── cache_ratio.go
│   │   ├── compact_suffix.go
│   │   ├── expose_ratio.go
│   │   ├── exposed_cache.go
│   │   ├── group_ratio.go
│   │   ├── matching_test.go
│   │   └── model_ratio.go
│   ├── reasoning/
│   │   ├── suffix.go
│   │   └── suffix_test.go
│   ├── sensitive.go
│   ├── system_setting/
│   │   ├── cas.go
│   │   ├── discord.go
│   │   ├── fetch_setting.go
│   │   ├── group_assignment.go
│   │   ├── ldap.go
│   │   ├── legal.go
│   │   ├── oidc.go
│   │   ├── oidc_test.go
│   │   ├── passkey.go
│   │   ├── system_setting_old.go
│   │   ├── task_artifact.go
│   │   ├── task_artifact_store.go
│   │   ├── task_artifact_store_test.go
│   │   ├── task_artifact_test.go
│   │   └── telegram.go
│   ├── task_plugin.go
│   ├── task_plugin_test.go
│   ├── task_pricing_setting/
│   │   ├── config.go
│   │   └── config_test.go
│   └── user_usable_group.go
├── sources.list
├── types/
│   ├── price_data.go
│   ├── rw_map.go
│   ├── set.go
│   └── task_artifact.go
└── web/
    ├── classic/
    │   ├── .eslintrc.cjs
    │   ├── .gitignore
    │   ├── .prettierrc.mjs
    │   ├── Makefile
    │   ├── bun.lock
    │   ├── i18next.config.js
    │   ├── index.html
    │   ├── jsconfig.json
    │   ├── package.json
    │   ├── postcss.config.js
    │   ├── public/
    │   │   ├── azure_model_name.png
    │   │   ├── cover-4.webp
    │   │   ├── favicon.ico
    │   │   ├── logo.png
    │   │   ├── pay-apple.png
    │   │   ├── pay-card.png
    │   │   ├── pay-google.png
    │   │   ├── ratio.png
    │   │   └── robots.txt
    │   ├── src/
    │   │   ├── App.jsx
    │   │   ├── components/
    │   │   │   ├── auth/
    │   │   │   │   ├── LoginForm.jsx
    │   │   │   │   ├── OAuth2Callback.jsx
    │   │   │   │   ├── PasswordResetConfirm.jsx
    │   │   │   │   ├── PasswordResetForm.jsx
    │   │   │   │   ├── RegisterForm.jsx
    │   │   │   │   └── TwoFAVerification.jsx
    │   │   │   ├── common/
    │   │   │   │   ├── DocumentRenderer/
    │   │   │   │   │   └── index.jsx
    │   │   │   │   ├── ErrorBoundary.jsx
    │   │   │   │   ├── logo/
    │   │   │   │   │   ├── LinuxDoIcon.jsx
    │   │   │   │   │   ├── OIDCIcon.jsx
    │   │   │   │   │   └── WeChatIcon.jsx
    │   │   │   │   ├──kdown/
    │   │   │   │   │   ├── MarkdownRenderer.jsx
    │   │   │   │   │   └──kdown.css
    │   │   │   │   ├── modals/
    │   │   │   │   │   ├── RiskAcknowledgementModal.jsx
    │   │   │   │   │   └── SecureVerificationModal.jsx
    │   │   │   │   └── ui/
    │   │   │   │       ├── CardPro.jsx
    │   │   │   │       ├── CardTable.jsx
    │   │   │   │       ├── ChannelKeyDisplay.jsx
    │   │   │   │       ├── CompactModeToggle.jsx
    │   │   │   │       ├── JSONEditor.jsx
    │   │   │   │       ├── Loading.jsx
    │   │   │   │       ├── RenderUtils.jsx
    │   │   │   │       ├── ScrollableContainer.jsx
    │   │   │   │       └── SelectableButtonGroup.jsx
    │   │   │   ├── dashboard/
    │   │   │   │   ├── AnnouncementsPanel.jsx
    │   │   │   │   ├── ApiInfoPanel.jsx
    │   │   │   │   ├── ChartsPanel.jsx
    │   │   │   │   ├── DashboardHeader.jsx
    │   │   │   │   ├── FaqPanel.jsx
    │   │   │   │   ├── StatsCards.jsx
    │   │   │   │   ├── UptimePanel.jsx
    │   │   │   │   ├── index.jsx
    │   │   │   │   └── modals/
    │   │   │   │       └── SearchModal.jsx
    │   │   │   ├── layout/
    │   │   │   │   ├── Footer.jsx
    │   │   │   │   ├── NoticeModal.jsx
    │   │   │   │   ├── PageLayout.jsx
    │   │   │   │   ├── SetupCheck.js
    │   │   │   │   ├── SiderBar.jsx
    │   │   │   │   ├── components/
    │   │   │   │   │   └── SkeletonWrapper.jsx
    │   │   │   │   └── headerbar/
    │   │   │   │       ├── ActionButtons.jsx
    │   │   │   │       ├── HeaderLogo.jsx
    │   │   │   │       ├── LanguageSelector.jsx
    │   │   │   │       ├── MobileMenuButton.jsx
    │   │   │   │       ├── Navigation.jsx
    │   │   │   │       ├── NewYearButton.jsx
    │   │   │   │       ├── NotificationButton.jsx
    │   │   │   │       ├── ThemeToggle.jsx
    │   │   │   │       ├── UserArea.jsx
    │   │   │   │       └── index.jsx
    │   │   │   ├── model-deployments/
    │   │   │   │   └── DeploymentAccessGuard.jsx
    │   │   │   ├── playground/
    │   │   │   │   ├── ChatArea.jsx
    │   │   │   │   ├── CodeViewer.jsx
    │   │   │   │   ├── ConfigManager.jsx
    │   │   │   │   ├── CustomInputRender.jsx
    │   │   │   │   ├── CustomRequestEditor.jsx
    │   │   │   │   ├── DebugPanel.jsx
    │   │   │   │   ├── FloatingButtons.jsx
    │   │   │   │   ├── ImageUrlInput.jsx
    │   │   │   │   ├── MessageActions.jsx
    │   │   │   │   ├── MessageContent.jsx
    │   │   │   │   ├── OptimizedComponents.js
    │   │   │   │   ├── ParameterControl.jsx
    │   │   │   │   ├── SSEViewer.jsx
    │   │   │   │   ├── SettingsPanel.jsx
    │   │   │   │   ├── ThinkingContent.jsx
    │   │   │   │   └── configStorage.js
    │   │   │   ├── settings/
    │   │   │   │   ├── ChannelSelectorModal.jsx
    │   │   │   │   ├── ChatsSetting.jsx
    │   │   │   │   ├── CustomOAuthSetting.jsx
    │   │   │   │   ├── DashboardSetting.jsx
    │   │   │   │   ├── DrawingSetting.jsx
    │   │   │   │   ├── HttpStatusCodeRulesInput.jsx
    │   │   │   │   ├── ModelDeploymentSetting.jsx
    │   │   │   │   ├── ModelSetting.jsx
    │   │   │   │   ├── OperationSetting.jsx
    │   │   │   │   ├── OtherSetting.jsx
    │   │   │   │   ├── PaymentSetting.jsx
    │   │   │   │   ├── PerformanceSetting.jsx
    │   │   │   │   ├── PersonalSetting.jsx
    │   │   │   │   ├── RateLimitSetting.jsx
    │   │   │   │   ├── RatioSetting.jsx
    │   │   │   │   ├── SystemSetting.jsx
    │   │   │   │   └── personal/
    │   │   │   │       ├── cards/
    │   │   │   │       │   ├── AccountManagement.jsx
    │   │   │   │       │   ├── CheckinCalendar.jsx
    │   │   │   │       │   ├── NotificationSettings.jsx
    │   │   │   │       │   └── PreferencesSettings.jsx
    │   │   │   │       ├── components/
    │   │   │   │       │   ├── TwoFASetting.jsx
    │   │   │   │       │   └── UserInfoHeader.jsx
    │   │   │   │       └── modals/
    │   │   │   │           ├── AccountDeleteModal.jsx
    │   │   │   │           ├── ChangePasswordModal.jsx
    │   │   │   │           ├── EmailBindModal.jsx
    │   │   │   │           └── WeChatBindModal.jsx
    │   │   │   ├── setup/
    │   │   │   │   ├── SetupWizard.jsx
    │   │   │   │   ├── components/
    │   │   │   │   │   ├── StepNavigation.jsx
    │   │   │   │   │   └── steps/
    │   │   │   │   │       ├── AdminStep.jsx
    │   │   │   │   │       ├── CompleteStep.jsx
    │   │   │   │   │       ├── DatabaseStep.jsx
    │   │   │   │   │       └── UsageModeStep.jsx
    │   │   │   │   └── index.jsx
    │   │   │   ├── table/
    │   │   │   │   ├── channels/
    │   │   │   │   │   ├── ChannelsActions.jsx
    │   │   │   │   │   ├── ChannelsColumnDefs.jsx
    │   │   │   │   │   ├── ChannelsFilters.jsx
    │   │   │   │   │   ├── ChannelsTable.jsx
    │   │   │   │   │   ├── ChannelsTabs.jsx
    │   │   │   │   │   ├── index.jsx
    │   │   │   │   │   └── modals/
    │   │   │   │   │       ├── BatchTagModal.jsx
    │   │   │   │   │       ├── ChannelUpstreamUpdateModal.jsx
    │   │   │   │   │       ├── CodexOAuthModal.jsx
    │   │   │   │   │       ├── CodexUsageModal.jsx
    │   │   │   │   │       ├── ColumnSelectorModal.jsx
    │   │   │   │   │       ├── EditChannelModal.jsx
    │   │   │   │   │       ├── EditTagModal.jsx
    │   │   │   │   │       ├── ModelSelectModal.jsx
    │   │   │   │   │       ├── ModelTestModal.jsx
    │   │   │   │   │       ├── MultiKeyManageModal.jsx
    │   │   │   │   │       ├── OllamaModelModal.jsx
    │   │   │   │   │       ├── ParamOverrideEditorModal.jsx
    │   │   │   │   │       ├── SingleModelSelectModal.jsx
    │   │   │   │   │       ├── StatusCodeRiskGuardModal.jsx
    │   │   │   │   │       └── statusCodeRiskGuard.js
    │   │   │   │   ├── mj-logs/
    │   │   │   │   │   ├── MjLogsActions.jsx
    │   │   │   │   │   ├── MjLogsColumnDefs.jsx
    │   │   │   │   │   ├── MjLogsFilters.jsx
    │   │   │   │   │   ├── MjLogsTable.jsx
    │   │   │   │   │   ├── index.jsx
    │   │   │   │   │   └── modals/
    │   │   │   │   │       ├── ColumnSelectorModal.jsx
    │   │   │   │   │       └── ContentModal.jsx
    │   │   │   │   ├── model-deployments/
    │   │   │   │   │   ├── DeploymentsActions.jsx
    │   │   │   │   │   ├── DeploymentsColumnDefs.jsx
    │   │   │   │   │   ├── DeploymentsFilters.jsx
    │   │   │   │   │   ├── DeploymentsTable.jsx
    │   │   │   │   │   ├── index.jsx
    │   │   │   │   │   └── modals/
    │   │   │   │   │       ├── ColumnSelectorModal.jsx
    │   │   │   │   │       ├── ConfirmationDialog.jsx
    │   │   │   │   │       ├── CreateDeploymentModal.jsx
    │   │   │   │   │       ├── EditDeploymentModal.jsx
    │   │   │   │   │       ├── ExtendDurationModal.jsx
    │   │   │   │   │       ├── UpdateConfigModal.jsx
    │   │   │   │   │       ├── ViewDetailsModal.jsx
    │   │   │   │   │       └── ViewLogsModal.jsx
    │   │   │   │   ├── model-pricing/
    │   │   │   │   │   ├── filter/
    │   │   │   │   │   │   ├── PricingDisplaySettings.jsx
    │   │   │   │   │   │   ├── PricingEndpointTypes.jsx
    │   │   │   │   │   │   ├── PricingGroups.jsx
    │   │   │   │   │   │   ├── PricingQuotaTypes.jsx
    │   │   │   │   │   │   ├── PricingTags.jsx
    │   │   │   │   │   │   └── PricingVendors.jsx
    │   │   │   │   │   ├── layout/
    │   │   │   │   │   │   ├── PricingPage.jsx
    │   │   │   │   │   │   ├── PricingSidebar.jsx
    │   │   │   │   │   │   ├── content/
    │   │   │   │   │   │   │   ├── PricingContent.jsx
    │   │   │   │   │   │   │   └── PricingView.jsx
    │   │   │   │   │   │   └── header/
    │   │   │   │   │   │       ├── PricingTopSection.jsx
    │   │   │   │   │   │       ├── PricingVendorIntro.jsx
    │   │   │   │   │   │       ├── PricingVendorIntroSkeleton.jsx
    │   │   │   │   │   │       ├── PricingVendorIntroWithSkeleton.jsx
    │   │   │   │   │   │       └── SearchActions.jsx
    │   │   │   │   │   ├── modal/
    │   │   │   │   │   │   ├── ModelDetailSideSheet.jsx
    │   │   │   │   │   │   ├── PricingFilterModal.jsx
    │   │   │   │   │   │   └── components/
    │   │   │   │   │   │       ├── FilterModalContent.jsx
    │   │   │   │   │   │       ├── FilterModalFooter.jsx
    │   │   │   │   │   │       ├── ModelBasicInfo.jsx
    │   │   │   │   │   │       ├── ModelEndpoints.jsx
    │   │   │   │   │   │       ├── ModelHeader.jsx
    │   │   │   │   │   │       └── ModelPricingTable.jsx
    │   │   │   │   │   └── view/
    │   │   │   │   │       ├── card/
    │   │   │   │   │       │   ├── PricingCardSkeleton.jsx
    │   │   │   │   │       │   └── PricingCardView.jsx
    │   │   │   │   │       └── table/
    │   │   │   │   │           ├── PricingTable.jsx
    │   │   │   │   │           └── PricingTableColumns.jsx
    │   │   │   │   ├── models/
    │   │   │   │   │   ├── ModelsActions.jsx
    │   │   │   │   │   ├── ModelsColumnDefs.jsx
    │   │   │   │   │   ├── ModelsFilters.jsx
    │   │   │   │   │   ├── ModelsTable.jsx
    │   │   │   │   │   ├── ModelsTabs.jsx
    │   │   │   │   │   ├── components/
    │   │   │   │   │   │   └── SelectionNotification.jsx
    │   │   │   │   │   ├── index.jsx
    │   │   │   │   │   └── modals/
    │   │   │   │   │       ├── EditModelModal.jsx
    │   │   │   │   │       ├── EditPrefillGroupModal.jsx
    │   │   │   │   │       ├── EditVendorModal.jsx
    │   │   │   │   │       ├── MissingModelsModal.jsx
    │   │   │   │   │       ├── PrefillGroupManagement.jsx
    │   │   │   │   │       ├── SyncWizardModal.jsx
    │   │   │   │   │       └── UpstreamConflictModal.jsx
    │   │   │   │   ├── redemptions/
    │   │   │   │   │   ├── RedemptionsActions.jsx
    │   │   │   │   │   ├── RedemptionsColumnDefs.jsx
    │   │   │   │   │   ├── RedemptionsDescription.jsx
    │   │   │   │   │   ├── RedemptionsFilters.jsx
    │   │   │   │   │   ├── RedemptionsTable.jsx
    │   │   │   │   │   ├── index.jsx
    │   │   │   │   │   └── modals/
    │   │   │   │   │       ├── DeleteRedemptionModal.jsx
    │   │   │   │   │       └── EditRedemptionModal.jsx
    │   │   │   │   ├── subscriptions/
    │   │   │   │   │   ├── SubscriptionsActions.jsx
    │   │   │   │   │   ├── SubscriptionsColumnDefs.jsx
    │   │   │   │   │   ├── SubscriptionsDescription.jsx
    │   │   │   │   │   ├── SubscriptionsTable.jsx
    │   │   │   │   │   ├── index.jsx
    │   │   │   │   │   └── modals/
    │   │   │   │   │       └── AddEditSubscriptionModal.jsx
    │   │   │   │   ├── task-logs/
    │   │   │   │   │   ├── TaskLogsActions.jsx
    │   │   │   │   │   ├── TaskLogsColumnDefs.jsx
    │   │   │   │   │   ├── TaskLogsFilters.jsx
    │   │   │   │   │   ├── TaskLogsTable.jsx
    │   │   │   │   │   ├── index.jsx
    │   │   │   │   │   └── modals/
    │   │   │   │   │       ├── AudioPreviewModal.jsx
    │   │   │   │   │       ├── ColumnSelectorModal.jsx
    │   │   │   │   │       └── ContentModal.jsx
    │   │   │   │   ├── tokens/
    │   │   │   │   │   ├── TokensActions.jsx
    │   │   │   │   │   ├── TokensColumnDefs.jsx
    │   │   │   │   │   ├── TokensDescription.jsx
    │   │   │   │   │   ├── TokensFilters.jsx
    │   │   │   │   │   ├── TokensTable.jsx
    │   │   │   │   │   ├── index.jsx
    │   │   │   │   │   └── modals/
    │   │   │   │   │       ├── CCSwitchModal.jsx
    │   │   │   │   │       ├── CopyTokensModal.jsx
    │   │   │   │   │       ├── DeleteTokensModal.jsx
    │   │   │   │   │       └── EditTokenModal.jsx
    │   │   │   │   ├── usage-logs/
    │   │   │   │   │   ├── UsageLogsActions.jsx
    │   │   │   │   │   ├── UsageLogsColumnDefs.jsx
    │   │   │   │   │   ├── UsageLogsFilters.jsx
    │   │   │   │   │   ├── UsageLogsTable.jsx
    │   │   │   │   │   ├── components/
    │   │   │   │   │   │   └── ParamOverrideEntry.jsx
    │   │   │   │   │   ├── index.jsx
    │   │   │   │   │   └── modals/
    │   │   │   │   │       ├── ChannelAffinityUsageCacheModal.jsx
    │   │   │   │   │       ├── ColumnSelectorModal.jsx
    │   │   │   │   │       ├── ParamOverrideModal.jsx
    │   │   │   │   │       └── UserInfoModal.jsx
    │   │   │   │   └── users/
    │   │   │   │       ├── UsersActions.jsx
    │   │   │   │       ├── UsersColumnDefs.jsx
    │   │   │   │       ├── UsersDescription.jsx
    │   │   │   │       ├── UsersFilters.jsx
    │   │   │   │       ├── UsersTable.jsx
    │   │   │   │       ├── index.jsx
    │   │   │   │       └── modals/
    │   │   │   │           ├── AddUserModal.jsx
    │   │   │   │           ├── DeleteUserModal.jsx
    │   │   │   │           ├── DemoteUserModal.jsx
    │   │   │   │           ├── EditUserModal.jsx
    │   │   │   │           ├── EnableDisableUserModal.jsx
    │   │   │   │           ├── PromoteUserModal.jsx
    │   │   │   │           ├── ResetPasskeyModal.jsx
    │   │   │   │           ├── ResetTwoFAModal.jsx
    │   │   │   │           ├── UserBindingManagementModal.jsx
    │   │   │   │           └── UserSubscriptionsModal.jsx
    │   │   │   └── topup/
    │   │   │       ├── InvitationCard.jsx
    │   │   │       ├── RechargeCard.jsx
    │   │   │       ├── SubscriptionPlansCard.jsx
    │   │   │       ├── index.jsx
    │   │   │       └── modals/
    │   │   │           ├── PaymentConfirmModal.jsx
    │   │   │           ├── SubscriptionPurchaseModal.jsx
    │   │   │           ├── TopupHistoryModal.jsx
    │   │   │           └── TransferModal.jsx
    │   │   ├── constants/
    │   │   │   ├── channel-affinity-template.constants.js
    │   │   │   ├── channel.constants.js
    │   │   │   ├── common.constant.js
    │   │   │   ├── console.constants.js
    │   │   │   ├── dashboard.constants.js
    │   │   │   ├── index.js
    │   │   │   ├── playground.constants.js
    │   │   │   ├── redemption.constants.js
    │   │   │   ├── toast.constants.js
    │   │   │   └── user.constants.js
    │   │   ├── context/
    │   │   │   ├── Status/
    │   │   │   │   ├── index.jsx
    │   │   │   │   └── reducer.js
    │   │   │   ├── Theme/
    │   │   │   │   └── index.jsx
    │   │   │   └── User/
    │   │   │       ├── index.jsx
    │   │   │       └── reducer.js
    │   │   ├── contexts/
    │   │   │   └── PlaygroundContext.jsx
    │   │   ├── helpers/
    │   │   │   ├── api.js
    │   │   │   ├── auth.jsx
    │   │   │   ├── base64.js
    │   │   │   ├── boolean.js
    │   │   │   ├── dashboard.jsx
    │   │   │   ├── data.js
    │   │   │   ├── history.js
    │   │   │   ├── index.js
    │   │   │   ├── log.js
    │   │   │   ├── passkey.js
    │   │   │   ├── quota.js
    │   │   │   ├── render.jsx
    │   │   │   ├── secureApiCall.js
    │   │   │   ├── statusCodeRules.js
    │   │   │   ├── subscriptionFormat.js
    │   │   │   ├── token.js
    │   │   │   └── utils.jsx
    │   │   ├── hooks/
    │   │   │   ├── channels/
    │   │   │   │   ├── upstreamUpdateUtils.js
    │   │   │   │   ├── useChannelUpstreamUpdates.jsx
    │   │   │   │   └── useChannelsData.jsx
    │   │   │   ├── chat/
    │   │   │   │   └── useTokenKeys.js
    │   │   │   ├── common/
    │   │   │   │   ├── useContainerWidth.js
    │   │   │   │   ├── useHeaderBar.js
    │   │   │   │   ├── useIsMobile.js
    │   │   │   │   ├── useMinimumLoadingTime.js
    │   │   │   │   ├── useNavigation.js
    │   │   │   │   ├── useNotifications.js
    │   │   │   │   ├── useSecureVerification.jsx
    │   │   │   │   ├── useSidebar.js
    │   │   │   │   ├── useSidebarCollapsed.js
    │   │   │   │   ├── useTableCompactMode.js
    │   │   │   │   └── useUserPermissions.js
    │   │   │   ├── dashboard/
    │   │   │   │   ├── useDashboardCharts.jsx
    │   │   │   │   ├── useDashboardData.js
    │   │   │   │   └── useDashboardStats.jsx
    │   │   │   ├── mj-logs/
    │   │   │   │   └── useMjLogsData.js
    │   │   │   ├── model-deployments/
    │   │   │   │   ├── useDeploymentsData.jsx
    │   │   │   │   └── useModelDeploymentSettings.js
    │   │   │   ├── model-pricing/
    │   │   │   │   ├── useModelPricingData.jsx
    │   │   │   │   └── usePricingFilterCounts.js
    │   │   │   ├── models/
    │   │   │   │   └── useModelsData.jsx
    │   │   │   ├── playground/
    │   │   │   │   ├── useApiRequest.jsx
    │   │   │   │   ├── useDataLoader.js
    │   │   │   │   ├── useMessageActions.jsx
    │   │   │   │   ├── useMessageEdit.jsx
    │   │   │   │   ├── usePlaygroundState.js
    │   │   │   │   └── useSyncMessageAndCustomBody.js
    │   │   │   ├── redemptions/
    │   │   │   │   └── useRedemptionsData.jsx
    │   │   │   ├── subscriptions/
    │   │   │   │   └── useSubscriptionsData.jsx
    │   │   │   ├── task-logs/
    │   │   │   │   └── useTaskLogsData.js
    │   │   │   ├── tokens/
    │   │   │   │   └── useTokensData.jsx
    │   │   │   ├── usage-logs/
    │   │   │   │   └── useUsageLogsData.jsx
    │   │   │   └── users/
    │   │   │       └── useUsersData.jsx
    │   │   ├── i18n/
    │   │   │   ├── i18n.js
    │   │   │   ├── language.js
    │   │   │   └── locales/
    │   │   │       ├── en.json
    │   │   │       ├── fr.json
    │   │   │       ├── ja.json
    │   │   │       ├── ru.json
    │   │   │       ├── vi.json
    │   │   │       ├── zh-CN.json
    │   │   │       └── zh-TW.json
    │   │   ├── index.css
    │   │   ├── index.jsx
    │   │   ├── pages/
    │   │   │   ├── About/
    │   │   │   │   └── index.jsx
    │   │   │   ├── Channel/
    │   │   │   │   └── index.jsx
    │   │   │   ├── Chat/
    │   │   │   │   └── index.jsx
    │   │   │   ├── Chat2Link/
    │   │   │   │   └── index.jsx
    │   │   │   ├── Dashboard/
    │   │   │   │   └── index.jsx
    │   │   │   ├── Forbidden/
    │   │   │   │   └── index.jsx
    │   │   │   ├── Home/
    │   │   │   │   └── index.jsx
    │   │   │   ├── Log/
    │   │   │   │   └── index.jsx
    │   │   │   ├── Midjourney/
    │   │   │   │   └── index.jsx
    │   │   │   ├── Model/
    │   │   │   │   └── index.jsx
    │   │   │   ├── ModelDeployment/
    │   │   │   │   └── index.jsx
    │   │   │   ├── NotFound/
    │   │   │   │   └── index.jsx
    │   │   │   ├── Playground/
    │   │   │   │   └── index.jsx
    │   │   │   ├── Pricing/
    │   │   │   │   └── index.jsx
    │   │   │   ├── PrivacyPolicy/
    │   │   │   │   └── index.jsx
    │   │   │   ├── Redemption/
    │   │   │   │   └── index.jsx
    │   │   │   ├── Setting/
    │   │   │   │   ├── Chat/
    │   │   │   │   │   └── SettingsChats.jsx
    │   │   │   │   ├── Dashboard/
    │   │   │   │   │   ├── SettingsAPIInfo.jsx
    │   │   │   │   │   ├── SettingsAnnouncements.jsx
    │   │   │   │   │   ├── SettingsDataDashboard.jsx
    │   │   │   │   │   ├── SettingsFAQ.jsx
    │   │   │   │   │   └── SettingsUptimeKuma.jsx
    │   │   │   │   ├── Drawing/
    │   │   │   │   │   └── SettingsDrawing.jsx
    │   │   │   │   ├── Model/
    │   │   │   │   │   ├── SettingClaudeModel.jsx
    │   │   │   │   │   ├── SettingGeminiModel.jsx
    │   │   │   │   │   ├── SettingGlobalModel.jsx
    │   │   │   │   │   ├── SettingGrokModel.jsx
    │   │   │   │   │   └── SettingModelDeployment.jsx
    │   │   │   │   ├── Operation/
    │   │   │   │   │   ├── SettingsChannelAffinity.jsx
    │   │   │   │   │   ├── SettingsCheckin.jsx
    │   │   │   │   │   ├── SettingsCreditLimit.jsx
    │   │   │   │   │   ├── SettingsGeneral.jsx
    │   │   │   │   │   ├── SettingsHeaderNavModules.jsx
    │   │   │   │   │   ├── SettingsLog.jsx
    │   │   │   │   │   ├── SettingsMonitoring.jsx
    │   │   │   │   │   ├── SettingsSensitiveWords.jsx
    │   │   │   │   │   ├── SettingsSidebarModulesAdmin.jsx
    │   │   │   │   │   └── SettingsSyslog.jsx
    │   │   │   │   ├── Payment/
    │   │   │   │   │   ├── SettingsGeneralPayment.jsx
    │   │   │   │   │   ├── SettingsPaymentGateway.jsx
    │   │   │   │   │   ├── SettingsPaymentGatewayCreem.jsx
    │   │   │   │   │   ├── SettingsPaymentGatewayStripe.jsx
    │   │   │   │   │   ├── SettingsPaymentGatewayWaffo.jsx
    │   │   │   │   │   └── SettingsPaymentGatewayWaffoPancake.jsx
    │   │   │   │   ├── Performance/
    │   │   │   │   │   └── SettingsPerformance.jsx
    │   │   │   │   ├── RateLimit/
    │   │   │   │   │   └── SettingsRequestRateLimit.jsx
    │   │   │   │   ├── Ratio/
    │   │   │   │   │   ├── GroupRatioSettings.jsx
    │   │   │   │   │   ├── ModelPricingCombined.jsx
    │   │   │   │   │   ├── ModelRatioSettings.jsx
    │   │   │   │   │   ├── ModelRationNotSetEditor.jsx
    │   │   │   │   │   ├── ModelSettingsVisualEditor.jsx
    │   │   │   │   │   ├── UpstreamRatioSync.jsx
    │   │   │   │   │   ├── components/
    │   │   │   │   │   │   ├── AutoGroupList.jsx
    │   │   │   │   │   │   ├── GroupChannelBinding.jsx
    │   │   │   │   │   │   ├── GroupGroupRatioRules.jsx
    │   │   │   │   │   │   ├── GroupSpecialUsableRules.jsx
    │   │   │   │   │   │   ├── GroupTable.jsx
    │   │   │   │   │   │   └── ModelPricingEditor.jsx
    │   │   │   │   │   └── hooks/
    │   │   │   │   │       └── useModelPricingEditorState.js
    │   │   │   │   └── index.jsx
    │   │   │   ├── Setup/
    │   │   │   │   └── index.jsx
    │   │   │   ├── Subscription/
    │   │   │   │   └── index.jsx
    │   │   │   ├── Task/
    │   │   │   │   └── index.jsx
    │   │   │   ├── Token/
    │   │   │   │   └── index.jsx
    │   │   │   ├── TopUp/
    │   │   │   │   └── index.js
    │   │   │   ├── User/
    │   │   │   │   └── index.jsx
    │   │   │   └── UserAgreement/
    │   │   │       └── index.jsx
    │   │   └── services/
    │   │       └── secureVerification.js
    │   ├── tailwind.config.js
    │   ├── vercel.json
    │   └── vite.config.js
    └── default/
        ├── .gitignore
        ├── .node-version
        ├── .npmrc
        ├── .oxfmtrc.json
        ├── .oxlintrc.json
        ├── AGENTS.md
        ├── Makefile
        ├── bun.lock
        ├── components.json
        ├── cz.yaml
        ├── index.html
        ├── knip.config.ts
        ├── netlify.toml
        ├── package.json
        ├── public/
        │   ├── favicon.ico
        │   ├── logo.png
        │   ├── pay-apple.png
        │   ├── pay-card.png
        │   ├── pay-google.png
        │   ├── waffo-logo-dark.svg
        │   └── waffo-logo-light.svg
        ├── rsbuild.config.ts
        ├── scripts/
        │   ├── add-copyright.mjs
        │   ├── format-with-protected-headers.mjs
        │   ├── oxlint/
        │   │   ├── __tests__/
        │   │   │   └── intl-locale.test.ts
        │   │   └── intl-locale.mjs
        │   └── sync-i18n.mjs
        ├── src/
        │   ├── assets/
        │   │   ├── brand-icons/
        │   │   │   ├── icon-discord.tsx
        │   │   │   ├── icon-docker.tsx
        │   │   │   ├── icon-facebook.tsx
        │   │   │   ├── icon-figma.tsx
        │   │   │   ├── icon-github.tsx
        │   │   │   ├── icon-gitlab.tsx
        │   │   │   ├── icon-gmail.tsx
        │   │   │   ├── icon-linuxdo.tsx
        │   │   │   ├── icon-medium.tsx
        │   │   │   ├── icon-notion.tsx
        │   │   │   ├── icon-skype.tsx
        │   │   │   ├── icon-slack.tsx
        │   │   │   ├── icon-stripe.tsx
        │   │   │   ├── icon-telegram.tsx
        │   │   │   ├── icon-trello.tsx
        │   │   │   ├── icon-wechat.tsx
        │   │   │   ├── icon-whatsapp.tsx
        │   │   │   ├── icon-zoom.tsx
        │   │   │   ├── index.ts
        │   │   │   └── sglang.svg
        │   │   ├── clerk-full-logo.tsx
        │   │   ├── clerk-logo.tsx
        │   │   ├── custom/
        │   │   │   ├── icon-dir.tsx
        │   │   │   ├── icon-layout-compact.tsx
        │   │   │   ├── icon-layout-default.tsx
        │   │   │   ├── icon-layout-full.tsx
        │   │   │   ├── icon-sidebar-floating.tsx
        │   │   │   ├── icon-sidebar-inset.tsx
        │   │   │   ├── icon-sidebar-sidebar.tsx
        │   │   │   ├── icon-sub2api.tsx
        │   │   │   ├── icon-theme-dark.tsx
        │   │   │   ├── icon-theme-light.tsx
        │   │   │   ├── icon-theme-system.tsx
        │   │   │   ├── icon-wan.tsx
        │   │   │   └── wan.png
        │   │   └── logo.tsx
        │   ├── components/
        │   │   ├── activity-time-cell.tsx
        │   │   ├── ai-elements/
        │   │   │   ├── __tests__/
        │   │   │   │   ├── code-block-editor.test.tsx
        │   │   │   │   ├── response-fade-render.test.tsx
        │   │   │   │   └── response-fade.test.ts
        │   │   │   ├── actions.tsx
        │   │   │   ├── artifact.tsx
        │   │   │   ├── branch.tsx
        │   │   │   ├── canvas.tsx
        │   │   │   ├── chain-of-thought.tsx
        │   │   │   ├── code-block.tsx
        │   │   │   ├── confirmation.tsx
        │   │   │   ├── connection.tsx
        │   │   │   ├── context.tsx
        │   │   │   ├── controls.tsx
        │   │   │   ├── conversation.tsx
        │   │   │   ├── edge.tsx
        │   │   │   ├── image.tsx
        │   │   │   ├── inline-citation.tsx
        │   │   │   ├── loader.tsx
        │   │   │   ├── message.tsx
        │   │   │   ├── node.tsx
        │   │   │   ├── open-in-chat.tsx
        │   │   │   ├── panel.tsx
        │   │   │   ├── plan.tsx
        │   │   │   ├── prompt-input.tsx
        │   │   │   ├── queue.tsx
        │   │   │   ├── reasoning.tsx
        │   │   │   ├── response-content.ts
        │   │   │   ├── response-fade.ts
        │   │   │   ├── response-node-guards.ts
        │   │   │   ├── response-renderer-alert.tsx
        │   │   │   ├── response-renderer-blocks.tsx
        │   │   │   ├── response-renderer-details.tsx
        │   │   │   ├── response-renderer-footnotes.tsx
        │   │   │   ├── response-renderer-image.tsx
        │   │   │   ├── response-renderer-inline.tsx
        │   │   │   ├── response-renderer-table.tsx
        │   │   │   ├── response-renderer.tsx
        │   │   │   ├── response-types.ts
        │   │   │   ├── response.tsx
        │   │   │   ├── shimmer.tsx
        │   │   │   ├── sources.tsx
        │   │   │   ├── suggestion.tsx
        │   │   │   ├── task.tsx
        │   │   │   ├── tool.tsx
        │   │   │   ├── toolbar.tsx
        │   │   │   └── web-preview.tsx
        │   │   ├── animate-in-view.tsx
        │   │   ├── auto-skeleton.tsx
        │   │   ├── coming-soon.tsx
        │   │   ├── command-menu.tsx
        │   │   ├── config-drawer.tsx
        │   │   ├── confirm-dialog.tsx
        │   │   ├── copy-button.tsx
        │   │   ├── data-table/
        │   │   │   ├── README.md
        │   │   │   ├── core/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   └── pagination.test.tsx
        │   │   │   │   ├── badge-cell.tsx
        │   │   │   │   ├── badge-list-cell.tsx
        │   │   │   │   ├── column-header.tsx
        │   │   │   │   ├── column-pinning.ts
        │   │   │   │   ├── content-sized-columns.ts
        │   │   │   │   ├── data-table-colgroup.tsx
        │   │   │   │   ├── data-table-header.tsx
        │   │   │   │   ├── data-table-row.tsx
        │   │   │   │   ├── data-table-view.tsx
        │   │   │   │   ├── pagination.tsx
        │   │   │   │   ├── row-action-menu.tsx
        │   │   │   │   ├── table-empty.tsx
        │   │   │   │   ├── table-sizing.ts
        │   │   │   │   ├── table-skeleton.tsx
        │   │   │   │   ├── truncated-cell.tsx
        │   │   │   │   └── types.ts
        │   │   │   ├── hooks/
        │   │   │   │   ├── use-data-table-view-mode.ts
        │   │   │   │   ├── use-data-table.ts
        │   │   │   │   └── use-debounced-column-filter.ts
        │   │   │   ├── index.ts
        │   │   │   ├── layout/
        │   │   │   │   ├── card-cell-utils.ts
        │   │   │   │   ├── card-grid.tsx
        │   │   │   │   ├── card-row-content.tsx
        │   │   │   │   ├── data-table-page.tsx
        │   │   │   │   └── mobile-card-list.tsx
        │   │   │   ├── static/
        │   │   │   │   ├── static-data-table-classnames.ts
        │   │   │   │   ├── static-data-table.tsx
        │   │   │   │   └── static-row-actions.tsx
        │   │   │   └── toolbar/
        │   │   │       ├── __tests__/
        │   │   │       │   └── mobile-filter.test.tsx
        │   │   │       ├── bulk-actions.tsx
        │   │   │       ├── faceted-filter.tsx
        │   │   │       ├── mobile-filter-panel.tsx
        │   │   │       ├── toolbar.tsx
        │   │   │       ├── view-mode-toggle.tsx
        │   │   │       └── view-options.tsx
        │   │   ├── date-picker.tsx
        │   │   ├── datetime-picker.tsx
        │   │   ├── dialog.tsx
        │   │   ├── drawer-layout.ts
        │   │   ├── empty-state.tsx
        │   │   ├── error-state.tsx
        │   │   ├── floating-window/
        │   │   │   └── __tests__/
        │   │   │       └── floating-window.test.tsx
        │   │   ├── floating-window.tsx
        │   │   ├── group-badge.tsx
        │   │   ├── html-content.tsx
        │   │   ├── json-code-editor/
        │   │   │   ├── __tests__/
        │   │   │   │   ├── json-code-editor-utils.test.ts
        │   │   │   │   └── json-code-editor.test.tsx
        │   │   │   └── json-code-editor-utils.ts
        │   │   ├── json-code-editor.tsx
        │   │   ├── json-editor.tsx
        │   │   ├── language-switcher.tsx
        │   │   ├── layout/
        │   │   │   ├── components/
        │   │   │   │   ├── app-header.tsx
        │   │   │   │   ├── app-sidebar.tsx
        │   │   │   │   ├── authenticated-layout.tsx
        │   │   │   │   ├── chat-presets-item.tsx
        │   │   │   │   ├── footer.tsx
        │   │   │   │   ├── glow.tsx
        │   │   │   │   ├── header-logo.tsx
        │   │   │   │   ├── header.tsx
        │   │   │   │   ├── logo.tsx
        │   │   │   │   ├── main.tsx
        │   │   │   │   ├── mobile-drawer.tsx
        │   │   │   │   ├── mockup.tsx
        │   │   │   │   ├── nav-group.tsx
        │   │   │   │   ├── nav-link-item.tsx
        │   │   │   │   ├── navbar.tsx
        │   │   │   │   ├── page-footer.tsx
        │   │   │   │   ├── public-header.tsx
        │   │   │   │   ├── public-layout.tsx
        │   │   │   │   ├── public-navigation.tsx
        │   │   │   │   ├── section-page-layout.tsx
        │   │   │   │   ├── section.tsx
        │   │   │   │   ├── sidebar-view-header.tsx
        │   │   │   │   ├── system-brand.tsx
        │   │   │   │   └── top-nav.tsx
        │   │   │   ├── config/
        │   │   │   │   ├── system-settings.config.ts
        │   │   │   │   └── top-nav.config.ts
        │   │   │   ├── constants.ts
        │   │   │   ├── index.ts
        │   │   │   ├── lib/
        │   │   │   │   ├── sidebar-view-registry.ts
        │   │   │   │   └── url-utils.ts
        │   │   │   └── types.ts
        │   │   ├── learn-more.tsx
        │   │   ├── loading-state.tsx
        │   │   ├── lobe-icon-field.tsx
        │   │   ├── long-text.tsx
        │   │   ├── masked-value-display.tsx
        │   │   ├── model-group-selector/
        │   │   │   ├── __tests__/
        │   │   │   │   └── layout.test.ts
        │   │   │   └── layout.ts
        │   │   ├── model-group-selector.tsx
        │   │   ├── multi-select/
        │   │   │   └── __tests__/
        │   │   │       ├── option-hint.test.tsx
        │   │   │       └── selection.test.tsx
        │   │   ├── multi-select.tsx
        │   │   ├── navigation-progress.tsx
        │   │   ├── notification-popover.tsx
        │   │   ├── page-transition.tsx
        │   │   ├── password-input.tsx
        │   │   ├── profile-dropdown.tsx
        │   │   ├── provider-badge.tsx
        │   │   ├── quota-details-popover.tsx
        │   │   ├── react-icon-by-name.tsx
        │   │   ├── rich-content.tsx
        │   │   ├── risk-acknowledgement-dialog.tsx
        │   │   ├── search.tsx
        │   │   ├── sign-out-dialog.tsx
        │   │   ├── skip-to-main.tsx
        │   │   ├── status-badge.tsx
        │   │   ├── table-id.tsx
        │   │   ├── tag-input.tsx
        │   │   ├── theme-quick-switcher.tsx
        │   │   ├── theme-switch.tsx
        │   │   ├── truncated-text.tsx
        │   │   ├── turnstile.tsx
        │   │   └── ui/
        │   │       ├── __tests__/
        │   │       │   ├── combobox.test.tsx
        │   │       │   ├── label.test.tsx
        │   │       │   └── portal-container.test.tsx
        │   │       ├── accordion.tsx
        │   │       ├── alert-dialog.tsx
        │   │       ├── alert.tsx
        │   │       ├── aspect-ratio.tsx
        │   │       ├── avatar.tsx
        │   │       ├── badge.tsx
        │   │       ├── breadcrumb.tsx
        │   │       ├── button-group.tsx
        │   │       ├── button.tsx
        │   │       ├── calendar.tsx
        │   │       ├── card.tsx
        │   │       ├── carousel.tsx
        │   │       ├── chart.tsx
        │   │       ├── checkbox.tsx
        │   │       ├── collapsible.tsx
        │   │       ├── combobox-input.tsx
        │   │       ├── combobox.tsx
        │   │       ├── command.tsx
        │   │       ├── context-menu.tsx
        │   │       ├── dialog.tsx
        │   │       ├── direction.tsx
        │   │       ├── drawer.tsx
        │   │       ├── dropdown-menu-events.ts
        │   │       ├── dropdown-menu.test.tsx
        │   │       ├── dropdown-menu.tsx
        │   │       ├── empty.tsx
        │   │       ├── field.tsx
        │   │       ├── form.tsx
        │   │       ├── hover-card.tsx
        │   │       ├── icon-badge.tsx
        │   │       ├── input-group.tsx
        │   │       ├── input-otp.tsx
        │   │       ├── input.tsx
        │   │       ├── item.tsx
        │   │       ├── kbd.tsx
        │   │       ├── label.tsx
        │   │       ├──kdown.tsx
        │   │       ├── menubar.tsx
        │   │       ├── native-select.tsx
        │   │       ├── navigation-menu.tsx
        │   │       ├── pagination.tsx
        │   │       ├── popover.tsx
        │   │       ├── portal-container.ts
        │   │       ├── progress.tsx
        │   │       ├── radio-group.tsx
        │   │       ├── resizable.tsx
        │   │       ├── scroll-area.tsx
        │   │       ├── select.tsx
        │   │       ├── separator.tsx
        │   │       ├── sheet.tsx
        │   │       ├── sidebar.tsx
        │   │       ├── skeleton.tsx
        │   │       ├── slider.tsx
        │   │       ├── sonner.tsx
        │   │       ├── spinner.tsx
        │   │       ├── switch.tsx
        │   │       ├── table.tsx
        │   │       ├── tabs.tsx
        │   │       ├── textarea.tsx
        │   │       ├── titled-card.tsx
        │   │       ├── toggle-group.tsx
        │   │       ├── toggle.tsx
        │   │       └── tooltip.tsx
        │   ├── config/
        │   │   └── fonts.ts
        │   ├── context/
        │   │   ├── direction-provider.tsx
        │   │   ├── font-provider.tsx
        │   │   ├── layout-provider.tsx
        │   │   ├── search-provider.tsx
        │   │   ├── theme-customization-provider.tsx
        │   │   └── theme-provider.tsx
        │   ├── env.d.ts
        │   ├── features/
        │   │   ├── about/
        │   │   │   ├── __tests__/
        │   │   │   │   └── about-page.test.tsx
        │   │   │   ├── api.ts
        │   │   │   ├── index.tsx
        │   │   │   └── types.ts
        │   │   ├── auth/
        │   │   │   ├── api.test.ts
        │   │   │   ├── api.ts
        │   │   │   ├── auth-layout.tsx
        │   │   │   ├── components/
        │   │   │   │   ├── legal-consent.tsx
        │   │   │   │   ├── oauth-callback-screen.tsx
        │   │   │   │   ├── oauth-providers.tsx
        │   │   │   │   ├── telegram-login-dialog.tsx
        │   │   │   │   └── terms-footer.tsx
        │   │   │   ├── constants.ts
        │   │   │   ├── forgot-password/
        │   │   │   │   ├── components/
        │   │   │   │   │   └── forgot-password-form.tsx
        │   │   │   │   └── index.tsx
        │   │   │   ├── hooks/
        │   │   │   │   ├── use-auth-redirect.ts
        │   │   │   │   ├── use-email-verification.ts
        │   │   │   │   ├── use-oauth-login.ts
        │   │   │   │   └── use-turnstile.ts
        │   │   │   ├── index.ts
        │   │   │   ├── lib/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   ├── oauth-callback-mode.test.ts
        │   │   │   │   │   └── oauth-popup.test.ts
        │   │   │   │   ├── auth-redirect.test.ts
        │   │   │   │   ├── auth-redirect.ts
        │   │   │   │   ├── oauth-bind-window.test.ts
        │   │   │   │   ├── oauth-bind-window.ts
        │   │   │   │   ├── oauth-callback-mode.ts
        │   │   │   │   ├── oauth-popup.ts
        │   │   │   │   ├── oauth.ts
        │   │   │   │   ├── password-encryption.ts
        │   │   │   │   ├── storage.ts
        │   │   │   │   ├── telegram-login.test.ts
        │   │   │   │   ├── telegram-login.ts
        │   │   │   │   └── validation.ts
        │   │   │   ├── otp/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   └── login-verification.test.tsx
        │   │   │   │   ├── components/
        │   │   │   │   │   └── otp-form.tsx
        │   │   │   │   └── index.tsx
        │   │   │   ├── passkey/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   └── domain-verification.test.tsx
        │   │   │   │   ├── api.ts
        │   │   │   │   ├── assertion.ts
        │   │   │   │   ├── components/
        │   │   │   │   │   └── passkey-domain-selector.tsx
        │   │   │   │   ├── hooks/
        │   │   │   │   │   └── use-passkey-management.ts
        │   │   │   │   ├── index.ts
        │   │   │   │   └── types.ts
        │   │   │   ├── reset-password-confirm/
        │   │   │   │   └── index.tsx
        │   │   │   ├── secure-verification/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   ├── verification-api.test.ts
        │   │   │   │   │   └── verification-flow.test.tsx
        │   │   │   │   ├── api.ts
        │   │   │   │   ├── components/
        │   │   │   │   │   └── secure-verification-dialog.tsx
        │   │   │   │   ├── hooks/
        │   │   │   │   │   └── use-secure-verification.ts
        │   │   │   │   ├── index.ts
        │   │   │   │   └── types.ts
        │   │   │   ├── sign-in/
        │   │   │   │   ├── components/
        │   │   │   │   │   └── user-auth-form.tsx
        │   │   │   │   └── index.tsx
        │   │   │   ├── sign-up/
        │   │   │   │   ├── components/
        │   │   │   │   │   └── sign-up-form.tsx
        │   │   │   │   └── index.tsx
        │   │   │   └── types.ts
        │   │   ├── channels/
        │   │   │   ├── api.ts
        │   │   │   ├── components/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   ├── channel-configuration.test.tsx
        │   │   │   │   │   ├── channel-quick-options.test.tsx
        │   │   │   │   │   ├── inference-status.test.tsx
        │   │   │   │   │   ├── model-mapping-batch-dialog.test.tsx
        │   │   │   │   │   ├── model-mapping-editor.test.tsx
        │   │   │   │   │   ├── provider-selection.test.tsx
        │   │   │   │   │   ├── responses-websocket-setting.test.tsx
        │   │   │   │   │   └── upstream-model-selection.test.tsx
        │   │   │   │   ├── channel-card.tsx
        │   │   │   │   ├── channel-plugin-extensions.tsx
        │   │   │   │   ├── channel-quick-options.tsx
        │   │   │   │   ├── channel-row-actions-context.ts
        │   │   │   │   ├── channel-type-badge.test.tsx
        │   │   │   │   ├── channel-type-badge.tsx
        │   │   │   │   ├── channels-columns.tsx
        │   │   │   │   ├── channels-dialogs.tsx
        │   │   │   │   ├── channels-primary-buttons.tsx
        │   │   │   │   ├── channels-provider.tsx
        │   │   │   │   ├── channels-table.tsx
        │   │   │   │   ├── data-table-bulk-actions.tsx
        │   │   │   │   ├── data-table-row-actions.tsx
        │   │   │   │   ├── data-table-tag-row-actions.tsx
        │   │   │   │   ├── dialogs/
        │   │   │   │   │   ├── advanced-custom-editor-dialog.tsx
        │   │   │   │   │   ├── balance-query-dialog.tsx
        │   │   │   │   │   ├── channel-test-dialog.tsx
        │   │   │   │   │   ├── codex-usage-dialog.tsx
        │   │   │   │   │   ├── configure-models-dialog.tsx
        │   │   │   │   │   ├── copy-channel-dialog.tsx
        │   │   │   │   │   ├── edit-tag-dialog.tsx
        │   │   │   │   │   ├── fetch-models-dialog.tsx
        │   │   │   │   │   ├── inference-status-dialog.tsx
        │   │   │   │   │   ├── missing-models-confirmation-dialog.tsx
        │   │   │   │   │   ├── multi-key-manage-dialog.tsx
        │   │   │   │   │   ├── multi-key-statistics-card.tsx
        │   │   │   │   │   ├── multi-key-table-row-actions.tsx
        │   │   │   │   │   ├── ollama-models-dialog.tsx
        │   │   │   │   │   ├── param-override-editor-dialog.tsx
        │   │   │   │   │   ├── passthrough-warning-dialog.tsx
        │   │   │   │   │   ├── status-code-risk-dialog.tsx
        │   │   │   │   │   ├── tag-batch-edit-dialog.tsx
        │   │   │   │   │   └── upstream-update-dialog.tsx
        │   │   │   │   ├── drawers/
        │   │   │   │   │   ├── channel-configuration.tsx
        │   │   │   │   │   ├── channel-mutate-drawer.tsx
        │   │   │   │   │   ├── channel-provider-picker.tsx
        │   │   │   │   │   └── sections/
        │   │   │   │   │       ├── channel-advanced-section.tsx
        │   │   │   │   │       ├── channel-api-access-section.tsx
        │   │   │   │   │       ├── channel-auth-section.tsx
        │   │   │   │   │       ├── channel-basic-section.tsx
        │   │   │   │   │       ├── channel-editor-loading-state.tsx
        │   │   │   │   │       ├── channel-models-section.tsx
        │   │   │   │   │       └── index.ts
        │   │   │   │   ├── model-mapping-batch-dialog.tsx
        │   │   │   │   ├── model-mapping-editor.tsx
        │   │   │   │   ├── model-redirect-panel.tsx
        │   │   │   │   ├── numeric-spinner-input.tsx
        │   │   │   │   ├── responses-websocket-setting.tsx
        │   │   │   │   └── upstream-model-selection.tsx
        │   │   │   ├── constants.ts
        │   │   │   ├── hooks/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   ├── channel-key-disclosure.test.tsx
        │   │   │   │   │   └── use-redirect-panel-placement.test.tsx
        │   │   │   │   ├── use-channel-key-disclosure.ts
        │   │   │   │   ├── use-channel-model-discovery.ts
        │   │   │   │   ├── use-channel-mutate-form.ts
        │   │   │   │   ├── use-channel-upstream-updates.ts
        │   │   │   │   └── use-redirect-panel-placement.ts
        │   │   │   ├── index.tsx
        │   │   │   ├── lib/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   ├── channel-field-update.test.ts
        │   │   │   │   │   ├── channel-table-row-id.test.ts
        │   │   │   │   │   ├── channel-type-options.test.ts
        │   │   │   │   │   ├── header-passthrough.test.ts
        │   │   │   │   │   ├── inference-status.test.ts
        │   │   │   │   │   ├── model-mapping-rules.test.ts
        │   │   │   │   │   ├── model-naming-patterns.test.ts
        │   │   │   │   │   ├── new-api-channel.test.ts
        │   │   │   │   │   └── task-plugin-base-url.test.ts
        │   │   │   │   ├── advanced-custom.ts
        │   │   │   │   ├── channel-actions.ts
        │   │   │   │   ├── channel-configuration.ts
        │   │   │   │   ├── channel-field-update.ts
        │   │   │   │   ├── channel-form-errors.ts
        │   │   │   │   ├── channel-form.ts
        │   │   │   │   ├── channel-plugin-extensions.ts
        │   │   │   │   ├── channel-type-config.ts
        │   │   │   │   ├── channel-utils.ts
        │   │   │   │   ├── header-passthrough.ts
        │   │   │   │   ├── index.ts
        │   │   │   │   ├── inference-status.ts
        │   │   │   │   ├── model-categories.ts
        │   │   │   │   ├── model-mapping-rules.ts
        │   │   │   │   ├── model-mapping-validation.ts
        │   │   │   │   ├── model-naming-patterns.ts
        │   │   │   │   ├── multi-key-utils.ts
        │   │   │   │   ├── ollama-utils.ts
        │   │   │   │   ├── responses-websocket.ts
        │   │   │   │   ├── status-code-risk-guard.ts
        │   │   │   │   ├── task-plugin-base-url.ts
        │   │   │   │   └── upstream-update-utils.ts
        │   │   │   └── types.ts
        │   │   ├── chat/
        │   │   │   ├── hooks/
        │   │   │   │   ├── use-active-chat-key.ts
        │   │   │   │   └── use-chat-presets.ts
        │   │   │   └── lib/
        │   │   │       ├── chat-links.ts
        │   │   │       └── send-to-fluent.ts
        │   │   ├── dashboard/
        │   │   │   ├── api.ts
        │   │   │   ├── components/
        │   │   │   │   ├── flow/
        │   │   │   │   │   ├── flow-charts.tsx
        │   │   │   │   │   └── flow-node-filter.tsx
        │   │   │   │   ├── models/
        │   │   │   │   │   ├── consumption-distribution-chart.tsx
        │   │   │   │   │   ├── log-stat-cards.tsx
        │   │   │   │   │   ├── model-charts.tsx
        │   │   │   │   │   ├── models-chart-preferences.tsx
        │   │   │   │   │   ├── models-filter-dialog.tsx
        │   │   │   │   │   └── performance-overview.tsx
        │   │   │   │   ├── overview/
        │   │   │   │   │   ├── __tests__/
        │   │   │   │   │   │   └── setup-guide.test.tsx
        │   │   │   │   │   ├── announcement-detail-dialog.tsx
        │   │   │   │   │   ├── announcements-panel.tsx
        │   │   │   │   │   ├── api-info-item.tsx
        │   │   │   │   │   ├── api-info-panel.tsx
        │   │   │   │   │   ├── faq-panel.tsx
        │   │   │   │   │   ├── overview-dashboard.tsx
        │   │   │   │   │   ├── performance-health-panel.tsx
        │   │   │   │   │   ├── summary-cards.tsx
        │   │   │   │   │   └── uptime-panel.tsx
        │   │   │   │   ├── ui/
        │   │   │   │   │   ├── panel-wrapper.tsx
        │   │   │   │   │   └── stat-card.tsx
        │   │   │   │   └── users/
        │   │   │   │       └── user-charts.tsx
        │   │   │   ├── constants.ts
        │   │   │   ├── hooks/
        │   │   │   │   ├── use-dashboard-config.tsx
        │   │   │   │   └── use-status-data.ts
        │   │   │   ├── index.tsx
        │   │   │   ├── lib/
        │   │   │   │   ├── api-info.ts
        │   │   │   │   ├── charts.ts
        │   │   │   │   ├── filters.ts
        │   │   │   │   ├── flow-selection.test.ts
        │   │   │   │   ├── flow-selection.ts
        │   │   │   │   ├── flow.test.ts
        │   │   │   │   ├── flow.ts
        │   │   │   │   ├── index.ts
        │   │   │   │   ├── stats.ts
        │   │   │   │   └── text.ts
        │   │   │   ├── section-registry.tsx
        │   │   │   └── types.ts
        │   │   ├── errors/
        │   │   │   ├── forbidden.tsx
        │   │   │   ├── general-error.tsx
        │   │   │   ├── maintenance-error.tsx
        │   │   │   ├── not-found-error.tsx
        │   │   │   └── unauthorized-error.tsx
        │   │   ├── home/
        │   │   │   ├── api.ts
        │   │   │   ├── components/
        │   │   │   │   ├── connection-line.tsx
        │   │   │   │   ├── feature-item.tsx
        │   │   │   │   ├── gateway-card.tsx
        │   │   │   │   ├── hero-buttons.tsx
        │   │   │   │   ├── hero-terminal-demo.tsx
        │   │   │   │   ├── icon-card.tsx
        │   │   │   │   ├── index.ts
        │   │   │   │   ├── scrolling-icons.tsx
        │   │   │   │   ├── sections/
        │   │   │   │   │   ├── cta.tsx
        │   │   │   │   │   ├── features.tsx
        │   │   │   │   │   ├── hero.tsx
        │   │   │   │   │   ├── how-it-works.tsx
        │   │   │   │   │   └── stats.tsx
        │   │   │   │   └── stat-item.tsx
        │   │   │   ├── constants.ts
        │   │   │   ├── hooks/
        │   │   │   │   ├── index.ts
        │   │   │   │   └── use-home-page-content.ts
        │   │   │   ├── index.tsx
        │   │   │   ├── lib/
        │   │   │   │   └── icon-mapper.tsx
        │   │   │   └── types.ts
        │   │   ├── groups/
        │   │   │   ├── __tests__/
        │   │   │   │   ├── group-user-transfer.test.tsx
        │   │   │   │   └── group-users.test.ts
        │   │   │   ├── api.ts
        │   │   │   ├── components/
        │   │   │   │   ├── group-settings-panel.tsx
        │   │   │   │   ├── group-user-transfer.tsx
        │   │   │   │   └── group-users-panel.tsx
        │   │   │   ├── index.tsx
        │   │   │   ├── lib/
        │   │   │   │   └── group-users.ts
        │   │   │   └── types.ts
        │   │   ├── keys/
        │   │   │   ├── api.ts
        │   │   │   ├── components/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   ├── api-key-group-cell.test.tsx
        │   │   │   │   │   ├── api-key-group-combobox.test.tsx
        │   │   │   │   │   ├── api-key-listing.test.tsx
        │   │   │   │   │   ├── api-keys-mutate-drawer.test.tsx
        │   │   │   │   │   └── auto-group-order-editor.test.tsx
        │   │   │   │   ├── api-key-group-cell.tsx
        │   │   │   │   ├── api-key-group-combobox.tsx
        │   │   │   │   ├── api-key-quota-cell.tsx
        │   │   │   │   ├── api-key-timestamp-cell.tsx
        │   │   │   │   ├── api-keys-cells.tsx
        │   │   │   │   ├── api-keys-columns.tsx
        │   │   │   │   ├── api-keys-delete-dialog.tsx
        │   │   │   │   ├── api-keys-dialogs.tsx
        │   │   │   │   ├── api-keys-multi-delete-dialog.tsx
        │   │   │   │   ├── api-keys-mutate-drawer.tsx
        │   │   │   │   ├── api-keys-primary-buttons.tsx
        │   │   │   │   ├── api-keys-provider.tsx
        │   │   │   │   ├── api-keys-table.tsx
        │   │   │   │   ├── auto-group-order-editor.tsx
        │   │   │   │   ├── auto-group-visuals.tsx
        │   │   │   │   ├── data-table-bulk-actions.tsx
        │   │   │   │   ├── data-table-row-actions.tsx
        │   │   │   │   └── dialogs/
        │   │   │   │       └── cc-switch-dialog.tsx
        │   │   │   ├── constants.ts
        │   │   │   ├── index.tsx
        │   │   │   ├── lib/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   └── auto-group-form.test.ts
        │   │   │   │   ├── api-key-form.ts
        │   │   │   │   └── index.ts
        │   │   │   └── types.ts
        │   │   ├── legal/
        │   │   │   ├── api.ts
        │   │   │   ├── index.ts
        │   │   │   ├── legal-document.tsx
        │   │   │   ├── privacy-policy.tsx
        │   │   │   ├── types.ts
        │   │   │   └── user-agreement.tsx
        │   │   ├── model-pricing/
        │   │   │   ├── __tests__/
        │   │   │   │   ├── editor-currency.test.tsx
        │   │   │   │   ├── editor-layout.test.tsx
        │   │   │   │   ├── pricing.test.ts
        │   │   │   │   └── save-errors.test.tsx
        │   │   │   ├── api.ts
        │   │   │   ├── currency.ts
        │   │   │   ├── model-pricing-panel.tsx
        │   │   │   ├── pricing-amount-input.tsx
        │   │   │   ├── pricing-conversion-dialog.tsx
        │   │   │   ├── pricing-currency-selector.tsx
        │   │   │   └── pricing.ts
        │   │   ├── models/
        │   │   │   ├── __tests__/
        │   │   │   │   ├── metadata-editing.test.tsx
        │   │   │   │   ├── metadata-sync.test.tsx
        │   │   │   │   ├── model-deletion.test.tsx
        │   │   │   │   ├── model-listing.test.tsx
        │   │   │   │   └── vendor-management.test.tsx
        │   │   │   ├── api.ts
        │   │   │   ├── components/
        │   │   │   │   ├── data-table-bulk-actions.tsx
        │   │   │   │   ├── data-table-row-actions.tsx
        │   │   │   │   ├── deployment-access-guard.tsx
        │   │   │   │   ├── deployments-columns.tsx
        │   │   │   │   ├── deployments-table.tsx
        │   │   │   │   ├── description-cell.tsx
        │   │   │   │   ├── dialogs/
        │   │   │   │   │   ├── create-deployment-drawer.tsx
        │   │   │   │   │   ├── description-dialog.tsx
        │   │   │   │   │   ├── extend-deployment-dialog.tsx
        │   │   │   │   │   ├── missing-models-dialog.tsx
        │   │   │   │   │   ├── model-delete-dialog.tsx
        │   │   │   │   │   ├── prefill-group-management-dialog.tsx
        │   │   │   │   │   ├── prefill-group-management.tsx
        │   │   │   │   │   ├── price-sync-dialog.tsx
        │   │   │   │   │   ├── rename-deployment-dialog.tsx
        │   │   │   │   │   ├── sync-wizard-dialog.tsx
        │   │   │   │   │   ├── update-config-dialog.tsx
        │   │   │   │   │   ├── upstream-conflict-dialog.tsx
        │   │   │   │   │   ├── vendor-mutate-dialog.tsx
        │   │   │   │   │   ├── vendor-operation-dialog.tsx
        │   │   │   │   │   ├── vendors-management-dialog.tsx
        │   │   │   │   │   ├── view-details-dialog.tsx
        │   │   │   │   │   └── view-logs-dialog.tsx
        │   │   │   │   ├── drawers/
        │   │   │   │   │   ├── model-mutate-drawer.tsx
        │   │   │   │   │   └── prefill-group-form-drawer.tsx
        │   │   │   │   ├── model-connections.tsx
        │   │   │   │   ├── model-square-status.tsx
        │   │   │   │   ├── models-columns.tsx
        │   │   │   │   ├── models-dialogs.tsx
        │   │   │   │   ├── models-primary-buttons.tsx
        │   │   │   │   ├── models-provider.tsx
        │   │   │   │   ├── models-table.tsx
        │   │   │   │   ├── prefill-group-shared.ts
        │   │   │   │   ├── vendor-linked-models.tsx
        │   │   │   │   └── vendors-table.tsx
        │   │   │   ├── constants.ts
        │   │   │   ├── hooks/
        │   │   │   │   └── use-model-deployment-settings.ts
        │   │   │   ├── index.tsx
        │   │   │   ├── lib/
        │   │   │   │   ├── deployments-utils.ts
        │   │   │   │   ├── index.ts
        │   │   │   │   ├── model-actions.ts
        │   │   │   │   ├── model-form.ts
        │   │   │   │   ├── model-utils.ts
        │   │   │   │   ├── query-keys.ts
        │   │   │   │   └── vendor-actions.ts
        │   │   │   ├── section-registry.tsx
        │   │   │   ├── types.ts
        │   │   │   └── vendor-api.ts
        │   │   ├── performance-metrics/
        │   │   │   ├── __tests__/
        │   │   │   │   └── summary.test.tsx
        │   │   │   ├── api.ts
        │   │   │   ├── lib/
        │   │   │   │   └── format.ts
        │   │   │   └── types.ts
        │   │   ├── playground/
        │   │   │   ├── api.ts
        │   │   │   ├── components/
        │   │   │   │   ├── chat/
        │   │   │   │   │   ├── playground-chat.tsx
        │   │   │   │   │   └── playground-empty-state.tsx
        │   │   │   │   ├── input/
        │   │   │   │   │   ├── playground-input-controls.tsx
        │   │   │   │   │   ├── playground-input-tools.tsx
        │   │   │   │   │   ├── playground-input.tsx
        │   │   │   │   │   └── playground-parameter-panel.tsx
        │   │   │   │   ├── message/
        │   │   │   │   │   ├── __tests__/
        │   │   │   │   │   │   └── playground-message-editor.test.tsx
        │   │   │   │   │   ├── message-action-button.tsx
        │   │   │   │   │   ├── message-actions.tsx
        │   │   │   │   │   ├── message-error-actions.tsx
        │   │   │   │   │   ├── message-error.tsx
        │   │   │   │   │   ├── message-metadata.tsx
        │   │   │   │   │   ├── playground-message-content.tsx
        │   │   │   │   │   └── playground-message-editor.tsx
        │   │   │   │   ├── message-action-button.tsx
        │   │   │   │   ├── message-actions.tsx
        │   │   │   │   ├── message-error.tsx
        │   │   │   │   ├── playground-chat.tsx
        │   │   │   │   └── playground-input.tsx
        │   │   │   ├── constants.ts
        │   │   │   ├── hooks/
        │   │   │   │   ├── index.ts
        │   │   │   │   ├── use-chat-handler.ts
        │   │   │   │   ├── use-message-action-guard.ts
        │   │   │   │   ├── use-playground-conversation.ts
        │   │   │   │   ├── use-playground-options.ts
        │   │   │   │   ├── use-playground-state.ts
        │   │   │   │   ├── use-stream-request.test.ts
        │   │   │   │   └── use-stream-request.ts
        │   │   │   ├── index.tsx
        │   │   │   ├── lib/
        │   │   │   │   ├── index.ts
        │   │   │   │   ├── input/
        │   │   │   │   │   ├── input-control-utils.ts
        │   │   │   │   │   └── input-tool-utils.ts
        │   │   │   │   ├── message/
        │   │   │   │   │   ├── conversation-message-utils.ts
        │   │   │   │   │   ├── message-action-utils.ts
        │   │   │   │   │   ├── message-content-utils.ts
        │   │   │   │   │   ├── message-editor-utils.ts
        │   │   │   │   │   ├── message-error-utils.ts
        │   │   │   │   │   ├── message-layout-utils.ts
        │   │   │   │   │   ├── message-reasoning-utils.ts
        │   │   │   │   │   ├── message-streaming-utils.ts
        │   │   │   │   │   ├── message-styles.ts
        │   │   │   │   │   ├── message-timing-utils.ts
        │   │   │   │   │   ├── message-update-utils.ts
        │   │   │   │   │   └── message-utils.ts
        │   │   │   │   ├── message-styles.ts
        │   │   │   │   ├── message-utils.ts
        │   │   │   │   ├── options/
        │   │   │   │   │   └── playground-option-utils.ts
        │   │   │   │   ├── parameters/
        │   │   │   │   │   └── playground-parameters.ts
        │   │   │   │   ├── payload-builder.ts
        │   │   │   │   ├── state/
        │   │   │   │   │   └── playground-state-utils.ts
        │   │   │   │   ├── storage/
        │   │   │   │   │   ├── storage-schema.ts
        │   │   │   │   │   └── storage.ts
        │   │   │   │   ├── storage.ts
        │   │   │   │   └── streaming/
        │   │   │   │       ├── payload-builder.ts
        │   │   │   │       ├── request-error-utils.ts
        │   │   │   │       └── stream-utils.ts
        │   │   │   └── types.ts
        │   │   ├── pricing/
        │   │   │   ├── __tests__/
        │   │   │   │   ├── breakdown-tier-match.test.ts
        │   │   │   │   ├── dynamic-price.test.ts
        │   │   │   │   ├── model-cards.test.tsx
        │   │   │   │   ├── pricing-controls.test.tsx
        │   │   │   │   ├── task-expr.test.ts
        │   │   │   │   ├── task-matrix-display.test.ts
        │   │   │   │   ├── task-matrix.test.ts
        │   │   │   │   └── task-price-display.test.tsx
        │   │   │   ├── api.ts
        │   │   │   ├── components/
        │   │   │   │   ├── cached-price-cell.tsx
        │   │   │   │   ├── dynamic-pricing-breakdown.tsx
        │   │   │   │   ├── empty-state.tsx
        │   │   │   │   ├── index.ts
        │   │   │   │   ├── loading-skeleton.tsx
        │   │   │   │   ├── model-billing-mode-badge.tsx
        │   │   │   │   ├── model-card-grid.tsx
        │   │   │   │   ├── model-card.tsx
        │   │   │   │   ├── model-details-api.tsx
        │   │   │   │   ├── model-details-apps.tsx
        │   │   │   │   ├── model-details-charts.tsx
        │   │   │   │   ├── model-details-performance.tsx
        │   │   │   │   ├── model-details-uptime-sparkline.tsx
        │   │   │   │   ├── model-details.tsx
        │   │   │   │   ├── model-perf-badge.tsx
        │   │   │   │   ├── model-price-cell.tsx
        │   │   │   │   ├── pricing-columns.tsx
        │   │   │   │   ├── pricing-sidebar.tsx
        │   │   │   │   ├── pricing-table.tsx
        │   │   │   │   ├── pricing-toolbar.tsx
        │   │   │   │   └── search-bar.tsx
        │   │   │   ├── constants.ts
        │   │   │   ├── hooks/
        │   │   │   │   ├── index.ts
        │   │   │   │   ├── use-billing-time.ts
        │   │   │   │   ├── use-filters.ts
        │   │   │   │   └── use-pricing-data.ts
        │   │   │   ├── index.tsx
        │   │   │   ├── lib/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   ├── billing-expression.test.ts
        │   │   │   │   │   └── time-rule-expr.test.ts
        │   │   │   │   ├── billing-expr.ts
        │   │   │   │   ├── billing-expression/
        │   │   │   │   │   ├── README.md
        │   │   │   │   │   ├── condition-display.ts
        │   │   │   │   │   ├── display.ts
        │   │   │   │   │   ├── parser.ts
        │   │   │   │   │   ├── runtime.ts
        │   │   │   │   │   ├── structure.ts
        │   │   │   │   │   ├── types.ts
        │   │   │   │   │   └── visual.ts
        │   │   │   │   ├── billing-mode.ts
        │   │   │   │   ├── breakdown-tier-match.ts
        │   │   │   │   ├── dynamic-price.ts
        │   │   │   │   ├── filters.ts
        │   │   │   │   ├── index.ts
        │   │   │   │   ├── mock-stats.ts
        │   │   │   │   ├── model-helpers.ts
        │   │   │   │   ├── plugin-pricing.ts
        │   │   │   │   ├── price.ts
        │   │   │   │   ├── seed.ts
        │   │   │   │   ├── task-expr.ts
        │   │   │   │   ├── task-matrix-display.ts
        │   │   │   │   ├── task-price-display.ts
        │   │   │   │   └── tier-expr.ts
        │   │   │   └── types.ts
        │   │   ├── profile/
        │   │   │   ├── __tests__/
        │   │   │   │   └── settings.test.tsx
        │   │   │   ├── api.ts
        │   │   │   ├── components/
        │   │   │   │   ├── checkin-calendar-card.tsx
        │   │   │   │   ├── language-preferences-card.tsx
        │   │   │   │   ├── profile-header.tsx
        │   │   │   │   ├── profile-settings-card.tsx
        │   │   │   │   ├── sidebar-modules-card.tsx
        │   │   │   │   └── tabs/
        │   │   │   │       └── notification-tab.tsx
        │   │   │   ├── constants.ts
        │   │   │   ├── hooks/
        │   │   │   │   ├── index.ts
        │   │   │   │   └── use-profile.ts
        │   │   │   ├── index.tsx
        │   │   │   ├── lib/
        │   │   │   │   ├── format.ts
        │   │   │   │   ├── index.ts
        │   │   │   │   └── user-settings.ts
        │   │   │   └── types.ts
        │   │   ├── rankings/
        │   │   │   ├── api.ts
        │   │   │   ├── components/
        │   │   │   │   ├── entity-links.tsx
        │   │   │   │   ├── growth-text.tsx
        │   │   │   │   ├── index.ts
        │   │   │   │   ├──ket-share-section.tsx
        │   │   │   │   ├── model-leaderboard.tsx
        │   │   │   │   ├── models-section.tsx
        │   │   │   │   ├── pulse-section.tsx
        │   │   │   │   └── rankings-hero.tsx
        │   │   │   ├── hooks/
        │   │   │   │   └── use-rankings.ts
        │   │   │   ├── index.tsx
        │   │   │   ├── lib/
        │   │   │   │   ├── format.ts
        │   │   │   │   └── index.ts
        │   │   │   └── types.ts
        │   │   ├── redemption-codes/
        │   │   │   ├── api.ts
        │   │   │   ├── components/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   ├── bulk-delete.test.tsx
        │   │   │   │   │   ├── redemptions-export.test.tsx
        │   │   │   │   │   └── redemptions-mutate-drawer.test.tsx
        │   │   │   │   ├── data-table-bulk-actions.tsx
        │   │   │   │   ├── data-table-row-actions.tsx
        │   │   │   │   ├── redemptions-columns.tsx
        │   │   │   │   ├── redemptions-delete-dialog.tsx
        │   │   │   │   ├── redemptions-dialogs.tsx
        │   │   │   │   ├── redemptions-export-dialog.tsx
        │   │   │   │   ├── redemptions-mobile-list.tsx
        │   │   │   │   ├── redemptions-mutate-drawer.tsx
        │   │   │   │   ├── redemptions-primary-buttons.tsx
        │   │   │   │   ├── redemptions-provider.tsx
        │   │   │   │   └── redemptions-table.tsx
        │   │   │   ├── constants.ts
        │   │   │   ├── index.tsx
        │   │   │   ├── lib/
        │   │   │   │   ├── index.ts
        │   │   │   │   ├── redemption-form.ts
        │   │   │   │   └── utils.ts
        │   │   │   └── types.ts
        │   │   ├── security/
        │   │   │   ├── __tests__/
        │   │   │   │   ├── account-security.test.tsx
        │   │   │   │   ├── enrollment.test.tsx
        │   │   │   │   └── page.test.tsx
        │   │   │   ├── api.ts
        │   │   │   ├── components/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   ├── access-token-card.test.tsx
        │   │   │   │   │   ├── login-session-utils.test.ts
        │   │   │   │   │   └── privacy-card.test.tsx
        │   │   │   │   ├── access-token-card.tsx
        │   │   │   │   ├── account-action-card.tsx
        │   │   │   │   ├── account-bindings.tsx
        │   │   │   │   ├── dialogs/
        │   │   │   │   │   ├── access-token-dialog.tsx
        │   │   │   │   │   ├── change-password-dialog.tsx
        │   │   │   │   │   ├── delete-account-dialog.tsx
        │   │   │   │   │   ├── email-bind-dialog.tsx
        │   │   │   │   │   ├── telegram-bind-dialog.tsx
        │   │   │   │   │   ├── two-fa-backup-dialog.tsx
        │   │   │   │   │   ├── two-fa-disable-dialog.tsx
        │   │   │   │   │   ├── two-fa-setup-dialog.tsx
        │   │   │   │   │   └── wechat-bind-dialog.tsx
        │   │   │   │   ├── login-session-dialogs.tsx
        │   │   │   │   ├── login-session-item.tsx
        │   │   │   │   ├── login-session-utils.ts
        │   │   │   │   ├── login-sessions-card.tsx
        │   │   │   │   ├── passkey-card.tsx
        │   │   │   │   ├── privacy-card.tsx
        │   │   │   │   └── two-fa-card.tsx
        │   │   │   ├── hooks/
        │   │   │   │   ├── use-access-token.ts
        │   │   │   │   ├── use-account-security.ts
        │   │   │   │   ├── use-two-fa-setup.ts
        │   │   │   │   └── use-two-fa.ts
        │   │   │   └── index.tsx
        │   │   ├── setup/
        │   │   │   ├── api.ts
        │   │   │   ├── components/
        │   │   │   │   ├── admin-step.tsx
        │   │   │   │   ├── complete-step.tsx
        │   │   │   │   ├── database-step.tsx
        │   │   │   │   ├── step-navigation.tsx
        │   │   │   │   └── usage-mode-step.tsx
        │   │   │   ├── index.ts
        │   │   │   ├── setup-wizard.tsx
        │   │   │   └── types.ts
        │   │   ├── subscriptions/
        │   │   │   ├── api.ts
        │   │   │   ├── components/
        │   │   │   │   ├── data-table-row-actions.tsx
        │   │   │   │   ├── dialogs/
        │   │   │   │   │   ├── reset-subscriptions-dialog.tsx
        │   │   │   │   │   ├── subscription-purchase-dialog.tsx
        │   │   │   │   │   ├── toggle-status-dialog.tsx
        │   │   │   │   │   └── user-subscriptions-dialog.tsx
        │   │   │   │   ├── subscriptions-columns.tsx
        │   │   │   │   ├── subscriptions-dialogs.tsx
        │   │   │   │   ├── subscriptions-mutate-drawer.tsx
        │   │   │   │   ├── subscriptions-primary-buttons.tsx
        │   │   │   │   ├── subscriptions-provider.tsx
        │   │   │   │   └── subscriptions-table.tsx
        │   │   │   ├── constants.ts
        │   │   │   ├── index.tsx
        │   │   │   ├── lib/
        │   │   │   │   ├── format.ts
        │   │   │   │   ├── index.ts
        │   │   │   │   └── plan-form.ts
        │   │   │   └── types.ts
        │   │   ├── system-info/
        │   │   │   ├── __tests__/
        │   │   │   │   └── system-tasks-panel.test.tsx
        │   │   │   ├── api.ts
        │   │   │   ├── components/
        │   │   │   │   ├── system-instances-panel.tsx
        │   │   │   │   ├── system-task-history.tsx
        │   │   │   │   ├── system-tasks-panel.tsx
        │   │   │   │   └── system-tasks-table.tsx
        │   │   │   ├── constants.ts
        │   │   │   ├── index.tsx
        │   │   │   └── types.ts
        │   │   ├── system-settings/
        │   │   │   ├── __tests__/
        │   │   │   │   ├── pricing-sync.test.tsx
        │   │   │   │   └── task-public-address.test.ts
        │   │   │   ├── api.ts
        │   │   │   ├── auth/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   └── passkey-settings.test.tsx
        │   │   │   │   ├── basic-auth-section.tsx
        │   │   │   │   ├── bot-protection-section.tsx
        │   │   │   │   ├── cas-section.tsx
        │   │   │   │   ├── custom-oauth/
        │   │   │   │   │   ├── api.ts
        │   │   │   │   │   ├── components/
        │   │   │   │   │   │   ├── access-policy-templates.ts
        │   │   │   │   │   │   ├── discovery-button.tsx
        │   │   │   │   │   │   ├── preset-selector.tsx
        │   │   │   │   │   │   ├── provider-form-dialog.tsx
        │   │   │   │   │   │   └── provider-table.tsx
        │   │   │   │   │   ├── custom-oauth-section.tsx
        │   │   │   │   │   ├── hooks/
        │   │   │   │   │   │   ├── use-custom-oauth-mutations.ts
        │   │   │   │   │   │   └── use-custom-oauth-providers.ts
        │   │   │   │   │   └── types.ts
        │   │   │   │   ├── group-assignment-rules-editor.tsx
        │   │   │   │   ├── group-assignment-rules.ts
        │   │   │   │   ├── index.tsx
        │   │   │   │   ├── ldap-section.tsx
        │   │   │   │   ├── oauth-callback-url.ts
        │   │   │   │   ├── oauth-section.tsx
        │   │   │   │   ├── passkey-section.tsx
        │   │   │   │   └── section-registry.tsx
        │   │   │   ├── billing/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   └── group-channel-binding.test.tsx
        │   │   │   │   ├── group-channel-binding-section.tsx
        │   │   │   │   ├── index.tsx
        │   │   │   │   └── section-registry.tsx
        │   │   │   ├── components/
        │   │   │   │   ├── form-dirty-indicator.tsx
        │   │   │   │   ├── form-navigation-guard.tsx
        │   │   │   │   ├── settings-accordion.tsx
        │   │   │   │   ├── settings-card.tsx
        │   │   │   │   ├── settings-form-layout.tsx
        │   │   │   │   ├── settings-page-context.tsx
        │   │   │   │   ├── settings-page.tsx
        │   │   │   │   └── settings-section.tsx
        │   │   │   ├── content/
        │   │   │   │   ├── announcements-section.tsx
        │   │   │   │   ├── api-info-section.tsx
        │   │   │   │   ├── chat-dialog.tsx
        │   │   │   │   ├── chat-settings-section.tsx
        │   │   │   │   ├── chat-settings-visual-editor.tsx
        │   │   │   │   ├── dashboard-section.tsx
        │   │   │   │   ├── drawing-settings-section.tsx
        │   │   │   │   ├── faq-section.tsx
        │   │   │   │   ├── index.tsx
        │   │   │   │   ├── json-toggle-section.tsx
        │   │   │   │   ├── section-registry.tsx
        │   │   │   │   ├── uptime-kuma-section.tsx
        │   │   │   │   └── utils.ts
        │   │   │   ├── general/
        │   │   │   │   ├── channel-affinity/
        │   │   │   │   │   ├── api.ts
        │   │   │   │   │   ├── cache-stats-dialog.tsx
        │   │   │   │   │   ├── constants.ts
        │   │   │   │   │   ├── index.tsx
        │   │   │   │   │   ├── rule-editor-dialog.tsx
        │   │   │   │   │   ├── session-rules-table.tsx
        │   │   │   │   │   └── types.ts
        │   │   │   │   ├── checkin-settings-section.tsx
        │   │   │   │   ├── pricing-section.tsx
        │   │   │   │   ├── quota-settings-section.tsx
        │   │   │   │   ├── system-behavior-section.tsx
        │   │   │   │   ├── system-info-section.tsx
        │   │   │   │   └── task-public-address.ts
        │   │   │   ├── hooks/
        │   │   │   │   ├── use-accordion-state.ts
        │   │   │   │   ├── use-form-dirty-guard.ts
        │   │   │   │   ├── use-reset-form.ts
        │   │   │   │   ├── use-safe-json-state.ts
        │   │   │   │   ├── use-settings-form.ts
        │   │   │   │   ├── use-system-options.ts
        │   │   │   │   └── use-update-option.ts
        │   │   │   ├── index.tsx
        │   │   │   ├── integrations/
        │   │   │   │   ├── amount-discount-dialog.tsx
        │   │   │   │   ├── amount-discount-visual-editor.tsx
        │   │   │   │   ├── amount-options-visual-editor.tsx
        │   │   │   │   ├── creem-product-dialog.tsx
        │   │   │   │   ├── creem-products-visual-editor.tsx
        │   │   │   │   ├── email-settings-section.tsx
        │   │   │   │   ├── ionet-deployment-settings-section.tsx
        │   │   │   │   ├── monitoring-settings-section.tsx
        │   │   │   │   ├── payment-method-dialog.tsx
        │   │   │   │   ├── payment-methods-visual-editor.tsx
        │   │   │   │   ├── payment-settings-section.tsx
        │   │   │   │   ├── utils.ts
        │   │   │   │   ├── waffo-pancake-api.ts
        │   │   │   │   ├── waffo-pancake-settings-section.tsx
        │   │   │   │   ├── waffo-settings-section.tsx
        │   │   │   │   └── worker-settings-section.tsx
        │   │   │   ├── maintenance/
        │   │   │   │   ├── config.ts
        │   │   │   │   ├── header-navigation-section.tsx
        │   │   │   │   ├── log-settings-section.tsx
        │   │   │   │   ├── notice-section.tsx
        │   │   │   │   ├── performance-section.tsx
        │   │   │   │   ├── sidebar-modules-section.tsx
        │   │   │   │   ├── syslog-settings-section.tsx
        │   │   │   │   └── update-checker-section.tsx
        │   │   │   ├── models/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   ├── group-auto-limit-validation.test.ts
        │   │   │   │   │   ├── plugin-pricing.test.tsx
        │   │   │   │   │   ├── request-simulation.test.tsx
        │   │   │   │   │   ├── task-pricing-copy.test.tsx
        │   │   │   │   │   ├── tool-price-validation.test.tsx
        │   │   │   │   │   └── visual-billing-editor.test.tsx
        │   │   │   │   ├── billing-time-fields.tsx
        │   │   │   │   ├── channel-selector-dialog.tsx
        │   │   │   │   ├── claude-settings-card.tsx
        │   │   │   │   ├── conflict-confirm-dialog.tsx
        │   │   │   │   ├── constants.ts
        │   │   │   │   ├── draft-number-input.tsx
        │   │   │   │   ├── gemini-settings-card.tsx
        │   │   │   │   ├── global-settings-card.tsx
        │   │   │   │   ├── grok-settings-card.tsx
        │   │   │   │   ├── group-ratio-form.tsx
        │   │   │   │   ├── group-ratio-visual-editor.tsx
        │   │   │   │   ├── group-special-usable-editor.tsx
        │   │   │   │   ├── index.tsx
        │   │   │   │   ├── model-pricing-core.ts
        │   │   │   │   ├── model-pricing-inputs.tsx
        │   │   │   │   ├── model-pricing-sheet.tsx
        │   │   │   │   ├── model-pricing-snapshots.ts
        │   │   │   │   ├── model-ratio-form.tsx
        │   │   │   │   ├── model-ratio-table-columns.tsx
        │   │   │   │   ├── model-ratio-visual-editor.tsx
        │   │   │   │   ├── pricing-format.ts
        │   │   │   │   ├── ratio-settings-card.tsx
        │   │   │   │   ├── request-simulation.tsx
        │   │   │   │   ├── section-registry.tsx
        │   │   │   │   ├── task-plugin-pricing-editor.tsx
        │   │   │   │   ├── task-pricing-matrix.tsx
        │   │   │   │   ├── task-usage-pricing-editor.tsx
        │   │   │   │   ├── tier-price-fields.tsx
        │   │   │   │   ├── tiered-pricing-editor.tsx
        │   │   │   │   ├── tool-price-settings.tsx
        │   │   │   │   ├── upstream-price-cells.tsx
        │   │   │   │   ├── upstream-ratio-sync-columns.tsx
        │   │   │   │   ├── upstream-ratio-sync-helpers.ts
        │   │   │   │   ├── upstream-ratio-sync-table.tsx
        │   │   │   │   ├── upstream-ratio-sync.tsx
        │   │   │   │   ├── utils.ts
        │   │   │   │   ├── visual-billing-document-editor.tsx
        │   │   │   │   ├── visual-condition-tree.css
        │   │   │   │   └── visual-condition-tree.tsx
        │   │   │   ├── operations/
        │   │   │   │   ├── index.tsx
        │   │   │   │   └── section-registry.tsx
        │   │   │   ├── request-limits/
        │   │   │   │   ├── rate-limit-dialog.tsx
        │   │   │   │   ├── rate-limit-section.tsx
        │   │   │   │   ├── rate-limit-visual-editor.tsx
        │   │   │   │   ├── ssrf-section.tsx
        │   │   │   │   └── token-limit-section.tsx
        │   │   │   ├── request-policies/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   ├── channel-health-layout.test.tsx
        │   │   │   │   │   ├── routing.test.tsx
        │   │   │   │   │   └── settings.test.tsx
        │   │   │   │   ├── api.ts
        │   │   │   │   ├── channel-health-section.tsx
        │   │   │   │   ├── decision-record.tsx
        │   │   │   │   ├── defaults.ts
        │   │   │   │   ├── index.tsx
        │   │   │   │   ├── policy-label.ts
        │   │   │   │   ├── related-policy-link.tsx
        │   │   │   │   ├── request-checks-section.tsx
        │   │   │   │   ├── retry-section.tsx
        │   │   │   │   ├── routing-form.ts
        │   │   │   │   ├── routing-section.tsx
        │   │   │   │   ├── section-registry.tsx
        │   │   │   │   └── use-save-policy.ts
        │   │   │   ├── security/
        │   │   │   │   ├── index.tsx
        │   │   │   │   └── section-registry.tsx
        │   │   │   ├── site/
        │   │   │   │   ├── index.tsx
        │   │   │   │   └── section-registry.tsx
        │   │   │   ├── types.ts
        │   │   │   └── utils/
        │   │   │       ├── json-parser.ts
        │   │   │       ├── json-validators.ts
        │   │   │       ├── numeric-field.ts
        │   │   │       ├── route-config.ts
        │   │   │       └── section-registry.ts
        │   │   ├── system-update/
        │   │   │   ├── __tests__/
        │   │   │   │   ├── releases.test.ts
        │   │   │   │   └── update-checking.test.tsx
        │   │   │   ├── api.ts
        │   │   │   ├── releases.ts
        │   │   │   ├── store.ts
        │   │   │   ├── system-update-action.tsx
        │   │   │   ├── system-update-dialog.tsx
        │   │   │   └── use-system-update.ts
        │   │   ├── task-plugins/
        │   │   │   ├── __tests__/
        │   │   │   │   ├── enabled-option.test.ts
        │   │   │   │   ├──ketplace-install-dialog.test.tsx
        │   │   │   │   ├──ketplace-panel.test.tsx
        │   │   │   │   ├──ketplace-plugin-logo.test.tsx
        │   │   │   │   ├──ketplace.test.ts
        │   │   │   │   ├── plugin-card.test.tsx
        │   │   │   │   ├── plugin-changelog-panel.test.tsx
        │   │   │   │   ├── plugin-changelog.test.ts
        │   │   │   │   ├── plugin-detail-sheet.test.tsx
        │   │   │   │   ├── plugin-icon-file.test.ts
        │   │   │   │   ├── plugin-icon-image.test.tsx
        │   │   │   │   ├── plugin-icon.test.ts
        │   │   │   │   ├── plugin-meta-preview.test.ts
        │   │   │   │   ├── plugin-url.test.ts
        │   │   │   │   ├── plugins-table.test.tsx
        │   │   │   │   ├── upload-dialog.test.tsx
        │   │   │   │   ├── usage-schema-table.test.tsx
        │   │   │   │   └── website.test.tsx
        │   │   │   ├── api.ts
        │   │   │   ├── components/
        │   │   │   │   ├── javascript-viewer.tsx
        │   │   │   │   ├──ketplace-capabilities.tsx
        │   │   │   │   ├──ketplace-install-dialog.tsx
        │   │   │   │   ├──ketplace-panel.tsx
        │   │   │   │   ├──ketplace-plugin-card.tsx
        │   │   │   │   ├──ketplace-sources-dialog.tsx
        │   │   │   │   ├── plugin-card.tsx
        │   │   │   │   ├── plugin-changelog-panel.tsx
        │   │   │   │   ├── plugin-detail-sheet.tsx
        │   │   │   │   ├── plugin-endpoints.tsx
        │   │   │   │   ├── plugin-icon.tsx
        │   │   │   │   ├── plugin-integrity-check.tsx
        │   │   │   │   ├── plugin-metadata-card.tsx
        │   │   │   │   ├── plugin-model-list.tsx
        │   │   │   │   ├── plugin-sandbox.tsx
        │   │   │   │   ├── plugin-source-picker.tsx
        │   │   │   │   ├── plugin-url-import-field.tsx
        │   │   │   │   ├── plugin-website-link.tsx
        │   │   │   │   ├── plugins-table.tsx
        │   │   │   │   ├── source-diff.tsx
        │   │   │   │   ├── upload-dialog.tsx
        │   │   │   │   └── usage-schema-table.tsx
        │   │   │   ├── index.tsx
        │   │   │   ├── lib/
        │   │   │   │   ├── host-protocols.ts
        │   │   │   │   ├──ketplace.ts
        │   │   │   │   ├── plugin-changelog.ts
        │   │   │   │   ├── plugin-icon-file.ts
        │   │   │   │   ├── plugin-icon.ts
        │   │   │   │   ├── plugin-meta-preview.ts
        │   │   │   │   ├── plugin-url.ts
        │   │   │   │   └── plugin-website.ts
        │   │   │   └── types.ts
        │   │   ├── usage-logs/
        │   │   │   ├── __tests__/
        │   │   │   │   ├── access.test.ts
        │   │   │   │   ├── artifacts.test.ts
        │   │   │   │   ├── mobile-layout.test.ts
        │   │   │   │   └── task-details.test.ts
        │   │   │   ├── api.ts
        │   │   │   ├── audit/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   ├── details.test.tsx
        │   │   │   │   │   └── viewer.test.tsx
        │   │   │   │   ├── api.ts
        │   │   │   │   ├── components/
        │   │   │   │   │   ├── audit-detail-fields.tsx
        │   │   │   │   │   ├── audit-detail-value.tsx
        │   │   │   │   │   ├── audit-log-columns.tsx
        │   │   │   │   │   ├── audit-log-details-dialog.tsx
        │   │   │   │   │   ├── audit-log-filter-bar.tsx
        │   │   │   │   │   └── audit-log-viewer.tsx
        │   │   │   │   ├── index.tsx
        │   │   │   │   └── lib/
        │   │   │   │       └── audit-details.ts
        │   │   │   ├── components/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   ├── cost-display.test.tsx
        │   │   │   │   │   ├── detail-preview.test.tsx
        │   │   │   │   │   ├── group-filter.test.tsx
        │   │   │   │   │   ├── log-type-filter.test.tsx
        │   │   │   │   │   ├── mobile-card.test.tsx
        │   │   │   │   │   ├── mobile-filter.test.tsx
        │   │   │   │   │   ├── model-badge.test.tsx
        │   │   │   │   │   ├── quota-adjustment.test.tsx
        │   │   │   │   │   ├── reject-reason.test.tsx
        │   │   │   │   │   └── usage-facts.test.tsx
        │   │   │   │   ├── columns/
        │   │   │   │   │   ├── column-helpers.tsx
        │   │   │   │   │   ├── common-logs-columns.tsx
        │   │   │   │   │   ├── drawing-logs-columns.tsx
        │   │   │   │   │   └── task-logs-columns.tsx
        │   │   │   │   ├── common-log-mobile-card.tsx
        │   │   │   │   ├── common-logs-filter-bar.tsx
        │   │   │   │   ├── common-logs-header-actions.tsx
        │   │   │   │   ├── common-logs-stats.tsx
        │   │   │   │   ├── compact-date-time-range-picker.tsx
        │   │   │   │   ├── dialogs/
        │   │   │   │   │   ├── audio-preview-dialog.tsx
        │   │   │   │   │   ├── details-dialog.tsx
        │   │   │   │   │   ├── fail-reason-dialog.tsx
        │   │   │   │   │   ├── image-dialog.tsx
        │   │   │   │   │   ├── log-detail-layout.tsx
        │   │   │   │   │   ├── prompt-dialog.tsx
        │   │   │   │   │   ├── task-details-dialog.tsx
        │   │   │   │   │   └── user-info-dialog.tsx
        │   │   │   │   ├── log-cost-display.tsx
        │   │   │   │   ├── logs-filter-toolbar.tsx
        │   │   │   │   ├── model-badge.tsx
        │   │   │   │   ├── plugin-author-link.tsx
        │   │   │   │   ├── task-artifacts.tsx
        │   │   │   │   ├── task-logs-filter-bar.tsx
        │   │   │   │   ├── timing-metrics-cell.tsx
        │   │   │   │   ├── usage-logs-mobile-card.tsx
        │   │   │   │   ├── usage-logs-provider.tsx
        │   │   │   │   └── usage-logs-table.tsx
        │   │   │   ├── constants.ts
        │   │   │   ├── data/
        │   │   │   │   └── schema.ts
        │   │   │   ├── index.tsx
        │   │   │   ├── lib/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   └── tool-surcharge.test.ts
        │   │   │   │   ├── columns.ts
        │   │   │   │   ├── filter.ts
        │   │   │   │   ├── format.ts
        │   │   │   │   ├── index.ts
        │   │   │   │   ├── mappers.ts
        │   │   │   │   ├── query-params.ts
        │   │   │   │   ├── quota-audit-operation.ts
        │   │   │   │   ├── status.ts
        │   │   │   │   ├── task-artifacts.ts
        │   │   │   │   ├── task-details.ts
        │   │   │   │   ├── task-mobile-layout.ts
        │   │   │   │   └── utils.ts
        │   │   │   ├── section-registry.tsx
        │   │   │   └── types.ts
        │   │   ├── users/
        │   │   │   ├── api.ts
        │   │   │   ├── components/
        │   │   │   │   ├── __tests__/
        │   │   │   │   │   ├── permissions.test.tsx
        │   │   │   │   │   └── quota-display.test.tsx
        │   │   │   │   ├── data-table-bulk-actions.tsx
        │   │   │   │   ├── data-table-row-actions.tsx
        │   │   │   │   ├── dialogs/
        │   │   │   │   │   ├── __tests__/
        │   │   │   │   │   │   ├── user-binding-dialog.test.tsx
        │   │   │   │   │   │   └── users-assign-group-dialog.test.tsx
        │   │   │   │   │   ├── user-binding-dialog.tsx
        │   │   │   │   │   └── users-assign-group-dialog.tsx
        │   │   │   │   ├── user-quota-cell.tsx
        │   │   │   │   ├── user-quota-dialog.tsx
        │   │   │   │   ├── users-columns.tsx
        │   │   │   │   ├── users-delete-dialog.tsx
        │   │   │   │   ├── users-mutate-drawer.tsx
        │   │   │   │   ├── users-primary-buttons.tsx
        │   │   │   │   ├── users-provider.tsx
        │   │   │   │   └── users-table.tsx
        │   │   │   ├── constants.ts
        │   │   │   ├── index.tsx
        │   │   │   ├── lib/
        │   │   │   │   ├── index.ts
        │   │   │   │   ├── user-actions.ts
        │   │   │   │   └── user-form.ts
        │   │   │   └── types.ts
        │   │   └── wallet/
        │   │       ├── api.ts
        │   │       ├── components/
        │   │       │   ├── affiliate-rewards-card.tsx
        │   │       │   ├── creem-products-section.tsx
        │   │       │   ├── dialogs/
        │   │       │   │   ├── billing-history-dialog.tsx
        │   │       │   │   ├── creem-confirm-dialog.tsx
        │   │       │   │   ├── payment-confirm-dialog.tsx
        │   │       │   │   └── transfer-dialog.tsx
        │   │       │   ├── recharge-form-card.tsx
        │   │       │   ├── subscription-plans-card.tsx
        │   │       │   └── wallet-stats-card.tsx
        │   │       ├── constants.ts
        │   │       ├── hooks/
        │   │       │   ├── index.ts
        │   │       │   ├── use-affiliate.ts
        │   │       │   ├── use-billing-history.ts
        │   │       │   ├── use-creem-payment.ts
        │   │       │   ├── use-payment.test.ts
        │   │       │   ├── use-payment.ts
        │   │       │   ├── use-redemption.ts
        │   │       │   ├── use-topup-info.ts
        │   │       │   ├── use-waffo-pancake-payment.ts
        │   │       │   └── use-waffo-payment.ts
        │   │       ├── index.tsx
        │   │       ├── lib/
        │   │       │   ├── affiliate.ts
        │   │       │   ├── billing.ts
        │   │       │   ├── format.ts
        │   │       │   ├── index.ts
        │   │       │   ├── payment.test.ts
        │   │       │   ├── payment.ts
        │   │       │   └── ui.tsx
        │   │       └── types.ts
        │   ├── hooks/
        │   │   ├── __tests__/
        │   │   │   └── sidebar-config.test.tsx
        │   │   ├── index.ts
        │   │   ├── use-admin.ts
        │   │   ├── use-copy-to-clipboard.ts
        │   │   ├── use-countdown.ts
        │   │   ├── use-debounce.ts
        │   │   ├── use-dialog.ts
        │   │   ├── use-hidden-click-unlock.ts
        │   │   ├── use-media-query.ts
        │   │   ├── use-minimum-loading-time.ts
        │   │   ├── use-mobile.ts
        │   │   ├── use-mobile.tsx
        │   │   ├── use-notifications.ts
        │   │   ├── use-sidebar-config.ts
        │   │   ├── use-sidebar-data.ts
        │   │   ├── use-sidebar-view.ts
        │   │   ├── use-status.ts
        │   │   ├── use-system-config.ts
        │   │   ├── use-table-compact-mode.ts
        │   │   ├── use-table-url-state.ts
        │   │   ├── use-top-nav-links.ts
        │   │   └── use-user-display.ts
        │   ├── i18n/
        │   │   ├── config.ts
        │   │   ├── languages.ts
        │   │   ├── locales/
        │   │   │   ├── en.json
        │   │   │   ├── fr.json
        │   │   │   ├── ja.json
        │   │   │   ├── ru.json
        │   │   │   ├── vi.json
        │   │   │   ├── zh-TW.json
        │   │   │   └── zh.json
        │   │   └── static-keys.ts
        │   ├── lib/
        │   │   ├── __tests__/
        │   │   │   ├── http-cache.test.ts
        │   │   │   ├── localized-text.test.ts
        │   │   │   ├── server-error-notifications.test.ts
        │   │   │   └── status-query.test.tsx
        │   │   ├── admin-permissions.ts
        │   │   ├── api.ts
        │   │   ├── auth-session-sync.ts
        │   │   ├── auth-session.test.ts
        │   │   ├── auth-session.ts
        │   │   ├── avatar.ts
        │   │   ├── build-metadata.ts
        │   │   ├── channel-connection-info.ts
        │   │   ├── colors.ts
        │   │   ├── constants.ts
        │   │   ├── content-format.ts
        │   │   ├── cookies.ts
        │   │   ├── copy-to-clipboard.ts
        │   │   ├── currency.ts
        │   │   ├── dayjs.ts
        │   │   ├── dom-utils.ts
        │   │   ├── format.ts
        │   │   ├── frontend-cache.ts
        │   │   ├── handle-server-error.ts
        │   │   ├── http-client.ts
        │   │   ├── http-status-code-rules.ts
        │   │   ├── legacy-route.test.ts
        │   │   ├── legacy-route.ts
        │   │   ├── lobe-icon.tsx
        │   │   ├── localized-text.ts
        │   │   ├── model-provider.ts
        │   │   ├── motion.ts
        │   │   ├── nav-modules.ts
        │   │   ├── oauth.ts
        │   │   ├── passkey.ts
        │   │   ├── password-policy.ts
        │   │   ├── query-client.ts
        │   │   ├── roles.ts
        │   │   ├── secure-verification.ts
        │   │   ├── server-error-message.test.ts
        │   │   ├── server-error-message.ts
        │   │   ├── session-flag.ts
        │   │   ├── show-submitted-data.tsx
        │   │   ├── status-query.ts
        │   │   ├── theme-customization.ts
        │   │   ├── theme-radius.ts
        │   │   ├── time.ts
        │   │   ├── use-chart-theme.ts
        │   │   ├── use-controllable-state.ts
        │   │   ├── utils.ts
        │   │   └── vchart.ts
        │   ├── main.tsx
        │   ├── routeTree.gen.ts
        │   ├── routes/
        │   │   ├── (auth)/
        │   │   │   ├── forgot-password.tsx
        │   │   │   ├── oauth.tsx
        │   │   │   ├── otp.tsx
        │   │   │   ├── register.tsx
        │   │   │   ├── reset.tsx
        │   │   │   ├── route.tsx
        │   │   │   ├── sign-in.tsx
        │   │   │   ├── sign-up.tsx
        │   │   │   └── user/
        │   │   │       └── reset.tsx
        │   │   ├── (errors)/
        │   │   │   ├── 401.tsx
        │   │   │   ├── 403.tsx
        │   │   │   ├── 404.tsx
        │   │   │   ├── 500.tsx
        │   │   │   └── 503.tsx
        │   │   ├── __root.tsx
        │   │   ├── _authenticated/
        │   │   │   ├── channels/
        │   │   │   │   └── index.tsx
        │   │   │   ├── chat/
        │   │   │   │   └── $chatId.tsx
        │   │   │   ├── chat2link.tsx
        │   │   │   ├── dashboard/
        │   │   │   │   ├── $section.tsx
        │   │   │   │   └── index.tsx
        │   │   │   ├── errors/
        │   │   │   │   └── $error.tsx
        │   │   │   ├── groups/
        │   │   │   │   └── index.tsx
        │   │   │   ├── keys/
        │   │   │   │   └── index.tsx
        │   │   │   ├── models/
        │   │   │   │   ├── $section.tsx
        │   │   │   │   └── index.tsx
        │   │   │   ├── playground/
        │   │   │   │   └── index.tsx
        │   │   │   ├── profile/
        │   │   │   │   └── index.tsx
        │   │   │   ├── redemption-codes/
        │   │   │   │   └── index.tsx
        │   │   │   ├── route.tsx
        │   │   │   ├── security/
        │   │   │   │   └── index.tsx
        │   │   │   ├── subscriptions/
        │   │   │   │   └── index.tsx
        │   │   │   ├── system-info/
        │   │   │   │   └── index.tsx
        │   │   │   ├── system-settings/
        │   │   │   │   ├── auth/
        │   │   │   │   │   ├── $section.tsx
        │   │   │   │   │   └── index.tsx
        │   │   │   │   ├── billing/
        │   │   │   │   │   ├── $section.tsx
        │   │   │   │   │   └── index.tsx
        │   │   │   │   ├── content/
        │   │   │   │   │   ├── $section.tsx
        │   │   │   │   │   └── index.tsx
        │   │   │   │   ├── index.tsx
        │   │   │   │   ├── models/
        │   │   │   │   │   ├── $section.tsx
        │   │   │   │   │   └── index.tsx
        │   │   │   │   ├── operations/
        │   │   │   │   │   ├── $section.tsx
        │   │   │   │   │   └── index.tsx
        │   │   │   │   ├── request-policies/
        │   │   │   │   │   ├── $section.tsx
        │   │   │   │   │   └── index.tsx
        │   │   │   │   ├── route.tsx
        │   │   │   │   ├── security/
        │   │   │   │   │   ├── $section.tsx
        │   │   │   │   │   └── index.tsx
        │   │   │   │   └── site/
        │   │   │   │       ├── $section.tsx
        │   │   │   │       └── index.tsx
        │   │   │   ├── task-plugins/
        │   │   │   │   └── index.tsx
        │   │   │   ├── usage-logs/
        │   │   │   │   ├── $section.tsx
        │   │   │   │   ├── audit.tsx
        │   │   │   │   └── index.tsx
        │   │   │   ├── users/
        │   │   │   │   └── index.tsx
        │   │   │   └── wallet/
        │   │   │       └── index.tsx
        │   │   ├── about/
        │   │   │   └── index.tsx
        │   │   ├── index.tsx
        │   │   ├── oauth/
        │   │   │   └── $provider.tsx
        │   │   ├── pricing/
        │   │   │   ├── $modelId/
        │   │   │   │   └── index.tsx
        │   │   │   └── index.tsx
        │   │   ├── privacy-policy.tsx
        │   │   ├── rankings/
        │   │   │   └── index.tsx
        │   │   ├── setup/
        │   │   │   └── index.tsx
        │   │   └── user-agreement.tsx
        │   ├── stores/
        │   │   ├── auth-store.ts
        │   │   ├── notification-store.ts
        │   │   ├── pricing-preferences-store.ts
        │   │   └── system-config-store.ts
        │   ├── styles/
        │   │   ├── index.css
        │   │   ├── theme-presets.css
        │   │   └── theme.css
        │   ├── tanstack-table.d.ts
        │   └── test-setup.ts
        ├── tsconfig.app.json
        ├── tsconfig.json
        ├── tsconfig.node.json
        └── vitest.config.ts
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
