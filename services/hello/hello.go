package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/fariqodri/itfest/pkg"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse


// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {

	var buf bytes.Buffer
	//sess := session.Must(session.NewSession(&aws.Config{Region:&REGION}))
	//awsCreds := stscreds.NewCredentials(sess, "arn:aws:iam::585040772542:user/serverless")
	//authToken, err := rdsutils.BuildAuthToken(DB_HOST, REGION, DB_USER, awsCreds)

	db := database.GetDatabase()
	defer db.Close()

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
