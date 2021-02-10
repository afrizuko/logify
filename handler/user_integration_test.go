package handler

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/afrizuko/logify/model"
	"github.com/stretchr/testify/assert"
)

func TestGetCreateUpdateUsers(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	createdList := make([]model.User, 2)
	handler := NewUserHandler(model.NewUserRepository())

	for i := 0; i < len(createdList); i++ {
		t.Run("it creates and returns a new user", func(t *testing.T) {
			r := httptest.NewRequest("POST", "/users", newUser(t, "Peter"))
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, r)
			assert.Equal(t, 200, w.Result().StatusCode)

			var user model.User
			assert.NoError(t, json.NewDecoder(w.Body).Decode(&user))
			assert.NotEmpty(t, user)
			createdList[i] = user
		})
	}

	t.Run("it returns a list of users", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/users", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(t, 200, w.Result().StatusCode)

		var users []model.User
		err := json.NewDecoder(w.Body).Decode(&users)
		assert.NoError(t, err)
		assert.Equal(t, len(createdList), len(users))
	})

	t.Run("it returns a single user by id", func(t *testing.T) {
		r := httptest.NewRequest("GET", fmt.Sprintf("/users/%d", createdList[0].ID), nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(t, 200, w.Result().StatusCode)

		var user model.User
		err := json.NewDecoder(w.Body).Decode(&user)
		assert.NoError(t, err)
		assert.NotEmpty(t, user)
		assert.Equal(t, createdList[0].ID, user.ID)
	})

	t.Run("it returns 202 for update on a user", func(t *testing.T) {
		r := httptest.NewRequest("PUT", fmt.Sprintf("/users/%d", createdList[0].ID), newUser(t, "John"))
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(t, 202, w.Result().StatusCode)

		var user model.User
		assert.NoError(t, json.NewDecoder(w.Body).Decode(&user))
		assert.Equal(t, "John", user.Username)
	})

	t.Cleanup(func() {
		for _, user := range createdList {
			r := httptest.NewRequest("DELETE", fmt.Sprintf("/users/%d", user.ID), nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, r)
			assert.Equal(t, 202, w.Result().StatusCode)
		}
	})
}
