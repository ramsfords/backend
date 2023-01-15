package api

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/utils"
)

func BuildStatic(context echo.Context) error {
	fmt.Println("BuildStatic is called")
	home, _ := os.UserHomeDir()
	err := utils.ReplaceEnvValue("RAINBOW", filepath.Join(home, "projects/menuloom-print/static/.env"))
	if err != nil {
		fmt.Println(err.Error())
	}
	cmd := exec.Command("npm", "run", "build")
	cmd.Dir = "/Users/surendrakandel/projects/menuloom-print/static"
	cmd.Run()
	cmd.Wait()
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(stdout))
	// restore the original .env file
	err = utils.RestoreEnvValue("HORIZON", filepath.Join(home, "projects/menuloom-print/static/.env"))
	if err != nil {
		fmt.Println(err.Error())
	}
	return context.JSON(200, "BuildStatic is called")
}
