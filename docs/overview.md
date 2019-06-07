# Overview

A mattermost plugin to send surveys to the user for Riff edu meetings.

- A survey is sent to the meeting participants just after they leave a meeting by closing the browser tab or through the leave call button.

- The survey is sent through a bot user configured in system console settings.

- For each survey, sent to the user, the reminder settings can be configured.
  - If the reminders are enabled, a reminder post would be sent to the meeting participant to fill the survey.
  - The reminder would only be sent if the user has not responded to the survey.
  - The message for the reminder, interval between reminders and the maximum number of reminders sent to the user can be configured in the system console settings for the plugin.

- Each time the plugin is enabled, the plugin checks for the latest version of the survey with id: `f298903f8a80054ba09e342d0d9780635d3675a2` in the DB.

- If the survey does not exist in the DB, a new entry is created with survey version set to one.

- Each time you update one or more fields for the survey in the system console settings, a new version would be created for the save SurveyID.

- Support for multiple surveys with different IDs is planned for v2.

## System Console Settings

The plugin can be configured by these settings. These settings are visible through the `/admin_console/plugins/custom/survey` page for your mattermost where this plugin is installed.

The plugin configurations are validated in the server. If any one of these are invalid, the plugin would fail to start.

The information about which validation is failing can be seen in the server logs for your mattermost instance.

1. **Bot Username**:
    - Set the user which will send the survey to meeting participants.
    - This mattermost-user can be created manually through user sign up page.
    - If the corresponding settings are enabled then for the surveys sent to the meeting participants, the username would be: "Riff Bot" and the profile image would be the Riff Logo.
    - These settings are: `Enable integrations to override usernames` and `Enable integrations to override profile picture icons` in the integrations page of system console settings.

2. **Survey**:
    - This is the schema of the survey sent to the meeting participants.
    - It is a JSON object with the following fields:
        - title: (string) This is the title of the survey displayed on the header of the survey modal.
        - description: (string) This is a descriptive text visible on opening the survey modal.
        - questions (array) This is an array of questions. Each question is a JSON object with properties: type and text.
            - type: The type of the question. Currently supported types are:
                1. five-point-likert-scale: A five point likert type question with options ranging from strongly agree to strongly disagree.
                2. open: An open ended question with a textarea
            - text: The question text.

    - Example Schema:

        ```json
            {
                "title": "Survey Title",
                "description": "Survey Description",
                "questions": [
                    {
                        "type": "five-point-likert-scale",
                        "text": "What is your rating?"
                    },
                    {
                        "type": "open",
                        "text": "Please add any other comments."
                    }
                ]
            }
        ```

3. **Dashboard Path**:
    - The relative path to the Riff Stats page. This includes the path after the `team-name`.
    - It must start with a leading slash and must not have a trailing slash.

4. **Reminder Text**:
    - The message sent to the user as a reminder to fill the survey.
    - A reminder will only be sent to a user who has not yet submitted the survey yet.

5. **Max Reminder Count**:
    - The maximum number of times a reminder would be sent to the user.
    - This can be set to `0` to disable reminders.

6. **Reminder Interval**:
    - The time interval in minutes after the survey, or the last reminder for this survey, is sent to the user, after which the next reminder would be sent.
