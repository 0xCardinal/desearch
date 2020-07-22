#!/usr/bin/bash

# Download Research Papers (because information should be free for all)
# A wrapper around sci-hub.tw (credit: Alexandra Elbakyan)
# Author - Kumar Ashwin

if [ -z $1 ]; then
    echo "Specify download link...and try again!"
else
       
    # Fetching URL
    url=`curl https://sci-hub.tw/$1 2> /dev/null | grep download=true | grep -Eo "(http|https)://[a-zA-Z0-9./?=_%:-]*"`
    if [[ -z $url ]]; then 
        url=`curl https://sci-hub.tw/$1 2> /dev/null | grep download=true | grep -Eo "//[a-zA-Z0-9./?=_%:-]*"`
        url=https:${url}
    fi
       
    # Downloading
    echo "Downloading Research Paper..."
    curl -O $url 2> /dev/null

    # Renaming
    renamed=`ls | grep ?download=true | sed "s/?download=true//"`
    to_be_renamed=`ls | grep ?download=true`
    mv $to_be_renamed $renamed

    echo "Done."
fi

# Problem in regex - Site generates URL that cannot be read by the used regex. 
# and other issues xD
