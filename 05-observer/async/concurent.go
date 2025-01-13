package main

import (
	"fmt"
	"sync"
)

// Follower is an observer interface
type Follower interface {
	Update(post string) // React to a new post
}

// Influencer is a subject interface
type Influencer interface {
	AddFollower(f Follower)
	RemoveFollower(f Follower)
	NotifyFollowers(post string)
}

// Celebrity is a concrete influencer (Subject)
type Celebrity struct {
	name      string
	followers []Follower
	mu        sync.Mutex // Protect the follower list
}

func (c *Celebrity) AddFollower(f Follower) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.followers = append(c.followers, f)
	fmt.Printf("%s started following %s!\n", f, c.name)
}

func (c *Celebrity) RemoveFollower(f Follower) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i := range c.followers {
		if c.followers[i] == f {
			c.followers = append(c.followers[:i], c.followers[i+1:]...)
			fmt.Printf("%s unfollowed %s.\n", f, c.name)
			return
		}
	}
}

func (c *Celebrity) NotifyFollowers(post string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Printf("%s posts: %s\n", c.name, post)

	var wg sync.WaitGroup
	for _, follower := range c.followers {
		wg.Add(1)
		// Notify each follower in a goroutine
		go func(f Follower) {
			defer wg.Done()
			f.Update(post)
		}(follower)
	}
	wg.Wait()
}

// User is a concrete follower (Observer)
type User struct {
	username string
}

func (u *User) Update(post string) {
	fmt.Printf("%s received notification: New post -> %s\n", u.username, post)
}

func (u *User) String() string {
	return u.username
}

func main() {
	// Create an influencer
	influencer := &Celebrity{name: "TechGuru"}

	// Add followers
	user1 := &User{username: "Alice"}
	user2 := &User{username: "Bob"}
	user3 := &User{username: "Charlie"}

	influencer.AddFollower(user1)
	influencer.AddFollower(user2)
	influencer.AddFollower(user3)

	// Influencer posts something
	influencer.NotifyFollowers("Check out my latest tech review!")

	// Remove a follower and post again
	influencer.RemoveFollower(user2)
	influencer.NotifyFollowers("Unboxing the newest gadget!")
}
