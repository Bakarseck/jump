mkdir $1
cd $1
wget https://raw.githubusercontent.com/01-edu/public/master/subjects/$1/README.md
echo 'package main

func main() {

}' > main.go