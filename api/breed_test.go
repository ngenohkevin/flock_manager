package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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

func TestCreateBreed(t *testing.T) {
	breed := randomBreed()

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"breed": breed.BreedName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := breed.BreedName
				store.EXPECT().CreateBreed(gomock.Any(), gomock.Eq(arg)).Times(1).Return(breed, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchBreed(t, recorder.Body, breed)
			},
		},
		//{
		//	name: "NoAuthorization",
		//	body: gin.H{
		//		"breed": breed.BreedName,
		//	},
		//	buildStubs: func(store *mockdb.MockStore) {
		//		store.EXPECT().CreateBreed(gomock.Any(), gomock.Any()).Times(0)
		//	},
		//	checkResponse: func(recorder *httptest.ResponseRecorder) {
		//		require.Equal(t, http.StatusUnauthorized, recorder.Code)
		//	},
		//},
		{
			name: "InternalError",
			body: gin.H{
				"breed": breed.BreedName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateBreed(gomock.Any(), gomock.Any()).Times(1).Return(db.Breed{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
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

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/breeds"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}

}
func TestListBreedsAPI(t *testing.T) {
	n := 5
	breeds := make([]db.Breed, n)
	for i := 0; i < n; i++ {
		breeds[i] = randomBreed()
	}
	type Query struct {
		pageID   int
		pageSize int
	}

	testCases := []struct {
		name          string
		query         Query
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListBreedsParams{
					Limit:  int32(n),
					Offset: 0,
				}
				store.EXPECT().ListBreeds(gomock.Any(), gomock.Eq(arg)).
					Times(1).Return(breeds, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchBreeds(t, recorder.Body, breeds)
			},
		},
		//{
		//	name: "NoAuthorization",
		//	query: Query{
		//		pageID:   1,
		//		pageSize: n,
		//	},
		//	buildStubs: func(store *mockdb.MockStore) {
		//		store.EXPECT().ListBreeds(gomock.Any(), gomock.Any()).Times(0)
		//	},
		//	checkResponse: func(recorder *httptest.ResponseRecorder) {
		//		require.Equal(t, http.StatusUnauthorized, recorder.Code)
		//	},
		//},
		{
			name: "InternalError",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().ListBreeds(gomock.Any(), gomock.Any()).
					Times(1).Return([]db.Breed{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidPageID",
			query: Query{
				pageID:   -1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListBreeds(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidPageSize",
			query: Query{
				pageID:   1,
				pageSize: 100000,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListBreeds(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
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

			request, err := http.NewRequest(http.MethodGet, "/breeds", nil)
			require.NoError(t, err)

			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
			request.URL.RawQuery = q.Encode()

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
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

func requireBodyMatchBreeds(t *testing.T, body *bytes.Buffer, accounts []db.Breed) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccounts []db.Breed
	err = json.Unmarshal(data, &gotAccounts)
	require.NoError(t, err)
	require.Equal(t, accounts, gotAccounts)
}