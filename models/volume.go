package models

import (
	xml2json "github.com/samuelhug/goxml2json"
	"os/exec"
	"strings"
	"log"
	"bytes"
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
		return "", err
	}
	xml := strings.NewReader(string(output))
	json, err := xml2json.Convert(xml)
	if err != nil {
		return "", err
	}
	return json.String(), err
}

func CreateVolume(v *Volume) (bool, error) {
	var dsl = bytes.Buffer{}
	dsl.WriteString("volume create " + v.Name)
	dsl.WriteString("replica 2 transport tcp ")
	dsl.WriteString("10.9.30.201:/data/" + v.BrickPath)
	dsl.WriteString("10.9.31.112:/data/"+ v.BrickPath)
	log.Println(dsl)
	_, err := Gluster(dsl.String())
	if (err != nil) {
		return false, err
	}
	return true,nil
}

func DeleteVolume(name string) (bool, error) {
	dsl := "volume delete " + name
	_, err := Gluster(dsl)
	if (err != nil) {
		return false, err
	}
	return true,nil
}

func ListVolume() (string, error) {
	var dsl = "volume info all "
	result, err := Gluster(dsl)
	if (err != nil) {
		return "", err
	}
	return result,nil
}

func QueryVolume(name string) (string, error) {
	dsl := "volume info all " + name
	result, err := Gluster(dsl)
	if (err != nil) {
		return "", err
	}
	return result,nil
}
