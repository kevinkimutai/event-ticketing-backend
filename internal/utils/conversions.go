package utils

import (
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
