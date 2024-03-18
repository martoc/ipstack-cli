if [ -z "$TAG_VERSION" ]; then
  export TAG_VERSION="1.0.0"
fi

BINARY_PATH="./target/builds/ipstack-cli-$(go env GOOS)-$(go env GOARCH)"
