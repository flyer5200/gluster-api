package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["gluster-api/controllers:VolumeController"] = append(beego.GlobalControllerRouter["gluster-api/controllers:VolumeController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gluster-api/controllers:VolumeController"] = append(beego.GlobalControllerRouter["gluster-api/controllers:VolumeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["gluster-api/controllers:VolumeController"] = append(beego.GlobalControllerRouter["gluster-api/controllers:VolumeController"],
		beego.ControllerComments{
			Method: "Query",
			Router: `/:name`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gluster-api/controllers:VolumeController"] = append(beego.GlobalControllerRouter["gluster-api/controllers:VolumeController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
