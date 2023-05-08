pipeline {
    agent any

    environment {
        def workspace = env.WORKSPACE
        def jobName = env.JOB_NAME
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
                sh 'cd ${workspace}/${jobName}'
                sh 'echo ${workspace}'
                sh 'echo ${jobName}'
                sh 'cp config.yaml credential.yaml lelo-user ${buildDestination}'
            }
        }
    }
}
