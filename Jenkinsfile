pipeline {

    agent any

    stages {
        stage('Prepare Environment') {
            steps {
                sh 'mkdir -p $GOPATH/src/github.com/mtenrero/ATQ-Director'
                sh 'ln -s $WORKSPACE $GOPATH/src/github.com/mtenrero/ATQ-Director'
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
            steps {
                sh './build.sh'
            }
        }
    }
    
}