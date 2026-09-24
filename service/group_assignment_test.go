package service

import (
	"testing"

	"github.com/QuantumNous/new-api/setting/ratio_setting"
	"github.com/QuantumNous/new-api/setting/system_setting"
)

func TestResolveAssignedGroupMatchesFirstValidRule(t *testing.T) {
	originalRatios := ratio_setting.GroupRatio2JSONString()
	t.Cleanup(func() {
		if err := ratio_setting.UpdateGroupRatioByJSONString(originalRatios); err != nil {
			t.Fatalf("restore group ratios: %v", err)
		}
	})
	if err := ratio_setting.UpdateGroupRatioByJSONString(`{"default":1,"student":1,"teacher":1}`); err != nil {
		t.Fatalf("set group ratios: %v", err)
	}

	rules := []system_setting.GroupAssignmentRule{
		{Attribute: "uid", Pattern: `^stu-`, Group: "student"},
		{Attribute: "uid", Pattern: `^tea-`, Group: "teacher"},
	}
	got := ResolveAssignedGroup(rules, DirectoryAttributes("uid", "stu-1001"))
	if got != "student" {
		t.Fatalf("expected student, got %q", got)
	}
	got = ResolveAssignedGroup(rules, DirectoryAttributes("uid", "tea-2002"))
	if got != "teacher" {
		t.Fatalf("expected teacher, got %q", got)
	}
	got = ResolveAssignedGroup(rules, DirectoryAttributes("uid", "admin-1"))
	if got != "" {
		t.Fatalf("expected empty group, got %q", got)
	}
}

func TestResolveAssignedGroupSkipsUnknownAndAutoGroups(t *testing.T) {
	originalRatios := ratio_setting.GroupRatio2JSONString()
	t.Cleanup(func() {
		if err := ratio_setting.UpdateGroupRatioByJSONString(originalRatios); err != nil {
			t.Fatalf("restore group ratios: %v", err)
		}
	})
	if err := ratio_setting.UpdateGroupRatioByJSONString(`{"default":1,"vip":1}`); err != nil {
		t.Fatalf("set group ratios: %v", err)
	}

	rules := []system_setting.GroupAssignmentRule{
		{Attribute: "uid", Pattern: `.*`, Group: "auto"},
		{Attribute: "uid", Pattern: `.*`, Group: "missing"},
		{Attribute: "uid", Pattern: `^vip-`, Group: "vip"},
	}
	got := ResolveAssignedGroup(rules, DirectoryAttributes("UID", "vip-9"))
	if got != "vip" {
		t.Fatalf("expected vip, got %q", got)
	}
}
