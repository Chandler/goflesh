echo "!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"
echo "Dropping database, ctrl+c to cancel"
echo "!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"
if [[ -z "$FLESHLOCATION" ]]
    then
    FLESHLOCATION=`pwd`
fi
set -x
SCHEMA=$FLESHLOCATION/db/schema.sql
BASEDATA=$FLESHLOCATION/db/base_data.sql
if [[ -f "$SCHEMA" ]]
    then
    sleep 3
    psql -p 5454 -U postgres -c "DROP DATABASE flesh_local;"
    createdb -p 5454 -O postgres -U postgres flesh_local
    psql -p 5454 -U postgres -d flesh_local < $SCHEMA
    psql -p 5454 -U postgres -d flesh_local < $BASEDATA
    (revel run flesh dev &) && sleep 2 && curl http://localhost:9000/@tests/Generator/TestGenerateData && killall revel && killall flesh
fi
