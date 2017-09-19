package task

import (
	"fmt"

	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
)

//决策组件使用规则：
//   Text中存放所有待判定的条件，决策计算过程中，遍历Text条件
//         成立的条件填充到SupportArguments中，并计数SupportNum+1
//       不成立的条件填充到AgainstArguments中，并计数AgainstNum+1
//   最终SupportNum 和 AgainstNum比较得出决策结果 Decision
type DecisionCandidate struct {
	Enquiry
	Text             []string `json:"Text"`
	SupportArguments []string `json:"SupportArguments"`
	AgainstArguments []string `json:"AgainstArguments"`
	SupportNum       int      `json:"SupportNum"`
	AgainstNum       int      `json:"AgainstNum"`
	Result           int      `json:"Result"`
}

const (
	_Text             = "_Text"
	_SupportArguments = "_SupportArguments"
	_AgainstArguments = "_AgainstArguments"
	_SupportNum       = "_SupportNum"
	_AgainstNum       = "_AgainstNum"
	_Result           = "_Result"
)

func NewDecisionCandidate() *DecisionCandidate {
	d := &DecisionCandidate{}
	return d
}

//===============接口实现===================
func (dc DecisionCandidate) SetContract(p_contract inf.ICognitiveContract) {
	dc.Enquiry.SetContract(p_contract)
}

func (dc DecisionCandidate) GetContract() inf.ICognitiveContract {
	return dc.Enquiry.GetContract()
}

func (dc DecisionCandidate) CleanValueInProcess() {
	dc.Enquiry.CleanValueInProcess()
	dc.ResetDecisionCandidate()
}

//===============描述态=====================

//===============运行态=====================
func (dc *DecisionCandidate) InitDecisionCandidate() error {
	var err error = nil
	err = dc.InitEnquriy()
	if err != nil {
		uniledgerlog.Error("InitDecisionCandidate fail[" + err.Error() + "]")
		return err
	}
	dc.SetCtype(constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_DecisionCandidate])
	//text
	if dc.Text == nil {
		dc.Text = make([]string, 0)
	}
	map_Text := make(map[string]string, 0)
	for _, p_text := range dc.Text {
		map_Text[p_text] = p_text
	}
	common.AddProperty(dc, dc.PropertyTable, _Text, map_Text)
	//supportArguments
	if dc.SupportArguments == nil {
		dc.SupportArguments = make([]string, 0)
	}
	map_supportArgument := make(map[string]string, 0)
	for _, p_support := range dc.SupportArguments {
		map_supportArgument[p_support] = p_support
	}
	common.AddProperty(dc, dc.PropertyTable, _SupportArguments, map_supportArgument)
	//againstArguments
	if dc.AgainstArguments == nil {
		dc.AgainstArguments = make([]string, 0)
	}
	map_againstArgument := make(map[string]string, 0)
	for _, p_against := range dc.AgainstArguments {
		map_againstArgument[p_against] = p_against
	}
	common.AddProperty(dc, dc.PropertyTable, _AgainstArguments, map_againstArgument)
	//SupportNum
	common.AddProperty(dc, dc.PropertyTable, _SupportNum, dc.SupportNum)
	//AgainstNum
	common.AddProperty(dc, dc.PropertyTable, _AgainstNum, dc.AgainstNum)
	//Decision
	common.AddProperty(dc, dc.PropertyTable, _Result, dc.Result)
	return err
}

//==========Getter方法
func (dc *DecisionCandidate) GetSupportArguments() []string {
	if dc.PropertyTable[_SupportArguments] == nil {
		return nil
	}
	supportargument_property, ok := dc.PropertyTable[_SupportArguments].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	supportargument_value, ok := supportargument_property.GetValue().(map[string]string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}

	var supportargument []string
	for _, v := range supportargument_value {
		supportargument = append(supportargument, v)
	}
	return supportargument
}

