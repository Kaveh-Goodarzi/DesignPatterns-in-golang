package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

type iterator interface {
	hasNext() bool
	next()
	getCurrentItem() *User
}

type iterable interface {
	getIterator() iterator
}

type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *UserIterator) next() {
	if u.hasNext() {
		u.index++
	}
}

func (u *UserIterator) getCurrentItem() *User {
	if u.hasNext() {
		return u.users[u.index]
	}

	return nil
}

type userIterableCollection struct {
	users []*User
}

func (u *userIterableCollection) getIterator() iterator {
	return &UserIterator{
		index: 0,
		users: u.users,
	}
}

func main() {
	usr1 := &User{
		name: "Mark",
		age: 30,
	}

	usr2 := &User{
		name: "John",
		age: 25,
	}

	usr3 := &User{
		name: "gEnIuS",
		age: 13,
	}

	userIterableCollection := &userIterableCollection{
		users: []*User{usr1, usr2, usr3},
	}

	iterator := userIterableCollection.getIterator()

	for iterator.hasNext() {
		elm := iterator.getCurrentItem()
		fmt.Println("Current Item is: ", elm)
		iterator.next()
	}
}