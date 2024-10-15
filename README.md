# GenesysCloud Audio Downloader
Quick little tool to automate exporting audio prompts from the platform while also documenting them into the CSV

### Requirements
* Organization oauth set up on the platform and info passed to .env<br>
* oauth profile needs to have ***architect:userPrompt:view*** permission assigned
* Correct the page size in the API url based on your requirements, default is **?pageSize=220**
* Current region set to USW2, change to correct one before use

### API
https://developer.genesys.cloud/devapps/api-explorer#get-api-v2-architect-prompts
