# [DEMO] mattermost-plugin-mimic-user

A demo mattermost plugin to allow creating mattermost post mimicing any user.

Note: only for demo purposes. Not intended for use in production.

## Installation and setup

### Platform & tools

- Make sure you have following components installed:

  - Go - v1.14 - [Getting Started](https://golang.org/doc/install)
    > **Note:** If you have installed Go to a custom location, make sure the `$GOROOT` variable is set properly. Refer [Installing to a custom location](https://golang.org/doc/install#install).

  - NodeJS - v12.18 and NPM - [Downloading and installing Node.js and npm](https://docs.npmjs.com/getting-started/installing-node).

  - Make

## Building the plugins

- Run the following commands to prepare a compiled, distributable plugin zip:

```bash
$ mkdir -p ${GOPATH}/src/github.com/Brightscout
$ cd ${GOPATH}/src/github.com/Brightscout
$ git clone git@github.com:Brightscout/mattermost-plugin-mimic-user.git
$ cd mattermost-plugin-mimic-user
$ make dist
```

- This will produce a `.tar.gz` file in `/dist` directory that can be uploaded to mattermost.

## Installation

1. Upload this file in the Mattermost **System Console > Plugins > Management** page to install the plugin. To learn more about how to upload a plugin, [see the documentation](https://docs.mattermost.com/administration/plugins.html#plugin-uploads).

## Using the plugin

When you’ve tested the plugin and confirmed it’s working, Use the following curl command to create post on behalf of any user.

```bash
curl --location --request POST 'http://<Your-Mattermost-URL>/plugins/mattermost-plugin-mimic-user/api/v1/create-post' \
--header 'X-CSRF-Token: <X-CSRF-TOKEN>' \
--header 'X-Requested-With: XMLHttpRequest' \
--header 'Content-Type: application/json' \
--header 'Cookie: <Cookie>' \
--data-raw '{
    "message": <your-message>,
    "channel_id": <channel-id>,
    "user_id": <user-id>
 }'
```

1. Replace `<Your-Mattermost-URL>` with your Mattermost URL.
2. For `<X-CSRF_TOKEN>` and `<Cookie>` go to your browser and open mattermost, then open your browser developer console and go to the Network tab.
3. Now create a post in any channel, group or dm, In the network tab, there with be a new XHR call with a name `posts`, click on it and go to headers.
4. In headers there will be request headers, and under request headers you will see X-CSRF-TOKEN and Cookie, copy both of them.
4. Now replace `<X-CSRF-TOKEN>` and `<Cookie>` in curl command with the copied value in previous step.
5. Replace `<your-message>` with the message you want to post.
6. Replace `<channel-id>` with the channel id in which you want to create the post.
7. Replace `<user id>` with the user id you want to create the post. 
