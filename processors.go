package main

var currentId int

var todos QosClientA

// Give us some seed data
func init() {
	RepoCreateTodo(QosClientT{Name: "Write presentation"})
	RepoCreateTodo(QosClientT{Name: "Host meetup"})
}

//this is bad, I don't think it passes race condtions
func RepoCreateTodo(t QosClientT) QosClientT {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}
