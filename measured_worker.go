package main

type MeasuredWorker struct {
	Worker
	value int
}

func (m *MeasuredWorker) Work() {
	m.value++
}

func (m *MeasuredWorker) Value() int {
	return m.value
}
