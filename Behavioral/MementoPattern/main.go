package main

import "fmt"

type Editor struct {
	content string
}

func (e *Editor) AddContent(content string) {

}

func (e *Editor) Type(words string) {
	e.content += words
}

func (e *Editor) getContent() string {
	return e.content
}

func (e *Editor) save() *EditorMemento {
	return &EditorMemento{content: e.content}
}

type EditorMemento struct {
	content string
}

func (m *EditorMemento) getContent() string {
	return m.content
}

func (e *Editor) restore(m *EditorMemento) {
	e.content = m.getContent()
}

type EditorHistory struct {
	history []EditorMemento
}

func (h *EditorHistory) addMemento(m EditorMemento) {
	h.history = append(h.history, m)
}

func (h *EditorHistory) getMemento(index int) EditorMemento {
	return h.history[index]
}

func main() {
	e := Editor{}
	e.Type("This is the first sentence.")
	e.Type("This is the second sentence.")

	h := EditorHistory{}
	h.addMemento(*e.save())

	e.Type("This is third")

	e.restore(&h.history[0])
	fmt.Println(e.getContent())
}