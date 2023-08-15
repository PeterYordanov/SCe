package core

import "fmt"

func HandleError(err error, msg string) error {
	if err != nil {
		return fmt.Errorf("%s", msg)
	}
	return nil
}

func ChainError(err error, msg string) error {
	if err != nil {
		return fmt.Errorf("%s: %w", msg, err)
	}
	return nil
}

func PropagateError(err error) error {
	if err != nil {
		return err
	}
	return nil
}
