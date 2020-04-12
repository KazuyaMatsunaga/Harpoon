package controller

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// コンテナ一覧取得
func ContaList() (contas []types.Container) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containerListOptions := types.ContainerListOptions{
		All: true,
	}

	containers, err := cli.ContainerList(context.Background(), containerListOptions)
	if err != nil {
		panic(err)
	}

	return containers
}

// コンテナを停止
func ContaStop() {

}

//　コンテナを削除
func ContaRemove() {

}
