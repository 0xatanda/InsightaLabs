package seed

import (
	"encoding/json"
	"os"

	"github.com/0xatanda/InsightaLabs/internal/model"
	"github.com/0xatanda/InsightaLabs/internal/repository"
)

func Run(repo *repository.Repo) error {

	file, err := os.Open("seed.json")
	if err != nil {
		return err
	}
	defer file.Close()

	var profiles []model.Profile

	if err := json.NewDecoder(file).Decode(&profiles); err != nil {
		return err
	}

	for _, p := range profiles {
		_ = repo.CreateOrGetProfile(p.Name)
	}

	return nil
}
