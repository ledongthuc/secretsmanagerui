clean:
	rm -f ./output/build/*
	rm -f ./output/dist/*
build:
	go build -o ./output/build/secretmanagerui
run: build
	./output/build/secretmanagerui
package: build
	fyne package -executable ./output/build/secretmanagerui -icon Icon.png
