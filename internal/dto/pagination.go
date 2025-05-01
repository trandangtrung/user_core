package dto

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"encoding/json"

	"gorm.io/gorm"
)

// WithinParam represents a range filter with optional to value.
// e.g. from_date and to_date become Name="date", From and To strings.
type WithinParam struct {
	Name string  `json:"name"`
	From string  `json:"from"`
	To   *string `json:"to,omitempty"`
}

// Filter represents a single name/value filter pair.
type Filter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// PaginationDto holds pagination and filtering parameters.
type PaginationDto struct {
	Page         int           `json:"page"`
	Limit        int           `json:"limit"`
	Keyword      string        `json:"keyword"`
	OrderDir     string        `json:"order_dir,omitempty"`
	OrderBy      string        `json:"order_by"`
	Filters      []Filter      `json:"filters,omitempty"`
	CatIDs       []int         `json:"cat_id,omitempty"`
	WithinParams []WithinParam `json:"within_params,omitempty"`
	FromDate     string        `json:"from_date,omitempty"`
	ToDate       string        `json:"to_date,omitempty"`
}

// NewPaginationDto creates a PaginationDto from HTTP request query parameters.
func NewPaginationDto(q url.Values) *PaginationDto {

	dto := &PaginationDto{
		Keyword:  q.Get("keyword"),
		OrderBy:  firstNonEmpty(q.Get("order_by"), "id"),
		OrderDir: strings.ToUpper(firstNonEmpty(q.Get("order_dir"), "ASC")),
	}

	// Parse page and limit
	if page, err := strconv.Atoi(q.Get("page")); err == nil {
		dto.Page = page
	} else {
		dto.Page = 1
	}
	if limit, err := strconv.Atoi(q.Get("limit")); err == nil {
		dto.Limit = limit
	} else {
		dto.Limit = 10
	}

	// Parse category IDs
	for _, s := range q["cat_id"] {
		if id, err := strconv.Atoi(s); err == nil {
			dto.CatIDs = append(dto.CatIDs, id)
		}
	}

	dto.setWithinParams(q)

	// Lấy các filter khác
	for key, values := range q {
		if strings.HasPrefix(key, "from_") || strings.HasPrefix(key, "to_") ||
			key == "page" || key == "limit" || key == "order_by" || key == "order_dir" || key == "keyword" || key == "cat_id" {
			continue
		}
		for _, v := range values {
			dto.PushFilter(key, v)
		}
	}

	return dto
}

// firstNonEmpty returns the first non-empty string among inputs.
func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}

// setWithinParams processes query values for keys starting with "from_" or "to_", grouping them by suffix.
func (p *PaginationDto) setWithinParams(q map[string][]string) {
	ranges := make(map[string]WithinParam)

	// collect from and to values
	for key, vals := range q {
		switch {
		case strings.HasPrefix(key, "from_"):
			name := strings.TrimPrefix(key, "from_")
			r := ranges[name]
			r.Name = name
			r.From = vals[0]
			ranges[name] = r

		case strings.HasPrefix(key, "to_"):
			name := strings.TrimPrefix(key, "to_")
			r := ranges[name]
			r.Name = name
			val := vals[0]
			r.To = &val
			ranges[name] = r
		}
	}

	// append unique ranges
	for _, r := range ranges {
		p.WithinParams = append(p.WithinParams, r)
	}
}

// PushFilter appends a new Filter to the DTO.
func (p *PaginationDto) PushFilter(name, value string) {
	p.Filters = append(p.Filters, Filter{Name: name, Value: value})
}

// UpdateFilterName replaces the filter name of the first matching Filter.
func (p *PaginationDto) UpdateFilterName(oldName, newName string) {
	for i, f := range p.Filters {
		if f.Name == oldName {
			p.Filters[i].Name = newName
			return
		}
	}
}

// RemoveFilter removes the first Filter with the given name.
func (p *PaginationDto) RemoveFilter(name string) {
	for i, f := range p.Filters {
		if f.Name == name {
			p.Filters = append(p.Filters[:i], p.Filters[i+1:]...)
			return
		}
	}
}

// GetValue returns the filter value for a given filter name, unmarshalling JSON if applicable.
// Returns (value, true) if found, otherwise (nil, false).
func (p *PaginationDto) GetValue(name string) (interface{}, bool) {
	for _, f := range p.Filters {
		if f.Name == name {
			var out interface{}
			if err := json.Unmarshal([]byte(f.Value), &out); err == nil {
				return out, true
			}
			return f.Value, true
		}
	}
	return nil, false
}

// BuildWhere applies filter_name/filter_value, keyword search, and withinParams to the provided GORM DB.
func BuildWhere(db *gorm.DB, dto PaginationDto, table string, searchFields []string) *gorm.DB {
	// 1) Apply Filters
	for _, f := range dto.Filters {
		col := fmt.Sprintf("%s.%s", table, f.Name)
		db = db.Where(fmt.Sprintf("%s = ?", col), f.Value)
	}

	// 2) Apply keyword search
	keyword := strings.TrimSpace(dto.Keyword)
	if keyword != "" && len(searchFields) > 0 {
		keyword = strings.ToLower(keyword)
		var clauses []string
		var args []interface{}
		for _, field := range searchFields {
			col := fmt.Sprintf("%s.%s", table, field)
			clauses = append(clauses, fmt.Sprintf("lower(%s) LIKE ?", col))
			args = append(args, "%"+keyword+"%")
		}
		db = db.Where("("+strings.Join(clauses, " OR ")+")", args...)
	}

	// 3) Apply date ranges
	for _, wp := range dto.WithinParams {
		col := fmt.Sprintf("DATE(%s.%s)", table, wp.Name)
		if wp.To != nil {
			db = db.Where(fmt.Sprintf("%s BETWEEN DATE(?) AND DATE(?)", col), wp.From, *wp.To)
		} else {
			db = db.Where(fmt.Sprintf("%s >= DATE(?)", col), wp.From)
		}
	}

	return db
}

// Paginate applies pagination to the GORM DB.
func Paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}
