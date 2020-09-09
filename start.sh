#!/bin/sh
nohup ./bin/api > log/api.log 2>&1 &
echo "start api"
nohup ./bin/scheduler > log/scheduler.log 2>&1 &
echo "start scheduler.exe"
nohup ./bin/streamserver > log/streamserver.log 2>&1 &
echo "start streamserver"
nohup ./bin/web > log/web.log 2>&1 &
echo "start web"