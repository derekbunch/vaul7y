package view

import (
	"testing"

	"github.com/dkyanakiev/vaul7y/internal/models"
	"github.com/stretchr/testify/require"
)

func TestResolveDefaultPath(t *testing.T) {
	mounts := map[string]*models.MountOutput{
		"secret/":          {},
		"database/secret/": {},
	}

	t.Run("splits mount from path", func(t *testing.T) {
		mount, path, ok := resolveDefaultPath("secret/team/data-engineering", mounts)

		require.True(t, ok)
		require.Equal(t, "secret/", mount)
		require.Equal(t, "team/data-engineering", path)
	})

	t.Run("uses longest matching mount", func(t *testing.T) {
		mount, path, ok := resolveDefaultPath("database/secret/team/data-engineering", mounts)

		require.True(t, ok)
		require.Equal(t, "database/secret/", mount)
		require.Equal(t, "team/data-engineering", path)
	})

	t.Run("supports mount root", func(t *testing.T) {
		mount, path, ok := resolveDefaultPath("secret", mounts)

		require.True(t, ok)
		require.Equal(t, "secret/", mount)
		require.Equal(t, "", path)
	})

	t.Run("returns false when no mount matches", func(t *testing.T) {
		_, _, ok := resolveDefaultPath("missing/team", mounts)

		require.False(t, ok)
	})
}
