package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/phongtv-1971/go-skeleton/constants"
	mockdb "github.com/phongtv-1971/go-skeleton/db/mock"
	db "github.com/phongtv-1971/go-skeleton/db/sqlc"
	"github.com/phongtv-1971/go-skeleton/util"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserAPI(t *testing.T) {
	user := randomUser()

	testCases := []struct {
		name string
		userID int64
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			userID: user.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "NotFound",
			userID: user.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalError",
			userID: user.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			userID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			/* build stub */
			tc.buildStubs(store)

			/* Start test server and send request */
			server := NewServer(store, constants.Test)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/v1/users/%d", tc.userID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			/* check response */
			tc.checkResponse(t, recorder)
		})
	}
}

func randomUser() db.User {
	return db.User{
		ID: util.RandomInt(1, 1000),
		Email: util.RandomEmail(),
		Name: util.RandomName(),
	}
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User)  {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var response struct{
		Success bool    `json:"success"`
		Data    db.User `json:"data"`
	}
	err = json.Unmarshal(data, &response)
	require.NoError(t, err)
	require.Equal(t, user, response.Data)
}