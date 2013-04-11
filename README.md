# flesh

## Requirements
go - http://golang.org/
revel - http://robfig.github.io/revel/
postgres

## Setup
    brew install postgres
    brew install go --devel
    # install revel as per site instructions
    initdb .db -U postgres
    postgres -D .db -r /tmp/flesh-postgres.log -p 5454 & # start the server, hit enter twice
    createdb -p 5454 -O postgres -U postgres flesh_local
    ln -s ~/flesh/flesh $GOPATH/src/flesh # assuming you cloned to ~/flesh
    revel run flesh # if you added $GOPATH/bin to your path as per revel install instructions

