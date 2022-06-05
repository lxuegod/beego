package chapter02

import (
	"fmt"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type UploadController struct {
	beego.Controller
}

func (u *UploadController) Get() {
	u.TplName = "chapter02/upload.html"
}
func (u *UploadController) Post() {
	//获取上传的文件
	f, h, _ := u.GetFile("file")
	defer f.Close()
	fmt.Println(h.Filename)

	//保存上传的文件
	time_unix := time.Now().Unix()
	file_path := fmt.Sprintf("%d_%s", time_unix, h.Filename)
	u.SaveToFile("file", "upload/"+file_path)

	u.Ctx.WriteString("上传成功！")
}
