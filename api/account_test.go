// package api

// import (
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	mock_sqlc "github.com/muhammadsaman77/simplebank/db/mock"
// 	db "github.com/muhammadsaman77/simplebank/db/sqlc"
// 	"github.com/muhammadsaman77/simplebank/util"
// 	"github.com/stretchr/testify/require"
// 	"go.uber.org/mock/gomock"
// )

// func TestGetAccountAPI(t *testing.T) {
// 	account := randomAccount()
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	store:= mock_sqlc.NewMockStore(ctrl)
// 	store.EXPECT().GetAccount(gomock.Any(),gomock.Eq(account.ID)).Times(1).Return(account, nil)
// 	server := NewServer(util.Config{} ,store)
// 	recorder := httptest.NewRecorder()
// 	url := fmt.Sprintf("/accounts/%d",account.ID)
// 	request, err := http.NewRequest(http.MethodGet, url, nil)
// 	require.NoError(t, err)
// 	server.router.ServeHTTP(recorder, request)
// 	require.Equal(t, http.StatusOK, recorder.Code)
// }

// func randomAccount()db.Account{
// 	return db.Account{
// 		ID: util.RandomInt(1,1000),
// 		Owner: util.RandomOwner(),
// 		Balance: util.RandomMoney(),
// 		Currency: util.RandomCurrency(),
		
		
// 	}
// }