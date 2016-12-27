package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Apk struct {
	Id          int       `orm:"auto"`
	VersionName string    `orm:"size(64)"`           //版本号
	VersionCode string    `orm:"size(64)"`           //编译版本号
	Updatelog   string    `orm:"size(1024)"`         //更新日志
	Filename    string    `orm:"size(64)"`           //文件名
	DownloadUrl string    `orm:"size(128)"`          //下载地址
	Size        string    `orm:"size(64)"`           //文件大小
	CreatedAt   time.Time `orm:"index;auto_now_add"` //创建时间
}

func init() {
	orm.RegisterModel(new(Apk))
}

func AddApk(m *Apk) (bool, error) {
	o := orm.NewOrm()

	qs := o.QueryTable("apk")
	ex := qs.Filter("version_name", m.VersionName).Exist()

	if !ex {
		_, err := o.Insert(m)
		return false, err
	} else {
		return true, nil
	}
}

func GetApkAll() ([]*Apk, error) {
	o := orm.NewOrm()
	apks := make([]*Apk, 0)

	qs := o.QueryTable("apk")

	_, err := qs.OrderBy("-id").All(&apks)

	return apks, err
}

func GetLatestApk() (*Apk, error) {
	o := orm.NewOrm()

	apk := new(Apk)

	qs := o.QueryTable("apk")
	err := qs.OrderBy("-id").Limit(1, 0).One(apk)
	if err != nil {
		return nil, err
	}

	return apk, err
}

func DeleteApk(aid int) (string, error) {
	o := orm.NewOrm()

	apk := new(Apk)

	qs := o.QueryTable("apk")
	err := qs.Filter("id", aid).One(apk)
	if err != nil {
		return "", err
	} else {
		_, err2 := o.Delete(&Apk{Id: aid})
		if err2 != nil {
			return "", err2
		} else {
			return apk.Filename, err2
		}
	}

}
