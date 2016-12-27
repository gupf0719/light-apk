package controllers

import (
	"fmt"
	"light-apk/models"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// 获取文件大小的接口
type Size interface {
	Size() int64
}

func (c *MainController) Get() {
	c.TplName = "index.html"

	apks, merr := models.GetApkAll()
	if merr != nil {
		beego.Error(merr)
	} else {
		c.Data["Apks"] = apks
	}
}

func (c *MainController) GetLatest() {
	callback_data := make(map[string]interface{}) //返回数据

	apk, merr := models.GetLatestApk()
	if merr != nil {
		beego.Error(merr)
	}
	callback_data["apk"] = apk
	c.Data["json"] = callback_data
	c.ServeJSON()
}

func (c *MainController) Post() {
	name := c.GetString("name")
	code := c.GetString("code")
	ulog := c.GetString("log")
	f, h, ferr := c.GetFile("file")
	if ferr != nil {
		beego.Error(ferr)
	} else {
		fmt.Println(h.Filename)
		serr := c.SaveToFile("file", path.Join("static/files", h.Filename))
		if serr != nil {
			beego.Error(serr)
		}
		f.Close()
	}

	var fileSize int64
	if sizeInterface, ok := f.(Size); ok {
		fileSize = sizeInterface.Size()
	}
	fsize := strconv.FormatInt(fileSize, 10)

	server_url := beego.AppConfig.String("server_url")

	apk := models.Apk{VersionName: name, VersionCode: code, Updatelog: ulog, Filename: h.Filename, DownloadUrl: server_url + "/download/" + h.Filename,
		Size: fsize, CreatedAt: time.Now().Local()}

	_, merr := models.AddApk(&apk)
	if merr != nil {
		beego.Error(merr)
	}

	c.Redirect("/", 302)
}

func (c *MainController) Delete() {
	id, _ := c.GetInt("id")

	fname, merr := models.DeleteApk(id)
	if merr != nil {
		beego.Error(merr)
	}

	file := "static/files/" + fname //源文件路径
	err := os.Remove(file)          //删除文件test.txt
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/", 302)
}

//func (c *MainController) Download() {
//	name := c.GetString("name")

//	c.Ctx.Output.Download("files/" + name)

//	//	c.Redirect("/", 302)
//}
