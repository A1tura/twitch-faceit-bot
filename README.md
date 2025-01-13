# Python Bot for FACEIT Data Retrieval
## Commands
- **!stats Day**: Returns the number of wins and losses for the day.
- **!stats Month**: Returns the number of wins and losses for the current month.

## Setup
### 1. Clone this repository to your local machine:
 ```git clone https://github.com/A1tura/twitch-faceit-bot && cd twitch-faceit-bot ```
### 2. Set **ENV** variables in the **Dockerfile**

In the `Dockerfile`, you'll need to configure the following environment variables for proper authentication and functionality.

#### ENV Variables

- **`AUTH`**: Twitch OAuth Access Token
    - Visit [Twitch Token Generator](https://twitchtokengenerator.com/).
    - Log in with your Twitch account and copy the **Access Token**. Paste it in the `AUTH` variable.

- **`CLIENT_ID`**: Twitch Client ID
    - Visit [Twitch Token Generator](https://twitchtokengenerator.com/) to get the **Client ID**.
    - Copy the **Client ID** and paste it in the `CLIENT_ID` variable.

- **`BOT_USERNAME`**: Twitch bot account username
    - Enter the username of your bot Twitch account here (this should be the account your bot uses to interact with Twitch).

- **`STREAMER_USERNAME`**: Streamer username
    - Enter the username of the Twitch streamer account you want the bot to interact with.

- **`COOLDOWN`**: Command cooldown time
    - The recommended value is **10-20 seconds**. This value is to prevent the bot from hitting Twitch API rate limits.
    - Reference the [Twitch API limits](https://dev.twitch.tv/docs/chat/) for more details.

- **`FACEIT_ID`**: Faceit player ID
    - To get your Faceit player ID, visit [Faceit API](https://www.faceit.com/api/users/v1/nicknames/*nickname*) and fetch the payload. The ID will be under the `id` field.
    - Paste this ID into the `FACEIT_ID` variable.

- **`FACEIT_API`**: Faceit API Key
    - To get your Faceit API key, visit [Faceit Developer Portal](https://developers.faceit.com/apps) and create an app.
    - Once your app is created, click on it and go to **API Keys**.
    - Click the **+ button** to create a new API key. Select **Server Side** in the "type" field.
    - After creation, copy the key and paste it into the `FACEIT_API` variable.


### 3. Build and start Docker container
- **Build image**:  ``` docker build -t twitch-faceit-bot . ```
- **Run container**: ```docker run --name twitch-faceit-bot -d twitch-faceit-bot```

### 4. DONE

