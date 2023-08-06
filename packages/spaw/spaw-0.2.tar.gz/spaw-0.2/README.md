# Streamable API Wrapper (SPAW)
SPAW is an **Unoffical**   Streamable API wrapper for Python.  

## Streamable API Documentation
https://streamable.com/documentation

## Requirments
Python 3.5+  

The following modules are also used:
```
requests
json
```

## Usage
Install through pip  
```
pip3 install spaw
```  
To initialize an instance:  
``` 
import spaw  

spaw.SPAW() 
```

### Authentication
Streamable uses Basic Auth for their API.  
**As such the Streamable API requires Authentication to Upload/Import a video or to Retrieve a video in raw mp4 embed format.  
Only Retrieving a video in oEmbed format is allowed without Authentication.**  
To authenticate:
```
SPAW.auth(email, password)
```  
Where the arguments are the email and password of your Streamable account.
### Import a Video
Takes in given URL of a video and imports/uploads it onto Streamable.  
```
SPAW.videoImport(url)
```
Returns JSON with Shortcode and Status
```
{'status': 1, 'shortcode': 'code'}
```
The link to the video will be https://streamable.com/code

### Upload a Video
Takes in given filename for a video and uploads it onto Streamable.
```
SPAW.videoUpload(filename)
```
Returns JSON with Shortcode and Status just like importVideo()

### Retrieve a Video
Retrieves a Streamable video as embed code given the shortcode of video URL.  
The embed video can be retrieved as oEmbed or a raw mp4. by default set to 'raw'. 
```
SPAW.retrieve(shortcode,format)
```
Returns JSON for the embed code for the video.  
for e.g.  
`SPAW.retrieve('moo', 'oEmbed')`  
Returns
```
{"provider_url": "https://streamable.com", "html": "<iframe class=\"streamable-embed\" src=\"https://streamable.com/o/moo\" frameborder=\"0\" scrolling=\"no\" width=\"852\" height=\"480\" allowfullscreen></iframe>", "version": "1.0", "title": "\"Please don't eat me!\"", "type": "video", "provider_name": "Streamable", "thumbnail_url": "//images.streamable.com/west/image/f6441ae0c84311e4af010bc47400a0a4.jpg?height=100", "width": 852, "height": 480}
```
