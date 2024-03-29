// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"time"
)

type Breed struct {
	BreedID   int64     `json:"breed_id"`
	BreedName string    `json:"breed_name"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type Hatchery struct {
	ID           int64     `json:"id"`
	ProductionID int64     `json:"production_id"`
	Infertile    int64     `json:"infertile"`
	Early        int64     `json:"early"`
	Middle       int64     `json:"middle"`
	Late         int64     `json:"late"`
	DeadChicks   int64     `json:"dead_chicks"`
	AliveChicks  int64     `json:"alive_chicks"`
	CreatedAt    time.Time `json:"created_at"`
}

type Premise struct {
	ID        int64     `json:"id"`
	BreedID   int64     `json:"breed_id"`
	Farm      string    `json:"farm"`
	House     string    `json:"house"`
	CreatedAt time.Time `json:"created_at"`
}

type Production struct {
	ID           int64     `json:"id"`
	BreedID      int64     `json:"breed_id"`
	Eggs         int64     `json:"eggs"`
	Dirty        int64     `json:"dirty"`
	WrongShape   int64     `json:"wrong_shape"`
	WeakShell    int64     `json:"weak_shell"`
	Damaged      int64     `json:"damaged"`
	HatchingEggs int64     `json:"hatching_eggs"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}
