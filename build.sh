SHA="$(git rev-parse --short=10 HEAD)"
docker build . -t gowork:sha-"$SHA"
docker build . --build-arg SLIM=true -t gowork:sha-"$SHA"-slim
