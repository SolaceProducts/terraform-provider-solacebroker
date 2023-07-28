String GIT_REPO_PROVIDER = 'git@github.com:SolaceDev/terraform-provider-solacebroker.git'
String JENKINSCRED_GH_ROBOT_ID = 're-github-bot-1'

PROVIDER_NAME = 'terraform-provider-solacebroker'
TF_ORGANIZATION_NAME = 'SolaceDev-CRE'
TF_REGISTRY_TYPE = 'private'
TF_REGISTRY_NAME = 'solacebroker'
PROVIDER_BINARY_UPLOAD_URLS = [:]
PROVIDER_VERSION = ''

library 'jenkins-pipeline-library@main'

BINARIES = [
  'darwin_amd64', 
  'darwin_arm64', 
  
  'freebsd_386', 
  'freebsd_arm', 
  'freebsd_arm64', 
  'freebsd_amd64', 

  'linux_386', 
  'linux_arm', 
  'linux_arm64', 
  'linux_amd64', 

  'windows_386', 
  'windows_arm', 
  'windows_arm64', 
  'windows_amd64', 
]

def extractSemanticVersion(String inputString) {
	def pattern = /\d+\.\d+\.\d+/
	def matcher = (inputString =~ pattern)
	return matcher.find() ? matcher.group() : null
}

def isTerraformVersionExists(version){
  httpStatusCode = sh(
  script: """
      curl -s -o /dev/null -w '%{http_code}' \
      --header "Authorization: Bearer ${TF_BEARER_TOKEN}" \
      --header "Content-Type: application/vnd.api+json" \
      --request GET \
      https://app.terraform.io/api/v2/organizations/${TF_ORGANIZATION_NAME}/registry-providers/${TF_REGISTRY_TYPE}/${TF_ORGANIZATION_NAME}/${TF_REGISTRY_NAME}/versions/${version}
    """, 
    returnStdout: true
  ).trim()
  return (httpStatusCode == '200')
}

