package assertions

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisements/models"
)

func ContainsAdvertisementByTitle(t *testing.T, items []models.Advertisement, title string) {
	t.Helper()

	found := slices.ContainsFunc(items, func(item models.Advertisement) bool {
		return item.Title == title
	})
	require.True(t, found, "среди %d объявлений нет объявления с названием %q", len(items), title)
}
