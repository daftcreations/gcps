# Google Cloud Profile Switcher (GCPS)

[![Go](https://img.shields.io/badge/--00ADD8?logo=go&logoColor=ffffff)](https://golang.org/)  [![GitHub license](https://img.shields.io/github/license/Naman2706/gcps)](https://github.com/Naman2706/gcps/blob/master/LICENSE) [![GitHub release](https://img.shields.io/github/v/tag/Naman2706/gcps)](https://github.com/Naman2706/gcps/releases)


GCPS allows to switch between google cloud configuration profiles easily.

## Installation

```bash
# for linux systems
curl -o gcps https://github.com/Naman2706/gcps/releases/download/latest/gcps-all-{os}-{architecture}

chmod +x ./gcps

# move to your user bin to access it globally
mv ./gcps /usr/local/bin/
```

## Usage


<img src="demo.gif" width="100%">

- Switch through list `gcps` will return configured profile.

- Switch to particular profile `gcps {profile_name}`

- Switch to previous list `gcps -`


## Future Developement

Currently we have too many ideas to take ahead this profile swticher and make it more convenient and easy
