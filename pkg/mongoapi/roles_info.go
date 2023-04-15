package mongoapi

import (
	"fmt"
	"github.com/SuCrayon/crayontool-go/pkg/logger"
	"github.com/SuCrayon/crayontool-go/pkg/mongoapi/typereq"
	"github.com/SuCrayon/crayontool-go/pkg/mongoapi/typeresp"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// reference: https://www.mongodb.com/docs/manual/reference/command/rolesInfo/

const (
	KeyShowAuthenticationRestrictions = "showAuthenticationRestrictions"
	KeyShowBuiltinRoles               = "showBuiltinRoles"
	KeyShowPrivileges                 = "showPrivileges"
)

type RolesInfoReq struct {
	iCommandReq
	// RolesInfo string, document, array, or integer
	RolesInfo                      typereq.IRolesInfo
	ShowAuthenticationRestrictions bool
	ShowBuiltinRoles               bool
	ShowPrivileges                 bool
}

func (r *RolesInfoReq) SetRolesInfo(rolesInfo typereq.IRolesInfo) *RolesInfoReq {
	r.RolesInfo = rolesInfo
	return r
}

func (r *RolesInfoReq) SetShowAuthenticationRestrictions(showAuthenticationRestrictions bool) *RolesInfoReq {
	r.ShowAuthenticationRestrictions = showAuthenticationRestrictions
	return r
}

func (r *RolesInfoReq) SetShowBuiltinRoles(showBuiltinRoles bool) *RolesInfoReq {
	r.ShowBuiltinRoles = showBuiltinRoles
	return r
}

func (r *RolesInfoReq) SetShowPrivileges(showPrivileges bool) *RolesInfoReq {
	r.ShowPrivileges = showPrivileges
	return r
}

type RolesInfoResult struct {
	CommandResult `bson:",inline"`
	Roles         []typeresp.TRole `bson:"roles" json:"roles" yaml:"roles" xml:"roles"`
}

func (r *RolesInfoReq) ParseToBSONCmd() (bsonx.IDoc, error) {
	/*return bsonx.MDoc{
	    CmdRolesInfo: bsonx.Int64(1),
	    KeyShowAuthenticationRestrictions: bsonx.Boolean(r.ShowAuthenticationRestrictions),
	    KeyShowBuiltinRoles: bsonx.Boolean(r.ShowBuiltinRoles),
	    KeyShowPrivileges: bsonx.Boolean(r.ShowPrivileges),
	}, nil*/
	// 不能用MDoc（无序文档），本身这个命令需要的文档是有序的
	/*
	   这是在mongo命令行执行的结果，说明文档本身的键值对需要有序，取的第一个key作为命令的
	   > db.runCommand({rolesInfo: 1, showAuthenticationRestrictions: true})
	   { "roles" : [ ], "ok" : 1 }

	   > db.runCommand({showAuthenticationRestrictions: true, rolesInfo: 1})
	   {
	           "ok" : 0,
	           "errmsg" : "no such command: 'showAuthenticationRestrictions'",
	           "code" : 59,
	           "codeName" : "CommandNotFound"
	   }
	*/
	iDoc, err := r.iCommandReq.ParseToBSONCmd()
	if err != nil {
		return nil, err
	}

	doc := bsonx.Document(iDoc).Document()
	doc = doc.Prepend(KeyShowPrivileges, bsonx.Boolean(r.ShowPrivileges))
	doc = doc.Prepend(KeyShowBuiltinRoles, bsonx.Boolean(r.ShowBuiltinRoles))
	doc = doc.Prepend(KeyShowAuthenticationRestrictions, bsonx.Boolean(r.ShowAuthenticationRestrictions))
	doc = doc.Prepend(r.GetCommandStr(), r.RolesInfo.ToBSON())

	return doc, nil
}

func (r *RolesInfoReq) GetResult() (*RolesInfoResult, error) {
	resp, err := r.getResp()
	if err != nil {
		return nil, err
	}
	if resp.Error() != nil {
		return nil, resp.Error()
	}
	var ret RolesInfoResult
	err = resp.Decode(&ret)
	return &ret, err
}

func (r *RolesInfoReq) Do() *RolesInfoReq {
	defer r.iCommandReq.Do()
	cmd, err := r.ParseToBSONCmd()
	if err != nil {
		logger.Errorf("%s err: %+v\n", ParseToBSONCmdErr.Error(), err)
		return r
	}
	logger.Debugf("command doc: %v\n", cmd)
	result, err := r.GetCtl().RunCommand(r.GetDatabase(), cmd)
	if err != nil {
		logger.Errorf("%s err: %+v\n", RunCommandErr.Error(), err)
		return r
	}
	if result.Err() != nil {
		logger.Errorf("%s err: %+v\n", RunCommandResultErr.Error(), result.Err())
		return r
	}
	logger.Debug(fmt.Sprint(result.DecodeBytes()))
	r.setRespSingleResult(result)
	return r
}
