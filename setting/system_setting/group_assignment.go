package system_setting

type GroupAssignmentRule struct {
	Attribute string `json:"attribute"`
	Pattern   string `json:"pattern"`
	Group     string `json:"group"`
}
