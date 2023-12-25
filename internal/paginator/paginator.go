package paginator

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Filter struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

type Paginator struct {
	DB      *gorm.DB
	OrderBy []string
	Page    int
	PerPage int
	Filter  []Filter
}

type Data struct {
	TotalRecords int64       `json:"total_records"`
	Records      interface{} `json:"records"`
	CurrentPage  int         `json:"current_page"`
	TotalPages   int         `json:"total_pages"`
}

func New(tx *gorm.DB, c *fiber.Ctx) *Paginator {
	return &Paginator{
		DB:      tx,
		OrderBy: []string{convertOrder(getKey("sort", "-updated_at", c))},
		Page:    getIntKey("page", 1, c),
		PerPage: getIntKey("per_page", 15, c),
		Filter:  getFilters("filter", c),
	}
}

func (p *Paginator) Paginate(dataSource interface{}) (*Data, error) {
	db := p.DB

	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	var output Data
	var count int64
	var offset int

	for _, filter := range p.Filter {
		// We do this because of the type of datasource is not model, it's []model
		// Double elem is used to get the model type
		v := reflect.TypeOf(dataSource).Elem().Elem()
		for fieldIterator := 0; fieldIterator < v.NumField(); fieldIterator++ {
			// We need this check for invalid filter keys
			if v.Field(fieldIterator).Tag.Get("json") == filter.Key {
				rt := reflect.TypeOf(filter.Value)
				switch rt.Kind() {
				case reflect.Slice:
					db = db.Where(
						filter.Key+" IN ?",
						filter.Value.([]interface{}),
					)
				default:
					db = db.Where(
						"LOWER("+filter.Key+"::text) LIKE LOWER(?)",
						fmt.Sprintf("%%%s%%", filter.Value),
					)
				}
			}
		}
	}

	// Disable preloading temporarily for count query
	tmp := db.Statement.Preloads
	db.Statement.Preloads = map[string][]interface{}{}
	err := db.Model(dataSource).Count(&count).Error
	db.Statement.Preloads = tmp
	if err != nil {
		return nil, err
	}

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.PerPage
	}

	db = db.Limit(p.PerPage).Offset(offset)

	// Run final query
	err = db.Find(dataSource).Error
	if err != nil {
		return nil, err
	}

	output.TotalRecords = count
	output.Records = dataSource
	output.CurrentPage = p.Page
	output.TotalPages = getTotalPages(p.PerPage, count)

	return &output, nil
}

func getTotalPages(perPage int, totalRecords int64) int {
	return int(math.Ceil(float64(totalRecords) / float64(perPage)))
}

func getKey(k string, d string, c *fiber.Ctx) string {
	if c.FormValue(k) != "" {
		return c.FormValue(k)
	}
	return d
}

func getIntKey(k string, d int, c *fiber.Ctx) int {
	key := getKey(k, fmt.Sprintf("%d", d), c)
	value, err := strconv.Atoi(key)
	if err != nil {
		return d
	}
	return value
}

func convertOrder(order string) string {
	if order[0] == '-' {
		return order[1:] + " desc"
	} else if order[0] == '+' {
		return order[1:] + " asc"
	}
	return order + " asc"
}

func getFilters(field_name string, c *fiber.Ctx) []Filter {
	filters := []Filter{}
	if c.Query(field_name) != "" {
		json.Unmarshal(
			[]byte(strings.TrimSpace(c.Query(field_name))),
			&filters,
		)
	}

	return filters
}
