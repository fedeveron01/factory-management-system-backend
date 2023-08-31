package entities

import "github.com/fedeveron01/golang-base/cmd/core"

type PurchaseOrder struct {
	core.EntitiesBase
	Number               int
	Description          string
	PurchaseOrderDetails []PurchaseOrderDetail
}
