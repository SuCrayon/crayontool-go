package mongoapi

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

/*type CommandResp interface {
    Error() error
    Decode(v interface{}) error
    DecodeBytes() (bson.Raw, error)
}*/

type commandResp struct {
	done         bool
	mutex        sync.RWMutex
	err          error
	singleResult *mongo.SingleResult
	cursor       *mongo.Cursor
}

type Signature struct {
	Hash  primitive.Binary `bson:"hash" json:"hash" yaml:"hash" xml:"hash"`
	KeyID int64            `bson:"keyId" json:"keyId" yaml:"keyId" xml:"keyId"`
}

// ClusterTime Only for replica sets and sharded clusters. For internal use only.
type ClusterTime struct {
	ClusterTime primitive.Timestamp `bson:"clusterTime" json:"clusterTime" yaml:"clusterTime" xml:"clusterTime"`
	Signature   Signature           `bson:"signature" json:"signature" yaml:"signature" xml:"signature"`
}

// CommandResultErr 通用错误返回
type CommandResultErr struct {
	ErrMsg   string `bson:"errmsg" json:"errmsg" yaml:"errmsg" xml:"errmsg"`
	Code     int64  `bson:"code" json:"code" yaml:"code" xml:"code"`
	CodeName string `bson:"codeName" json:"codeName" yaml:"codeName" xml:"codeName"`
}

// CommandResult 通用字段返回 reference: https://www.mongodb.com/docs/manual/reference/method/db.runCommand/#response
type CommandResult struct {
	CommandResultErr `bson:",inline"`
	OK               uint8 `bson:"ok" json:"ok" yaml:"ok" xml:"ok"`
	// OperationTime Only for replica sets and sharded clusters.
	OperationTime primitive.Timestamp `bson:"operationTime" json:"operationTime" yaml:"operationTime" xml:"operationTime"`
	// ClusterTime Only for replica sets and sharded clusters. For internal use only.
	ClusterTime ClusterTime `bson:"$clusterTime" json:"$clusterTime" yaml:"$clusterTime" xml:"$clusterTime"`
}

func (c *commandResp) Error() error {
	return c.err
}

func (c *commandResp) DecodeBytes() (bson.Raw, error) {
	return c.singleResult.DecodeBytes()
}

func (c *commandResp) Decode(v interface{}) error {
	return c.singleResult.Decode(v)
}

func (c *commandResp) Done() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.done
}
