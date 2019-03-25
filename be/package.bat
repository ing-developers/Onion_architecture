:: AUTOR MALDRU
:: BAT PARA LA INSTALACION DE DEPENDENCIAS EN GO
@echo off
echo "OBTENIENDO DEPENDENCIAS..."

::DEPENDENCIAS DEL PROYECTO (SOLO AGREGAR EL NOMBRE DE LA DEPENDENCIA O LIBRERIA DE GITHUB)
call :goGet "go-sql-driver/mysql"
call :goGet "lib/pq"
call :goGet "ing-developers/go-tools"
call :goGet "ing-developers/go-dal"
::DEPENDENCIAS DEL PROYECTO

echo "FIN DE PROCESO"
pause
EXIT /B
::FUNCION PARA DESCARGAR DEPENDENCIAS
:goGet
go get -u github.com/%~1