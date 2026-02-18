package php

import (
	"github/MaminirinaEdwino/api-maker-cli/src/utils"
	"os"
)
func CreateTemplateRendererFile(projectName string){
	file, err := os.Create(projectName+"/src/templaterenderer/templateRenderer.php")
	utils.ErrorChecker(err)
	file.WriteString(`
<?php

class TemplateRender{

    public static function render(string $path, $param){
        $params = $param;
        include_once "./src/views".$path;
    }
}	
`)
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

func CreateFileUploaderFile(projectName string){
	file, err := os.Create(projectName+"/src/core/fileUploader.php")
	utils.ErrorChecker(err)
	file.WriteString(`
<?php
function FileUploader(string $file_name): string
{
    if (
        isset($_FILES[$file_name]) and $_FILES[$file_name]['error'] == 0
    ) {
        if ($_FILES[$file_name]['size'] <= 1000000) {
            $infosfichier =
                pathinfo($_FILES[$file_name]['name']);
            $extension_upload = $infosfichier['extension'];
            $extensions_autorisees = array(
                'jpg',
                'jpeg',
                'gif',
                'png'
            );
            if (in_array(
                $extension_upload,
                $extensions_autorisees
            )) {
                move_uploaded_file($_FILES[$file_name]['tmp_name'], 'uploads/' .
                    basename($_FILES[$file_name]['name']));
                echo "L'envoi a bien été effectué !";
                return 'uploads/' .
                    basename($_FILES[$file_name]['name']);
            }
        }
    }
    return "";
}
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