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
	BrickPath string
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

func CreateVolume(masterAddr string, slaveAddr string, v *Volume) (string, error) {
	_, err := Gluster("volume", "create", v.Name, "replica","2", "transport","tcp", masterAddr + v.BrickPath, slaveAddr + v.BrickPath)
	if (err != nil) {
		return "", err
	}
	_,err = StartVolume(v.Name)
	if (err != nil) {
		return "", err
	}
	result, err := Gluster("volume", "set", v.Name, "nfs.disable","on")
	if (err != nil) {
		return "", err
	}
	return result, nil
}

func StartVolume(name string) (string, error) {
	result, err := Gluster("volume", "start",name,"force","--mode=script")
	if (err != nil) {
		return "", err
	}
	return result, nil
}

func DeleteVolume(name string) (bool, error) {
	_, err := StopVolume(name)
	if (err == nil) {
		_, err := Gluster("volume", "delete", name,"--mode=script")
		if (err != nil) {
			return false, err
		}
	}
	return true, nil
}

func StopVolume(name string) (bool, error) {
	_, err := Gluster("volume", "stop",name,"force","--mode=script")
	if (err != nil) {
		return false, err
	}
	return true, nil
}

func ListVolume() (string, error) {
	result, err := Gluster("volume", "info", "all")
	if (err != nil) {
		return result, err
	}
	return result, nil
}

func QueryVolume(name string) (string, error) {
	result, err := Gluster("volume", "info", name)
	if (err != nil) {
		return result, err
	}
	return result, nil
}