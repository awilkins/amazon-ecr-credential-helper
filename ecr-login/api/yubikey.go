package api

import (
	"os/exec"
	"strings"
)

type YubikeyTokenProvider struct {
	KeyName string
}

func (p YubikeyTokenProvider) Provide() (string, error) {

	cmd := exec.Command("ykman", "oath", "code", "-s", p.KeyName)

	if output, err := cmd.Output(); err != nil {
		return "", err
	} else {
		totpCode := string(output)
		totpCode = strings.TrimSpace(totpCode)
		return totpCode, nil
	}

}
