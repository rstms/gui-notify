package notify

import (
	"github.com/stretchr/testify/require"
	"log"
	"path/filepath"
	"testing"
)

func init() {
	log.Println("init")
	Init("notify", Version, filepath.Join("testdata", "config.yaml"))
}

func TestNotify(t *testing.T) {
	err := Send("this is the title", "this is the message")
	require.Nil(t, err)
}
