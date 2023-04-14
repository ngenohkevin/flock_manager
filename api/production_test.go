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

func TestGetProductionAPI(t *testing.T) {
	prod := randomProduction()

	testCases := []struct {
		name          string
		ID            int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			ID:   prod.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProduction(gomock.Any(), gomock.Eq(prod.ID)).Times(1).
					Return(prod, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProd(t, recorder.Body, prod)
			},
		},
		{
			name: "NotFound",
			ID:   prod.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProduction(gomock.Any(), gomock.Eq(prod.ID)).Times(1).
					Return(db.Production{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "internalError",
			ID:   prod.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProduction(gomock.Any(), gomock.Eq(prod.ID)).Times(1).
					Return(db.Production{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			ID:   0,
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

			url := fmt.Sprintf("/production/%d", tc.ID)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(t, recorder)
		})
	}
}

// Randomize production
func randomProduction() db.Production {

	return db.Production{
		ID:           util.RandomInt(1, 1000),
		Eggs:         util.RandomProduction(),
		Dirty:        util.RandomProduction(),
		WrongShape:   util.RandomProduction(),
		Damaged:      util.RandomProduction(),
		HatchingEggs: util.RandomProduction(),
	}
}

func requireBodyMatchProd(t *testing.T, body *bytes.Buffer, prod db.Production) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotProd db.Production
	err = json.Unmarshal(data, &gotProd)
	require.NoError(t, err)
	require.Equal(t, prod, gotProd)
}

func requireBodyMatchProduction(t *testing.T, body *bytes.Buffer, production []db.Production) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotProduction []db.Production
	err = json.Unmarshal(data, &gotProduction)
	require.NoError(t, err)
	require.Equal(t, production, gotProduction)
}
