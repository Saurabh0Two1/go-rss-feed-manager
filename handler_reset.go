package main

import (
	"context"
	"fmt"
)

func ResetHandler(s *State, cmd Command) error {

	ctx := context.Background()
	err := s.db.ResetDb(ctx)

	if err != nil {
		errorStr := fmt.Errorf("failed to reset DB due to: %w", err)
		fmt.Println(errorStr)
		return errorStr
	}

	fmt.Printf("DB reset successful.")

	return nil

}
