#!/bin/bash

#set enviroment keys
source /root/keys/set_env_keys.sh

 curl -sS \
  -d "auth_token=$HIPCHAT_TOKEN&room_id=$HIPCHAT_ROOM&from=prodbot&color=purple&message=Deploying! user: `whoami`&notify=1" \
  https://api.hipchat.com/v1/rooms/message

# Assumes other environment variables are set
set -x
cd $FLESH_SYNC_REPO_DIR
git pull origin master:master
export FLESH_COMMIT=`git rev-parse HEAD`
NEW_FLESH_LOC=$FLESH_ROOT_DIR/$FLESH_COMMIT
mkdir -p $NEW_FLESH_LOC
git archive --format=tar $FLESH_COMMIT | (cd $NEW_FLESH_LOC && tar xf -)
cd $NEW_FLESH_LOC
cat goPackages.txt | xargs -t go get -u
# rm $GOPATH/src/flesh
npm install
bower install --allow-root
grunt compile
cd $FLESH_SYNC_REPO_DIR
ln -sfn $NEW_FLESH_LOC $FLESHLOCATION
killall flesh
supervisorctl restart revel
