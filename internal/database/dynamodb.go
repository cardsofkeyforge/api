package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/kadekutama/dynamodb"
	"github.com/pkg/errors"
	"os"
)

type Filter struct {
	Expression       string
	ExpressionValues []interface{}
}

type QueryRequest struct {
	TableName      string
	Filter         *Filter
	PartitionKey   string
	PartitionValue interface{}
	RangeKey       string
	RangeOperator  dynamo.Operator
	RangeValue     interface{}
}

type ScanRequest struct {
	TableName string
	Filter    *Filter
}

func New() (dynamodb.DB, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create dynamodb session")
	}
	db := dynamodb.New(sess)
	return db, nil
}

func Scan(sr *ScanRequest, result interface{}) error {
	if db, err := New(); err != nil {
		return err
	} else {
		table := db.Table(sr.TableName)
		switch v := result.(type) {
		default:
			scan := table.Scan()
			if sr.Filter != nil {
				scan = scan.Filter(sr.Filter.Expression, sr.Filter.ExpressionValues...)
			}
			err = scan.All(v)
			return err
		}
	}
}

func Query(qr *QueryRequest, result interface{}) error {
	if db, err := New(); err != nil {
		return err
	} else {
		table := db.Table(qr.TableName)
		switch v := result.(type) {
		default:
			query := table.Get(qr.PartitionKey, qr.PartitionValue)
			if qr.RangeKey != "" && qr.RangeOperator != "" && qr.RangeValue != "" {
				query = query.Range(qr.RangeKey, qr.RangeOperator, qr.RangeValue)
			}
			if qr.Filter != nil {
				query = query.Filter(qr.Filter.Expression, qr.Filter.ExpressionValues...)
			}
			err = query.All(v)
			return err
		}
	}
}
