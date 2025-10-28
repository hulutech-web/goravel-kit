package excel

import (
	"fmt"
	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/path"
	"goravel/app/models"
	"strings"
	"time"
)

type ExcelService struct{}

func NewExcelService() *ExcelService {
	return &ExcelService{}
}

func (e *ExcelService) ReadStreamExcel(ctx http.Context, file filesystem.File) (*models.File, error) {

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
	if err2 := e.ImportUserData(att.Uri); err2 != nil {
		return nil, err2
	}
	if err2 := facades.Orm().Query().Model(&models.File{}).Create(&att); err2 != nil {
		return nil, err2
	}

	//直接导入users

	return &att, err
}

func (e *ExcelService) ImportUserData(pathStr string) error {
	// 1. 打开Excel文件
	target_path := path.Public(pathStr)
	f, err := excelize.OpenFile(target_path)
	if err != nil {
		return err
	}
	defer f.Close()

	// 获取第一个工作表名
	sheetName := f.GetSheetName(0)

	// 2. 解析头信息（假设第一行是表头）
	headers, err := f.GetRows(sheetName, excelize.Options{RawCellValue: true})
	if err != nil {
		return err
	}

	if len(headers) == 0 {
		return err
	}

	// 打印表头
	fmt.Println("=== 表头信息 ===")
	for i, header := range headers[0] {
		fmt.Printf("列 %d: %s\n", i+1, header)
	}
	facades.Orm().Query().Exec("truncate table `user_histories`")
	users := []models.UserHistory{}
	// 3. 解析内容信息（从第二行开始）
	fmt.Println("\n=== 内容数据 ===")
	for rowIdx, row := range headers {
		if rowIdx == 0 {
			continue // 跳过表头行
		}
		fmt.Printf("第 %d 行数据:\n", rowIdx+1)
		user := models.UserHistory{}
		for colIdx, cell := range row {
			headerName := headers[0][colIdx] // 获取对应的列名
			fmt.Printf("  %s: %s\n", headerName, cell)
			if colIdx == 1 {
				user.Realname = cell
			}
			if colIdx == 2 {
				user.Sex = cell
			}
			if colIdx == 3 {
				user.Sex = cell
			}
		}
		users = append(users, user)
	}
	return facades.Orm().Query().Create(&users)
}
