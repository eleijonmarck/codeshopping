#!/usr/bin/env groovy

node {
    checkout scm
        stage('Build') {
                echo 'Building......'
                echo 'wtf.....'
        }
        stage('Test') {
            echo "Running make test"
            runMake('test')
            echo "ran the tests"
        }

        stage('Deploy') {
                echo 'Deploying....'
        }
}

def runMake(command){
    node {
        sh 'make ${command}'
}
