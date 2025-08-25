package cmd

import (
	"github.com/stretchr/testify/require"
	"log"
	"path/filepath"
	"testing"
)

func initTestConfig(t *testing.T) {
	Init("notify", rootCmd.Version, filepath.Join("testdata", "config.yaml"))
	//viper.WriteConfigTo(os.Stdout)
}

func TestViperGet(t *testing.T) {
	initTestConfig(t)
	testValue := ViperGetString("test_value")
	require.Equal(t, "testing123", testValue)
}

func TestLog(t *testing.T) {
	log.Println("log message")
	Shutdown()
}

func TestDebugLog(t *testing.T) {
	ViperSet("debug", true)
	initTestConfig(t)
	log.Println("log message")
	Shutdown()
}
