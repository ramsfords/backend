package utils

import (
	"fmt"
	"os"
	"strings"
)

func RestoreEnvValue(replaceData string, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	strData := string(data)
	strData = strings.ReplaceAll(strData, fmt.Sprintf("\nPUBLIC_THEME=%s", replaceData), "\nPUBLIC_THEME={THEME_NAME}")
	fmt.Println("current data is after change", string(data))
	err = os.WriteFile(path, []byte(strData), 0644)
	return err
}
