pipeline {
    agent any

    environment {
        def buildDestination = '/data/app/lelo/lelo-user'
        def workspace = '/var/lib/jenkins'
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
                sh 'cd ${workspace}/${env.JOB_BASE_NAME}'
                sh 'cp config.yaml credential.yaml lelo-user ${buildDestination}'
            }
        }
    }
}
