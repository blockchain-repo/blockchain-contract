#############################################
##########  表达式解析规则：   ##############
#############################################
1. 表达式分类-大类： 在engine/execengine/constdef/ExpressionType.go中定义
   1）常量表达式
   2）变量表达式
   3）条件表达式
   4）函数表达式
   5）决策候选条件表示式
2. 表达式分类-小类： 在engine/execengine/constdef/ExpressionClassify.go中定义
   1）实数表达式
   2）浮点数表达式
   3）布尔表达式
   4）字符串表达式
   5）日期表达式
   6）数组表达式
   7）条件表达式
   8）函数表达式
   9）变量表达式
   10）决策候选条件表达式
3. 变量表达式命名规则：
   规则：componentName.propertyName.attributeName[.idx].attributeName
            componentName：组件名称，    即组件定义的Name名称，通过GetComponentItem(str_name)获取
            propertyName ：组件属性名称，组件中的属性变量名
            attributeName：嵌套属性：slice\array\struct\map\string\int\float等结构属性名称
            idx          ：嵌套属性：slice\array 数组结构的索引号
   样例：task_enquiry_test.PreCondition.0.ExpressionResult.Code
   注： Struct结构      =》 对应struct属性名称
        Map结构         =》 对应Map的key名称
        slice或array结构=》 对应数组元素的下标

    1）合约组件名前缀        contract_
    2）查询任务名前缀        task_enquiry_
    3）动作任务名前缀        task_action_
    4）决策任务名前缀        task_decision_
    5）计划任务名前缀        task_plan_
    6）决策候选条件名前缀    task_candidate_
    7）整型数据名前缀        data_intdata_
    8）无符号整型数据名前缀  data_uintdata_
    9）浮点型数据名前缀      data_uintdata_
    10）文本型数据名前缀     data_text_
    11）日期型数据名前缀     data_date_
    12）数组型数据名前缀     data_array_
    13）复数型数据名前缀     data_compound_
    14）矩阵型数据名前缀     data_matrix_
    15）操作结果型数据名前缀 data_operateresult_
    16）函数表达式名前缀     expression_function_
    17）逻辑表达式名前缀     expression_logicargument_

4. 函数表达式
   规则：Func[a-zA-Z_0-9]*（[variablename,]）

5. 条件判断表达式
   规则：
  1）常量式条件判断表达式
     true  false
  2）逻辑式条件判断表达式
     ! && ||
     > >= < <= == !=  ( )
  3）包含变量或函数的条件表达式