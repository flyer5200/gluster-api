package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"gluster-api/models"
)

type VolumeController struct {
	beego.Controller
}

// @router / [post]
func (c *VolumeController) Create() {
	var ob *models.Volume
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	result, err := models.CreateVolume(ob)
	if(err == nil){
		c.Data["json"] = map[string]bool{"status": result}
	}
	if(err != nil){
		c.Data["json"] = map[string]string{"error": err.Error()}
	}
	c.ServeJSON()
}

// @router / [delete]
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
		c.Data["json"] = map[string]string{"status": result}
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
// @router / [get]
func (c *VolumeController) List() {
	result, err := models.ListVolume()
	if(err == nil){
		c.Data["json"] = map[string]string{"status": result}
	}
	if(err != nil){
		c.Data["json"] = map[string]string{"error": err.Error()}
	}
}