docker build . -t gowork:sha-"$(git rev-parse --short=8 HEAD)"
