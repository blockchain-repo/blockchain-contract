package constdef

const (
	Task_Unknown = iota
	Task_Enquiry
	Task_Action
	Task_Decision
	Task_DecisionCandidate
	Task_Plan
	Task_DelegateTaskGroup
)

var TaskType = map[int]string{
	Task_Unknown : "Task_Unknown",
	Task_Enquiry: "Task_Enquiry",
	Task_Action: "Task_Action",
	Task_Decision: "Task_Decision",
	Task_DecisionCandidate: "Task_DecisionCandidate",
	Task_Plan: "Task_Plan",
	Task_DelegateTaskGroup: "Task_DelegateTaskGroup",
}