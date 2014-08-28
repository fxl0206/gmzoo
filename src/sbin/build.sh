cd ../
GOPATH=`pwd`
export GOPATH
cd $GOPATH
git pull
cp -R $GOPATH/src/views ./
cp -R $GOPATH/src/conf ./
go build -o .gmzoo $GOPATH/src/main.go