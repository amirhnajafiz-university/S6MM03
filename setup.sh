#!/bin/sh

echo "Rolling in ..."

# docker start
docker compose up -d 

# python setup
cd python

# installing dependencies
if [ $1 == "p3" ] then
    pip3 install -r requirements.txt
else
    pip install -r requirements.txt
fi

# running python for videos
if [ $1 == "p3" ] then
    python3 script.py videos/1.mp4 1 && python3 script.py videos/2.mp4 2
else
    python script.py videos/1.mp4 1 && python script.py videos/2.mp4 2
fi

# send files to docker
make d-push

echo "Rolling out ..."
