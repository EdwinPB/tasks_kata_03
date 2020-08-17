package models

func (ms ModelSuite) Test_Task_Add_Task_Storage() {
	taskStorage := TaskStorage{}
	taskStorage.Add(Task{})
	ms.Equal(1, len(taskStorage))
	taskStorage.Add(Task{})
	taskStorage.Add(Task{})
	ms.Equal(3, len(taskStorage))

}
