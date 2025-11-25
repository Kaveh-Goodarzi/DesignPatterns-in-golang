package main

import "fmt"

type Character interface {
	Render(name string)
}

type CharacterFactory struct {
	characters map[string]Character
}

func newCharacterFactory() *CharacterFactory {
	return &CharacterFactory{characters: make(map[string]Character)}
}

func (cf *CharacterFactory) GetSharedCharacters(height, weight, hairColor string) Character {
	key := height + weight + hairColor

	if _, exists := cf.characters[key]; !exists {
		fmt.Println("Creating new character with key:", key)

		cf.characters[key] = &SharedCharacter{
			height: height, weight: weight, hairColor: hairColor,
		}
	}
	return cf.characters[key]
}

type SharedCharacter struct {
	height    string
	weight    string
	hairColor string
}

func (sc *SharedCharacter) Render(name string) {
	fmt.Printf("Rendering character %s with height %s, weight %s, hair color %s\n", name, sc.height, sc.weight, sc.hairColor)
}

type CharacterClient struct {
	name      string
	character Character
}

func NewCharacterClient(name, height, weight, hairColor string, factory *CharacterFactory) *CharacterClient {
	character := factory.GetSharedCharacters(height, weight, hairColor)
	return &CharacterClient{name: name, character: character}
}

func (cc *CharacterClient) Render() {
	cc.character.Render(cc.name)
}

func main() {
	factory := newCharacterFactory()

	character1 := NewCharacterClient("Hero1", "6ft", "180lbs", "Black", factory)
	character2 := NewCharacterClient("Hero2", "6ft", "180lbs", "Black", factory)
	character3 := NewCharacterClient("Hero3", "5ft", "150lbs", "Blonde", factory)

	character1.Render()
	character2.Render()
	character3.Render()
}