package upload

import (
	"fmt"
	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/spf13/cast"
	"goravel/app/models"
	"strings"
	"time"
)

type UploadService struct {
}

func NewUploadService() *UploadService {
	return &UploadService{}
}
func (*UploadService) Upload(ctx http.Context, file filesystem.File, cate_id int) (*models.File, error) {
	/*	if cate_id == 0 {
		return nil, fmt.Errorf("请选择分类，为默认分类")
	}*/
	//获取原始的名字
	name := file.GetClientOriginalName()
	size, _ := file.Size()
	extension := file.GetClientOriginalExtension()
	//获取当前年月2022-02
	yearMonth := fmt.Sprintf("%d-%02d", time.Now().Year(), time.Now().Month())
	putFile, err := facades.Storage().PutFile(yearMonth, file)
	path_str := strings.ReplaceAll(putFile, "\\", "/")

	//根据extension判断是file,image,video,audio
	if err != nil {
		return nil, err
	}
	typeName := ""
	switch extension {
	case "jpg", "jpeg", "png", "gif", "bmp", "webp":
		typeName = "image"
	case "mp4", "avi", "mkv", "mov", "wmv", "flv", "m4v", "mpg", "mpeg", "3gp", "ogg", "webm":
		typeName = "video"
	case "mp3", "wav", "aac", "m4a", "wma", "flac", "alac", "aiff", "ape", "midi", "amr":
		typeName = "audio"
	default:
		typeName = "file"
	}
	//保存文件，返回保存路径
	user := models.User{}
	if err1 := facades.Auth(ctx).User(&user); err1 != nil {
		return nil, err1
	} else {
		att := models.File{
			UserID: user.ID,
			Type:   typeName,
			CID:    cast.ToUint(cate_id),
			Name:   name,
			//文件路径
			Uri:      "/uploads/" + path_str,
			Ext:      extension,
			Size:     size,
			Engine:   "local",
			Path:     path_str,
			TenantID: 1,
		}
		if err2 := facades.Orm().Query().Model(&models.File{}).Create(&att); err2 != nil {
			return nil, err2
		}
		return &att, err
	}
}

func (*UploadService) UploadAvatar(ctx http.Context, file filesystem.File) (*models.File, error) {
	/*	if cate_id == 0 {
		return nil, fmt.Errorf("请选择分类，为默认分类")
	}*/
	//获取原始的名字
	name := file.GetClientOriginalName()
	size, _ := file.Size()
	extension := file.GetClientOriginalExtension()
	//获取当前年月2022-02
	yearMonth := fmt.Sprintf("%d-%02d", time.Now().Year(), time.Now().Month())
	putFile, err := facades.Storage().PutFile(yearMonth, file)
	// 统一转换为正斜杠
	path_str := strings.ReplaceAll(putFile, "\\", "/")
	//根据extension判断是file,image,video,audio
	if err != nil {
		return nil, err
	}
	typeName := ""
	switch extension {
	case "jpg", "jpeg", "png", "gif", "bmp", "webp":
		typeName = "image"
	case "mp4", "avi", "mkv", "mov", "wmv", "flv", "m4v", "mpg", "mpeg", "3gp", "ogg", "webm":
		typeName = "video"
	case "mp3", "wav", "aac", "m4a", "wma", "flac", "alac", "aiff", "ape", "midi", "amr":
		typeName = "audio"
	default:
		typeName = "file"
	}
	//保存文件，返回保存路径
	att := models.File{
		Type: typeName,
		CID:  10,
		Name: name,
		//文件路径
		Uri:      "/uploads/" + path_str,
		Ext:      extension,
		Size:     size,
		Engine:   "local",
		Path:     path_str,
		TenantID: 1,
	}
	if err2 := facades.Orm().Query().Model(&models.File{}).Create(&att); err2 != nil {
		return nil, err2
	}
	return &att, err
}

type Ret struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	ExtName string `json:"extname"`
	Url     string `json:"url"`
}

/*认证信息上传*/
func (*UploadService) UploadCert(ctx http.Context, file filesystem.File) (*Ret, error) {
	/*	if cate_id == 0 {
		return nil, fmt.Errorf("请选择分类，为默认分类")
	}*/
	//获取原始的名字
	name := file.GetClientOriginalName()
	size, _ := file.Size()
	extension := file.GetClientOriginalExtension()
	//获取当前年月2022-02
	yearMonth := fmt.Sprintf("%d-%02d", time.Now().Year(), time.Now().Month())
	putFile, err := facades.Storage().PutFile(yearMonth, file)
	// 统一转换为正斜杠
	path_str := strings.ReplaceAll(putFile, "\\", "/")
	//根据extension判断是file,image,video,audio
	if err != nil {
		return nil, err
	}
	typeName := ""
	switch extension {
	case "jpg", "jpeg", "png", "gif", "bmp", "webp":
		typeName = "image"
	case "mp4", "avi", "mkv", "mov", "wmv", "flv", "m4v", "mpg", "mpeg", "3gp", "ogg", "webm":
		typeName = "video"
	case "mp3", "wav", "aac", "m4a", "wma", "flac", "alac", "aiff", "ape", "midi", "amr":
		typeName = "audio"
	default:
		typeName = "file"
	}
	//保存文件，返回保存路径
	att := models.File{
		Type: typeName,
		CID:  10,
		Name: name,
		//文件路径
		Uri:      "/uploads/" + path_str,
		Ext:      extension,
		Size:     size,
		Engine:   "local",
		Path:     path_str,
		TenantID: 1,
	}
	if err2 := facades.Orm().Query().Model(&models.File{}).Create(&att); err2 != nil {
		return nil, err2
	}

	ret := Ret{
		ID:      att.ID,
		Name:    att.Name,
		ExtName: att.Ext,
		Url:     att.Uri,
	}
	return &ret, err
}
