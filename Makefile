all:
	go generate ./...
	go build
check_go_style:
	bash -c "mkdir -p checkstyle; cd checkstyle && export GOPATH=`pwd` && go get github.com/qiniu/checkstyle/gocheckstyle"
	checkstyle/bin/gocheckstyle -config=.go_style ./...
