package slice_test

import (
	"testing"

	"github.com/albimukti/Tranning-golang/session-4-crud-user/entity"
	"github.com/albimukti/Tranning-golang/session-4-crud-user/repository/slice"
	"github.com/stretchr/testify/require"
)

func TestUserRepository(t *testing.T) {
	repo := slice.NewUserRepository([]entity.User{})

	t.Run("CreateUser", func(t *testing.T) {
		newUser := entity.User{Name: "albi", Email: "albiaja@example.com", Password: "gituaja123"}
		createdUser := repo.CreateUser(&newUser)

		require.Equal(t, 0, createdUser.ID)
		require.Equal(t, "albi", createdUser.Name)
		require.Equal(t, "albiaja@example.com", createdUser.Email)
		require.NotZero(t, createdUser.CreatedAt)
		require.NotZero(t, createdUser.UpdatedAt)
	})

	t.Run("GetUserByID", func(t *testing.T) {
		user, found := repo.GetUserByID(0)

		require.True(t, found)
		require.Equal(t, "albi", user.Name)

		_, notFound := repo.GetUserByID(99)
		require.False(t, notFound)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		update := entity.User{Name: "Albi Updated", Email: "albi.updated@example.com", Password: "albi123"}
		updatedUser, found := repo.UpdateUser(0, update)

		require.True(t, found)
		require.Equal(t, "Albi Updated", updatedUser.Name)
		require.Equal(t, "albi.updated@example.com", updatedUser.Email)

		_, notFound := repo.UpdateUser(99, update)
		require.False(t, notFound)
	})

	t.Run("DeleteUser", func(t *testing.T) {
		deleted := repo.DeleteUser(0)
		require.True(t, deleted)

		notDeleted := repo.DeleteUser(99)
		require.False(t, notDeleted)
	})

	t.Run("GetAllUsers", func(t *testing.T) {
		allUsers := repo.GetAllUsers()
		require.Empty(t, allUsers) // Tidak tersisa pengguna setelah penghapusan
	})
}
