@echo off

.\xorm.exe reverse mysql root:@(127.0.0.1:3306)/lovely?charset=utf8 D:\job\GO\GOPATH\src\github.com\go-xorm\cmd\xorm\templates\goxorm models

pause