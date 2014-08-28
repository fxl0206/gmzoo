cd ../../
GOPATH=`pwd`
export GOPATH
cd $GOPATH
git pull
cp -R $GOPATH/src/sbin ./
rm -rf $GOPATH/sbin/install.sh