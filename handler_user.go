package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/carlosbueloni/gator-rss/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	userName := cmd.Args[0]

	if _, err := s.db.GetUser(context.Background(), userName); err != nil {
		os.Exit(1)
	}

	err := s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User %v has been set\n", userName)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	ctx := context.Background()
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	user, err := s.db.CreateUser(ctx, userParams)
	if err != nil {
		return err
	}
	s.cfg.SetUser(user.Name)
	fmt.Printf("User: %v with UUID: %v was created at %v, and last updated at: %v", user.Name, user.ID, user.CreatedAt, user.UpdatedAt)

	return nil

}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	if len(users) == 0 {
		fmt.Println("No users registered")
		return nil
	}
	currentUser := s.cfg.CurrentUserName
	for _, user := range users {
		if user.Name == currentUser {
			fmt.Printf("* %v (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %v \n", user.Name)
	}
	return nil
}
