package mongoapi

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type MongoCtl interface {
	Client() (*mongo.Client, error)
	Connect() error
	Disconnect() error
	ApplyURI(uri string) error
	ApplyURIAndConnect(uri string) error
	RunCommand(database string, runCommand interface{}, opts ...*options.RunCmdOptions) (*mongo.SingleResult, error)
	RunCommandCursor(database string, runCommand interface{}, opts ...*options.RunCmdOptions) (*mongo.Cursor, error)
	DiagnosticCommander() DiagnosticCommander
	RoleCommander() RoleCommander
}

type mongoCtl struct {
	ctx    context.Context
	once   sync.Once
	client *mongo.Client
}

/*type IType interface {
    ToBSON() bsonx.Val
}*/

var (
	ClientNotInitErr = errors.New("mongo client not init")
)

func NewMongoCtl(ctx context.Context) MongoCtl {
	return &mongoCtl{
		ctx: ctx,
	}
}

func (ctl *mongoCtl) Client() (*mongo.Client, error) {
	if ctl.client == nil {
		return nil, ClientNotInitErr
	}
	return ctl.client, nil
}

func (ctl *mongoCtl) Connect() error {
	if ctl.client == nil {
		return ClientNotInitErr
	}
	return ctl.client.Connect(ctl.ctx)
}

func (ctl *mongoCtl) Disconnect() error {
	if ctl.client == nil {
		return ClientNotInitErr
	}
	return ctl.client.Disconnect(ctl.ctx)
}

func (ctl *mongoCtl) ApplyURI(uri string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	ctl.client = client
	return nil
}

func (ctl *mongoCtl) ApplyURIAndConnect(uri string) error {
	if err := ctl.ApplyURI(uri); err != nil {
		return err
	}
	return ctl.Connect()
}

func (ctl *mongoCtl) DiagnosticCommander() DiagnosticCommander {
	return GetDiagnosticCommander(ctl)
}

func (ctl *mongoCtl) RoleCommander() RoleCommander {
	return GetRoleCommander(ctl)
}

func (ctl *mongoCtl) RunCommand(database string, runCommand interface{}, opts ...*options.RunCmdOptions) (*mongo.SingleResult, error) {
	client, err := ctl.Client()
	if err != nil {
		return nil, err
	}
	return client.Database(database).RunCommand(ctl.ctx, runCommand, opts...), nil
}

func (ctl *mongoCtl) RunCommandCursor(database string, runCommand interface{}, opts ...*options.RunCmdOptions) (*mongo.Cursor, error) {
	client, err := ctl.Client()
	if err != nil {
		return nil, err
	}
	return client.Database(database).RunCommandCursor(ctl.ctx, runCommand, opts...)
}
