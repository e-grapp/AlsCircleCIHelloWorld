set -o xtrace

./aloha
sleep 5s
./aloha
sleep 5s
./aloha
sleep 5s
./aloha
./ksleep 2 1.5
# ./ksshgen -a rsa --override -b 4096
./kpwdgen -l 16 -s 0 -c 1