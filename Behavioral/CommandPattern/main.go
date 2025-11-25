package main

import "fmt"

type Command interface {
	Execute()
	Undo()
}

type AddCommand struct {
	Receiver *Receiver
	Value    int
}

func (c *AddCommand) Execute() {
	c.Receiver.Add(c.Value)
}

func (c *AddCommand) Undo() {
	c.Receiver.Subtract(c.Value)
}

func ReceiveCommand(receiver Receiver, value int) {

}

type Receiver struct {
	value int
}

func (r *Receiver) Add(value int) {
	r.value += value
}

func (r *Receiver) Subtract(value int) {
	r.value -= value
}

func (r *Receiver) GetValue() int {
	return r.value
}

type SubtractCommand struct {
	Receiver *Receiver
	Value    int
}

func (c *SubtractCommand) Execute() {
	c.Receiver.Subtract(c.Value)
}

func (c *SubtractCommand) Undo() {
	c.Receiver.Add(c.Value)
}

type Invoker struct {
	commands []Command
	undoStack []Command
}

func (i *Invoker) addCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) executeCommands() {
	for _, command := range i.commands {
		command.Execute()
		i.undoStack = append(i.undoStack, command)
	}
	i.commands = []Command{}
}

func (i *Invoker) undo() {
	if len(i.undoStack) != 0 {
		lastCommand := i.undoStack[len(i.undoStack)-1]
		lastCommand.Undo()
		i.undoStack = i.undoStack[:len(i.undoStack)-1]
	}
}

func main() {
	Receiver := &Receiver{}
	invoker := &Invoker{}

	invoker.addCommand(&AddCommand{Receiver: Receiver, Value: 5})
	invoker.addCommand(&SubtractCommand{Receiver: Receiver, Value: 3})

	invoker.executeCommands()

	fmt.Println("Current Value:", Receiver.GetValue())

	invoker.undo()

	fmt.Println("Current Value after undo:", Receiver.GetValue())
}