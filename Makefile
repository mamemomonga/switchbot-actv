build:
	mkdir -p dist/windows-amd64
	GOOS=windows GOARCH=amd64 go build -o dist/switchbot-actv-windows-amd64/switchbot-actv.exe ./
	cp helpers/start.bat dist/switchbot-actv-windows-amd64/
	cd dist && 7za a -tzip ./switchbot-actv-windows-amd64.zip switchbot-actv-windows-amd64