func (dc *DecisionCandidate) GetAgainstArguments() []string {
	if dc.PropertyTable[_AgainstArguments] == nil {
		return nil
	}
	againstargument_property, ok := dc.PropertyTable[_AgainstArguments].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	againstargument_value, ok := againstargument_property.GetValue().(map[string]string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}

	var againstargument []string
	for _, v := range againstargument_value {
		againstargument = append(againstargument, v)
	}
	return againstargument
}

func (dc *DecisionCandidate) GetSupportNum() int {
	if dc.PropertyTable[_SupportNum] == nil {
		return 0
	}
	supportnum_property, ok := dc.PropertyTable[_SupportNum].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0
	}
	supportnum_value, ok := supportnum_property.GetValue().(int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0
	}
	return supportnum_value
}

func (dc *DecisionCandidate) GetAgainstNum() int {
	if dc.PropertyTable[_AgainstNum] == nil {
		return 0
	}
	againstnum_property, ok := dc.PropertyTable[_AgainstNum].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0
	}
	againstnum_value, ok := againstnum_property.GetValue().(int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0
	}
	return againstnum_value
}

func (dc *DecisionCandidate) GetResult() int {
	if dc.PropertyTable[_Result] == nil {
		return 0
	}
	decision_property, ok := dc.PropertyTable[_Result].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0
	}
	decision_value, ok := decision_property.GetValue().(int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0
	}
	return decision_value
}

//==========Setter方法
func (dc *DecisionCandidate) SetSupportArguments(arr_supportarguments []string) {
	dc.SupportArguments = arr_supportarguments
	supportargument_property, ok := dc.PropertyTable[_SupportArguments].(property.PropertyT)
	if !ok {
		supportargument_property = *property.NewPropertyT(_SupportArguments)
	}
	supportargument_property.SetValue(arr_supportarguments)
	dc.PropertyTable[_SupportArguments] = supportargument_property
}

func (dc *DecisionCandidate) SetAgainstArguments(arr_againstarguments []string) {
	dc.AgainstArguments = arr_againstarguments
	againstargument_property, ok := dc.PropertyTable[_AgainstArguments].(property.PropertyT)
	if !ok {
		againstargument_property = *property.NewPropertyT(_AgainstArguments)
	}
	againstargument_property.SetValue(arr_againstarguments)
	dc.PropertyTable[_AgainstArguments] = againstargument_property
}

func (dc *DecisionCandidate) SetSupportNum(int_supportnum int) {
	dc.SupportNum = int_supportnum
	supportnum_property, ok := dc.PropertyTable[_SupportNum].(property.PropertyT)
	if !ok {
		supportnum_property = *property.NewPropertyT(_SupportNum)
	}
	supportnum_property.SetValue(int_supportnum)
	dc.PropertyTable[_SupportNum] = supportnum_property
}

func (dc *DecisionCandidate) SetAgainstNum(int_againstnum int) {
	dc.AgainstNum = int_againstnum
	againstnum_property, ok := dc.PropertyTable[_AgainstNum].(property.PropertyT)
	if !ok {
		againstnum_property = *property.NewPropertyT(_AgainstNum)
	}
	againstnum_property.SetValue(int_againstnum)
	dc.PropertyTable[_AgainstNum] = againstnum_property
}

func (dc *DecisionCandidate) SetResult(int_decision int) {
	dc.Result = int_decision
	decision_property, ok := dc.PropertyTable[_Result].(property.PropertyT)
	if !ok {
		decision_property = *property.NewPropertyT(_Result)
	}
	decision_property.SetValue(int_decision)
	dc.PropertyTable[_Result] = decision_property
}

//==========Add 方法
func (dc *DecisionCandidate) AddText(p_strarr []string) {
	if p_strarr != nil {
		text_property, ok := dc.PropertyTable[_Text].(property.PropertyT)
		if !ok {
			text_property = *property.NewPropertyT(_Text)
		}
		if text_property.GetValue() == nil {
			text_property.SetValue(make([]string, 0))
		}
		map_text, ok := text_property.GetValue().(map[string]string)
		if !ok {
			map_text = make(map[string]string, 0)
		}
		for _, v_Text := range p_strarr {
			map_text[v_Text] = v_Text
		}
		text_property.SetValue(map_text)
		dc.PropertyTable[_Text] = text_property
	}
}

