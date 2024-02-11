Example RestAPI endpoints for retrieving / sending a JSON structure.

# START:
    

# POST:

curl http://localhost:8080/books \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"title": "book title","author": "author name","release_date": "00.00.00","pages": 1}'


# GET:
    curl http:\\localhost:8080