pipeline {
    agent any

    stages {
        stage('Checkout') {      
            steps {
                checkout scm   
            }
        }
        
        stage('Static Code Analysiz') {
            steps {
                echo 'static analysis...'
            }
        }
        
        stage('Test') {
            steps {
                echo 'testler calisti'
            }
        }
        
        stage('Build') {
            steps {
                sh 'docker build -t first-build .'
            }
        }
        
        stage('Deploy') {
            steps {
                echo 'deploying...'
            }
        }
    }
}

