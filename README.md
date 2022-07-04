<div align="center">
    <h1>In The Name Of GOD</h1>    
</div>

<p align="center">
    <img src="assets/logo.png" width="300" />
</p> 

This is our final project for **Multimedia Systems** course at _CEIT/AUT_.

## Application
Start the application:
```console
chmod +x setup.sh 
./setup.sh
```

### Front-end
Start front-end application:
```console
cd front-end
npm start
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

