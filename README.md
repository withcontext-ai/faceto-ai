# ✨ Introduction
Welcome to FaceTo-AI, we're excited to see what you build!

# 🧑‍ What is FaceTo-AI
FaceTo-AI is a free, open source (Apache 2.0), opinionated, end-to-end real-time communications stack with first-party SDKs across every major software platform. It offers numerous features:
* Horizontally-scalable WebRTC Selective Forwarding Unit
* Modern, full-featured client SDKs
* JWT-based token authentication
* Backend APIs for complex workflows and low-latency coordination with the SFU
* Robust connectivity and networking over UDP, TCP, and built-in TURN
* A single, pure Go SFU binary for easy deployment
* Real-time speaker detection
* Simulcast, selective subscriptions and other automatic bandwidth management optimizations
* A universal egress and recording system for stream export
* Metadata, moderation, and data message APIs

# 🤔 Why We Build FaceTo-AI
AI video communication is a rapidly growing industry with immense potential for growth and innovation. With the increasing globalization of businesses and the rise of remote work, video conferencing has become an essential tool for communication and collaboration. The recent COVID-19 pandemic has further accelerated the adoption of video communication technologies, as more and more people are forced to work from home.

AI-powered video communication is particularly promising because it offers a range of advanced features that can enhance the user experience and improve productivity. For example, AI can be used to automatically transcribe and translate conversations in real-time, making it easier for people who speak different languages to communicate effectively. It can also be used to analyze facial expressions and body language, providing valuable insights into the emotions and reactions of participants.

Despite its many benefits, the AI video communication industry also faces a number of challenges. One of the biggest challenges is ensuring the security and privacy of user data, as video communication platforms are often targeted by hackers and other malicious actors. Another challenge is the need to continually improve the technology to keep up with changing user needs and preferences.

Overall, the AI video communication industry is poised for significant growth in the coming years, as more businesses and individuals recognize the benefits of this powerful technology. However, to fully realize its potential, the industry must continue to innovate and adapt to meet the changing needs of its users.

# 🔑 About Token
If necessary, please contact me via email and tell me your purpose. I will give you feedback in time and provide you with a testable token. Please keep it confidential. My email is **zhangxiubo@elihr.cn**
All endpoints require a signed access token. This token should be set via HTTP header:
```
Authorization: Bearer <token>
```
# 🔥 API Docs
Now, we have opened the API for accessing video links. Please refer to the interface description in the **API Docs** directory for specific usage. Please feel free to contact us if you have any questions.
[Get FaceTo-AI Room Link](https://apifox.com/apidoc/shared-1fbfe214-d536-41b7-8209-bb504e876572/api-85139473)
# 🔥 About ChatAPI
If you want to know how chatapi is designed, please go here.[About ChatAPI](./server/docs/about_chatapi.md)
# 🔥 Flow
Please see our overall design flowchart here.[Flow](./server/docs/flow.md)

---
# 👉 About Project


## Directory structure
```
# web     # frontend folder
.
# server  # backend folder
.
├── api                           # api proto file
│   ├── error                     # error proto
│   ├── faceto                    # api proto
│   └── room                      # room proto
├── api_gen                       # api proto gen folder
├── assists                       # assists folder
├── bin                           # bin
├── cmd                           # start command
│   └── faceto-ai
├── configs                       # config
├── internal                      # business processing 
│   ├── biz                       # biz handler
│   │   └── liveGPT               # liveGPT Service
│   ├── conf                      # config file
│   ├── data                      # data handler
│   │   ├── ent                   # ent db file handler
│   │   └── schema                # db table config file
│   ├── pkg                       # the third pkg folder
│   │   ├── event                 # event handler
│   │   ├── middleware            # middleware handler
│   │   └── utils                 # common utils pkg
│   │       ├── crypt
│   │       ├── helper
│   │       └── log
│   ├── server                    # server handler
│   │   └── handler 
│   └── service                   # implement proto service api
├── test
└── third_party                   # thrid proto pkg
```

# ⏩ About Server Kratos
[https://go-kratos.dev/](https://go-kratos.dev/)
## Install Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

# ⏩ Running Locally
## Run Front
```
# From the web/ directory
yarn install && yarn dev
```

## Run Server
```
# From the server/ directory
# dev
make dev

# prod
make prod
```

## Run Docker
```
# From the / directory
# makefile
make docker-all
```

```
# From the / directory
# docker-compose
docker-compose up
```

## Local Join Room
Once both services are running you can navigate to http://localhost:3000. There's one more step needed when running locally. When deployed, KITT is spawned via a LiveKit webhook, but locally - the webhook will have no way of reaching your local lkgpt-service that's running. So you'll have to manually call an API to spawn KITT:

```
# <room_name> comes from the url slug when you enter a room in the UI
curl -XPOST http://localhost:8001/join/<room_name>
```

# 🔑 About Env Key
You need to modify the `config.env.tpl` file in the configs folder to `config.env`, and then add your own key and configuration information.
The following are comments for all variables:

```
# env dev or prod
export FACETOAI_ENV=
# debug dev for true
export DEBUG=

# mysql connect string
export FACETO_DATA_DATABASE_SOURCE=""

# openai key or host
export OPENAI_KEY=
export OPENAI_HOST=

# storage 
# Since our service is deployed on Azure, there are Azure related configurations here.
export AZURE_BLOB_HOST=
export AZURE_BLOB_CDN_HOST=
export AZURE_ACCOUNT_NAME=
export AZURE_ACCOUNT_KEY=
export AZURE_BLOB_CONTAINER_NAME=
export AZURE_CONNECTION_STRING=""
export INDEX_QUEUE_NAME=

# livekit
# see https://docs.livekit.io/realtime/
export LIVEKIT_NAME=
export LIVEKIT_URL=
export LIVEKIT_API_KEY=
export LIVEKIT_SECRET_KEY=

# google gcp json file
# This is the permission authentication file of Google Cloud. You need to put the file on CDN or anywhere else. Just fill in the URL address here.
export GOOGLE_GCP_CREDENTIALS=

# eleventlabs key
# see https://docs.elevenlabs.io/welcome/introduction
export ELEVEN_API_KEY=

# logsnag key
# see https://docs.logsnag.com/get-started/quick-start
export LOGSNAG_API_KEY=
```