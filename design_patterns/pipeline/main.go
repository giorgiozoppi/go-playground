package main

type ComputeFunc func(int) int
type PipelineExecutor struct {
	Sink      chan int
	Stages    int
	Operators []ComputeFunc
}

func (*PipelineExecutor) AddStage(f ComputeFunc) {

}
func (*PipelineExecutor) Execute(values []int) int {
	return 0
}
func main() {
}
