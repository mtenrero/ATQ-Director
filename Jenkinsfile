pipeline {

    agent any

    stages {
        stage('Cleanup') {
            steps {
                sh 'rm -fr $GOPATH/src/github.com/mtenrero/ATQ-Director'
            }
        }
        
        stage('Prepare Environment') {
            steps {
                sh 'echo $GOPATH'
                sh 'mkdir -p $GOPATH/src/github.com/mtenrero/'
                sh 'ln -s "$WORKSPACE" "$GOPATH/src/github.com/mtenrero/ATQ-Director"'
                sh 'ls -a $GOPATH/src/github.com/mtenrero/ATQ-Director'
                sh 'go get -u github.com/golang/dep/cmd/dep'
                sh 'go get -u github.com/golang/lint/golint'
                sh 'go get -u github.com/tebeka/go2xunit'
                sh 'go get -u golang.org/x/tools/cmd/cover'
                sh 'go get -u github.com/mattn/goveralls'
            }
        }

        stage('Download Vendor') {
            steps {
                sh 'cd $GOPATH/src/github.com/mtenrero/ATQ-Director && dep ensure'
            }
        }

        stage('Test') {
            steps {
                sh 'cd $GOPATH/src/github.com/mtenrero/ATQ-Director && go test ./... -coverprofile=coverage.txt -covermode=atomic'
            }
        }

        stage('Code Coverage') {
            environment { 
                COVERTOKEN = credentials('coveralls-token') 
            }
            
            steps {
                sh 'cd $GOPATH/src/github.com/mtenrero/ATQ-Director && goveralls -coverprofile=coverage.txt -repotoken $COVERTOKEN'
            }
        }

        stage('Build') {
            when { expression { env.BRANCH_NAME == 'master' } }
            steps {
                sh 'cd $GOPATH/src/github.com/mtenrero/ATQ-Director && ./build.sh'
            }
        }
    }
    
}