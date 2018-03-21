package agollo

import (
	"github.com/go-apollo/agollo/test"
	"testing"
)

func TestStart(t *testing.T) {
	go runMockConfigServer(onlyNormalConfigResponse)
	go runMockNotifyServer(onlyNormalResponse)
	defer closeMockConfigServer()

	Start()

	value := getValue("key1")
	test.Equal(t, "value1", value)
}
