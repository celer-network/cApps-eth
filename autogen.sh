#!/usr/bin/env bash
# Add a commit with generated abi/bin files to the branch
# This script should only be run by build bot for abi/bin consistency
# email/name are just placeholder

PRID="$1"
BRANCH="$2"

TRUFFLE_VER=`node_modules/.bin/truffle version`

setup_git() {
  echo "setup git"
  git config --global user.email "build@celer.network"
  git config --global user.name "Build Bot"
}

get_pr() {
  echo "get pull request"
  git fetch origin pull/$PRID/head:$BRANCH
  git checkout $BRANCH
}

# $1 folder, $2 sol filename, $3 package name
generate_files() {
  echo "generate binding files"
  jq .abi build/contracts/$2.json > genfiles/$1/$2.abi
  jq -r .bytecode build/contracts/$2.json > genfiles/$1/$2.bin
  mkdir -p genfiles/$1/go/$3
  .bin/abigen -abi genfiles/$1/$2.abi -bin genfiles/$1/$2.bin -pkg $3 -type $2 -out genfiles/$1/go/$3/$3.go
  # web3j/bin/web3j solidity generate -a=genfiles/$1/$2.abi -b=genfiles/$1/$2.bin -o=genfiles/$1/java -p=$3
}

update_genfiles() {
  echo "update binding files"

  mkdir -p genfiles/simple-app/java
  mkdir -p genfiles/simple-app/go
  generate_files simple-app SimpleSingleSessionApp singlesessionapp
  generate_files simple-app SimpleMultiSessionApp multisessionapp

  mkdir -p genfiles/gomoku/java
  mkdir -p genfiles/gomoku/go
  generate_files gomoku SingleGomoku singlegomoku
  generate_files gomoku MultiGomoku multigomoku

  mkdir -p genfiles/interface/java
  mkdir -p genfiles/interface/go
  generate_files interface ISingleSession isinglesession
  generate_files interface IMultiSession imultisession
  generate_files interface IBooleanResult ibooleanresult
  generate_files interface INumericResult inumericresult
}

# append a new commit with generated files to current PR
commit_and_push() {
  echo "commit and push"
  git add genfiles/
  git commit -m "Update genfiles by ci" -m "$TRUFFLE_VER"
  # gh_token is an env in ci project setting
  git push https://${GH_TOKEN}@github.com/celer-network/cApps.git $BRANCH &> /dev/null
}

setup_git
get_pr
update_genfiles
if [[ `git status --porcelain` ]]; then
  commit_and_push
fi
