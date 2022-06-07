package files

import (
	"fmt"
	"strings"
)

func generateEvent(args Args) {
	template := "templates/Event.php"
	file := args.AppPath + "/src/Domain/" + args.Domain + "/EventSourcing/" + args.EventName +"Event.php"

	if !Exist(template) {
		fmt.Println("Le template Event " + template + " n'existe pas.")
	}

	if Exist(file) {
		fmt.Println("Le fichier " + file + " existe deja.")
	}
	params := strings.ReplaceAll(args.Params, "-", ";\n")+";"
	GenerateFiles(template, file, args, params)
}

func appendEventListener(args Args) {
	filePath := args.AppPath + "/config/services/event_listener.yaml"
	if Exist(filePath) {
		content := "   Macompta\\Api\\" + args.BasePath + "\\Domain\\" + args.Domain + "\\EventSourcing\\" + args.EventName +"Event:\n        -Macompta\\Api\\" + args.AppPath + "\\Infrastructure\\" + args.Domain + "\\EventListener\\Projection"+args.Domain+"EventListener"
		Append(filePath, content)
	} else {
		fmt.Println("Le fichier " + filePath + " n'existe pas.")
	}
}

func GenerateEventFiles(args Args) {
	generateEvent(args)
	appendEventListener(args)
}
