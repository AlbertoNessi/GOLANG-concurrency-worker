package main

type MeasuredWorker struct {
	Worker
	value int
}

func (m *MeasuredWorker) Work() {
	// m.Worker.Work()
	// removed the method from the interface Work because it slows the executing x 5. 
	// Just increase the value by 1
	m.value++
}

func (m *MeasuredWorker) Value() int {
	return m.value
}
