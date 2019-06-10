.PHONY serve:
serve:
	go run application.go

application.zip:
	zip -r application.zip .

.PHONY environment:
environment:
	@echo Restoring ozone layer...

.PHONY deployment:
deployment:
	@echo downloading RAM to remote server...
