package env

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetKVpairsEnv(envName string) map[string]string {
	str := Getenv(envName)
	const (
		pairsDelim = ","
		kvDelim = "="
	)
	if len(str) == 0 {
		return nil
	}
	pairs := strings.Split(str, pairsDelim)
	var result = make(map[string]string)
	for _, pair := range pairs {
		kvs := strings.Split(pair, kvDelim)
		if len(kvs) != 2 {
			logrus.Warnf("not correct amount of parts %+v for '%+v'", len(kvs), SecurePrint(pair))
			continue
		}
		k := kvs[0]
		v := kvs[1]
		result[k] = v
	}
	return result
}

func Getenv(name string) string {
	env := strings.TrimSpace(os.Getenv(name))
	logrus.Infof(`env %+v="%+v"`, name, SecurePrint(env))
	return env
}

func Header(r *http.Request, name string) string {
	h := strings.TrimSpace(r.Header.Get(name))
	logrus.Infof(`-H %+v="%+v"`, name, SecurePrint(h))
	return h
}

func GetenvDefault(envName, defaultVal string) string {
	v := Getenv(envName)
	if len(v) > 0 {
		return v
	}
	logrus.Infof("using default value for %+v='%+v'", envName, SecurePrint(defaultVal))
	return defaultVal
}

func GetenvBoolDefault(envName string, defaultVal bool) bool {
	v := GetenvDefault(envName, fmt.Sprintf("%+v", defaultVal))
	r, err := strconv.ParseBool(v)
	if err != nil {
		logrus.Errorf("couldn't parse '%+v' to bool for env '%+v'", v, envName)
		return defaultVal
	}
	return r
}

func GetenvIntDefault(envName string, defaultVal int) int {
	v := GetenvDefault(envName, fmt.Sprintf("%+v", defaultVal))
	r, err := strconv.Atoi(v)
	if err != nil {
		logrus.Errorf("couldn't parse '%+v' to bool for env '%+v'", v, envName)
		return defaultVal
	}
	return r
}