package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReplaceEnvValue(replaceData string, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	strData := string(data)
	strData = strings.ReplaceAll(strData, "\nPUBLIC_THEME={THEME_NAME}", fmt.Sprintf("\nPUBLIC_THEME=%s", replaceData))
	err = os.WriteFile(path, []byte(strData), 0644)
	return err
}
