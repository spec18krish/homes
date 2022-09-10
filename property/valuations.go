package valuations

import (
	"fmt"
	"time"
)

type ValuationModel struct {
	Id            int64
	StreetAddress string
	Town          string
	ValuationDate time.Time
	Value         float64
}

func LastEncounteredRecord() {
	fmt.Printf("lastEncounteredRecord called")
}
func FirstEncounteredRecord() {
	fmt.Printf("firstEncounteredRecord called")
}

func FilterRecord() {
	fmt.Printf("filterRecord called")
}

func ProcessChunks() {
	fmt.Printf("processChunks called")
}
