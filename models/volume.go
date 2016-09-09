package models

import (
	xml2json "github.com/samuelhug/goxml2json"
	"os/exec"
	"strings"
	"log"
)

type Volume struct {
	Name      string
	Status    int
	Quota     Quota
	BrickPath string
}

type Quota struct {
	Path      string
	Hardlimit string
	Softlimit string
	Used      string
	Available string
}

func Gluster(vars ...string) (string, error) {
	if len(vars) < 1 {
		return "incorrect url", nil
	}
	args := append(vars, "--xml")
	gCmd := exec.Command("gluster", args...)
	log.Println(vars, args, gCmd.Path, gCmd.Args)
	output, err := gCmd.CombinedOutput()
	if err != nil {
		log.Fatalln(err.Error())
		return "", err
	}
	xml := strings.NewReader(string(output))
	json, err := xml2json.Convert(xml)
	if err != nil {
		log.Fatalln(err.Error())
		return "", err
	}
	return json.String(), err
}

func CreateVolume(v *Volume) (bool, error) {
	_, err := Gluster("volume","create ",v.Name,"replica 2","transport tcp","10.9.30.201:/data/" + v.BrickPath,"10.9.31.112:/data/"+ v.BrickPath)
	if (err != nil) {
		return false, err
	}
	return true,nil
}

func DeleteVolume(name string) (bool, error) {
	_, err := Gluster("volume","delete",name)
	if (err != nil) {
		return false, err
	}
	return true,nil
}

func ListVolume() (string, error) {
	result, err := Gluster("volume","info","all")
	if (err != nil) {
		return "", err
	}
	return result,nil
}

func QueryVolume(name string) (string, error) {
	result, err := Gluster("volume","info",name)
	if (err != nil) {
		return "", err
	}
	return result,nil
}
