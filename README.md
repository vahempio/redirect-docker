# Redirect

Tiny Go app to redirect the entry point to a chosen URL.

By default, no port are routed to the host. Only the :80 is exposed on the docker.

Usage example:
```sh
$ docker run --rm -p 8080:80 -e URL="http://foo.bar.com" vahempio/redirect
```
And open <http://localhost:8080/> to see the redirection to foo.bar.com.

Github source code: <https://github.com/vahempio/redirect-docker/>

DockerHub automated build: <https://hub.docker.com/r/vahempio/redirect/>
