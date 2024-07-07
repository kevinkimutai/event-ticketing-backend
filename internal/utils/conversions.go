package utils

import (
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func ConvertTimeToTimestamp(t time.Time) pgtype.Timestamptz {
	tstring := t.String()

	var timestamp pgtype.Timestamptz
	timestamp.Scan(tstring)

	return timestamp
}

func ConvertStringToText(s string) pgtype.Text {

	var text pgtype.Text
	text.Scan(s)

	return text
}

func ConvertFloat64ToNumeric(f float64) pgtype.Numeric {

	var numeric pgtype.Numeric
	numeric.Scan(f)

	return numeric

}

func ConvertNumericToFloat64(numeric pgtype.Numeric) float64 {
	fval, _ := numeric.Value()

	//Convert To Float64
	var floatVal float64
	if strVal, ok := fval.(string); ok {
		floatVal, _ = strconv.ParseFloat(strVal, 64)
	}
	return floatVal

}
