package models

type User struct {
	id   int
	name string
	salt string
}

type Token string
