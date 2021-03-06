package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"gluster-api/models"
	"bytes"
)

type QuotaController struct {
	beego.Controller
}

// @Title Create Quota
// @Description Create Quota
// @Param Quota
// @Success 200
// @Failure 400 body is empty
// @router / [post]
func (c *QuotaController) CreateQuota() {
	var ob *models.QuotaParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	ob.Percent = "100%"
	result, err := models.CreateQuota(ob)
	if(err == nil){
		c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		c.Ctx.Output.Body(bytes.NewBufferString(result).Bytes())
	}
	if(err != nil){
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
	}
}

// @Title Change Quota
// @Description Change Quota
// @Param Quota
// @Success 200
// @Failure 400 body is empty
// @router / [put]
func (c *QuotaController) ChangeQuota() {
	var ob *models.QuotaParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	initGlusterParam()
	result, err := models.ChangeQuota(ob)
	if(err == nil){
		c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		c.Ctx.Output.Body(bytes.NewBufferString(result).Bytes())
	}
	if(err != nil){
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
	}
}

// @Title Remove Quota
// @Description Remove Quota
// @Param Quota
// @Success 200
// @Failure 400 body is empty
// @router /:volume [delete]
func (c *QuotaController) RemoveQuota() {
	volume := c.GetString(":volume")
	result, err := models.RemoveQuota(volume)
	if(err == nil){
		c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		c.Ctx.Output.Body(bytes.NewBufferString(result).Bytes())
	}
	if(err != nil){
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
	}
}

// @Title Remove Quota
// @Description Remove Quota
// @Param Quota
// @Success 200
// @Failure 400 body is empty
// @router /:volume [get]
func (c *QuotaController) QueryQuota() {
	volume := c.GetString(":volume")
	result, err := models.QueryQuota(volume)
	if(err == nil){
		c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		c.Ctx.Output.Body(bytes.NewBufferString(result).Bytes())
	}
	if(err != nil){
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
	}
}