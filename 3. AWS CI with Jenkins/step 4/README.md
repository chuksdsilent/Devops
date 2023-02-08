# Slack Integration

1. Create a Workspace in slack

2. Create a channel

3. Go to https://api.slack.com

4. click on start building

5. create app >> Type the name and select the workspace >> create app

6. Go to OAuth & Permissions

7. Under Bot Token Scopes click on add

8. select chat-write

9. Scroll up >> under OAuth Tokens & Redirect URLs >> click on install app to workspance

10. click on allow

11. Copy the token

12. Go to slack channel

13. type @YourAppName >> press enter >> Invite the channel

14. Go to Jenkins Dashboard >> Manage Jenkins

15. Go to plugins

16. Go to available and search for slack then install without restart

17. Go to Manage Jenkins >> Manage Credentials >> Jenkins >> Globla Credentials >> Add Credentials >> under kind select secret text

18. copy the slack token an paste in secret box then type id and description

19. Go to manage jenkins >> configure systems

20. Scroll down and check custom slack app bot user >> save

21. Open your job >> configure >> scroll down

22. Under Post-build action >> select slack notifications

23. check mark the ones you want.

24. click on advance >> scroll down and select the slack token

25. Type the workspace name and the channel(#yourchaneelname)

26. Test Connection

27. Under post build action. select archive your artifact

28. type

```
**/*.war
```

29. then save and build your job.
