docker build . -t gowork:sha-"$(git rev-parse --short=10 HEAD)"
