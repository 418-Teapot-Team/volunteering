package repository

import (
	"errors"
	"github.com/go-sql-driver/mysql"
)

const UniqueViolationCode = 1062

var UniqueViolationError = errors.New("the user with this email is already exists")

func handleError(err error) error {
	ms, ok := err.(*mysql.MySQLError)
	if !ok {
		return err
	}
	if ms.Number == UniqueViolationCode {
		return nil
	}
	return err
}
