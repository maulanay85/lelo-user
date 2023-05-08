pipeline {
    agent any

    environment {
        def workspace = pwd()
        def buildDestination = '/data/app/lelo/lelo-user/'
    }
    tools {
        go '1.19'
    }

    stages {
        stage ('Test') {
            steps {
                echo 'Test Skipped'
            }
        }
        stage('Build') {
            steps {
                sh 'GOOS=linux GOARCH=amd64 go build .'
            }
        }
        stage('Deploy') {
            steps {
                sh 'cd ${buildDestination}'
                sh 'rm -f config.yaml credential.yaml lelo-user'
                sh 'cd ${workspace}'
                sh 'cp config.yaml credential.yaml lelo-user ${buildDestination}'
            }
        }
    }
}
