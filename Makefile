ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find rpc -name *.proto")
else
	API_PROTO_FILES=$(shell find . -name "*.proto")
endif

api:
	cd api
	goctl api go -api ./run.api -dir ../
	goctl api plugin -plugin goctl-swagger="swagger -filename api.json" -api run.api -dir ../swagger





