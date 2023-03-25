package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	mockdb "github.com/ngenohkevin/flock_manager/db/mock"
	db "github.com/ngenohkevin/flock_manager/db/sqlc"
	"github.com/ngenohkevin/flock_manager/db/util"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// GetBreed tests
func TestGetBreedAPI(t *testing.T) {
	breed := randomBreed()

	testCases := []struct {
		name          string
		BreedID       int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			BreedID: breed.BreedID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetBreed(gomock.Any(), gomock.Eq(breed.BreedID)).Times(1).
					Return(breed, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchBreed(t, recorder.Body, breed)
			},
		},
		{
			name:    "NotFound",
			BreedID: breed.BreedID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetBreed(gomock.Any(), gomock.Eq(breed.BreedID)).Times(1).
					Return(db.Breed{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:    "internalError",
			BreedID: breed.BreedID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetBreed(gomock.Any(), gomock.Eq(breed.BreedID)).Times(1).
					Return(db.Breed{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:    "InvalidID",
			BreedID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetBreed(gomock.Any(), gomock.Any()).Times(0)
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
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/breeds/%d", tc.BreedID)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(t, recorder)
		})
	}
}

// Randomise breeds
func randomBreed() db.Breed {
	return db.Breed{
		BreedID:   util.RandomInt(1, 1000),
		BreedName: util.RandomBreed(),
	}
}

func requireBodyMatchBreed(t *testing.T, body *bytes.Buffer, breed db.Breed) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotBreed db.Breed
	err = json.Unmarshal(data, &gotBreed)
	require.NoError(t, err)
	require.Equal(t, breed, gotBreed)
}
