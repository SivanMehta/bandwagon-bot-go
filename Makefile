.PHONY:
serve:
	go run application.go

application.zip:
	rm -rf deploy/*
	cp -r public/ deploy
	cp application.go deploy
	zip -r application.zip deploy/