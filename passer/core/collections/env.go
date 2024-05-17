package collections

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvCollection struct {
	SupabaseURL        string
	SupabaseAnonKey    string
	SupabaseServiceKey string
	Port               string
}

func NewEnv() *EnvCollection {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &EnvCollection{
		SupabaseURL:        os.Getenv("SUPABASE_URL"),
		SupabaseAnonKey:    os.Getenv("SUPABASE_ANON_KEY"),
		SupabaseServiceKey: os.Getenv("SUPABASE_SERVICE_KEY"),
		Port:               fmt.Sprintf(":%v", os.Getenv("PORT")),
	}
}

var Env *EnvCollection
