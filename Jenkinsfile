pipeline {
    agent any

    environment {
        def buildDestination = '/data/app/lelo/lelo-user'
        def workDirectory = '/var/lib/jenkins/workspace'
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
                dir(buildDestination) {
                    sh 'rm -f config.yaml credential.yaml lelo-user'                    
                }
                dir(workDirectory + '/' + env.JOB_BASE_NAME) {
                    sh 'cp lelo-user ${buildDestination}'    
                }
            }
        }
    }
}
