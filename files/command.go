package files

import (
	"fmt"
	"strings"
)

func generateCommand(args Args) {
	template := "templates/Command.php"
	file := args.AppPath + "/src/Domain/" + args.Domain + "/Command/" + args.Name + "Command.php"

	if !Exist(template) {
		fmt.Println("Le template Command " + template + " n'existe pas.")
	}

	if Exist(file) {
		fmt.Println("Le fichier " + file + " existe deja.")
	}

	params := strings.ReplaceAll(args.Params, "-", ";\n")+";"
	GenerateFiles(template, file , args, params)
}

func generateCommandHandler(args Args) {
	template := "templates/CommandHandler.php"
	file := args.AppPath + "/src/Domain/" + args.Domain + "/Command/" + args.Name + "CommandHandler.php"

	if !Exist(template) {
		fmt.Println("Le template CommandHandler " + template + " n'existe pas.")
	}

	if Exist(file) {
		fmt.Println("Le fichier " + file + " existe deja.")
	}
	params := strings.ReplaceAll(args.Params, "-", ",\n")

	GenerateFiles(template, file , args, params)

}

func appendCommandService(args Args) {
	filePath := args.AppPath+"/config/services/command_subscribe.yaml"
	if (Exist(filePath)) {
		content := "  Macompta\\Api\\"+args.BasePath+"\\Domain\\"+args.Domain+"\\Command\\"+args.Name+"Command: 'Macompta\\Api\\"+args.BasePath+"\\Domain\\"+args.Domain+"\\Command\\"+args.Name+"CommandHandler'"
		Append(filePath, content)
	} else {
		fmt.Println("Le fichier "+filePath+ "n'existe pas.")
	}
}


func GenerateCommandFiles(args Args) {
	generateCommand(args)
	generateCommandHandler(args)
	appendCommandService(args)

}

