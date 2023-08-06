# atnlp


A tool for natural language processing.

[Read the docs](https://wedavey.github.io/atnlp/)

## Technologies

atnlp uses the following technologies: 

- [conda](https://www.anaconda.com/) - datascience platform 
- [scipy](https://www.scipy.org/) - datascience tools
- [sphinx](http://www.sphinx-doc.org/en/master/) - documentaion
- [bumpversion](https://github.com/peritus/bumpversion) - semantic versioning 
- [github](https://github.com/) - software development platform 
- [travis](https://travis-ci.org/wedavey/atnlp) - continuous integration 
 


## Quick start
More details in [quickstart docs](https://wedavey.github.io/atnlp/quickstart.html)

### Prerequisites 

- **conda** (anaconda / miniconda) - follow the 
[installation instructions](https://conda.io/docs/user-guide/install/index.html) 
for your platform and select python 3 version.


### Install 
 
Install atnlp and its dependencies: 
```bash
conda update conda -c conda-forge
git clone git@github.com:wedavey/atnlp.git
conda env create -f atnlp/environment.yml -n atnlp
conda activate atnlp
pip install ./atnlp
```

## Running tests

TODO...


## Deployment 

Start [training topic models](https://wedavey.github.io/atnlp/topiclabel.html) straight away using scripts, or open a notebook and start hacking.


## Versioning

We use [SemVer](https://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/wedavey/atnlp/tags). 

## Authors

- **Will Davey** - main developer

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
