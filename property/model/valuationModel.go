package ValuationModel

import (
	"time"
)

type ValuationModel struct {
	Id            int64
	StreetAddress string
	Town          string
	ValuationDate time.Time
	Value         float64
}

func (model ValuationModel) Init(id int64, streetAddress string, town string, valuationDate time.Time, value float64) {

}
