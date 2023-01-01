package mongoapi

import (
	"crayontool-go/pkg/constant"
	"crayontool-go/pkg/logger"
	"fmt"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type PingReq struct {
	iCommandReq
}

type PingResult struct {
	commandResult
}

func (p *PingReq) ParseToBSONCmd() (bsonx.IDoc, error) {
	doc, err := p.iCommandReq.ParseToBSONCmd()
	if err != nil {
		return nil, err
	}
	doc = bsonx.Document(doc).Document().Prepend(p.GetCommandStr(), bsonx.Boolean(constant.True))
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
		logger.Errorf("%s err: %+v\n", RunCommandResultErr.Error(), err)
		/*return &commandResp{
		    err: result.Err(),
		}*/
		p.setRespErr(result.Err())
		return p
	}
	fmt.Println(result.DecodeBytes())
	//return &commandResp{singleResult: result}
	p.setRespSingleResult(result)
	return p
}
