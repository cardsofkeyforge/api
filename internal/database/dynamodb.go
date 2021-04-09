package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/kadekutama/dynamodb"
	"github.com/pkg/errors"
	"os"
)

func New() (dynamodb.DB, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create dynamodb session")
	}
	db := dynamodb.New(sess)
	return db, nil
}

func Scan(tableName string, expression *string, values *[]interface{}, result interface{}) error {
	if db, err := New(); err != nil {
		return err
	} else {
		table := db.Table(tableName)
		switch v := result.(type) {
		default:
			err := table.Scan().Filter(*expression, *values...).All(v)
			return err
		}
	}
}
