set -o xtrace

go run ./aloha
go run ./aloha
go run ./aloha
go run ./aloha
./aloha
./aloha
./aloha
./ksleep 2 1.5
# ./ksshgen -a rsa --override -b 4096
./kpwdgen -l 16 -s 0 -c 1