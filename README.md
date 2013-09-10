# flesh


## Technologies

* go - http://golang.org/
* revel - http://robfig.github.io/revel/
* ember - http://emberjs.com/
* ember data - https://github.com/emberjs/data


## Setup

### Environment variables

Configuration that shouldn't be committed (e.g. passwords) go in environment variables.

    export FLESH_EMAIL_OVERRIDE=youemailaddress@gmail.com # in development mode, send all emails to this address
    export FLESH_MANDRILL_KEY= # this has to be generated per-developer

### From the shell

    # you probably haven't updated brew in a while
      brew update

    # get the repo
      export FLESHLOCATION="~/flesh"
      git clone git@github.com:Chandler/flesh.git $FLESHLOCATION
      cd $FLESHLOCATION

    # node.js
      brew install npm
      npm -g install grunt-cli
      npm install
      bower install
      
    # database
      brew install postgres
      initdb .db -U postgres
      # Make postgres fast(er).
      # Do NOT use these settings on a server, only locally!
      echo "shared_buffers = 9MB"     >> .db/postgresql.conf
      echo "work_mem = 50MB"          >> .db/postgresql.conf
      echo "fsync = off"              >> .db/postgresql.conf
      echo "synchronous_commit = off" >> .db/postgresql.conf
      # Start postgres and create DB
      sudo sysctl -w kern.sysv.shmall=1073741824
      sudo sysctl -w kern.sysv.shmmax=67108864
      pg_ctl start -D .db -l /tmp/flesh-postgres.log -o "-p 5454"
      createdb -p 5454 -O postgres -U postgres flesh_local
      psql -p 5454 -U postgres -d flesh_local < $FLESHLOCATION/db/schema.sql

    # database connection pooling
      brew install pgpool-ii
      pgpool -f conf/pgpool.conf

    # revel
      mkdir ~/gocode
      export GOPATH=~/gocode # you should add this line to your .bash_profile too
      export PATH="$PATH:$GOPATH/bin" # you should add this line to your .bash_profile too
      brew install go
      brew upgrade go
      brew install mercurial
      cat goPackages.txt | xargs -t go get -u
      ln -s $FLESHLOCATION $GOPATH/src/flesh

    # Generate test data
      ./scripts/trashRebuildTestData.sh

## Running Locally
    #assets (check gruntfile.js for all the availiable tasks)
      grunt compile #build assets once
      grunt w #build assets once and then watch for file changes (3rd party libraries not watched)

    #server
      revel run flesh


## Other
    #useful syntax highlighting
      cd ~/Library/Application\ Support/Sublime\ Text\ 2/Packages #something close to this
      git clone git://github.com/jashkenas/coffee-script-tmbundle CoffeeScript
      git clone https://github.com/LearnBoost/stylus.git Stylus
    
    #mocking a tag via twilio
      curl -X POST --data "Body=5V4MR&From=12089912446&AccountSid=$TWILIO_ACCOUNT_SID" http://localhost:9000/api/sms


## Useful postgres commands
Make postgres fast(er).

Do NOT use these settings on a server, only locally!

    echo "shared_buffers = 9MB"     >> .db/postgresql.conf
    echo "work_mem = 50MB"          >> .db/postgresql.conf
    echo "fsync = off"              >> .db/postgresql.conf
    echo "synchronous_commit = off" >> .db/postgresql.conf
    pg_ctl -D .db restart

start/stop/restart

    pg_ctl start -D .db -o "-p 5454"
    pg_ctl stop -D .db -o "-p 5454"
    pg_ctl restart -D .db -o "-p 5454"
    pgpool -f conf/pgpool.conf

connect with psql

    psql -p 5454 -U postgres -d flesh_local

# Generating test data

drop/update schema/generate test data cycle:

     # from flesh root (not /scripts!)
     ./scripts/trashRebuildTestData.sh

generate test data:

    # visit localhost:9000/@tests
    # and click Generator > TestGenerateData > Run
    # This will create a bunch of dummy data for you

