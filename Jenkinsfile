String GIT_REPO_PROVIDER = 'git@github.com:SolaceDev/terraform-provider-solacebroker.git'
String JENKINSCRED_GH_ROBOT_ID = 're-github-bot-1'

library 'jenkins-pipeline-library@main'

PROVIDER_NAME = 'terraform-provider-solacebroker'
TF_ORGANIZATION_NAME = 'SolaceDev-CRE'
TF_REGISTRY_TYPE = 'private'
TF_REGISTRY_NAME = 'solacebroker'
GIT_TAG_EXISTS = ''
PROVIDER_BINARY_UPLOAD_URLS = [:]
PROVIDER_VERSION = ''

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

node(label: 'master') {
	def root = tool type: 'go', name: 'go120'

	withEnv([
	"GOROOT=${root}", 
	"PATH+GO=${root}/bin"
	]) {
		stage('Check for Git Tag'){
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

				gitTagCheck = sh(
						script: "git tag -l v${PROVIDER_VERSION}",
						returnStdout: true
				).trim()

				SHASUMS_FILE = "${PROVIDER_NAME}_${PROVIDER_VERSION}_SHA256SUMS"
				SHASUMS_SIG_FILE = "${PROVIDER_NAME}_${PROVIDER_VERSION}_SHA256SUMS.sig"

				GIT_TAG_EXISTS = (gitTagCheck == '') ? false : true

				if (GIT_TAG_EXISTS) {
					//If the git tag exists we want to clear the old binaries and build new ones
					withCredentials([
					string(credentialsId: 'terraform-github-token-secret', variable: 'GITHUB_TOKEN'), 
					]) {
							getReleaseResponse = sh(
								script: """
									curl -L \
									-H "Accept: application/vnd.github+json" \
									-H "Authorization: Bearer ${GITHUB_TOKEN}"\
									-H "X-GitHub-Api-Version: 2022-11-28" \
									https://api.github.com/repos/SolaceDev/${PROVIDER_NAME}/releases/tags/v${PROVIDER_VERSION}
								""", 
								returnStdout: true
							).trim()
							releaseJson = new groovy.json.JsonSlurperClassic().parseText(getReleaseResponse)
							releaseId = releaseJson['id']
							deleteReleaseResponse = sh(
								script: """
									curl -L \
									-X DELETE \
									-H "Accept: application/vnd.github+json" \
									-H "Authorization: Bearer ${GITHUB_TOKEN}"\
									-H "X-GitHub-Api-Version: 2022-11-28" \
									https://api.github.com/repos/SolaceDev/${PROVIDER_NAME}/releases/${releaseId}
								""", 
							returnStdout: true
							)
					}
				} else {
					sh "git tag v${PROVIDER_VERSION}"
					sh "git push origin v${PROVIDER_VERSION}"
				}
			}
		}

		stage ('Create binaries and SHASUMS'){
			withCredentials([
				string(credentialsId: 'terraform-github-token-secret', variable: 'GITHUB_TOKEN'), 
				string(credentialsId: 'tf-passphrase', variable: 'GPG_PASSPHRASE'), 
				file(credentialsId: 'tf-gpg-private-key-file', variable: 'privateKey')
			]) {
				env.GPG_TTY="\$(tty)"  
				env.GORELEASER_CURRENT_TAG="v${PROVIDER_VERSION}"  
				sh "gpg --import --passphrase ${GPG_PASSPHRASE} --batch --yes --no-tty ${privateKey}"
				
				sh 'curl -sfL https://goreleaser.com/static/run | bash'
				sh "cd dist/ && gpg --batch --pinentry-mode loopback --passphrase ${GPG_PASSPHRASE} --armor --detach-sign --output ${SHASUMS_SIG_FILE} ${SHASUMS_FILE}"
			}
		}

		stage ('Create Registry Version'){
			withCredentials([
			string(credentialsId: 'terraform-registry-key-id', variable: 'TF_REGISTRY_KEY_ID'), 
			string(credentialsId: 'terraform-bearer-token', variable: 'TF_BEARER_TOKEN')
			]) {
					if (GIT_TAG_EXISTS){
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

		stage ('Upload SHASUMS'){
			sh "curl -T dist/${SHASUMS_FILE} ${SHASUMS_UPLOAD_URL}"
			sh "curl -T dist/${SHASUMS_SIG_FILE} ${SHASUMS_SIG_UPLOAD_URL}"
		}


		stage ('Create Platforms for Binaries'){
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
													"filename": "${binary}"
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

		stage ('Upload Binaries'){
			PROVIDER_BINARY_UPLOAD_URLS.each{entry ->  
				sh "curl -T dist/${PROVIDER_NAME}_${PROVIDER_VERSION}_${entry.key}.zip ${entry.value}"
			}
		}        
	}
}

  