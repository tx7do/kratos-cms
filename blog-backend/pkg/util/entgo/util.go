package entgo

import (
	"encoding/json"
	"strings"

	"entgo.io/ent/dialect/sql"

	"kratos-blog/pkg/util/stringcase"
)

type WhereConditions []func(s *sql.Selector)
type OrderConditions []func(*sql.Selector)

// QueryCommandToWhereConditions 查询命令转换为选择条件
func QueryCommandToWhereConditions(queryCmd map[string]string) WhereConditions {
	if queryCmd == nil {
		return nil
	}

	queryStr, ok := queryCmd["query"]
	if !ok {
		return nil
	}

	var queryMap map[string]string
	err := json.Unmarshal([]byte(queryStr), &queryMap)
	if err != nil {
		return nil
	}

	var whereCond WhereConditions
	for k, v := range queryMap {
		key := stringcase.ToSnakeCase(k)
		value := v
		whereCond = append(whereCond, func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(key), value))
		})
	}

	return whereCond
}

// QueryCommandToOrderConditions 查询命令转换为排序条件
func QueryCommandToOrderConditions(orderByCmd map[string]string) OrderConditions {
	if orderByCmd == nil {
		return nil
	}

	orderByStr, ok := orderByCmd["orderBy"]
	if !ok {
		return nil
	}

	var orderByMap map[string]string
	err := json.Unmarshal([]byte(orderByStr), &orderByMap)
	if err != nil {
		return nil
	}

	var orderCond OrderConditions
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

	return orderCond
}

// QueryCommandToSelector 查询命令转换为选择器
func QueryCommandToSelector(queryCmd, orderByCmd map[string]string) (WhereConditions, OrderConditions) {
	return QueryCommandToWhereConditions(queryCmd), QueryCommandToOrderConditions(orderByCmd)
}
