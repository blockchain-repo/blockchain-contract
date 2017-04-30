// threadpool
package pipelines

type ThreadPool struct {
	funcNumber     int
	funcQueue      chan func() error
	funcResult     chan error
	finishCallback func()
}

// 初始化
func (pool *ThreadPool) Init(number int) {
	pool.funcQueue = make(chan func() error, number)
	pool.funcNumber = number
	pool.funcResult = make(chan error, number)
}

// 添加任务
func (pool *ThreadPool) AddTask(task func() error) {
	pool.funcQueue <- task
}

// 设置结束回调
func (pool *ThreadPool) SetFinishCallback(callback func()) {
	pool.finishCallback = callback
}

// 关闭
func (pool *ThreadPool) Stop() {
	close(pool.funcQueue)
	close(pool.funcResult)
}

// 开始
func (pool *ThreadPool) Start() []error {
	// 开启Number个goroutine
	for i := 0; i < pool.funcNumber; i++ {
		go func() {
			for {
				task, ok := <-pool.funcQueue
				if !ok {
					break
				}

				err := task()
				pool.funcResult <- err
			}
		}()
	}

	var slErr []error

	// 获得每个work的执行结果
	for j := 0; j < pool.funcNumber; j++ {
		res, ok := <-pool.funcResult
		if !ok {
			break
		}

		if res != nil {
			slErr = append(slErr, res)
		}
	}

	// 所有任务都执行完成，回调函数
	if pool.finishCallback != nil {
		pool.finishCallback()
	}

	return slErr
}
