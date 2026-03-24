
# function Build-Task {
# $env:GOOS = 'js' 
# $env:GOARCH = 'wasm' 
# go build -o game.wasm  
# Remove-Item Env:GOOS
# Remove-Item Env:GOARCH
# }

function Run {
    go run .
}

switch($Target){
    "Run" { Run }
}