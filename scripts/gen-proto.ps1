$WebClient = New-Object System.Net.WebClient

$WorkDir = (Get-Location)
$ProtoExec = Join-Path -Path $WorkDir -ChildPath "protoc.exe"
If (-Not (Test-Path $ProtoExec)){
    $ProtoFile = Join-Path -Path (Get-Location) -ChildPath "proto.zip"
    $WebClient.DownloadFile("https://github.com/protocolbuffers/protobuf/releases/download/v21.2/protoc-21.2-win64.zip",$ProtoFile)
    $TempDir = New-Item -Path (Join-Path -Path (Get-Location) -ChildPath "_TEMP") -ItemType Directory -Force
    Expand-Archive -Path $ProtoFile -DestinationPath $TempDir
    Join-Path -Path $TempDir -ChildPath "bin/protoc.exe"
    Move-Item -Path (Join-Path -Path $TempDir -ChildPath "bin/protoc.exe") -Destination $WorkDir
    Remove-Item -Path $TempDir -Force -Recurse
    Remove-Item -Path $ProtoFile
}
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
protoc --go_out=../pkg/ --go_opt=paths=source_relative --go-grpc_out=../pkg/ --go-grpc_opt=paths=source_relative --proto_path=.. proto/bgp.proto