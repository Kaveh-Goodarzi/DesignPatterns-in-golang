package main

import (
	"fmt"
	"strings"
)

type WritingState interface {
	Write(Text string)
}

type UpperCase struct {
	WritingState
}

func (u *UpperCase) Write(Text string) {
	fmt.Println(strings.ToUpper(Text))
}

type LowerCase struct {
	WritingState
}

func (l *LowerCase) Write(Text string) {
	fmt.Println(strings.ToLower(Text))
}

type DefaultText struct {
	WritingState
}

func (d *DefaultText) Write(Text string) {
	fmt.Println(Text)
}

type TextEditor struct {
	State WritingState
}

func (t *TextEditor) DefaultText() {
	t.State = &DefaultText{}
}

func (t *TextEditor) SetState(state WritingState) {
	t.State = state
}

func (t *TextEditor) Type(Text string) {
	t.State.Write(Text)
}

func main() {
	editor := TextEditor{}
	editor.SetState(&DefaultText{})
	editor.Type("Default text")
	fmt.Println("=========")

	editor.SetState(&UpperCase{})
	editor.Type("Upper case text")
	fmt.Println("=========")

	editor.SetState(&LowerCase{})
	editor.Type("Lower case text")
}