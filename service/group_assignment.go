package service

import (
	"regexp"
	"strings"

	"github.com/QuantumNous/new-api/setting/ratio_setting"
	"github.com/QuantumNous/new-api/setting/system_setting"
)

func ResolveAssignedGroup(rules []system_setting.GroupAssignmentRule, attributes map[string][]string) string {
	if len(rules) == 0 || len(attributes) == 0 {
		return ""
	}
	normalized := make(map[string][]string, len(attributes))
	for name, values := range attributes {
		key := strings.ToLower(strings.TrimSpace(name))
		if key == "" {
			continue
		}
		normalized[key] = append(normalized[key], values...)
	}
	for _, rule := range rules {
		attribute := strings.ToLower(strings.TrimSpace(rule.Attribute))
		pattern := strings.TrimSpace(rule.Pattern)
		group := strings.TrimSpace(rule.Group)
		if attribute == "" || pattern == "" || group == "" {
			continue
		}
		if group == "auto" || !ratio_setting.ContainsGroupRatio(group) {
			continue
		}
		compiled, err := regexp.Compile(pattern)
		if err != nil {
			continue
		}
		for _, value := range normalized[attribute] {
			if compiled.MatchString(value) {
				return group
			}
		}
	}
	return ""
}

func DirectoryAttributes(pairs ...string) map[string][]string {
	attributes := make(map[string][]string, len(pairs)/2+1)
	for i := 0; i+1 < len(pairs); i += 2 {
		name := strings.TrimSpace(pairs[i])
		value := strings.TrimSpace(pairs[i+1])
		if name == "" || value == "" {
			continue
		}
		key := strings.ToLower(name)
		attributes[key] = append(attributes[key], value)
	}
	return attributes
}

func MergeDirectoryAttributes(base map[string][]string, extra map[string][]string) map[string][]string {
	merged := make(map[string][]string, len(base)+len(extra))
	for name, values := range base {
		key := strings.ToLower(strings.TrimSpace(name))
		if key == "" {
			continue
		}
		merged[key] = append(merged[key], values...)
	}
	for name, values := range extra {
		key := strings.ToLower(strings.TrimSpace(name))
		if key == "" {
			continue
		}
		merged[key] = append(merged[key], values...)
	}
	return merged
}
