pipeline {
    agent any
    // tools {
    //     go 'go1.21.6'
    // }
        
    environment {
        // ANSIBLE_CONFIG = '/etc/ansible/ansible.cfg'
        // PRIVATE_KEY_PATH = '/etc/ansible/morokey.pem'
        // ANSIBLE_LOCAL_TEMP = '/tmp/ansible_tmp'
        // LOGIN_CREDS = credentials('docker_login_creds')
        // ANSIBLE_PRIVATE_KEY=credentials('SOLO-PRIVATE-KEY') 

        // GO114MODULE = 'on'
        // CGO_ENABLED = 0 
        // GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
        // LOGIN_CREDS = credentials('docker_login_creds')
        // HELM_NAME = "tutor-golang"
        DOCKER_REGISTRY="docker.io"
        DOCKER_CONTAINER="fismed-be-dev"
        DOCKER_IMAGE="fismed-be"
        // DOCKER_VERSION=sh (returnStdout: true, script: 'echo v$(./git-buildnumber)').trim()
        // HELM_REPO="oci://registry-1.docker.io/solusik8s"
        // HELM_CHART="steradian"
        // HELM_VERSION="1.0.2"
        // HELM_NS="helm-tutor"
        // HELM_VALUES="value-tutor-golang.yaml"

    }

    stages {

        stage('Initialize GO') {
            steps {
                script {
                        echo "initialize GO"
                        sh 'pwd'
                        sh 'echo "-------------"'
                        sh 'ls -al'
                        sh 'echo "-------------"'
                        sh 'whoami'

                    
                }
            }
        }
        stage("Build Docker image") {
            steps{
                echo "Build image"
                sh 'docker compose build'
                
            }
        }



        stage("Push Docker Image") {
            steps{
                echo "Push Image"
           
                sh '''
                    docker stop fismed-be-dev
                    docker compose down
                    docker rm fismed-be-dev
                    docker compose up -d
                '''
                
            }
        }
        // stage("Deploy"){
        //     steps{
        //         echo "Deploy"
        //         withKubeConfig([credentialsId: 'solo-k8s-token-helmtutor', serverUrl: 'https://solo-rke.fadjri.bid/k8s/clusters/local']) {
                    

        //             sh '''
        //                 echo running helm - HELM_NAME=${HELM_NAME}
        //                 echo image.repository=${DOCKER_REGISTRY}/${DOCKER_USER}/${DOCKER_IMAGE}
        //                 echo image.tag=${DOCKER_VERSION}
        //                 helm upgrade --install ${HELM_NAME} ${HELM_REPO}/${HELM_CHART} \
        //                     --set nameOverride=${HELM_NAME} \
        //                     --set fullnameOverride=${HELM_NAME} \
        //                     --version ${HELM_VERSION} \
        //                     --set image.repository=${DOCKER_REGISTRY}/${DOCKER_USER}/${DOCKER_IMAGE} \
        //                     --set image.tag=${DOCKER_VERSION} \
        //                     --namespace ${HELM_NS} \
        //                     --values ${HELM_VALUES} \
                            
        //             '''
        //         }
        //     }
        // }
// ansible-playbook -i hosts --private-key=$ANSIBLE_PRIVATE_KEY deploy-container.yaml
        // ansible -i hosts --private-key=$ANSIBLE_PRIVATE_KEY k8s-master -m ping
        // stage('Run Ansible Playbook') {
        //     steps {
        //         script {
        //             sh 'ansible-playbook ansible/deploy-container.yaml'
        //         }
        //     }
        // }
    }

    post {
        always {
            cleanWs()
        }
    }
}