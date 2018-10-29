$env:GOARCH="wasm";
$env:GOOS="js";
go build -o index.wasm main.go aStar.go node.go;
$env:GOARCH="";
$env:GOOS="";
