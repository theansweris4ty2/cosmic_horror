
# function Build-Task {
# $env:GOOS = 'js' 
# $env:GOARCH = 'wasm' 
# go build -o game.wasm  
# Remove-Item Env:GOOS
# Remove-Item Env:GOARCH
# }
function Build-Task {
    go build .
}
function Run-Task {
    go run .
}