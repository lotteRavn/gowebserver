[![Stories in Ready](https://badge.waffle.io/lotteRavn/gowebserver.png?label=ready&title=Ready)](https://waffle.io/lotteRavn/gowebserver)
[![Stories in Ready](https://badge.waffle.io/skbogner/gowebserver.png?label=ready&title=Ready)](https://waffle.io/skbogner/gowebserver)
[![Stories in Ready](https://badge.waffle.io/sthordall/gowebserver.png?label=ready&title=Ready)](https://waffle.io/sthordall/gowebserver)
[![Stories in Ready](https://badge.waffle.io/sthordall/gowebserver.png?label=ready&title=Ready)](https://waffle.io/sthordall/gowebserver)
[![Stories in Ready](https://badge.waffle.io/gowebserver/gowebserver.png?label=ready&title=Ready)](https://waffle.io/gowebserver/gowebserver)
[![Stories in Ready](https://badge.waffle.io/Brax94/gowebserver.png?label=ready&title=Ready)](https://waffle.io/Brax94/gowebserver)
[![Stories in Ready](https://badge.waffle.io/Brax94/gowebserver.png?label=ready&title=Ready)](https://waffle.io/Brax94/gowebserver)
[![Stories in Ready](https://badge.waffle.io/Brax94/gowebserver.png?label=ready&title=Ready)](https://waffle.io/Brax94/gowebserver)
[![Stories in Ready](https://badge.waffle.io/Brax94/gowebserver.png?label=ready&title=Ready)](https://waffle.io/Brax94/gowebserver)
[![Stories in Ready](https://badge.waffle.io/Brax94/gowebserver.png?label=ready&title=Ready)](https://waffle.io/Brax94/gowebserver)
Run the follwing command to run the simple web server:
$ go run http.go  

Open browser and goto http://localhost:8000

Replace localhost with the name or IP of your DockerHost.

## Run with docker


Build the image:

    docker build -t myapp .

Test that you can run it:

    docker run -d -p 8000:8000  --name myapp myapp:latest
    curl $(docker-machine ip code):8000


## Test Driven Development

Here's how to run the test suite:

    docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.6 go test -v
