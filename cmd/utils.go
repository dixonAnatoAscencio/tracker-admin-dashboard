package main

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port   string
	DBPath string
}

func loadConfig() Config {
	return Config{
		Port:   getEnv("PORT", "8080"),
		DBPath: getEnv("DATABASE_URL", "./data/orders.db"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func loadTemplates(router *gin.Engine) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	functions := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}

	tmpl, err := template.New("base.tmpl").
		Funcs(functions).
		ParseFiles(
			filepath.Join(wd, "templates", "base.tmpl"),
			filepath.Join(wd, "templates", "order.tmpl"),
		)
	if err != nil {
		return err
	}

	router.SetHTMLTemplate(tmpl)
	return nil
}
