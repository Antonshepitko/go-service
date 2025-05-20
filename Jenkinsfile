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
                 }
             }
             stage('Deploy') {
                 steps {
                     sshagent(credentials: ['jenkins-ssh-credentials']) {
                         sh '''
                             echo "Deploying to $DEPLOY_SERVER"
                             scp go-service deployer@45.144.52.219:/app/go-service
                             ssh deployer@45.144.52.219 "chmod +x /app/go-service"
                             ssh deployer@45.144.52.219 "pkill -f go-service || true"
                             ssh deployer@45.144.52.219 "nohup /app/go-service > /app/nohup.out 2>&1 & echo 'Started'"
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