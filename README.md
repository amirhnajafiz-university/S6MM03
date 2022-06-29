# Hollyworld

This is our final project for **Multimedia Systems** course at _CEIT/AUT_.

## Back-end
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

```

A single movie response:
```json

```

## Python
Make sure the back-end is running with _docker-compose_.

Install requirements:
```shell
pip/pip3 install -r requirements.txt
```

Run the script with movies inside the _videos_ dir:
```shell
python/python3 script.py [Input file name] [Output id]
```

Example:
```shell
python3 script.py 1.mp4 1
```

After that run the following command to export our files to docker container:
```shell
make d-push id=[Output id]
```

## Contributes
- Amirhossein Najafizadeh
- Farshid Nooshi
