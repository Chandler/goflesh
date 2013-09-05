export FLESH_ROOT_DIR=/root
export FLESH_SYNC_REPO_DIR=$FLESH_ROOT_DIR/flesh-deploy
export FLESH_DEPLOY_DIR=$FLESH_SYNC_REPO_DIR/deploy
export FLESHLOCATION=$FLESH_ROOT_DIR/flesh
export FLESH_DB_SCHEMA=$FLESH_DEPLOY_DIR/db/schema.sql
export FLESH_DB_BASEDATA=$FLESH_DEPLOY_DIR/db/base_data.sql
export PG_VERSION=9.2
export PG=/usr/lib/postgresql/$PG_VERSION/bin/
export GOPATH=$FLESH_ROOT_DIR/gocode
export PATH="$PATH:$GOPATH/bin"

cd $FLESH_ROOT_DIR

apt-get install -yf python-software-properties

add-apt-repository -y ppa:chris-lea/node.js
add-apt-repository -y ppa:chris-lea/nginx-devel
add-apt-repository -y ppa:nginx/stable
add-apt-repository -y ppa:duh/golang
$FLESH_DEPLOY_DIR/postgres-ppa.sh

apt-get update
apt-get upgrade -y

apt-get install -yf postgresql-$PG_VERSION postgresql-client-$PG_VERSION pgpool2 nginx nginx-common nginx-full git mercurial python-pip zsh nodejs golang make htop

# Node.js dependencies
npm -g install grunt-cli bower

# Python dependencies
pip install supervisor

# stop services and disable autostart
service pgpool2 stop
service postgres stop
service nginx stop
update-rc.d -f postgresql remove
update-rc.d -f pgpool2 remove
update-rc.d -f nginx remove

## Oh-my-zsh!
curl -L https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh | sh

## Clone the repo
git clone -b master https://github.com/Chandler/flesh.git $FLESH_SYNC_REPO_DIR

## Setup postgres and pgpool

cat - >> /etc/postgresql/$PG_VERSION/flesh/postgresql.conf <<EOF
# Settings for flesh running on Digital Ocean
work_mem = 5MB
synchronous_commit = off
wal_buffers = 4MB
shared_buffers = 108MB
effective_cache_size = 256MB
checkpoint_segments = 8
random_page_cost = 1.5
max_connections = 70

EOF
# postgres and pgpool connection information (trust everyone locally)
echo "local all all trust" >> /etc/postgresql/$PG_VERSION/flesh/pg_hba.conf
echo "local all all trust" >> /etc/pgpool2/pg_hba.conf

# increase shared memory limit in kernel
echo "kernel.shmmax = 125829120" >> /etc/sysctl.conf
sysctl -p

# create database cluster
pg_createcluster --start-conf manual $PG_VERSION flesh

# build revel DB tables
createdb -p 5454 -U postgres flesh
psql -p 5454 -U postgres -d flesh < $FLESH_DB_SCHEMA
psql -p 5454 -U postgres -d flesh < $FLESH_DB_BASEDATA

EOF
ln -s $FLESH_DEPLOY_DIR/apps/flesh/git-deploy.conf /etc/git-deploy.conf

# Set up go
mkdir -p $GOPATH

# set up environment
echo "# Environment generated at environment setup time:" | tee -a ~/.bashrc ~/.zshrc
env -u PWD env | xargs -L 1 echo "export" | sed 's/=\(.*\)/="\1"/g' | tee -a ~/.bashrc ~/.zshrc
