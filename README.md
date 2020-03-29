# Breaking Bad API 
<img align="center" width="600" src="/home/javierfigueroa/go/src/awesomeProject/markdown/GolangImg.png">

## Contents
[Descriptions](#Descriptions)  

[Stage](#Stage)  

[Endpoints - Characters](#GET---GetCharacters) 

[Endpoints - Quotes](#GET---GetQuotes)

[Endpoints - Episodes](#GET---GetEpisodes)

[Endpoints - Deaths](#GET---GetDeaths)  

# 
##  Descriptions
 This is a REST API developed in golang to obtain information about the episodes, quotes of the characters and how they died and see the characters.
 # 
 ## Stage
The project is carried out with a MVC structure (Model Views Controllers).

**Router:** This a interfaces between endpoints and its corresponding controllers.

**Controllers:** This is the intermediary between the router and the services. It handles the rest messages.

**Services** This contains the business logic.

**Domain**: This contains the structs and interfaces to development of the API.

# 
## Endpoints
#### GET-Characters()
>localhost:8086/getcharacters

**Response Code** 200 Ok

**Response Body** 
```
]

    {   
        "char_id":1,
        "name":"Walter White",
        "birthday":"09-07-1958",
        "occupation":["High School Chemistry Teacher","Meth King Pin"],
        "img":"https://images.amcnetworks.com/amc.com/wp-content/uploads/2015/04/cast_bb_700x1000_walter-white-lg.jpg",
        "status":"Presumed dead",
        "nickname":"Heisenberg",
        "appearance":[1,2,3,4,5],
        "portrayed":"Bryan Cranston",
        "category":"Breaking Bad",
        "better_call_saul_appearance":[]
    },
         ...
]
```
