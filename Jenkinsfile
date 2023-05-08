pipeline {
    agent any
    def project = env.JOB_NAME
    def workspace = pwd()

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
                echo project
                echo workspace
            }
        }
    }
}
