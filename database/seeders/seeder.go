package seeders

import "GDColumn/pkg/seed"

func Initialize() {

	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}
