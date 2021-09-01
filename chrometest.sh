for i in {1..20}
do; echo ${i}
curl -I -sS https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb -o /dev/null --fail; done