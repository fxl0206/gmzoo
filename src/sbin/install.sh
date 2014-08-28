GOPATH=../../
export GOPATH
cd $GOPATH
git pull
cp -R $GOPATH/src/sbin ./
./sbin/build.sh