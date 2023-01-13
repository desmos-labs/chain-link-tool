package nomic

import (
	"os"
	"path"

	"github.com/manifoldco/promptui"
)

type NomicGetter interface {
	GetPrivateKeyPath() (string, error)
}

type nomicGetter struct{}

func (g *nomicGetter) GetPrivateKeyPath() (string, error) {
	return g.getPrivateKeyPath()
}

func (g *nomicGetter) getPrivateKeyPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	prompt := promptui.Prompt{
		Label:   "Please insert the path of private key (fully qualified path)",
		Default: path.Join(wd, "privkey"),
	}
	return prompt.Run()
}
