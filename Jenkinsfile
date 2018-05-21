pipeline {

    agent {
        docker {
            image 'tenrero/golang-dep-alpine:1.10.2'
            reuseNode true
            args '-it -v /var/run/docker.sock:/var/run/docker.sock -v $WORKSPACE:/go/src/app -w /go/src/app'
        }
    }

    stages {        
        stage('Prepare Environment') {
            steps {
                sh 'echo $GOPATH'
                sh 'rm -fr /go/src/app/vendor && exit 0'
                sh 'go get -u github.com/golang/dep/cmd/dep'
                sh 'go get -u github.com/golang/lint/golint'
                sh 'go get -u github.com/tebeka/go2xunit'
                sh 'go get -u golang.org/x/tools/cmd/cover'
                sh 'go get -u github.com/mattn/goveralls'
            }
        }

        stage('Download Vendor') {
            steps {
                sh 'cd /go/src/app && dep ensure'
            }
        }

        stage('Test') {
            steps {
                sh 'cd /go/src/app && go test ./... -race -coverprofile=coverage.txt -covermode=atomic'
            }
        }

        stage('Build') {
            steps {
                sh 'cd /go/src/app && ./build.sh'
            }
        }
    }
    
}