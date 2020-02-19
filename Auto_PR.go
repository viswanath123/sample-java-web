function create_pr() {
    if [[ -z $1 || -z $2 || -z $3 ]]; then
      echo "Please provide a message and/or branch"
      exit 1
    fi
    repo=$1
    branch=$2
    message=$3
    # Add all the locally
    git add .
    
    # This will create the new commit to be merged
    git commit -m "$message"
    # This creates the new local branch
    git checkout -b pr/$branch
    # This pushes the local branch to the remote
    git push -u origin pr/$branch
    curl -XPOST -d '{"title":"${message}", "base":"master", "head":"pr/${branch}", "body":"${message}"}' \
https://api.github.com/repos/${repo}/pulls
    exit 0
}
