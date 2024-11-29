package mysql

import (
	"app/stores/models"
	"fmt"
)

type Conn struct {
	userDb map[int]*models.User
}

func NewConn() Conn {
	NewCon := Conn{userDb: map[int]*models.User{}}
	return NewCon
}

// Creating a map to act as a data store

func (c Conn) Create(u models.User) (*models.User, bool) {

	fmt.Println("Creating a user in ", " u : ", &u)
	//Need to check if user exists if yes then throw error else save
	c.userDb[u.Id] = &u

	return &u, true

}

// func (c Conn) CreateSimple(u string) error {
// 	if u != "" {
// 		fmt.Println("Creating a user in ", c.db, " u : ", u)
// 		return nil
// 	}
// 	return nil
// }

func (c Conn) Update(id int, name string) (*models.User, bool) {
	u, ok := c.userDb[id]

	if !ok {
		fmt.Println("User with id ", id, "Is not found for update")
		return nil, false
	}
	fmt.Println("Updaating a user in ", " u : ", u)

	u.Name = name
	return u, true
}

func (c Conn) Delete(id int) (*models.User, bool) {
	u, ok := c.userDb[id]

	if !ok {
		fmt.Println("User with id ", id, "Is not found for delete")
		return nil, false
	}
	fmt.Println("deleting a user in map", " u : ", u)

	delete(c.userDb, id)
	return nil, true
}

func (c Conn) FetchAll() (map[int]*models.User, bool) {
	if c.userDb == nil {
		fmt.Println("No value in db")
		return nil, false
	}
	return c.userDb, true
}

func (c Conn) FetchUser(id int) (*models.User, bool) {
	u, ok := c.userDb[id]

	if !ok {
		fmt.Println("User with id ", id, "Is not found")
		return nil, false
	}
	fmt.Println("fetching user")
	return u, true
}
