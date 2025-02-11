# music-librarycd backend

## Compile

__Compile backend__:
```
docker build -f backend.Dockerfile -t music-library-backend ./backend
```

__Compile frontend__:
```
docker build -f frontend.Dockerfile -t music-library-frontend ./frontend
```

## Launch

```
docker-compose up --build
```
