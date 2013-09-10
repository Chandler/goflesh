#!/bin/bash
pg_dump -U postgres -p 5455 -O flesh | gzip --best > /tmp/backups/backup-`date +"%s"`.sql.gz