node(label: 'master') {
	def root = tool type: 'go', name: 'go120'
	def nodeRoot = tool name: 'node19'
	withEnv([
	"GOROOT=${root}", 
	"PATH+GO=${root}/bin",
	"PATH+NODE=${nodeRoot}/bin"
	]) {
		stage('Get Semantic Version') {
			cleanWs()
			checkout ( [$class: 'GitSCM',
				branches: [[name: 'main' ]],
				userRemoteConfigs: [[
          credentialsId: JENKINSCRED_GH_ROBOT_ID, 
          url: GIT_REPO_PROVIDER]]
			])
			sshagent(credentials: [JENKINSCRED_GH_ROBOT_ID]) {
				sh "git checkout ${env.BRANCH_NAME}"
				PROVIDER_VERSION = extractSemanticVersion(env.BRANCH_NAME)
        sh 'npm install adm-zip' 

				SHASUMS_FILE = "${PROVIDER_NAME}_${PROVIDER_VERSION}_SHA256SUMS"
				SHASUMS_SIG_FILE = "${PROVIDER_NAME}_${PROVIDER_VERSION}_SHA256SUMS.sig"
			}
		}

    if (PROVIDER_VERSION == null) {
      echo '[UNSTABLE] No Semantic Version Found'
      currentBuild.result = 'UNSTABLE'
      return
    }

    stage ('Create binaries and SHASUMS') {
      withCredentials([
        string(credentialsId: 'terraform-github-token-secret', variable: 'GITHUB_TOKEN'), 
        string(credentialsId: 'tf-passphrase', variable: 'GPG_PASSPHRASE'), 
        string(credentialsId: 'terraform-registry-key-id', variable: 'TF_REGISTRY_KEY_ID'), 
        file(credentialsId: 'tf-gpg-private-key-file', variable: 'privateKey'),
        file(credentialsId: 'tf-gpg-public-key-file', variable: 'publicKey')
      ]) {
        env.GPG_TTY="\$(tty)"  
        env.PROVIDER_VERSION=PROVIDER_VERSION
        sh "gpg --import --passphrase ${GPG_PASSPHRASE} --batch --yes --no-tty ${privateKey}"
        sh "gpg --import --passphrase ${GPG_PASSPHRASE} --batch --yes --no-tty  ${publicKey}"
        sh './prepare_terraform_release.sh'
      }
    }

		stage ('Create Registry Version') {
			withCredentials([
			string(credentialsId: 'terraform-registry-key-id', variable: 'TF_REGISTRY_KEY_ID'), 
			string(credentialsId: 'terraform-bearer-token', variable: 'TF_BEARER_TOKEN')
			]) {
				if (isTerraformVersionExists(PROVIDER_VERSION)){
					deleteVersionResponse = sh(
						script: """
							curl \
							--header "Authorization: Bearer ${TF_BEARER_TOKEN}" \
							--header "Content-Type: application/vnd.api+json" \
							--request DELETE \
							https://app.terraform.io/api/v2/organizations/${TF_ORGANIZATION_NAME}/registry-providers/${TF_REGISTRY_TYPE}/${TF_ORGANIZATION_NAME}/${TF_REGISTRY_NAME}/versions/${PROVIDER_VERSION}
						""", 
						returnStdout: true
					).trim().tokenize("\n") 
				}
				request = """
					{
						"data": {
							"type": "registry-provider-versions",
							"attributes": {
								"version": "${PROVIDER_VERSION}",
								"key-id": "${TF_REGISTRY_KEY_ID}",
								"protocols": ["5.0"]
							}
						}
					}
				"""

				println ("creating release version")
				url = "https://app.terraform.io/api/v2/organizations/${TF_ORGANIZATION_NAME}/registry-providers/${TF_REGISTRY_TYPE}/${TF_ORGANIZATION_NAME}/${TF_REGISTRY_NAME}/versions"
				def (String response, int code) = sh(
					script: """
						curl \
							-X POST \
							-H 'Authorization: Bearer ${TF_BEARER_TOKEN}' \
							-H 'Content-Type: application/vnd.api+json' \
							-d '${request}' \
							'${url}'
					""", 
					returnStdout: true
				).trim().tokenize("\n") 

				json = new groovy.json.JsonSlurperClassic().parseText(response)
				SHASUMS_UPLOAD_URL = json['data']['links']['shasums-upload']
				SHASUMS_SIG_UPLOAD_URL = json['data']['links']['shasums-sig-upload']
			}
		}

		stage ('Upload SHASUMS') {
			sh "curl -T dist/${SHASUMS_FILE} ${SHASUMS_UPLOAD_URL}"
			sh "curl -T dist/${SHASUMS_SIG_FILE} ${SHASUMS_SIG_UPLOAD_URL}"
		}

		stage ('Create Platforms for Binaries') {
			for (binary in BINARIES) {
				osAndArch = binary.split('_')
				binaryFileName = "${PROVIDER_NAME}_${PROVIDER_VERSION}_${binary}.zip"
				
				fileShasum = sh(
					script: "grep ${binaryFileName} dist/${SHASUMS_FILE} | awk '{ print \$1 }'",
					returnStdout: true
				).trim()
				
				withCredentials([
						string(credentialsId: 'terraform-bearer-token', variable: 'TF_BEARER_TOKEN')
				]) {
						request = """
								{
									"data": {
											"type": "registry-provider-version-platforms",
											"attributes": {
													"os": "${osAndArch[0]}",
													"arch": "${osAndArch[1]}",
													"shasum": "${fileShasum}",
													"filename": "${binaryFileName}"
											}
									}
								}
						"""
						url = "https://app.terraform.io/api/v2/organizations/${TF_ORGANIZATION_NAME}/registry-providers/${TF_REGISTRY_TYPE}/${TF_ORGANIZATION_NAME}/${TF_REGISTRY_NAME}/versions/${PROVIDER_VERSION}/platforms"
						def (String response, int code) = sh(
								script: """curl \
									-X POST \
									-H 'Authorization: Bearer ${TF_BEARER_TOKEN}' \
									-H 'Content-Type: application/vnd.api+json' \
									-d '${request}' \
									'${url}'
								""", 
								returnStdout: true
						).trim().tokenize("\n")
						json = new groovy.json.JsonSlurperClassic().parseText(response)
						PROVIDER_BINARY_UPLOAD_URLS[binary] = json['data']['links']['provider-binary-upload']
				}
			}
		}

		stage ('Upload Binaries') {
			PROVIDER_BINARY_UPLOAD_URLS.each{entry ->  
				sh "curl -T dist/${PROVIDER_NAME}_${PROVIDER_VERSION}_${entry.key}.zip ${entry.value}"
			}
		}        
	}
}
