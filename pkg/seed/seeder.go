package seed

import (
	"GDColumn/pkg/console"
	"GDColumn/pkg/database"
	"gorm.io/gorm"
)

var seeders []Seeder

var orderedSeederNames []string

type SeederFunc func(*gorm.DB)

type Seeder struct {
	Func SeederFunc
	Name string
}

func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		Name: name,
		Func: fn,
	})
}

func SetRunOrder(names []string) {
	orderedSeederNames = names
}

// GetSeeder 通过名称来获取 Seeder 对象
func GetSeeder(name string) Seeder {
	for _, sdr := range seeders {
		if name == sdr.Name {
			return sdr
		}
	}
	return Seeder{}
}

func RunAll() {

	executed := make(map[string]string)
	for _, name := range orderedSeederNames {
		sdr := GetSeeder(name)
		if len(sdr.Name) > 0 {
			console.Warning("Running Odered Seeder: " + sdr.Name)
			sdr.Func(database.DB)
			executed[name] = name

		}
	}

	for _, sdr := range seeders {
		if _, ok := executed[sdr.Name]; !ok {
			console.Warning("Running Seeder: " + sdr.Name)
			sdr.Func(database.DB)
		}
	}
}

func RunSeeder(name string) {
	for _, sdr := range seeders {
		if name == sdr.Name {
			sdr.Func(database.DB)
			break
		}
	}
}