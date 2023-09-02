package entities

import (
	"github.com/fedeveron01/golang-base/cmd/core"
	"time"
)

type ProductionOrder struct {
	core.EntitiesBase
	StartDateTime         time.Time
	EndDateTime           time.Time
	ProductionOrderDetail []ProductionOrderDetail
}