package files

func GenerateFiles(template string, file string, args Args, params string) {

	var templateContent []string
	var templateContentParsed []string

	templateContent = Read(template)
	templateContentParsed = Parser(templateContent, args, params)

	Write(file, templateContentParsed)
}