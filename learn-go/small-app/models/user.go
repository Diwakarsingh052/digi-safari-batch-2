package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CacheStore map[string]User

type Conn struct {
	//db *sql.DB
	cache CacheStore
}

func NewConn() Conn {
	return Conn{cache: make(map[string]User, 100)}
}

func (c *Conn) CreateUser(n NewUser) (User, error) {
	// generate a passwordHash using bcrypt
	passHash, err := bcrypt.GenerateFromPassword([]byte(n.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	// assigning the newUser values to the User type
	us := User{
		Id:           uuid.NewString(),
		Email:        n.Email,
		Name:         n.Name,
		Age:          n.Age,
		PasswordHash: string(passHash),
	}
	// put the new user in the map
	c.cache = CacheStore{n.Email: us}

	return us, nil
}
