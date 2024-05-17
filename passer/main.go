package main

import (
	"fmt"
	"log"

	"github.com/KRTirtho/mess-backend/passer/core/collections"
	"github.com/supabase-community/supabase-go"
)

// Supabase client with service role key
var SupabaseServerClient *supabase.Client
var Env collections.Env

func main() {
	log.SetFlags(0)

	Env = *collections.NewEnv()

	client, err := supabase.NewClient(Env.SupabaseURL, Env.SupabaseServiceKey, nil)
	if err != nil {
		fmt.Println("cannot initialize client", err)
	}

	SupabaseServerClient = client
}
