mkdir -p ~/lista/
cd ~/lista/
git clone --recursive  git@github.com:$3.git lista-$1
cd lista-$1
git checkout $2
git reset --hard $1
sh deploy/start.sh
echo "DEPLOYED!"
