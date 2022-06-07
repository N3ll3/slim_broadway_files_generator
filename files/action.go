package files

import "fmt"


func GenerateActionFiles(args Args) {
	template := "templates/Action.php"
	file := args.AppPath + "/src/Application/Actions/" + args.Name + "Action.php"

	if !Exist(template) {
		fmt.Println("Le template Action "+template+" n'existe pas.")
	}

	if Exist(file) {
		fmt.Println("Le fichier "+file+" existe deja.")
	}

	GenerateFiles(template, file, args, "")

}
