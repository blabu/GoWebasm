##Test project golang + webasm
Requirements: 
1. Golang version 1.13 or above 
2. Specified GOOS=js and GOARCH=wasm in your build environment
3. Typed go build -o main.wasm ./ for build wasm file.
Thats file will be fetch in html when it load in browser.
Use some simple web server like GoLive or something else.
