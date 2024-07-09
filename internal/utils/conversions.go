package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func ConvertTimeToTimestamp(t time.Time) pgtype.Timestamptz {
	// parsedTime, err := time.Parse(time.RFC3339, t.String())
	// if err != nil {
	// 	fmt.Println("Error parsing time:", err)

	// }

	var timestamp pgtype.Timestamptz
	timestamp.Scan(t)

	return timestamp
}

func ConvertStringToText(s string) pgtype.Text {

	var text pgtype.Text
	text.Scan(s)

	return text
}

// func ConvertFloat64ToNumeric(f float64) pgtype.Numeric {

// 	// Convert float64 to string with desired precision (e.g., 2 decimal places)
// 	strVal := strconv.FormatFloat(f, 'f', -2, 64)

// 	var numeric pgtype.Numeric
// 	numeric.Scan(strVal)

// 	return numeric

// }

func ConvertFloat64ToNumeric(f float64) pgtype.Numeric {
	// Convert float64 to string with desired precision (e.g., 2 decimal places)
	strVal := strconv.FormatFloat(f, 'f', -2, 64)

	// Create a new pgtype.Numeric object
	numeric := pgtype.Numeric{}

	// Scan the string representation using pgtype.Scan
	err := numeric.Scan(strVal)
	if err != nil {
		fmt.Println(err)
	}

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

func ConvertFloat64ToInt8(f int64) pgtype.Int8 {

	int8Pgtype := pgtype.Int8{
		Int64: f,
		Valid: true,
	}

	fmt.Println(int8Pgtype)
	return int8Pgtype
}

func StringToInt32(s string) int32 {
	// Parse the string to an int64 first
	i64, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	// Convert the int64 to int32

	fmt.Println("PCONVERT")
	fmt.Println(int32(i64))
	return int32(i64)
}
