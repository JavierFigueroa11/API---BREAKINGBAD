# Breaking Bad API 

<img align="center" width="600" src="markdown/BrBad.png">


#
## Contents
[Descriptions](#Descriptions)  

[Stage](#Stage)  

[Testing](#Testing) 

[Tools](#Tools)  

[Endpoints - Characters](#GET---GetCharacters) 

[Endpoints - Quotes](#GET---GetQuotes)

[Endpoints - Episodes](#GET---GetEpisodes)

[Endpoints - Deaths](#GET---GetDeaths)  

# 
##  Descriptions
 This is a REST API developed in golang to obtain information about the episodes, quotes of the characters and how they died and see the characters.
 The file **"CharacterBB.txt"** contains the data of all the characters to use them in the REST API.
 

 # 
 ## Stage
The project is carried out with a MVC structure (Model Views Controllers).

**Router:** This a interfaces between endpoints and its corresponding controllers.

**Controllers:** This is the intermediary between the router and the services. It handles the rest messages.

**Services** This contains the business logic.

**Domain**: This contains the structs and interfaces to development of the API.


#
## Testing

To run the test in golang, you must first position yourself in the folder that contains the test, in this case it is  **`domain.go`**

**The commands to run the tests are :**

`go test`

`go test -cover`


#
## Tools

The selected API is about [Breaking Bad!](breakingbadapi.com), we also use **GoLand** as IDE and finally we use **POSTMAN** which serves as a REST API test to make requests in the api.



# 
## Endpoints
###Characters

#### GET-GetCharacters()
>localhost:8086/getcharacters

**Response Code** 200 Ok

**Response Body** 
```
]

        {
                "char_id": 0,
                "name": "",
                "birthday": "",
                "occupation": null,
                "img": "",
                "status": "",
                "nickname": "",
                "appearance": null,
                "portrayed": "",
                "category": "",
                "better_call_saul_appearance": null
            },
            {
                "char_id": 1,
                "name": "Walter White",
                "birthday": "09-07-1958",
                "occupation": [
                    "High School Chemistry Teacher",
                    "Meth King Pin"
                ],
                "img": "https://images.amcnetworks.com/amc.com/wp-content/uploads/2015/04/cast_bb_700x1000_walter-white-lg.jpg",
                "status": "Presumed dead",
                "nickname": "Heisenberg",
                "appearance": [
                    1,
                    2,
                    3,
                    4,
                    5
                ],
                "portrayed": "Bryan Cranston",
                "category": "Breaking Bad",
                "better_call_saul_appearance": []
            },
            {
                "char_id": 2,
                "name": "Jesse Pinkman",
                "birthday": "09-24-1984",
                "occupation": [
                    "Meth Dealer"
                ],
                "img": "https://upload.wikimedia.org/wikipedia/en/thumb/f/f2/Jesse_Pinkman2.jpg/220px-Jesse_Pinkman2.jpg",
                "status": "Alive",
                "nickname": "Cap n' Cook",
                "appearance": [
                    1,
                    2,
                    3,
                    4,
                    5
                ],
                "portrayed": "Aaron Paul",
                "category": "Breaking Bad",
                "better_call_saul_appearance": []
            },
         ...
]
```
#
#### GET-GetCharactersID()
>localhost:8086/getcharacters/:id

**Response Code** 200 Ok

**Response Body**
```
]

        {
                "char_id": 0,
                "name": "",
                "birthday": "",
                "occupation": null,
                "img": "",
                "status": "",
                "nickname": "",
                "appearance": null,
                "portrayed": "",
                "category": "",
                "better_call_saul_appearance": null
            },
    ...
]
```


# 
#### POST-PostCharacters()
>localhost:8086/createcharacters

**Response Code** 200 Ok

**Response Body** 
```
{
    "char_id": 2,
    "name": "Jesse Pinkman",
    "birthday": "09-24-1984",
    "occupation": [
        "Meth Dealer"
    ],
    "img": "https://upload.wikimedia.org/wikipedia/en/thumb/f/f2/Jesse_Pinkman2.jpg/220px-Jesse_Pinkman2.jpg",
    "status": "Alive",
    "nickname": "Cap n' Cook",
    "appearance": [
        1,
        2,
        3,
        4,
        5
    ],
    "portrayed": "Aaron Paul",
    "category": "Breaking Bad",
    "better_call_saul_appearance": []
}
```

#
#### PUT- PutCharacters()
>localhost:8086/updatecharacters

**Response Code** 200 Ok

**Response Body**
```
{
    "char_id": 0,
    "name": "Freddie Pickman",
    "birthday": "05-12-1996",
    "occupation": [
        "Student",
        "Father"
    ],
    "img": "Link",
    "status": "Alive",
    "nickname": "",
    "appearance": [
        1
    ],
    "portrayed": "Javier Figueroa",
    "category": "Breaking Bad",
    "better_call_saul_appearance": []
}
```

#
#### PUT- DeleteCharactersID()
>localhost:8086/deletecharacters/:id

**Response Code** 200 Ok

**Response Body**
```
[
    {
        "char_id": 0,
        "name": "Freddie Pickman",
        "birthday": "05-12-1996",
        "occupation": [
            "Student",
            "Father"
        ],
        "img": "Link",
        "status": "Alive",
        "nickname": "",
        "appearance": [
            1
        ],
        "portrayed": "Javier Figueroa",
        "category": "Breaking Bad",
        "better_call_saul_appearance": []
    },
    {
        "char_id": 1,
        "name": "Walter White",
        "birthday": "09-07-1958",
        "occupation": [
            "High School Chemistry Teacher",
            "Meth King Pin"
        ],
        "img": "https://images.amcnetworks.com/amc.com/wp-content/uploads/2015/04/cast_bb_700x1000_walter-white-lg.jpg",
        "status": "Presumed dead",
        "nickname": "Heisenberg",
        "appearance": [
            1,
            2,
            3,
            4,
            5
        ],
        "portrayed": "Bryan Cranston",
        "category": "Breaking Bad",
        "better_call_saul_appearance": []
    }
]
```

#

###Quotes
#### GET-GetQuotes()
>localhost:8086/getquotes

**Response Code** 200 Ok

**Response Body** 

```
[
    {
        "quote_id": 1,
        "quote": "I am not in danger, Skyler. I am the danger!",
        "author": "Walter White",
        "series": "Breaking Bad"
    },
    {
        "quote_id": 2,
        "quote": "Stay out of my territory.",
        "author": "Walter White",
        "series": "Breaking Bad"
    },
    {
        "quote_id": 3,
        "quote": "IFT",
        "author": "Skyler White",
        "series": "Breaking Bad"
    },
    {
        "quote_id": 4,
        "quote": "I watched Jane die. I was there. And I watched her die. I watched her overdose and choke to death. I could have saved her. But I didn’t.",
        "author": "Walter White",
        "series": "Breaking Bad"
    },
  ...
]
```

#
#### GET-GetQuotesID()
>localhost:8086/getquotes/:id

**Response Code** 200 Ok

**Response Body** 
```
[
    {
        "quote_id": 3,
        "quote": "IFT",
        "author": "Skyler White",
        "series": "Breaking Bad"
    }
]
```


#

###Episodes
#### GET-GetEpisodes()
>localhost:8086/getepisodes

**Response Code** 200 Ok

**Response Body** 
```
[
    {
        "episode_id": 1,
        "title": "Pilot",
        "season": "1",
        "air_date": "01-20-2008",
        "characters": [
            "Walter White",
            "Jesse Pinkman",
            "Skyler White",
            "Hank Schrader",
            "Marie Schrader",
            "Walter White Jr.",
            "Krazy-8",
            "Bogdan Wolynetz"
        ],
        "episode": "1",
        "series": "Breaking Bad"
    },
    {
        "episode_id": 2,
        "title": "Cat's in the Bag...",
        "season": "1",
        "air_date": "01-27-2008",
        "characters": [
            "Walter White",
            "Jesse Pinkman",
            "Skyler White",
            "Walter White Jr.",
            "Krazy-8"
        ],
        "episode": "2",
        "series": "Breaking Bad"
    },
    {
        "episode_id": 3,
        "title": "...And the Bag's in the River",
        "season": "1",
        "air_date": "02-10-2008",
        "characters": [
            "Walter White",
            "Jesse Pinkman",
            "Skyler White",
            "Hank Schrader",
            "Marie Schrader",
            "Walter White Jr.",
            "Krazy-8",
            "Gretchen Schwartz"
        ],
        "episode": "3",
        "series": "Breaking Bad"
    },
 ...
]
```
#
#### GET-GetEpisodesID()
>localhost:8086/getepisodes/:id

**Response Code** 200 Ok

**Response Body** 
```
[
    {
        "episode_id": 18,
        "title": "Mandala",
        "season": "2",
        "air_date": "05-17-2009",
        "characters": [
            "Walter White",
            "Jesse Pinkman",
            "Skyler White",
            "Saul Goodman",
            "Gustavo Fring",
            "Jane Margolis",
            "Ted Beneke",
            "Donald Margolis",
            "Combo"
        ],
        "episode": "11",
        "series": "Breaking Bad"
    }
]
```

#

###Deaths
#### GET-GetDeaths()
>localhost:8086/getdeaths

**Response Code** 200 Ok

**Response Body** 

```
[
    {
        "death_id": 40,
        "death": "Bodyguards of Gus Fring",
        "cause": "Multiple gunshots.",
        "responsible": "Walter White",
        "last_words": "What, you got a problem with stairs?",
        "season": 4,
        "episode": 13,
        "number_of_deaths": 2
    },
    {
        "death_id": 28,
        "death": "Cartel Assassins",
        "cause": "Shot in close range.",
        "responsible": "Mike Ehrmantraut",
        "last_words": "Unknown",
        "season": 4,
        "episode": 4,
        "number_of_deaths": 2
    },
    {
        "death_id": 23,
        "death": "Rival Dealers",
        "cause": "Ran over with a van, then shot in the head.",
        "responsible": "Walter White",
        "last_words": "Unknown",
        "season": 3,
        "episode": 12,
        "number_of_deaths": 2
    },
 ...
]
```

<img align="center" width="600" src="markdown/GOImg.png">
