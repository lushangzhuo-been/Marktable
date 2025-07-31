package common

import (
	"errors"
	"mark3/global"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

func TestToPdf(fileSrcPath string, outPath string, fileType string) (fileOutPath string, error error) {

	osName := runtime.GOOS //获取系统类型
	switch osName {
	case "darwin": //mac系统
		command := "/Applications/LibreOffice.app/Contents/MacOS/soffice"
		pdfFile, err := FuncDocs2Pdf(command, fileSrcPath, outPath, fileType)
		if err != nil {
			global.GVA_LOG.Error("转化异常：", err.Error())
			return "", errors.New("上传文件转PDF失败")
		}
		global.GVA_LOG.Info("转化后的文件：", pdfFile)
		return pdfFile, err
	case "linux":
		command := global.GVA_CONFIG.LibreOffice.Version
		pdfFile, err := FuncDocs2Pdf(command, fileSrcPath, outPath, fileType)
		if err != nil {
			global.GVA_LOG.Error("转化异常：", err.Error())
			return "", errors.New("上传文件转PDF失败")
		}
		global.GVA_LOG.Info("转化后的文件：", pdfFile)
		return pdfFile, err
	case "windows":
		command := "soffice"
		pdfFile, err := FuncDocs2Pdf(command, fileSrcPath, outPath, fileType)
		if err != nil {
			global.GVA_LOG.Error("转化异常：", err.Error())
			return "", errors.New("上传文件转PDF失败")
		}
		global.GVA_LOG.Info("转化后的文件：", pdfFile)
		return pdfFile, err
	default:
		global.GVA_LOG.Error("暂时不支持的系统转化:" + runtime.GOOS)
		return "", errors.New("暂时不支持的系统转化:" + runtime.GOOS)
	}
}

/**
*@tips libreoffice 转换指令：
* libreoffice6.2 invisible --convert-to pdf csDoc.doc --outdir /home/[转出目录]
*
* @function 实现文档类型转换为pdf或html或其他格式
* @param command:libreofficed的命令(具体以版本为准)；win：soffice； linux：libreoffice6.2
*     fileSrcPath:转换文件的路径
*         fileOutDir:转换后文件存储目录
*       converterType：转换的类型pdf/html
* @return fileOutPath 转换成功生成的文件的路径 error 转换错误
 */
func FuncDocs2Pdf(command string, fileSrcPath string, fileOutDir string, converterType string) (fileOutPath string, error error) {
	//校验fileSrcPath
	srcFile, erByOpenSrcFile := os.Open(fileSrcPath)
	if erByOpenSrcFile != nil && os.IsNotExist(erByOpenSrcFile) {
		return "", erByOpenSrcFile
	}
	//如文件输出目录fileOutDir不存在则自动创建
	outFileDir, erByOpenFileOutDir := os.Open(fileOutDir)
	if erByOpenFileOutDir != nil && os.IsNotExist(erByOpenFileOutDir) {
		erByCreateFileOutDir := os.MkdirAll(fileOutDir, os.ModePerm)
		if erByCreateFileOutDir != nil {
			global.GVA_LOG.Error("File ouput dir create error.....", erByCreateFileOutDir.Error())
			return "", erByCreateFileOutDir
		}
	}
	//关闭流
	defer func() {
		_ = srcFile.Close()
		_ = outFileDir.Close()
	}()
	//convert
	cmd := exec.Command(command, "--invisible", "--convert-to", converterType,
		fileSrcPath, "--outdir", fileOutDir)
	byteByStat, errByCmdStart := cmd.Output()
	//命令调用转换失败
	if errByCmdStart != nil {
		return "", errByCmdStart
	}
	//success
	fileOutPath = fileOutDir + "/" + strings.Split(path.Base(fileSrcPath), ".")[0]
	fileOutPath += "." + converterType
	global.GVA_LOG.Info("文件转换成功...", string(byteByStat))
	return fileOutPath, nil
}
