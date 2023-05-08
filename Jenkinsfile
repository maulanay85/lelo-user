pipeline {
    agent any

    environment {
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
                echo pwd()
                sh 'cd ${buildDestination}'
                sh 'rm -f config.yaml credential.yaml lelo-user'
                sh 'cd ${env.WORKSPACE}'
                sh 'echo ${env.WORKSPACE}'
                sh 'echo ${env.JOB_NAME}'
                sh 'cp config.yaml credential.yaml lelo-user ${buildDestination}'
            }
        }
    }
}
