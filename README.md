# flesh

## Requirements
go - http://golang.org/
revel - http://robfig.github.io/revel/


## Setup
    #node
      brew install npm
      npm -g install grunt-cli
      npm -g install jamjs #client package manager
      npm install

    #db
      brew install postgres

    #revel
      brew install go --devel
      #install revel as per site instructions
      initdb .db -U postgres
      postgres -D .db -r /tmp/flesh-postgres.log -p 5454 & # start the server, hit enter twice
      createdb -p 5454 -O postgres -U postgres flesh_local
      ln -s ~/flesh $GOPATH/src/flesh # assuming you cloned to ~/flesh
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
