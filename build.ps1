$env:GOOS='linux'

switch($args[1])
{
    'api' {
        .\boxgen.exe --action routergen
        $apiOut = ".\bin\api.linux." + $args[0]
        go build -o $apiOut .\app\api\home\main.go .\app\api\home\register.go
    }
    Default {
        "wrong"
    }
}

$env:GOOS='windows'
