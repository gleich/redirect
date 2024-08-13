# redirect

[![build](https://github.com/gleich/redirect/actions/workflows/build.yml/badge.svg)](https://github.com/gleich/redirect/actions/workflows/build.yml)
[![release](https://github.com/gleich/redirect/actions/workflows/release.yml/badge.svg)](https://github.com/gleich/redirect/actions/workflows/release.yml)
[![lint](https://github.com/gleich/redirect/actions/workflows/lint.yml/badge.svg)](https://github.com/gleich/redirect/actions/workflows/lint.yml)

Very basic redirect program

## How to use

Simply run the docker image `ghcr.io/gleich/redirect`. The environment variable `REDIRECT_URL` is set to the URL it should redirect to and `REDIRECT_STATUS_CODE` is the status code it should redirect with (defaults to status code of 307).
