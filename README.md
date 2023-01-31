<h1 align="center">
    S6MM03
</h1>    

<br />

This is our final project for **Multimedia Systems** course at _CEIT/AUT_.
This Project was for building an IMDB website clone using **Docker, Dash Protocol for streaming online and chunking data, Angular Framework for Frontend, SQLite for Database and GoLang for Backend**. For Chunking Data and streaming online we used Dash Protocol and Python FFmpeg Video Streaming Library.

#### Approaches for online streaming

* **plain VoD** for streaming videos 
* implementation of our own backend server for **Dash protocol**
* using **HLS protocol and VoD**


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

- [Amirhossein Najafizadeh](https://amirhnajafiz.github.io/)
- [Farshid Nooshi](https://farshidnooshi.github.io)

