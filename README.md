# Hollyworld

This is our final project for **Multimedia Systems** course at _CEIT/AUT_.

<p align="center">
    <img src="assets/logo.png" width="300" />
</p> 

## Application
Start the application:
```shell
chmod +x setup.sh 
./setup.sh
```

### Endpoints
- :8080/api/movies
  - Get top 6 movies
- :8080/api/movies/{id}
  - Get a single movie
- :8080/dash/{id}
  - Get a movie Dash File
- :8080/poster/{id}
  - Get a movie poster

### Responses
All movies response:
```json
[
  {
    "id": 1,
    "title": "title",
  },
]
```

A single movie response:
```json
{
  "id": 1,
  "title": "title",
  "director": "director",
  "score": 10,
  "description": "something",
  "type": "DASH1/DASH2/HLS",
  "link": "https://link"
}
```

## Contributes
- Amirhossein Najafizadeh
- Farshid Nooshi

