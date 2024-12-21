package internal

import (
	"os"
	"path"

	"github.com/GodlessLiu/tpl-cli/config"
	"github.com/go-git/go-git/v5"
	"github.com/pterm/pterm"
)

type InitConfig struct {
	Dirname      string
	TemplateName string
}

func getDirPath(dirname string) string {
	pa, _ := os.Getwd()
	return path.Join(pa, dirname)
}

func InitLogic() {
	var conf InitConfig
	dirName := pterm.DefaultInteractiveTextInput.WithDefaultText("Please enter the directory name: ").WithDefaultValue("my-app")
	conf.Dirname, _ = dirName.Show()

	options := config.GetTemplateKeys()
	conf.TemplateName, _ = pterm.DefaultInteractiveSelect.WithOptions(options).WithDefaultText("Please select a template: ").Show()

	if _, err := git.PlainClone(getDirPath(conf.Dirname), false, &git.CloneOptions{
		URL:      config.GetTemplate(conf.TemplateName),
		Progress: os.Stdout,
	}); err != nil {
		pterm.Error.Println(err.Error())
		panic(err.Error())
	}
	pterm.Info.Printfln("download template successfully, please check the directory %s", conf.Dirname)
}
