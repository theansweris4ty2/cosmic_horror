
function web {
$env:GOOS = 'js' 
$env:GOARCH = 'wasm' 
go build -o game.wasm  .
Remove-Item Env:GOOS
Remove-Item Env:GOARCH
}

function build {

    go build -o cosmic_horror.exe -ldflags="-w -s" .
}

function wasmserve {
    go run github.com/hajimehoshi/wasmserve@latest .
}






