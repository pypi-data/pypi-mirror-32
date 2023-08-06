# restcountries

A library that provides a Python wrapper to [Musixmatch API](https://developer.musixmatch.com/).

## Installation

The easiest way to install the latest version
is by using pip/easy_install to pull it from PyPI:

    pip install python-musixmatch

You may also use Git to clone the repository from
Github and install it manually:

    git clone https://github.com/yakupadakli/python-musixmatch.git
    cd python-musixmatch
    python setup.py install

Python 2.7, 3.4, 3.5 and 3.6, is supported for now.


## Usage

    from musixmatch.api import Musixmatch
    
    api_key = "1dbab33f7d05917dd776bcf0ae403c5b"
    musixmatch = Musixmatch(api_key)



{
    "message": {
        "header": {
            "status_code": 200,
            "execute_time": 0.0080959796905518
        },
        "body": {
            "artist": {
                "artist_id": 1039,
                "artist_mbid": "cc197bad-dc9c-440d-a5b5-d52ba2e14234",
                "artist_name": "Coldplay",
                "artist_name_translation_list": [],
                "artist_comment": "",
                "artist_country": "GB",
                "artist_alias_list": [
                    {
                        "artist_alias": "こーるどぷれい"
                    },
                    {
                        "artist_alias": "ku wan yue dui"
                    },
                    {
                        "artist_alias": "The Coldplay"
                    },
                    {
                        "artist_alias": "Cold Play"
                    }
                ],
                "artist_rating": 95,
                "primary_genres": {
                    "music_genre_list": [
                        {
                            "music_genre": {
                                "music_genre_id": 20,
                                "music_genre_parent_id": 34,
                                "music_genre_name": "Alternative",
                                "music_genre_name_extended": "Alternative",
                                "music_genre_vanity": "Alternative"
                            }
                        },
                        {
                            "music_genre": {
                                "music_genre_id": 21,
                                "music_genre_parent_id": 34,
                                "music_genre_name": "Rock",
                                "music_genre_name_extended": "Rock",
                                "music_genre_vanity": "Rock"
                            }
                        }
                    ]
                },
                "secondary_genres": {
                    "music_genre_list": [
                        {
                            "music_genre": {
                                "music_genre_id": 14,
                                "music_genre_parent_id": 34,
                                "music_genre_name": "Pop",
                                "music_genre_name_extended": "Pop",
                                "music_genre_vanity": "Pop"
                            }
                        }
                    ]
                },
                "artist_twitter_url": "https://twitter.com/coldplay",
                "artist_vanity_id": "Coldplay",
                "artist_edit_url": "https://www.musixmatch.com/artist/Coldplay?utm_source=application&utm_campaign=api&utm_medium=Yakup%3A1409612952973",
                "artist_share_url": "https://www.musixmatch.com/artist/Coldplay",
                "artist_credits": {
                    "artist_list": []
                },
                "restricted": 0,
                "managed": 0,
                "updated_time": "2013-11-05T11:24:57Z"
            }
        }
    }
}