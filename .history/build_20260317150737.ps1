$Env: GOOS = 'js'
$Env: GOARCH = 'wasm'
go build -o game.wasm 
Remove-Item Env:GOOS
Remove-Item Env:GOARCH