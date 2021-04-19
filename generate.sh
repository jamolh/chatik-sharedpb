# Файлы нужно генерировать в папке git.alifpay.tj/terminals/sharedpb
# Папка sharedpb - корневой каталог для всех прото-файлов

# -I – defining the root folder for protocol buffer
# --go_out – folder for output
protoc --go_out=$GOPATH/src type/user.proto
protoc --go_out=$GOPATH/src type/message.proto

# Генерация сервисов
protoc --go_out=plugins=grpc:$GOPATH/src service/chatik.proto