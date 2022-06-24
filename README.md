# slimfilesgenerator avec Go et Cobra

* Installer Go : https://go.dev/doc/install
* Ajouter dans $HOME/.profile
  
    ```
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
    ```
  
### Syntax

| root command  | Command | Resource | arguments | flags |
|--|--|--|--|--|
|  slimfilesgenerator| generate| action | PrefixDeMaClass | --appPath=chemin/absolu/api --basePath=base/path --domain=Domain|
|  slimfilesgenerator| generate| command | PrefixDeMaClass | --appPath=chemin/absolu/api --basePath=base/path --domain=Domain --params="param1, param2, param3"|
|  slimfilesgenerator| generate| event | PrefixDeMaClass | --appPath=chemin/absolu/api --basePath=base/path --domain=Domain --params="param1, param2, param3" --eventAction=""|
|  slimfilesgenerator| generate| all | PrefixDeMaClass | --appPath=chemin/absolu/api --basePath=base/path --domain=Domain --params="param1, param2, param3" --eventAction=""|

NB : echapper les slash et antislash

# Installation
```
go install slimfilesgenerator
```

# Play with slimfilesgenerator

## Prérequis

Lancer la ligne de commande dans un dossier qui dispose lui-même d'un dossier 'templates' avec les fichiers suivants présents :

* Action.php
* Command.php
* CommandHandler.php
* Event.php

Les variables dans les templates sont :

  * {%BASE_PATH%} => --BasePath
  * {%DOMAIN%} => --domain
  * {%PARAMS%} => --params
  * {%EVENT_NAME%} => --EventName
  * {%NAME%} => {arguments} PrefixDeMaClass
	
}


```
Genere Action, Command, Event, All
        exemple : slimfilesgenerator generate action test --appPath={} --basePath={} --domain={} --eventName={} --params="param1, param2, param3,..."

Usage:
  slimfilesgenerator generate <file> <name> [flags]
  slimfilesgenerator generate [command]

Available Commands:
  action      Generer le fichier Action
  all         Generer tous les fichiers
  command     Generer les fichiers command
  event       Generer le fichier event

Flags:
  -h, --help   help for generate

Global Flags:
  -a, --appPath string
  -b, --basePath string
  -d, --domain string
  -p, --parms string
  -e, --eventName string

Use "slimfilesgenerator generate [command] --help" for more information about a command.
 ```

 
## Example
```
$ slimfilesgenerator generate all ModifierCodeCouleur --appPath=$HOME/Sites/applications/api --basePath=Admin\\User --domain=User --params="codeCouleur" --eventName=AeteModifie
```


