package bsonkey

const (
	KeyComment      = "comment"
	KeyWriteConcern = "writeConcern"
)

// BSON Key Name
const (
	KeyRole        = "role"
	KeyDB          = "db"
	KeyCollection  = "collection"
	KeyCluster     = "cluster"
	KeyAnyResource = "anyResource"
)

// Query and Projection Operators

// Comparison Query Operators

const (
	OprEq  = "$eq"
	OprGt  = "$gt"
	OprGte = "$gte"
	OprIn  = "$in"
	OprLt  = "$lt"
	OprLte = "$lte"
	OprNe  = "$ne"
	OprNin = "$nin"
)

// Logical Query Operators

const (
	OprAnd = "$and"
	OprNot = "$not"
	OprNor = "$nor"
	OprOr  = "$or"
)
