package collection_endpoint

import (
	"errors"
	"github.com/rs/zerolog/log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type QueryInfo struct {
	spec *QuerySpec

	fields []string
	// field => desc|asc
	sorts map[string]string
	// limit
	limit  int
	offset int

	// field[op] => value
	filter map[string]string
}

func NewQueryInfo(spec *QuerySpec, query url.Values) (QueryInfo, error) {
	for k, v := range spec.Default {
		if !query.Has(k) {
			query.Set(k, v)
		}
	}

	filterRe := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*\[[a-zA-Z_][a-zA-Z_]*]$`)
	var err error
	info := QueryInfo{
		spec:   spec,
		sorts:  map[string]string{},
		filter: map[string]string{},
	}

	for k, v := range query {
		switch k {
		case "field", "fields":
			info.fields = strings.Split(v[0], ",")
		case "sort":
			for _, s := range strings.Split(v[0], ",") {
				if strings.HasPrefix(s, "-") {
					info.sorts[s[1:]] = "desc"
				} else if strings.HasPrefix(s, "+") {
					info.sorts[s[1:]] = ""
				} else {
					info.sorts[s] = ""
				}
			}
		case "limit":
			info.limit, err = strconv.Atoi(v[0])
			if err != nil {
				return info, err
			}
		case "offset":
			info.offset, err = strconv.Atoi(v[0])
			if err != nil {
				return info, err
			}
		default:
			if filterRe.Match([]byte(k)) {
				info.filter[k] = v[0]
			} else {
				log.Warn().Str("key", k).Msg("unknown key")
			}
		}
	}

	return info, info.verify()
}
func (q *QueryInfo) verify() error {
	if len(q.fields) == 0 {
		return errors.New("fields is empty")
	}
	if q.limit <= 0 {
		return errors.New("limit must be greater or equal to 0")
	}
	if q.offset < 0 {
		return errors.New("offset must be greater than 0")
	}
	return nil
}

func (q *QueryInfo) BuildSQL() (data, total string, err error) {
	sqlb := strings.Builder{}

	sqlb.WriteString("select " + strings.Join(q.fields, ", "))
	sqlb.WriteString(" from " + q.spec.Table)
	sqlb.WriteString(" where 1 = 1")

	if len(q.filter) > 0 {
		filters := []string{}
		for k, v := range q.filter {
			filters = append(filters, k+"="+v)
		}
		if err = parseAndWriteConditions(strings.Join(filters, "&"), q.spec.Filter, &sqlb); err != nil {
			return
		}
	}
	total = `select count(*) as total from (` + sqlb.String() + `) as t`

	if len(q.sorts) > 0 {
		sqlb.WriteString(" order by ")
		sorts := []string{}

		for k, v := range q.sorts {
			if v != "" {
				sorts = append(sorts, k+" "+v)
			} else {
				sorts = append(sorts, k)
			}
		}
		sqlb.WriteString(strings.Join(sorts, ", "))
	}

	if q.limit > 0 {
		sqlb.WriteString(" limit " + strconv.Itoa(q.limit))
		if q.offset > 0 {
			sqlb.WriteString(" offset " + strconv.Itoa(q.offset))
		}
	}

	jsonObjectArgs := []string{}
	for _, field := range q.fields {
		jsonObjectArgs = append(jsonObjectArgs, "'"+field+"'")
		jsonObjectArgs = append(jsonObjectArgs, "t."+field)
	}
	data = `select json_arrayagg(json_object(` + strings.Join(jsonObjectArgs, ", ") + `)) as data from (` + sqlb.String() + `) as t group by 1`
	return
}