package user_handlers

import (
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

//	func newTestServer(t *testing.T, store db.Store) *Server {
//		config := util.Config{
//			TokenSymmetricKey:   util.RandomPass(32),
//			AccessTokenDuration: time.Minute,
//		}
//
//		server, err := NewServer(config, store)
//		require.NoError(t, err)
//
//		return server
//	}
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
