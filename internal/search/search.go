package search

import (
	"strings"

	"github.com/limanmys/inventory-server/internal/database"
	"gorm.io/gorm"
)

func Search(query string, tx *gorm.DB) {
	words := strings.Fields(query)
	tx.Statement.Parse(tx.Statement.Model)
	fields := []string{}
	if tx.Statement.Table != "" {
		fields = append(fields, "\""+tx.Statement.Table+"\".*")
	}
	for _, join := range tx.Statement.Joins {
		if strings.Contains(join.Name, "JOIN") {
			continue
		}
		fields = append(fields, "\""+join.Name+"\".*")
	}
	concat := strings.Join(fields, ", ' ',")
	tmpTx := database.Connection()
	tmpTx = tmpTx.Where("CONCAT("+concat+") ilike ?", "%"+query+"%")
	for _, word := range words {
		tmpTx = tmpTx.Or("CONCAT("+concat+") ilike ?", "%"+word+"%")
	}
	tx.Where(tmpTx)
}
