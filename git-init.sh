#/bin/bash

email=barsaliou1998@gmail.com
username=serignmbaye
repo=https://learn.zone01dakar.sn/git/serignmbaye/piscine-go.git
gitConfigDir=./.gitconfig/
if [ -d -a "$gitConfigDir" ]
then
    echo "Removing git config directory : $gitConfigDir..."
	rm -rf $gitConfigDir
fi
gitCredentialDir=./.git-credentials
if [ -d -a "$gitCredentialDir" ]
then
    echo "Removing git credential directory : $gitCredentialDir..."
	rm -rf $gitCredentialDir
fi
repoDir=./piscine-go
if [ -d -a "$gitCredentialDir" ]
then
    echo "Removing repository directory : $repoDir..."
	rm -rf $repoDir
fi
git config --global credential.helper store
git config --global user.email "$email"
git config --global user.name "$username"
git config --list
git clone "$repo"
