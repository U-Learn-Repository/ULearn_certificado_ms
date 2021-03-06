package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/diagutierrezro/ULearn_certificado_ms/controllers:CertificadoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/ULearn_certificado_ms/controllers:CertificadoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/ULearn_certificado_ms/controllers:CertificadoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/ULearn_certificado_ms/controllers:CertificadoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/ULearn_certificado_ms/controllers:CertificadoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/ULearn_certificado_ms/controllers:CertificadoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/ULearn_certificado_ms/controllers:CertificadoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/ULearn_certificado_ms/controllers:CertificadoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/ULearn_certificado_ms/controllers:CertificadoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/ULearn_certificado_ms/controllers:CertificadoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/:name/:surname/:documento/:nombreCurso/:duracionCurso",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
