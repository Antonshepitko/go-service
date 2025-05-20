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
                        echo "Preparing deployment to 45.144.52.219"
                        ssh deployer@45.144.52.219 "rm -f /app/go-service || echo 'No file to remove'"
                        echo "Copying new binary"
                        scp go-service deployer@45.144.52.219:/app/go-service
                        ssh deployer@45.144.52.219 "echo 'Making executable'; chmod +x /app/go-service && echo 'chmod succeeded' || echo 'chmod failed'"
                        ssh deployer@45.144.52.219 "echo 'Restarting service'; sudo systemctl restart go-service && echo 'Service restarted' || echo 'Failed to restart service'"
                        ssh deployer@45.144.52.219 "echo 'Verifying service'; sudo systemctl status go-service || echo 'Service not running'"
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