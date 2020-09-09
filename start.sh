#!/bin/sh

nohup ./bin/api.exe > log/api.log 2>&1 &
echo "start api"
nohup ./bin/scheduler.exe > log/scheduler.log 2>&1 &
echo "start scheduler.exe"
nohup ./bin/streamserver.exe > log/streamserver.log 2>&1 &
echo "start streamserver"
nohup ./bin/web.exe > log/web.log 2>&1 &
echo "start web"