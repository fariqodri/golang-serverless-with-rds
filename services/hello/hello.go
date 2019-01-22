package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
	"github.com/jmoiron/sqlx"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

var (
	DB_NAME = os.Getenv("DB_NAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_PORT = os.Getenv("DB_PORT")
	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	REGION = os.Getenv("REGION")
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {

	var buf bytes.Buffer
	sess := session.Must(session.NewSession(&aws.Config{Region:&REGION}))
	awsCreds := stscreds.NewCredentials(sess, "arn:aws:iam::585040772542:user/serverless")
	authToken, err := rdsutils.BuildAuthToken(DB_HOST, REGION, DB_USER, awsCreds)
	dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true",
		DB_USER, authToken, DB_HOST, DB_NAME)

	var db *sqlx.DB
	db = sqlx.MustConnect("mysql", dnsStr)
	_ = db.MustExec("CREATE TABLE IF NOT EXISTS USER (user_id VARCHAR(50));")
	body, err := json.Marshal(map[string]interface{}{
		"message": "RDS Connected",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
