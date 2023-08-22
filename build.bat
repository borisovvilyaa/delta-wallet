@echo off

echo " ___   ____  _    _____   __        _        __    _     _     ____ _____ "
echo "| | \ | |_  | |    | |   / /\      \ \    / / /\  | |   | |   | |_   | |  "
echo "|_|_/ |_|__ |_|__  |_|  /_/--\      \_\/\/ /_/--\ |_|__ |_|__ |_|__  |_|  "

rem Обновляем зависимости
go mod tidy

rem Копируем зависимости в папку vendor
go mod vendor

rem Запускаем веб-браузер (Microsoft Edge) для открытия локального URL
start msedge http://localhost:8080/

rem Запускаем основное приложение
go run main.go
