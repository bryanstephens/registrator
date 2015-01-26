FROM progrium/busybox
MAINTAINER Jeff Lindsay <progrium@gmail.com>

ADD ./registrator /bin/

ENV DOCKER_HOST unix:///tmp/docker.sock
ENTRYPOINT ["/bin/registrator"]
