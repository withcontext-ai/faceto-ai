# What is ChatAPI
* It is a parameter to get the video link, [Get FaceTo-AI Room Link](https://apifox.com/apidoc/shared-1fbfe214-d536-41b7-8209-bb504e876572/api-85139473)
* ChatAPI is a third-party interface for interfacing with ChatGPT.
* This interface enables video custom dialog streams.
* Feel free to customize your interface with ChatGPT's prompt, temperature, top_p, n, etc., and wonderful things will happen.
* This interface needs to be developed by your engineers

# Why We Support ChatAPI
* Our FaceTo-AI talks directly to chatapi, but our prompt is special and can only be used in our scenario.
* In order to achieve the diversity of video conversations, we have specifically developed this interface, and through the customization of third-party interfaces, things get interesting.
* We hope that we can realize the diversity and infinite possibilities of video dialogue by opening third-party interfaces, and we will realize the scenarios you want you to think of.

# How To Design ChatAPI
### 1.1 Communication protocol
> HTTPS
### 1.2 Request method
> Only the POST method is supported to initiate requests.
### 1.3 Character encoding
> The BASE64 encoding of HTTP communication and message adopts UTF-8 character set encoding format.
### 1.4 Context-Type
> application/json
### 1.5 Parameter
> Request Body

| name | type | description  | mark |
| ---- | ---- | ---- | ---- |
| messages | array | video  conversations  | **required** |

> message

| name | type | description  | mark |
| ---- | ---- | ---- | ---- |
| role | string | Chat message role defined by the OpenAI API.  | **required** The role of the author of this message. One of `system`, `user`, or `assistant` |
| content | string | The contents of the message.  | **required** |
| name | string | This property isn't in the official documentation, but it's in the documentation for the official library for python: |  |

[the documentation for the official library for python:](https://github.com/openai/openai-python/blob/main/chatml.md)
[How_to_count_tokens_with_tiktoken.ipynb](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb)

> example

```
{
    "messages": [
        {
            "role": "user",
            "content": "hello"
        }
    ]
}
```

### 1.6 Return Value
> **The return value must be returned as a stream**
> Accept: text/event-stream

> example
```
data:{"content":"I"}

data:{"content":" am"}

data:{"content":" an"}

data:{"content":" AI"}

data:{"content":" language"}

data:{"content":" model"}

data:{"content":" created"}

data:{"content":" by"}

data:{"content":" Open"}

data:{"content":"AI"}

data:{"content":"."}

data:{"content":" I"}

data:{"content":" am"}

data:{"content":" designed"}

data:{"content":" to"}

data:{"content":" assist"}

data:{"content":" and"}

data:{"content":" communicate"}

data:{"content":" with"}

data:{"content":" human"}

data:{"content":" beings"}

data:{"content":" using"}

data:{"content":" natural"}

data:{"content":" language"}

data:{"content":" processing"}

data:{"content":"."}

data:[DONE]
```

> **Attention: the end must be 'data:[DONE]'**

`data:[DONE]`

### 1.7 Curl Test
```
curl --location --request POST '{Your ChatAPI URL}' \
--header 'Content-Type: application/json' \
--header 'Accept: */*' \
--header 'Connection: keep-alive' \
--data-raw '{
    "messages": [
        {
            "role": "user",
            "content": "who are you?"
        }
    ]
}'
```