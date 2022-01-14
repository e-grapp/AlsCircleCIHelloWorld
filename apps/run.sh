set -o xtrace

./aloha
sleep 5
./aloha
sleep 5
./aloha
sleep 5
./aloha
./ksleep 2 1.5
# ./ksshgen -a rsa --override -b 4096
./kpwdgen -l 16 -s 0 -c 1