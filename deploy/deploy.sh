# Assumes environment variables are set

set -x
cd $FLESH_SYNC_REPO_DIR
git pull origin master:master
export FLESH_COMMIT=`git rev-parse HEAD`
mkdir -p $FLESH_ROOT_DIR/$FLESH_COMMIT
git archive --format=tar $FLESH_COMMIT | (cd $FLESH_ROOT_DIR/$FLESH_COMMIT && tar xf -)
ln -sfn $FLESH_ROOT_DIR/$FLESH_COMMIT $FLESHLOCATION
cd $FLESHLOCATION
cat goPackages.txt | xargs -t go get -u
# rm $GOPATH/src/flesh
npm install
bower install --allow-root
grunt compile
cd $FLESH_SYNC_REPO_DIR
supervisorctl restart revel
