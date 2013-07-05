# flesh


## Technologies

* go - http://golang.org/
* revel - http://robfig.github.io/revel/
* ember - http://emberjs.com/
* ember data - https://github.com/emberjs/data


## Setup
    # get the repo
      export FLESHLOCATION="~/flesh"
      git clone git@github.com:Chandler/flesh.git $FLESHLOCATION
      cd $FLESHLOCATION

    # node.js
      brew install npm
      npm -g install grunt-cli
      npm -g install jamjs
      npm install
      jam install

      # (temporary)
      # from $FLESHLOCATION:
      git clone https://github.com/Chandler/grunt-ember-handlebars.git
      npm install grunt-ember-handlebars
      
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
      pg_ctl start -D .db -l /tmp/flesh-postgres.log -o "-p 5454"
      createdb -p 5454 -O postgres -U postgres flesh_local
      psql -p 5454 -U postgres -d flesh_local < $FLESHLOCATION/db/schema.sql

    # revel
      brew install go
      cat goPackages.txt | xargs -t go get -u
      ln -s $FLESHLOCATION $GOPATH/src/flesh

    # Generate test data
      # visit localhost:9000/@tests
      # and click Generator > TestGenerateData > Run
      # This will create a bunch of dummy data for you

    # JS tests
      rpm -g install jasmine-node
      # run tests using jasmine-node spec/ --coffee

## Running Locally
    #assets (check gruntfile.js for all the availiable tasks)
      grunt compile #build assets once
      grunt w #build assets once and then watch for file changes

    #server
      revel run flesh # if you added $GOPATH/bin to your path as per revel install instructions


## Other
    #useful syntax highlighting
      cd ~/Library/Application\ Support/Sublime\ Text\ 2/Packages #something close to this
      git clone git://github.com/jashkenas/coffee-script-tmbundle CoffeeScript
      git clone https://github.com/LearnBoost/stylus.git Stylus


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

connect with psql

    psql -p 5454 -U postgres -d flesh_local

drop/update schema/generate test data cycle:

    echo "DROP DATABASE flesh_local;" | psql -p 5454 -U postgres
    createdb -p 5454 -O postgres -U postgres flesh_local
    psql -p 5454 -U postgres -d flesh_local < $FLESHLOCATION/db/schema.sql
    # visit localhost:9000/@tests
    # and click Generator > TestGenerateData > Run
    # This will create a bunch of dummy data for you
