package handler

import (
	"comm/logger"
	"comm/mark"
	"comm/util"
	"context"
	"encoding/json"
	pbDtm "proto/dtm"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/workflow"
)

var (
	dtmServer = "http://localhost:36789/api/dtmsvr"
	qsBusi    = "http://127.0.0.1:8080/api/dtm"
)

func (h *Handler) WorkflowGrpc(ctx context.Context, req *pbDtm.BusiReq, rsp *pbDtm.BusiReply) error {
	var err error
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "WorkflowGrpc")()
	workflow.InitHTTP(dtmServer, qsBusi+"/transResume")
	wfName := "workflow-http"
	err = workflow.Register(wfName, func(wf *workflow.Workflow, data []byte) error {
		var req map[string]interface{}
		err := json.Unmarshal(data, &req)
		if err != nil {
			logger.Errorf("WorkflowGrpc json.Unmarshal failed. [%s]", err.Error())
			return err
		}
		// Register branches with rollback functions
		_, err = wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
			_, err := wf.NewRequest().SetBody(req).Post(qsBusi + "/transOutRevert")
			if err != nil {
				logger.Errorf("WorkflowGrpc transOutRevert.Post failed. [%s]", err.Error())
				return err
			}
			return nil
		}).NewRequest().SetBody(req).Post(qsBusi + "/transOut")
		if err != nil {
			logger.Errorf("WorkflowGrpc transOut.Post failed. [%s]", err.Error())
			return err
		}
		// Register branches with rollback functions
		_, err = wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
			_, err := wf.NewRequest().SetBody(req).Post(qsBusi + "/transInRevert")
			if err != nil {
				logger.Errorf("WorkflowGrpc transInRevert.Post failed. [%s]", err.Error())
				return err
			}
			return nil
		}).NewRequest().SetBody(req).Post(qsBusi + "/transIn")
		if err != nil {
			logger.Errorf("WorkflowGrpc transIn.Post failed. [%s]", err.Error())
			return err
		}
		return err
	})
	if err != nil {
		logger.Errorf("WorkflowGrpc workflow.Register failed. [%s]", err.Error())
		return err
	}

	gid := util.RandomString(12)
	if err != nil {
		logger.Errorf("WorkflowGrpc json.Marshal failed. [%s]", err.Error())
		return err
	}
	req1 := map[string]interface{}{"amount": 30}
	data, err := json.Marshal(req1)
	if err != nil {
		logger.Errorf("WorkflowGrpc json.Marshal failed. [%s]", err.Error())
		return err
	}
	err = workflow.Execute(wfName, gid, data)
	logger.Infof("workflow.Execute result is: %v", err)

	return nil
}
