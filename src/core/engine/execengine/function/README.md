#############################################
##########    函数解析规则：   ##############
#############################################
函数定义规则：
1. 函数定义样式
   ex：func FuncXXXXXXX(args...interface{})(common.OperateResult, error)
       args...interface{}          ：传递可变参数
       (common.OperateResult,error)：返回通用结果类型
2. 函数、方法命名规则
   ex: FuncXXXXX
       函数名称需要以Func开头
3. 函数定义文件命名规则
   ex：ContractFunctionForXXXXX
       XXXX为函数分类标记，如Common是合约通用方法，TIANJS是天安金交所特有的扩展方法
4. 扩展函数分类标识定义：在engine/execengine/constdef/FunctionSource.go中定义
	DEMO_SRC          ：demo 的合约执行方法
	TIANJS_SRC        ：天安金交所的合约执行方法
	GUANGXIBIANMAO_SRC：广西边贸的合约执行方法
5. 函数定义配置文件： conf/UCVM.yaml
   逗号分隔定义加载的合约方法标识串
   function_source: DEMO_SRC,GUANGXIBIANMAO_SRC

