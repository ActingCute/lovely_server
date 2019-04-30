@echo off

.\xorm.exe reverse mysql root:admin@(127.0.0.1:3306)/githair_master?charset=utf8 D:\GOPATH\src\github.com\go-xorm\cmd\xorm\templates\goxorm models

pause