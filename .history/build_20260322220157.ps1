$env:GOOS = 'js' 
$env:GOARCH = 'wasm' 
go build -o game.wasm  main.go
Remove-Item Env:GOOS
Remove-Item Env:GOARCH


