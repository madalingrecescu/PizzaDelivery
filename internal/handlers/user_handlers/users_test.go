package user_handlers

//
//import (
//	"bytes"
//	"database/sql"
//	"encoding/json"
//	"fmt"
//	"github.com/golang/mock/gomock"
//	mockdb_users "github.com/madalingrecescu/PizzaDelivery/internal/db/mock"
//	db "github.com/madalingrecescu/PizzaDelivery/internal/db/sqlc_users"
//	"github.com/madalingrecescu/PizzaDelivery/internal/util"
//	"github.com/stretchr/testify/require"
//	"io"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestGetAccount(t *testing.T) {
//	account, _ := randomAccount(t)
//
//	testCases := []struct {
//		name          string
//		accountUserId int32
//		buildStubs    func(store *mockdb_users.MockStore)
//		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
//	}{
//		{
//			name:          "OK",
//			accountUserId: account.UserID,
//			buildStubs: func(store *mockdb_users.MockStore) {
//				store.EXPECT().
//					GetAccount(gomock.Any(), gomock.Eq(account.UserID)).
//					Times(1).
//					Return(account, nil)
//			},
//			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
//				require.Equal(t, http.StatusOK, recorder.Code)
//				requireBodyMatchAccount(t, recorder.Body, account)
//			},
//		},
//		{
//			name:          "NotFound",
//			accountUserId: account.UserID,
//			buildStubs: func(store *mockdb_users.MockStore) {
//				store.EXPECT().
//					GetAccount(gomock.Any(), gomock.Eq(account.UserID)).
//					Times(1).
//					Return(db.User{}, sql.ErrNoRows)
//			},
//			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
//				require.Equal(t, http.StatusNotFound, recorder.Code)
//			},
//		},
//		{
//			name:          "InternalError",
//			accountUserId: account.UserID,
//			buildStubs: func(store *mockdb_users.MockStore) {
//				store.EXPECT().
//					GetAccount(gomock.Any(), gomock.Eq(account.UserID)).
//					Times(1).
//					Return(db.User{}, sql.ErrConnDone)
//			},
//			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
//				require.Equal(t, http.StatusInternalServerError, recorder.Code)
//			},
//		},
//		{
//			name:          "InvalidID",
//			accountUserId: 0,
//			buildStubs: func(store *mockdb_users.MockStore) {
//				store.EXPECT().
//					GetAccount(gomock.Any(), gomock.Any()).
//					Times(0)
//			},
//			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
//				require.Equal(t, http.StatusBadRequest, recorder.Code)
//			},
//		},
//	}
//
//	for i := range testCases {
//		tc := testCases[i]
//
//		t.Run(tc.name, func(t *testing.T) {
//			ctrl := gomock.NewController(t)
//			defer ctrl.Finish()
//
//			store := mockdb_users.NewMockStore(ctrl)
//			tc.buildStubs(store)
//
//			// strat test server and send request
//			server := newTestServer(t, store)
//			recorder := httptest.NewRecorder()
//
//			url := fmt.Sprintf("/user/%d", tc.accountUserId)
//			request, err := http.NewRequest(http.MethodGet, url, nil)
//			require.NoError(t, err)
//
//			server.router.ServeHTTP(recorder, request)
//			tc.checkResponse(t, recorder)
//		})
//	}
//}
//
////	func TestCreateAccount(t *testing.T) {
////		account, password := randomAccount(t)
////
////		testCases := []struct {
////			name          string
////			requestBody   gin.H
////			buildStubs    func(store *mockdb_users.MockStore)
////			checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
////		}{
////			{
////				name: "OK",
////				requestBody: gin.H{
////					"user_id":      account.UserID,
////					"username":     account.Username,
////					"password":     password,
////					"email":        account.Email,
////					"phone_number": account.PhoneNumber,
////				},
////				buildStubs: func(store *mockdb_users.MockStore) {
////					store.EXPECT().
////						CreateAccount(gomock.Any(), gomock.Any()).
////						Times(1).
////						Return(account, nil)
////				},
////				checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
////					require.Equal(t, http.StatusOK, recorder.Code)
////					requireBodyMatchAccount(t, recorder.Body, account)
////				},
////			},
////			{
////				name: "InternalError",
////				requestBody: gin.H{
////					"user_id":      account.UserID,
////					"username":     account.Username,
////					"password":     password,
////					"email":        account.Email,
////					"phone_number": account.PhoneNumber,
////				},
////				buildStubs: func(store *mockdb_users.MockStore) {
////					store.EXPECT().
////						CreateAccount(gomock.Any(), gomock.Any()).
////						Times(1).
////						Return(db.User{}, sql.ErrConnDone)
////				},
////
////				checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
////					require.Equal(t, http.StatusInternalServerError, recorder.Code)
////				},
////			},
////			{
////				name: "DuplicateUsername",
////				requestBody: gin.H{
////					"user_id":      account.UserID,
////					"username":     account.Username,
////					"password":     password,
////					"email":        account.Email,
////					"phone_number": account.PhoneNumber,
////				},
////				buildStubs: func(store *mockdb_users.MockStore) {
////					store.EXPECT().
////						CreateAccount(gomock.Any(), gomock.Any()).
////						Times(1).
////						Return(db.User{}, &pq.Error{Code: "23505"})
////				},
////
////				checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
////					require.Equal(t, http.StatusForbidden, recorder.Code)
////				},
////			},
////			{
////				name: "InvalidEmail",
////				requestBody: gin.H{
////					"user_id":      account.UserID,
////					"username":     account.Username,
////					"password":     password,
////					"email":        "invalid-email",
////					"phone_number": account.PhoneNumber,
////				},
////				buildStubs: func(store *mockdb_users.MockStore) {
////					store.EXPECT().
////						CreateAccount(gomock.Any(), gomock.Any()).
////						Times(0)
////				},
////
////				checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
////					require.Equal(t, http.StatusBadRequest, recorder.Code)
////				},
////			},
////			{
////				name: "PasswordTooShort",
////				requestBody: gin.H{
////					"user_id":      account.UserID,
////					"username":     account.Username,
////					"password":     "123",
////					"email":        account.Email,
////					"phone_number": account.PhoneNumber,
////				},
////				buildStubs: func(store *mockdb_users.MockStore) {
////					store.EXPECT().
////						CreateAccount(gomock.Any(), gomock.Any()).
////						Times(0)
////				},
////
////				checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
////					require.Equal(t, http.StatusBadRequest, recorder.Code)
////				},
////			},
////		}
////
////		for i := range testCases {
////			tc := testCases[i]
////
////			t.Run(tc.name, func(t *testing.T) {
////				ctrl := gomock.NewController(t)
////				defer ctrl.Finish()
////
////				store := mockdb_users.NewMockStore(ctrl)
////				tc.buildStubs(store)
////
////				// strat test server and send request
////				server := newTestServer(t, store)
////				recorder := httptest.NewRecorder()
////
////				requestBody, err := json.Marshal(tc.requestBody)
////				require.NoError(t, err)
////
////				request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewReader(requestBody))
////				require.NoError(t, err)
////
////				server.router.ServeHTTP(recorder, request)
////				tc.checkResponse(t, recorder)
////			})
////		}
////	}
//func randomAccount(t *testing.T) (db.User, string) {
//
//	password := util.RandomPass(6)
//	hashedPassword, err := util.HashPassword(password)
//	require.NoError(t, err)
//
//	return db.User{
//		Username:       util.RandomNameOrEmail(5, false),
//		Email:          util.RandomNameOrEmail(5, true),
//		HashedPassword: hashedPassword,
//		PhoneNumber:    util.RandomPhoneNumber(8),
//	}, hashedPassword
//}
//
//func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.User) {
//	data, err := io.ReadAll(body)
//	require.NoError(t, err)
//
//	var gotAccount db.User
//	err = json.Unmarshal(data, &gotAccount)
//	gotAccount.UserID = account.UserID
//	gotAccount.HashedPassword = account.HashedPassword
//	require.NoError(t, err)
//	require.Equal(t, account, gotAccount)
//
//}
