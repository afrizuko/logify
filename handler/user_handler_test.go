package handler

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/afrizuko/logify/model"
	"github.com/stretchr/testify/assert"
)

var handler *UserHandler

func init() {
	service := model.NewUserMockImpl(5)
	handler = NewUserHandler(service)
}

func TestGetUsers(t *testing.T) {
	r := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	assert.Equal(t, 200, w.Result().StatusCode)

	var users []model.User
	err := json.NewDecoder(w.Body).Decode(&users)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(users))
}

func TestGetUser(t *testing.T) {
	t.Run("it returns a single user by id", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/users/1", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(t, 200, w.Result().StatusCode)

		var user model.User
		err := json.NewDecoder(w.Body).Decode(&user)
		assert.NoError(t, err)
		assert.NotEmpty(t, user)
	})

	t.Run("it returns a user not found", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/users/0", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func TestCreateUser(t *testing.T) {

	t.Run("it returns 200 for creation of a user", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/users", newUser(t, "Peter"))
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(t, 200, w.Result().StatusCode)

		var user model.User
		assert.NoError(t, json.NewDecoder(w.Body).Decode(&user))
		assert.Equal(t, uint(6), user.ID)
	})

	t.Run("it returns a single user by id", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/users/6", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(t, 200, w.Result().StatusCode)

		var user model.User
		err := json.NewDecoder(w.Body).Decode(&user)
		assert.NoError(t, err)
		assert.Equal(t, uint(6), user.ID)
	})
}

func TestUpdateUser(t *testing.T) {

	t.Run("it returns 202 for update on a user", func(t *testing.T) {
		r := httptest.NewRequest("PUT", "/users/6", newUser(t, "John"))
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(t, 202, w.Result().StatusCode)

		var user model.User
		assert.NoError(t, json.NewDecoder(w.Body).Decode(&user))
		assert.Equal(t, "John", user.Username)
	})

	t.Run("it returns a single user by id", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/users/6", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(t, 200, w.Result().StatusCode)

		var user model.User
		err := json.NewDecoder(w.Body).Decode(&user)
		assert.NoError(t, err)
		assert.Equal(t, "John", user.Username)
	})
}

func TestDeleteUser(t *testing.T) {

	t.Run("it returns 200 for deleting a user", func(t *testing.T) {
		r := httptest.NewRequest("DELETE", "/users/6", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(t, 202, w.Result().StatusCode)
	})

	t.Run("it returns a single user by id", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/users/6", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func newUser(t *testing.T, name string) *bytes.Buffer {
	buf := new(bytes.Buffer)
	assert.NoError(t, json.NewEncoder(buf).Encode(&model.User{Username: name}))
	return buf
}
