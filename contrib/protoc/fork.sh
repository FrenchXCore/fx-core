#!/usr/bin/env bash

set -eo pipefail
#set -x

VERSION=$(buf --version)
echo "buf version: $VERSION"

if [ -z "$BUF_NAME" ]; then
  echo "buf name not found, please set BUF_NAME"
  exit 1
fi

if [ -z "$BUF_TOKEN" ]; then
  echo "buf token not found, please set BUF_TOKEN"
  exit 1
fi

if [ -z "$BUF_ORG" ]; then
  echo "buf org not found, please set BUF_ORG"
  exit 1
fi

echo "buf registry login $BUF_NAME with ******"
echo "$BUF_TOKEN" | buf registry login --username "$BUF_NAME" --token-stdin

#read -rp "Are you sure $BUF_NAME want to push proto to $BUF_ORG? [y/N] " input
#if [[ "$input" != "y" && "$input" != "Y" ]]; then
#  exit 1
#fi

echo "USER $BUF_NAME push proto to $BUF_ORG ..."

if [ ! -d build/fork ]; then
  mkdir -p build/fork
fi

# download cosmos-sdk proto
commit_hash=$(go list -m -f '{{.Version}}' github.com/cosmos/cosmos-sdk)
if [ ! -f "./build/fork/cosmos-sdk-proto.zip" ]; then
  echo "download cosmos-sdk $commit_hash"
  wget -c "https://github.com/cosmos/cosmos-sdk/archive/$commit_hash.zip" -O "./build/fork/cosmos-sdk-proto.zip"
fi

(
  cd build/fork
  rm -rf cosmos-sdk
  unzip -q -o "./cosmos-sdk-proto.zip"
  # shellcheck disable=SC2010
  mv "$(ls | grep cosmos-sdk | grep -v grep | grep -v zip)" cosmos-sdk
  rm -rf cosmos-sdk/.git

  # buf push
  cd cosmos-sdk/proto
  # replace buf.yaml buf.build/cosmos/cosmos-sdk => buf.build/functionx/cosmos-sdk
  sed -i 's/buf.build\/cosmos\/cosmos-sdk/buf.build\/'"$BUF_ORG"'\/cosmos-sdk/g' buf.yaml
  #sed -i '.bak' 's/buf.build\/cosmos\/cosmos-sdk/buf.build\/'"$BUF_ORG"'\/cosmos-sdk/g' buf.yaml
  #rm buf.yaml.bak

  echo "buf push cosmos-sdk proto with tag $commit_hash ..."
  buf push --tag "$commit_hash"
)

# download ibc-go proto
commit_hash=$(go list -m -f '{{.Version}}' github.com/cosmos/ibc-go/v6)
if [ ! -f "./build/fork/ibc-go-proto.zip" ]; then
  echo "download ibc-go $commit_hash"
  wget -c "https://github.com/cosmos/ibc-go/archive/$commit_hash.zip" -O "./build/fork/ibc-go-proto.zip"
fi
(
  cd build/fork
  rm -rf ibc-go
  unzip -q -o "./ibc-go-proto.zip"
  # shellcheck disable=SC2010
  mv "$(ls | grep ibc-go | grep -v grep | grep -v zip)" ibc-go
  rm -rf ibc-go/.git

  # buf push TODO v6.1.0 add name and deps
  append="version: v1\nname: buf.build/$BUF_ORG/ibc\ndeps:\n  - buf.build/cosmos/cosmos-sdk:8cb30a2c4de74dc9bd8d260b1e75e176\n  - buf.build/cosmos/cosmos-proto:1935555c206d4afb9e94615dfd0fad31\n  - buf.build/cosmos/gogo-proto:bee5511075b7499da6178d9e4aaa628b\n  - buf.build/googleapis/googleapis:783e4b5374fa488ab068d08af9658438"

  cd ibc-go
  cp third_party/proto/proofs.proto proto/proofs.proto
  cd proto
  sed -i 's|version: v1|'"${append}"'|g' buf.yaml
  #sed -i '.bak' 's|version: v1|'"${append}"'|g' buf.yaml
  #rm buf.yaml.bak

  buf mod update
  echo "buf push ibc-go proto with tag $commit_hash ..."
  buf push --tag "$commit_hash"
)
