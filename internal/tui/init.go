package tui

import (
	bubbleteaModel "test-cet-wp-plugin/internal/model/bubbletea"
	"test-cet-wp-plugin/internal/model/structs"
)

func InitializeTea(Environments *structs.Environments) {
	EnvironmentsRef = Environments

}

func initParent() *bubbleteaModel.MainModel {
	return &bubbleteaModel.MainModel{
		Hello: false,
	}
}
