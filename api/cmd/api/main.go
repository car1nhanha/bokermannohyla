package main

import (
	"context"
	"log"
	"os"

	"thoropa/internal/database"
	"thoropa/internal/handler"
	"thoropa/internal/repository"
	"thoropa/internal/router"
	"thoropa/internal/service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/joho/godotenv"
)

var ginLambda *ginadapter.GinLambdaV2

func init() {
	godotenv.Load()

	ctx := context.Background()

	db := database.NewDynamoClient(ctx)

	// camadas
	repo := repository.NewLinkRepository(db)
	service := service.NewLinkService(repo)
	handler := handler.NewLinkHandler(service)

	r := router.SetupRouter(handler)

	// Adapter do Gin pra Lambda
	ginLambda = ginadapter.NewV2(r)
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	if os.Getenv("IS_RUNNING_LOCAL") == "true" {
		log.Println("🚀 Rodando localmente em http://localhost:8080")

		ctx := context.Background()

		db := database.NewDynamoClient(ctx)
		repo := repository.NewLinkRepository(db)
		service := service.NewLinkService(repo)
		handler := handler.NewLinkHandler(service)

		r := router.SetupRouter(handler)

		if err := r.Run(":8080"); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("✅ Lambda inicializada")
		lambda.Start(Handler)
	}
}
