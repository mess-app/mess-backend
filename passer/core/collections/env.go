package collections

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	SupabaseURL        string
	SupabaseAnonKey    string
	SupabaseServiceKey string
}

func NewEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Env{
		SupabaseURL:        os.Getenv("SUPABASE_URL"),
		SupabaseAnonKey:    os.Getenv("SUPABASE_ANON_KEY"),
		SupabaseServiceKey: os.Getenv("SUPABASE_SERVICE_KEY"),
	}
}
