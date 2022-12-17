package types

// QueryAddition 查询请求附加参数
type QueryAddition struct {
	// Fields 选择字段，,分隔的字符串列表
	Fields string `json:"fields,optional" form:"fields,optional"`
	// Sort 排序方式，根据哪些字段排序，,分隔的字符串列表
	// eg: sort=+name,-age 按name升序，age降序
	Sort string `json:"sort,optional" form:"sort,optional"`
	// Offset 偏移，用于分页
	Offset int64 `json:"offset,optional" form:"offset,optional"`
	// Limit 查询条数，用于分页
	Limit int64 `json:"limit,optional" form:"limit,optional"`
}
