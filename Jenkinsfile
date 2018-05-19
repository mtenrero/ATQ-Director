pipeline {

    stages {
        stage('Prepare Environment') {
            steps {
                sh 'curl -L -s https://github.com/golang/dep/releases/download/v0.3.1/dep-linux-amd64 -o $GOPATH/bin/dep'
                sh 'chmod +x $GOPATH/bin/dep'
                sh 'go get golang.org/x/tools/cmd/cover'
                sh 'go get github.com/mattn/goveralls'
            }
        }

        stage('Download Vendor') {
            steps {
                sh 'dep ensure'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./... -race -coverprofile=coverage.txt -covermode=atomic'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -v ./...'
            }
        }
    }
    
}