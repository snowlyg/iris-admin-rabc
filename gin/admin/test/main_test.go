package test

import (
	"os"
	"testing"

	"github.com/snowlyg/helper/tests"
	rbac "github.com/snowlyg/iris-admin-rbac/gin"
	web_tests "github.com/snowlyg/iris-admin/server/web/tests"
	"github.com/snowlyg/iris-admin/server/web/web_gin"
)

var TestServer *web_gin.WebServer
var TestClient *tests.Client

func TestMain(m *testing.M) {
	mysqlPwd := os.Getenv("mysqlPwd")
	redisPwd := os.Getenv("redisPwd")
	var uuid string
	uuid, TestServer = web_tests.BeforeTestMainGin(mysqlPwd, redisPwd, 4, rbac.PartyFunc, rbac.SeedFunc)
	code := m.Run()
	web_tests.AfterTestMain(uuid)

	os.Exit(code)
}
