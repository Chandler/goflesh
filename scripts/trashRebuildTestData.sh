echo "!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"
echo "Dropping database, ctrl+c to cancel"
echo "!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"
echo "Be sure to run this from the root, not /scripts"
set -x
sleep 3
echo "DROP DATABASE flesh_local;" | psql -p 5454 -U postgres
createdb -p 5454 -O postgres -U postgres flesh_local
psql -p 5454 -U postgres -d flesh_local < $FLESHLOCATION/db/schema.sql
(revel run flesh dev &) && sleep 2 && curl http://localhost:9000/@tests/Generator/TestGenerateData  && killall revel && killall flesh