func (dc *DecisionCandidate) ShowText() {
	text_property, ok := dc.PropertyTable[_Text].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	if text_property.GetValue() != nil {
		map_text, ok := text_property.GetValue().(map[string]string)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return
		}
		for _, v_Text := range map_text {
			fmt.Println(v_Text)
		}
	}
}

func (dc *DecisionCandidate) ResetDecisionCandidate() {
	dc.SetSupportArguments(make([]string, 0))
	dc.SetAgainstArguments(make([]string, 0))
	dc.SetSupportNum(0)
	dc.SetAgainstNum(0)
	dc.SetResult(0)
}

func (dc *DecisionCandidate) AddSupportArgument(p_Support string) {
	if p_Support != "" {
		supports_property, ok := dc.PropertyTable[_SupportArguments].(property.PropertyT)
		if !ok {
			supports_property = *property.NewPropertyT(_SupportArguments)
		}
		if supports_property.GetValue() == nil {
			supports_property.SetValue(make(map[string]string, 0))
		}
		map_supports, ok := supports_property.GetValue().(map[string]string)
		if !ok {
			map_supports = make(map[string]string, 0)
		}
		map_supports[p_Support] = p_Support
		supports_property.SetValue(map_supports)
		dc.PropertyTable[_SupportArguments] = supports_property
	}
}

func (dc *DecisionCandidate) AddAgainstArgument(p_against string) {
	if p_against != "" {
		against_property, ok := dc.PropertyTable[_AgainstArguments].(property.PropertyT)
		if !ok {
			against_property = *property.NewPropertyT(_AgainstArguments)
		}
		if against_property.GetValue() == nil {
			against_property.SetValue(make(map[string]string, 0))
		}
		map_againsts, ok := against_property.GetValue().(map[string]string)
		if !ok {
			map_againsts = make(map[string]string, 0)
		}
		map_againsts[p_against] = p_against
		against_property.SetValue(map_againsts)
		dc.PropertyTable[_AgainstArguments] = against_property
	}
}

func (dc *DecisionCandidate) Eval() error {
	var support_sum int
	var against_sum int
	var err error

	text_property, ok := dc.PropertyTable[_Text].(property.PropertyT)
	if !ok {
		err = fmt.Errorf("Text Assert Error!")
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, err.Error()))
		return err
	}

	if text_property.GetValue() != nil {
		text_map, ok := text_property.GetValue().(map[string]string)
		if !ok {
			err = fmt.Errorf("Text's Value Assert Error!")
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, err.Error()))
			return err
		}
		for _, v_expression := range text_map {
			v_contract := dc.GetContract()

			uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), Decision evaluate expression is (%s)]",
				uniledgerlog.NO_ERROR, v_contract.GetContractId(), dc.GetName(), dc.GetTaskId(), v_expression))

			v_result, err := v_contract.EvaluateExpression(constdef.ExpressionType[constdef.Expression_Condition], v_expression)
			if err != nil {
				uniledgerlog.Error("DecisionCandidate.Eval fail[ %+v ]", err)
				return err
			}

			v_bool_result, ok := v_result.(bool)
			if !ok {
				err = fmt.Errorf("DecisionCandidate Result Assert Error!")
				uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, err.Error()))
				return err
			}

			if v_bool_result {
				support_sum++
				dc.AddSupportArgument(v_expression)
			} else {
				against_sum++
				dc.AddAgainstArgument(v_expression)
			}
		}
		uniledgerlog.Debug("support_sum is %d, against_sum is %d", support_sum, against_sum)
		dc.SetSupportNum(support_sum)
		dc.SetAgainstNum(against_sum)

		dc.SetResult(support_sum - against_sum)
	}
	return err
}
