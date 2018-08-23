# API TOKENS
https://confluence.atlassian.com/cloud/api-tokens-938839638.html

# JIRA Software REST API
https://developer.atlassian.com/cloud/jira/software/rest/#api-board-get

# リリースされているかどうか
https://developer.atlassian.com/cloud/jira/platform/rest/#api-api-2-project-projectIdOrKey-versions-get

```
curl --request GET --url 'https://your-domain.atlassian.net/rest/api/2/project/{projectIdOrKey}/versions'  --user 'user:api-token' --header 'Accept: application/json'
```
