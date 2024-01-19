package api_test

import (
	"datamanager/web"
	"testing"
)

func TestMain(m *testing.M) {
	web.SERVER_API_PORT = 3000
	go web.App()
	m.Run()
}
