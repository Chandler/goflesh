# flesh

## Requirements
go - http://golang.org/
revel - http://robfig.github.io/revel/


## Setup
    #node
      brew install npm
      npm -g install grunt-cli
      npm -g install jamjs
      npm install
      jam install

      (temporary)
      from root:
      git clone https://github.com/Chandler/grunt-ember-handlebars.git
      npm install grunt-ember-handlebars
      
    #db
      brew install postgres

    #revel
      brew install go --devel
      #install revel as per site instructions
      # then, install go packages
      cat goPackages.txt | xargs -t go get -u
      initdb .db -U postgres
      postgres -D .db -r /tmp/flesh-postgres.log -p 5454 & # start the server, hit enter twice
      createdb -p 5454 -O postgres -U postgres flesh_local
      ln -s ~/flesh $GOPATH/src/flesh # assuming you cloned to ~/flesh

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


## Heroku

We have a staging server at `flesh.herokuapp.com`. The config should look like this

    BUILDPACK_URL=https://github.com/robfig/heroku-buildpack-go-revel
    GOPATH=/app/.go
    GOROOT=/app/.goroot
    PATH=bin:node_modules/.bin:/usr/local/bin:/usr/bin:/bin:/app/.goroot/bin

##Useful postgres commands
Make postgres fast(er).

Do NOT use these settings on a server, only locally!

    echo "shared_buffers = 9MB"     >> .db/postgresql.conf
    echo "work_mem = 50MB"          >> .db/postgresql.conf
    echo "fsync = off"              >> .db/postgresql.conf
    echo "synchronous_commit = off" >> .db/postgresql.conf
    pg_ctl -D .db restart

connect with psql

    psql -p 5454 -U postgres -d flesh_local
