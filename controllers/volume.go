package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"gluster-api/models"
	"bytes"
)

var masterAddr string;
var slaveAddr string;

func initGlusterParam(){
	if(masterAddr == nil){
		masterAddr = beego.AppConfig.String("MasterAddr")
	}
	if(slaveAddr == nil){
		slaveAddr = beego.AppConfig.String("SlaveAddr")
	}
}
type VolumeController struct {
	beego.Controller
}
// @Title Delete
// @Description get volume by name
// @Param name
// @Success 200
// @Failure 400 :name is empty
// @router / [post]
func (c *VolumeController) Create() {
	var ob *models.Volume
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	initGlusterParam()
	result, err := models.CreateVolume(masterAddr,slaveAddr,ob)
	if(err == nil){
		c.Data["json"] = map[string]bool{"status": result}
	}
	if(err != nil){
		c.Data["json"] = map[string]string{"error": err.Error()}
	}
	c.ServeJSON()
}
// @Title Delete
// @Description get volume by name
// @Param name
// @Success 200
// @Failure 400 :name is empty
// @router /:name [delete]
func (c *VolumeController) Delete() {
	name := c.GetString(":name")
	result, err := models.DeleteVolume(name)
	if(err == nil){
		c.Data["json"] = map[string]bool{"status": result}
	}
	if(err != nil){
		c.Data["json"] = map[string]string{"error": err.Error()}
	}
	c.ServeJSON()
}

// @Title Get
// @Description get volume by name
// @Param name
// @Success 200
// @Failure 400 :name is empty
// @router /:name [get]
func (c *VolumeController) Query() {
	name := c.GetString(":name")
	result, err := models.QueryVolume(name)
	if(err == nil){
		c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		c.Ctx.Output.Body(bytes.NewBufferString(result).Bytes())
	}
	if(err != nil){
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
	}
}

// @Title Get
// @Description get volume by name
// @Param name
// @Success 200
// @router / [get]
func (c *VolumeController) List() {
	result, err := models.ListVolume()
	if(err == nil){
		c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		c.Ctx.Output.Body(bytes.NewBufferString(result).Bytes())
	}
	if(err != nil){
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
	}
}