pipeline {

    agent any

    stages {
        stage('Prepare Environment') {
            steps {
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