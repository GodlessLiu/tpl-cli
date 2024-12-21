package internal

import (
	"fmt"

	"github.com/GodlessLiu/tpl-cli/config"
	"github.com/pterm/pterm"
)

type InitConfig struct {
	Dirname      string
	TemplateName string
}

func InitLogic() {
	var conf InitConfig
	dirName := pterm.DefaultInteractiveTextInput.WithDefaultText("Please enter the directory name: ").WithDefaultValue("my-app")
	conf.Dirname, _ = dirName.Show()

	options := config.GetTemplateKeys()
	conf.TemplateName, _ = pterm.DefaultInteractiveSelect.WithOptions(options).WithDefaultText("Please select a template: ").Show()
	fmt.Println(conf)
}
