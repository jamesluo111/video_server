#!/bin/bash

# Build web UI
cd /wwwroot/video_server/web
go install
cp /go/bin/web /go/bin/video_server_web_ui/web
cp -R /wwwroot/video_server/templates /go/bin/video_server_web_ui/
