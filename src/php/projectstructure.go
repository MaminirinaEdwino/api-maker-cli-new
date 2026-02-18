package php

import (
	"github/MaminirinaEdwino/api-maker-cli/src/utils"
	"os"
)
func CreateCoreFile(){

}

func CreateTemplateRendererFile(){

}
func CreateUtilsFile(projectName string){
	file, err := os.Create(projectName+"/scr/core/utils/relativeRoutes.php")
	utils.ErrorChecker(err)
	file.WriteString(`
<?php
function To_relative_path($route ){
    $current_uri = $_SERVER['REQUEST_URI'];
    $depth = substr_count(trim(dirname($current_uri), '/'), '/');
    $relative_path_to_root = str_repeat('../', $depth); 
    return $relative_path_to_root.$route;
}
?>
	`)
}
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
	err = os.MkdirAll(Projectname+"/src/core/utils", os.ModePerm)
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