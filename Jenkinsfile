pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                git url: 'git@github.com:Antonshepitko/go-service.git', branch: 'master'
            }
        }
        stage('Build') {
            steps {
                sh '''
                    go mod init go-service || true
                    go build -o go-service
                    ls -l go-service
                '''
                stash name: 'go-service-binary', includes: 'go-service'
            }
        }
        stage('Deploy') {
            steps {
                unstash 'go-service-binary'
                sh '''
                    echo "Current directory: $(pwd)"
                    ls -l go-service
                '''
                sshagent(credentials: ['jenkins-ssh-credentials']) {
                    sh '''
                        echo "Deploying to 45.144.52.219"
                        scp go-service deployer@45.144.52.219:/app/go-service
                        ssh deployer@45.144.52.219 "echo 'Making executable'; chmod +x /app/go-service && echo 'chmod succeeded' || echo 'chmod failed'"
                        ssh deployer@45.144.52.219 "echo 'Stopping old process'; pkill -f go-service || echo 'No process to stop'"
                        ssh deployer@45.144.52.219 "echo 'Starting new process'; nohup /app/go-service > /app/nohup.out 2>&1 & echo 'Started'"
                        echo "Deploy completed"
                    '''
                }
            }
        }
    }
    post {
        always {
            echo "Cleaning up workspace"
            cleanWs()
        }
    }
}