package mongoapi

import (
	"go.mongodb.org/mongo-driver/bson"
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

type commandResult struct {
	OK int64 `bson:"ok"`
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
