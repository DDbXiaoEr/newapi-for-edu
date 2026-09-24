package model

import (
	"fmt"
	"testing"

	"github.com/QuantumNous/new-api/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func insertUsersForPaginationTest(t *testing.T, total int) {
	t.Helper()
	for id := 1; id <= total; id++ {
		user := &User{
			Id:          id,
			Username:    fmt.Sprintf("user%02d", id),
			Password:    "password123",
			DisplayName: fmt.Sprintf("User %02d", id),
			Email:       fmt.Sprintf("user%02d@example.com", id),
			Role:        common.RoleCommonUser,
			Status:      common.UserStatusEnabled,
			Group:       "default",
			AffCode:     fmt.Sprintf("aff%02d", id),
		}
		require.NoError(t, DB.Create(user).Error)
	}
}

func collectUserIDs(users []*User) []int {
	ids := make([]int, 0, len(users))
	for _, user := range users {
		ids = append(ids, user.Id)
	}
	return ids
}

func TestGetAllUsersSortsBeforePagination(t *testing.T) {
	truncateTables(t)
	insertUsersForPaginationTest(t, 42)

	pageOne, total, err := GetAllUsers(&common.PageInfo{Page: 1, PageSize: 20}, NewUserSortOptions("id", "asc"))
	require.NoError(t, err)
	assert.Equal(t, int64(42), total)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, collectUserIDs(pageOne))

	pageTwo, total, err := GetAllUsers(&common.PageInfo{Page: 2, PageSize: 20}, NewUserSortOptions("id", "asc"))
	require.NoError(t, err)
	assert.Equal(t, int64(42), total)
	assert.Equal(t, []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}, collectUserIDs(pageTwo))

	pageThree, total, err := GetAllUsers(&common.PageInfo{Page: 3, PageSize: 20}, NewUserSortOptions("id", "asc"))
	require.NoError(t, err)
	assert.Equal(t, int64(42), total)
	assert.Equal(t, []int{41, 42}, collectUserIDs(pageThree))
}

func TestSearchUsersSortsBeforePagination(t *testing.T) {
	truncateTables(t)
	insertUsersForPaginationTest(t, 42)

	users, total, err := SearchUsers("user", "", nil, nil, 20, 20, NewUserSortOptions("id", "asc"))
	require.NoError(t, err)
	assert.Equal(t, int64(42), total)
	assert.Equal(t, []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}, collectUserIDs(users))
}

func TestSearchUsersExcludeGroup(t *testing.T) {
	truncateTables(t)
	insertUsersForPaginationTest(t, 4)
	require.NoError(t, DB.Model(&User{}).Where("id IN ?", []int{1, 2}).Update("group", "vip").Error)

	users, total, err := SearchUsersWithOptions(SearchUsersOptions{
		ExcludeGroup:   "vip",
		ExcludeDeleted: true,
		StartIdx:       0,
		Num:            20,
		Sort:           NewUserSortOptions("id", "asc"),
	})
	require.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Equal(t, []int{3, 4}, collectUserIDs(users))
}

func TestAssignUsersGroupUpdatesEligibleUsers(t *testing.T) {
	setupUserUpdateTestState(t)
	insertUsersForPaginationTest(t, 3)
	require.NoError(t, DB.Model(&User{}).Where("id = ?", 2).Update("role", common.RoleAdminUser).Error)

	updated, skipped, err := AssignUsersGroup([]int{1, 2, 3, 3}, "vip", common.RoleAdminUser)
	require.NoError(t, err)
	assert.Equal(t, 1, skipped)
	assert.ElementsMatch(t, []int{1, 3}, updated)

	var groups []string
	require.NoError(t, DB.Model(&User{}).Where("id IN ?", []int{1, 2, 3}).Order("id").Pluck("group", &groups).Error)
	assert.Equal(t, []string{"vip", "default", "vip"}, groups)
}
