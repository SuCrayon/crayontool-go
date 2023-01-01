package mongoapi

import (
	"crayontool-go/pkg/logger"
	"fmt"
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
	RolesInfo                      []IRole
	ShowAuthenticationRestrictions bool
	ShowBuiltinRoles               bool
	ShowPrivileges                 bool
}

type RolesInfoResult struct {
	commandResult
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
	return bsonx.Doc{
		{
			Key:   r.GetCommandStr(),
			Value: bsonx.Int64(1),
		},
		{
			Key:   KeyShowAuthenticationRestrictions,
			Value: bsonx.Boolean(r.ShowAuthenticationRestrictions),
		},
		{
			Key:   KeyShowBuiltinRoles,
			Value: bsonx.Boolean(r.ShowBuiltinRoles),
		},
		{
			Key:   KeyShowPrivileges,
			Value: bsonx.Boolean(r.ShowPrivileges),
		},
	}, nil
}

func (r *RolesInfoReq) Do() *RolesInfoReq {
	cmd, err := r.ParseToBSONCmd()
	if err != nil {
		logger.Errorf("%s err: %+v\n", ParseToBSONCmdErr.Error(), err)
		return r
	}
	result, err := r.GetCtl().RunCommand(r.GetDatabase(), cmd)
	if err != nil {
		logger.Errorf("%s err: %+v\n", RunCommandErr.Error(), err)
		return r
	}
	if result.Err() != nil {
		logger.Errorf("%s err: %+v\n", RunCommandResultErr.Error(), err)
		return r
	}
	fmt.Println(result.DecodeBytes())
	return r
}
