package entgo

import (
	"encoding/json"
	"strings"

	"entgo.io/ent/dialect/sql"

	"kratos-blog/pkg/util/stringcase"
)

type WhereConditions []func(s *sql.Selector)
type OrderConditions []func(*sql.Selector)

// QueryCommandToSelector 查询命令转换为选择器
func QueryCommandToSelector(queryCmd, orderByCmd map[string]string) (WhereConditions, OrderConditions) {
	var whereCond WhereConditions
	var orderCond OrderConditions

	if queryCmd != nil {
		queryStr, ok := queryCmd["query"]
		if ok {
			var queryMap map[string]string
			err := json.Unmarshal([]byte(queryStr), &queryMap)
			if err == nil {
				for k, v := range queryMap {
					key := stringcase.ToSnakeCase(k)
					value := v
					whereCond = append(whereCond, func(s *sql.Selector) {
						s.Where(sql.EQ(s.C(key), value))
					})
				}
			}
		}
	}

	if orderByCmd != nil {
		orderByStr, ok := orderByCmd["orderBy"]
		if ok {
			var orderByMap map[string]string
			err := json.Unmarshal([]byte(orderByStr), &orderByMap)
			if err == nil {
				for k, v := range orderByMap {
					key := stringcase.ToSnakeCase(k)
					switch strings.ToLower(v) {
					case "desc", "descending", "descend":
						orderCond = append(orderCond, func(s *sql.Selector) {
							s.OrderBy(sql.Desc(s.C(key)))
						})
					case "asc", "ascending", "ascend":
						orderCond = append(orderCond, func(s *sql.Selector) {
							s.OrderBy(sql.Asc(s.C(key)))
						})
					}
				}
			}
		}
	}

	return whereCond, orderCond
}
