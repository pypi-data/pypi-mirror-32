
# Pleio Theme for Concierge
This is the Pleio theme for concierge.

## Quick start

### Install

    pip install concierge_theme_pleio

### Configure
Add "concierge_theme_pleio" to concierge's INSTALLED_APPS by adding it your
config.py's INSTALLED_APPS_PREFIX array.

    INSTALLED_APPS_PREFIX = [
        'concierge-theme-pleio',
        ...
    ]

## Development
If you plan on making changes to the theme, instead of installing it using
pip, clone the repo and execute the following commands.

    python setup.py develop
    cd concierge_theme_pleio
    yarn
    yarn build

If you want to make changes to the javascript or CSS files with live reload,
run the following in a terminal from the concierge_theme_pleio folder:

    yarn watch

## Building
To create a bundle ready for distribution, execute the following commands:

    (cd concierge_theme_pleio && yarn && yarn build)
    python setup.py sdist
