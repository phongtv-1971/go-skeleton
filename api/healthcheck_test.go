package api

import (
	"github.com/golang/mock/gomock"
	"github.com/phongtv-1971/go-skeleton/constants"
	mockdb "github.com/phongtv-1971/go-skeleton/db/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealCheckApi (t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	server := NewServer(store, constants.Test)

	recorder := httptest.NewRecorder()
	url := "/health_check"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
}
