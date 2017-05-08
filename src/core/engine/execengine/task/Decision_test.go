package task

import (
	"testing"
	"unicontract/src/core/engine/execengine/constdef"
)

func TestDecision(t *testing.T){
	des := &Decision{}
	des.InitDecision()
	des.SetCname("Decision")
    des.SetCaption("Decision")
	des.SetDescription("Decision Test")
	if des.GetName() != "Decision" {
		t.Error("Decision Error,GetName Error!")
	}
	if des.GetCaption() != "Decision" {
		t.Error("Decision Error,GetCaption Error!")
	}
	if des.GetDescription() != "Decision Test" {
		t.Error("Decision Error,GetDescription Error!")
	}
	if des.GetCtype() != constdef.ComponentType[constdef.Component_Task]+"."+constdef.TaskType[constdef.Task_Decision] {
		t.Error("Decision Error,GetCtype Error!")
	}
    //Test AddCandidate
    var cand_1 DecisionCandidate = DecisionCandidate{}
	cand_1.InitDecisionCandidate()
	cand_1.SetCname("Candidate_1")
	cand_1.SetCaption("Candidate_1")
	cand_1.SetDescription("Test Candidate_1")
	var cand_2 DecisionCandidate = DecisionCandidate{}
	cand_2.InitDecisionCandidate()
	cand_2.SetCname("Candidate_2")
	cand_2.SetCaption("Candidate_2")
	cand_2.SetDescription("Test Candidate_2")
	des.AddCandidate(cand_1)
	des.AddCandidate(cand_2)
	//Test GetCandidate
	if len(des.GetCandidateList()) != 2 {
		t.Error("Add DecisionCandidate Error!")
	}
	//Test GetCandidate
	if des.GetCandidate("Candidate_1").GetName() != "Candidate_1" {
		t.Error("GetCandidate Error, get second element Error!")
	}
	if des.GetCandidate("Candidate_2").GetName() != "Candidate_2" {
		t.Error("GetCandidate Error, get second element Error!")
	}
	//Test RemoveCandidate
	des.RemoveCandidate(cand_1)
	if len(des.GetCandidateList()) != 1 {
		t.Error("RemoveCandidate Error!")
	}

}
