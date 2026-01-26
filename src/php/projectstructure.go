package php

import (
	"github/MaminirinaEdwino/api-maker-cli/src/utils"
	"os"
)

func CreateProjectStructure(Projectname string){
	err := os.MkdirAll(Projectname+"/asset/style", os.ModePerm)
	utils.ErrorChecker(err)
	err = os.MkdirAll(Projectname+"/asset/script", os.ModePerm)
	utils.ErrorChecker(err)
	err = os.MkdirAll(Projectname+"/asset/image", os.ModePerm)
	utils.ErrorChecker(err)
	err = os.MkdirAll(Projectname+"/src/auth", os.ModePerm)
	utils.ErrorChecker(err)
	err = os.MkdirAll(Projectname+"/src/components", os.ModePerm)
	utils.ErrorChecker(err)
	err = os.MkdirAll(Projectname+"/src/config", os.ModePerm)
	utils.ErrorChecker(err)
	err = os.MkdirAll(Projectname+"/src/controllers", os.ModePerm)
	utils.ErrorChecker(err)
	err = os.MkdirAll(Projectname+"/src/core", os.ModePerm)
	utils.ErrorChecker(err)
	err = os.MkdirAll(Projectname+"/src/models", os.ModePerm)
	utils.ErrorChecker(err)
	err = os.MkdirAll(Projectname+"/src/repositories", os.ModePerm)
	utils.ErrorChecker(err)
	err = os.MkdirAll(Projectname+"/src/router", os.ModePerm)
	utils.ErrorChecker(err)
	err = os.MkdirAll(Projectname+"/src/templaterenderer", os.ModePerm)
	utils.ErrorChecker(err)
	err = os.MkdirAll(Projectname+"/src/views", os.ModePerm)
	utils.ErrorChecker(err)
}