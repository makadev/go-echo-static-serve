#!/bin/bash

set -ex

/usr/local/bin/check-config -dump

exec /usr/local/bin/_server
