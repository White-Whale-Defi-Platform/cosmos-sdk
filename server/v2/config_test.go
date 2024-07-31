package serverv2_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	serverv2 "cosmossdk.io/server/v2"
	grpc "cosmossdk.io/server/v2/api/grpc"
	store "cosmossdk.io/server/v2/store"
)

func TestReadConfig(t *testing.T) {
	currentDir, err := os.Getwd()
	require.NoError(t, err)
	configPath := filepath.Join(currentDir, "testdata")

	v, err := serverv2.ReadConfig(configPath)
	require.NoError(t, err)

	require.Equal(t, v.GetString("grpc.address"), grpc.DefaultConfig().Address)
	require.Equal(t, v.GetString("store.app-db-backend"), store.DefaultConfig().AppDBBackend)
	require.Equal(t, v.GetString("chain-id"), "simapp-v2-chain")
}

func TestUnmarshalSubConfig(t *testing.T) {
	currentDir, err := os.Getwd()
	require.NoError(t, err)
	configPath := filepath.Join(currentDir, "testdata")

	v, err := serverv2.ReadConfig(configPath)
	require.NoError(t, err)

	grpcConfig := grpc.DefaultConfig()
	err = serverv2.UnmarshalSubConfig(v, "grpc", &grpcConfig)
	require.NoError(t, err)

	require.True(t, grpc.DefaultConfig().Enable)
	require.False(t, grpcConfig.Enable)

	storeConfig := store.Config{}
	err = serverv2.UnmarshalSubConfig(v, "store", &storeConfig)
	require.NoError(t, err)
	require.Equal(t, *store.DefaultConfig(), storeConfig)
}
