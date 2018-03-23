# Curl commands
## Get all records, both commands return the same results
`curl -i 192.168.1.2:3000/todos`

`curl -i -X GET 192.168.1.2:3000/todos`

## Create record
`curl -i -X POST -d '{"description":"How to run Windows","complete":false}' 192.168.1.2:3000/todos`

## Update.  If record does not exist, creates new record.
`curl -i -X PUT -d '{"description":"How to run Linux","complete":false}' 192.168.1.2:3000/todos/2`

## Update record, but description field becomes blank
`curl -i -X PUT -d '{"complete":true}' 192.168.1.2:3000/todos/2`

## DELETE record
`curl -i -X DELETE 192.168.1.2:3000/todos/2`
