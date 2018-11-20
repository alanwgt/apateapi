package cache

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/alanwgt/apateapi/database"

	"github.com/alanwgt/apateapi/models"
	"github.com/patrickmn/go-cache"
)

// UserCache stores an user model with the public key in raw data
type UserCache struct {
	Model *models.User
	PubK  *[32]byte
}

var c *cache.Cache

func init() {
	c = cache.New(10*time.Minute, 20*time.Minute)
}

// GetUser will try to get the user from the cache, if it's not in
// the cache, it'll retrieve from DB and store in cache
func GetUser(un string) (*UserCache, error) {
	if u, ok := c.Get(un); ok {
		return (u.(*UserCache)), nil
	}

	u := &models.User{}
	c := database.GetOpenConnection()

	if c.First(u, &models.User{Username: un}).RecordNotFound() {
		return nil, errors.New("User '" + un + "' not found")
	}

	// store in cache
	uc := SetUser(u)

	return uc, nil
}

// SetUser stores the user in the cache with the default expiration
func SetUser(u *models.User) (uc *UserCache) {
	var bPubK [32]byte
	pubK, _ := base64.StdEncoding.DecodeString(u.PubKey)
	copy(bPubK[:], pubK)

	uc = &UserCache{
		u,
		&bPubK,
	}

	c.Set(u.Username, uc, cache.DefaultExpiration)
	return
}

// RemoveUser deletes an user entry from cache
func RemoveUser(u *models.User) {
	c.Delete(u.Username)
}
