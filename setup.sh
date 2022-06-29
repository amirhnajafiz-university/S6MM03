#!/bin/sh

echo "Rolling in ..."

# docker start
docker compose up -d 

# python setup
cd python

# installing dependencies
pip3 install -r requirements.txt

# running python for videos
python3 script.py videos/1.mp4 1 && python3 script.py videos/2.mp4 2

make d-pus

echo "Rolling out ..."
