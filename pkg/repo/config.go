package repo

type Config struct {
	RepoDSN string `env:"REPO_DSN,required"`
}
