while true
do
    curl -sS 'localhost:8088/api' | jq '.Version'
done