#!/usr/bin/env bash
read -p "Challenge name: " challenge_name
read -p "Language: " language

if [[ $challenge_name == "" ]]; then
    echo "Please supply challenge name"
    exit 1
elif [[ $language == "" || ! -d $language ]]; then
    echo "Directory for this language does not exist $2"
    exit 1
fi

pushd $language
pwd
cat ../skeleton"$language".txt > "$challenge_name"."$language"
cat ../skeleton"$language"-test.txt > "$challenge_name"_test."$language"
popd


