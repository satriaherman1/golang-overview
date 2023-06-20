package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/rs/cors"
)

type CalculationResult struct {
	TotalStock   int     `json:"totalStock"`
	TotalPrice   float64 `json:"totalPrice"`
	CurrentPrice float64 `json:"currentPrice"`
}

var testType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Test",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"no": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

func main() {
	// Definisikan skema GraphQL
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// Definisikan field query yang diperlukan
			"test": &graphql.Field{
				Type:        testType,
				Description: "A dummy query",
				Resolve:     resolveTest,
			},
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"calculatePrice": &graphql.Field{
				Type:        CalculationResultType,
				Description: "Calculate chicken stock",
				Args: graphql.FieldConfigArgument{
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.Float)),
					},
				},
				Resolve: calculatePriceResolver,
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Buat handler GraphQL
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Mengonfigurasi CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // CORS Origin
		AllowedMethods:   []string{"POST", "GET", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	})

	// Menggunakan CORS untuk handler GraphQL
	handlerWithCors := c.Handler(h)

	http.Handle("/graphql", handlerWithCors)

	// Jalankan server di port 8080
	log.Println("Running at  http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var CalculationResultType = graphql.NewObject(graphql.ObjectConfig{
	Name: "CalculationResult",
	Fields: graphql.Fields{
		"totalStock": &graphql.Field{
			Type: graphql.Int,
		},
		"totalPrice": &graphql.Field{
			Type: graphql.Float,
		},
		"currentPrice": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

func calculatePriceResolver(p graphql.ResolveParams) (interface{}, error) {
	prices, _ := p.Args["price"].([]interface{})

	totalStock := len(prices)
	totalPrice := 0.0

	for _, price := range prices {
		if p, ok := price.(float64); ok {
			totalPrice += p
		}
	}

	currentPrice := totalPrice / float64(totalStock)

	result := CalculationResult{
		TotalStock:   totalStock,
		TotalPrice:   totalPrice,
		CurrentPrice: currentPrice,
	}

	return result, nil
}

func resolveTest(params graphql.ResolveParams) (interface{}, error) {
	// Logika untuk mengambil data "id" dan "no"
	// Misalnya, mengambil data dari database atau sumber data lainnya
	id := "12345"
	no := 10

	// Mengembalikan objek yang berisi data "id" dan "no"
	return map[string]interface{}{
		"id": id,
		"no": no,
	}, nil
}
