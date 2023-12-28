package counter

import (
	"github.com/limanmys/inventory-server/internal/database"
)

func Get(table string) (map[string]int64, error) {
	var count int64
	if err := database.Connection().Table(table).Count(&count).Error; err != nil {
		return nil, err
	}

	return map[string]int64{"count": count}, nil
}
