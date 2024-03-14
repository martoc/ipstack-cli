#!/usr/bin/env ./bats/bin/bats

load 'test_helper/bats-support/load'
load 'test_helper/bats-assert/load'
load 'common.sh'

@test "Get Coordinates" {
  run $BINARY_PATH get-coordinates --ip 134.201.250.155 --access-key $ACCESS_KEY
  assert_success
  assert_equal $(echo $output | jq -r .type ) "ipv4"
}

