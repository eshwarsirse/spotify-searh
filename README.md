# spotify-search

A wrapper around the Spotify Search API to search Artists, Albums, Playlists based on given query.

The application takes care of renewing the token if its expired.
###Build app
go build -o application .

###Run app
./application

##curl command
curl --request GET 'localhost:8080/search?query=drake&type=artist,album'

##Docker
docker build -t spotify-search .

docker run --network=host  spotify-search