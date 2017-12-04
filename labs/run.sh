while true
do
    curl -sS 'localhost:8081/api' | jq '.Version'
done