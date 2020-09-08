docker build -t kasumi_dx_go .
docker run --rm -d -p 9000:9000 --link mysql_lara kasumi_dx_go