#!/bin/bash
nohup ./go-chat >> stdout.log 2>&1 &
while true ;do ls;sleep 100;done