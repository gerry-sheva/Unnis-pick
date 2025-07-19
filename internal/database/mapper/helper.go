package mapper

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func StringToUUID(id string) (uuid.UUID, error) {
	var uuid uuid.UUID
	return uuid, uuid.Scan(id)
}

func Float64ToNumeric(value float64) (pgtype.Numeric, error) {
	var numeric pgtype.Numeric
	str := strconv.FormatFloat(value, 'f', -1, 64)
	return numeric, numeric.Scan(str)
}
