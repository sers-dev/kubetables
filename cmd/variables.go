package main

import (
	"flag"
	"os"
)

type Variables struct {
	KubernetesConfigPath string
	InterfaceGroups map[string][]string
}

//var jsonStr = `
//{
//	"public": [
//		"enp0s3",
//		"enp0s8"
//	],
//	"internal": [
//		"wg0"
//	]
//}`

var variablesInstance *Variables
//m := map[string]map[string]string{}
//m["var1"] = map[string]string{}
//m["var1"]["var2"] = "something"
//fmt.Println(m["var1"]["var2"])

func (v *Variables) GetVariables() *Variables {
	if variablesInstance == nil {
		variablesInstance.KubernetesConfigPath = v.getStringFromEnvByKey("KUBERNETES_CONFIG_PATH")
		kubernetesConfigPathByArg := v.getStringFromArgs("k", "kubernetesconfigpath")
		if kubernetesConfigPathByArg != "" {
			variablesInstance.KubernetesConfigPath = kubernetesConfigPathByArg
		}

		//variablesInstance.InterfaceGroups = v.getMapFromEnvByPrefix("IF_GROUP_")
		//interfaceGroupsString := v.getMapFromArgs("i", "ifgroups")
	}
	return variablesInstance
}

func (v *Variables) getStringFromEnvByKey(envKey string) string {
	return os.Getenv(envKey)
}

func (v *Variables) getStringFromArgs(short string, long string) string {
	var value string

	flag.StringVar(&value, long, "", long)
	flag.StringVar(&value, short, "", long+" (shorthand)")
	flag.Parse()

	return value
}
/*
func (v *Variables) getMapFromEnvByPrefix(prefix string) map[string][]string {
	twodmap := map[string][]string{}
	envvars := os.Environ()
	for _, envvar := range envvars {
		if strings.HasPrefix(envvar, prefix) {
			envvarsplit := strings.Split(envvar, "=")
			values := strings.Split(envvarsplit[1], ",")
			twodmap[envvarsplit[0]] = values
		}
	}
	return twodmap
}

// multiple flags
func (v *Variables) getMapFromArgs(short string, long string) map[string][]string {
	var value string
	flag.StringVar(&value, long, "", long)
	flag.StringVar(&value, short, "", long+" (shorthand)")
	flag.Parse()
}
*/