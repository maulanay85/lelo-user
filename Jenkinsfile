pipeline {
    agent any

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
                echo pwd
            }
            // sh 'cd /data/app/lelo/lelo-user'
            // sh 'rm config.yaml credential.yaml lelo-user'
            // sh 'cp /'
        }
    }
}
