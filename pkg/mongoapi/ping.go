package mongoapi

import (
	"fmt"
	"github.com/SuCrayon/crayontool-go/pkg/logger"
	"github.com/SuCrayon/crayontool-go/pkg/mongoapi/typereq"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// reference: https://www.mongodb.com/docs/manual/reference/command/ping/#mongodb-dbcommand-dbcmd.ping

type PingReq struct {
	iCommandReq
}

type PingResult struct {
	CommandResult `bson:",inline"`
}

func (p *PingReq) ParseToBSONCmd() (bsonx.IDoc, error) {
	doc, err := p.iCommandReq.ParseToBSONCmd()
	if err != nil {
		return nil, err
	}
	oneVal := typereq.IntOne{}
	doc = bsonx.Document(doc).Document().Prepend(p.GetCommandStr(), oneVal.ToBSON())
	return doc, nil
}

func (p *PingReq) GetResult() (*PingResult, error) {
	resp, err := p.getResp()
	if err != nil {
		return nil, err
	}
	if resp.Error() != nil {
		return nil, resp.Error()
	}
	var ret PingResult
	err = resp.Decode(&ret)
	return &ret, err
}

func (p *PingReq) Do() *PingReq {
	defer p.iCommandReq.Do()
	cmd, err := p.ParseToBSONCmd()
	if err != nil {
		logger.Errorf("%s err: %+v\n", ParseToBSONCmdErr.Error(), err)
		/*return &commandResp{
		    err: err,
		}*/
		p.setRespErr(err)
		return p
	}
	logger.Debugf("command doc: %v\n", cmd)
	result, err := p.GetCtl().RunCommand(p.GetDatabase(), cmd)
	if err != nil {
		logger.Errorf("%s err: %+v\n", RunCommandErr.Error(), err)
		/*return &commandResp{
		    err: err,
		}*/
		p.setRespErr(err)
		return p
	}
	if result.Err() != nil {
		logger.Errorf("%s err: %+v\n", RunCommandResultErr.Error(), result.Err())
		/*return &commandResp{
		    err: result.Err(),
		}*/
		p.setRespErr(result.Err())
		return p
	}
	logger.Debug(fmt.Sprint(result.DecodeBytes()))
	p.setRespSingleResult(result)
	//return &commandResp{singleResult: result}
	return p
}
