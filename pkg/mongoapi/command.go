package mongoapi

import (
	"crayontool-go/pkg/constant"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

const (
	// defaultDatabase 默认使用数据库，参照mongo命令行连接后使用db命令查看的数据库
	defaultDatabase = "test"
)

var (
	ParseToBSONCmdErr   = errors.New("parse to bson doc failed")
	RunCommandErr       = errors.New("run db command failed")
	RunCommandResultErr = errors.New("run db command's result has error")
	CommandReqNotDoErr  = errors.New("command req not do, maybe forget to call method [do]")
)

type iCommandReq interface {
	getResp() (*commandResp, error)
	setCommandStr(commandStr string) iCommandReq
	setRespErr(err error)
	setRespSingleResult(singleResult *mongo.SingleResult)
	setRespCursor(cursor *mongo.Cursor)
	Do()
	GetCtl() MongoCtl
	GetDatabase() string
	GetCommandStr() string
	GetWriteConcern() *string
	GetComment() *string
	SetCtl(ctl MongoCtl) iCommandReq
	SetDatabase(database string) iCommandReq
	SetWriteConcern(writeConcern string) iCommandReq
	SetComment(comment string) iCommandReq
	ParseToBSONCmd() (bsonx.IDoc, error)
}

type commandReq struct {
	resp         commandResp
	Ctl          MongoCtl
	CommandStr   string
	Database     string
	WriteConcern *string
	Comment      *string
}

func (c *commandReq) Do() {
	c.resp.mutex.Lock()
	defer c.resp.mutex.Unlock()
	c.resp.done = constant.True
}

func (c *commandReq) setRespErr(err error) {
	c.resp.err = err
}

func (c *commandReq) setRespSingleResult(singleResult *mongo.SingleResult) {
	c.resp.singleResult = singleResult
}

func (c *commandReq) setRespCursor(cursor *mongo.Cursor) {
	c.resp.cursor = cursor
}

func (c *commandReq) getResp() (*commandResp, error) {
	if !c.resp.Done() {
		return nil, CommandReqNotDoErr
	}
	return &c.resp, nil
}

func (c *commandReq) ParseToBSONCmd() (bsonx.IDoc, error) {
	// default implements
	doc := bsonx.Doc{}
	if c.GetWriteConcern() != nil {
		doc.Append(KeyWriteConcern, bsonx.String(*c.GetWriteConcern()))
	}
	if c.GetComment() != nil {
		doc.Append(KeyComment, bsonx.String(*c.GetComment()))
	}
	return doc, nil
}

func (c *commandReq) SetDatabase(database string) iCommandReq {
	c.Database = database
	return c
}

func (c *commandReq) setCommandStr(commandStr string) iCommandReq {
	c.CommandStr = commandStr
	return c
}

func (c *commandReq) SetWriteConcern(writeConcern string) iCommandReq {
	c.WriteConcern = &writeConcern
	return c
}

func (c *commandReq) SetComment(comment string) iCommandReq {
	c.Comment = &comment
	return c
}

func (c *commandReq) SetCtl(ctl MongoCtl) iCommandReq {
	c.Ctl = ctl
	return c
}

func (c *commandReq) GetDatabase() string {
	return c.Database
}

func (c *commandReq) GetCommandStr() string {
	return c.CommandStr
}

func (c *commandReq) GetWriteConcern() *string {
	return c.WriteConcern
}

func (c *commandReq) GetComment() *string {
	return c.Comment
}

func (c *commandReq) GetCtl() MongoCtl {
	return c.Ctl
}

func NewDefaultCommandReq(ctl MongoCtl) iCommandReq {
	req := commandReq{
		Ctl:      ctl,
		Database: defaultDatabase,
	}
	return &req
}

type commander struct {
	ctl MongoCtl
}
