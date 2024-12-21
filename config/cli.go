package config

const Version = "0.0.1"

var Templates map[string]string

func init() {
	// TODO: modify to myself repo
	Templates = map[string]string{
		"ts":   "https://github.com/GodlessLiu/starter-ts.git",
		"js":   "https://github.com/GoAdminGroup/themes/tree/master/js",
		"nuxt": "https://github.com/GoAdminGroup/themes/tree/master/nuxt",
	}
}

func GetTemplate(name string) string {
	return Templates[name]
}

func GetTemplateKeys() []string {
	keys := make([]string, 0)
	for k := range Templates {
		keys = append(keys, k)
	}
	return keys
}
