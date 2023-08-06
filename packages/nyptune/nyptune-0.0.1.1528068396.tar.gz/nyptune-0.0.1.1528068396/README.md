# Nyptune

[![Build Status](https://travis-ci.org/fizx/nyptune.svg?branch=master)](https://travis-ci.org/fizx/nyptune)

Nyptune builds a distributed global cache of your datasets, models, Python environment, and intermediate results, enabling you to share **Jupyter notebooks that are completely reproducible**.  It's also **useful to not repeat work** when you have similar notebooks sharing datasets or intermediate results.

Nyptune can cache:

- [x] Conda environments
- [x]  PyPI environments
- [x]  data from files or URLs
- [ ] Numpy arrays
- [ ] PyTorch models
- [x] Pickleable variables
- [x] extensible for everything else!

Nyptune **builds on IPFS**, which uses technology similar to Bittorrent to host your files.  When you save your notebook, Nyptune hides IPFS links in the notebook's metadata.  You can then seed the cache from your own computer, or use a pinning service to keep them online.

How does it work?

### Installing nyptune

    :$ pip install nyptune
    :$ nyptune init

### In Jupyter

    # You need to load nyptune to take advantage of the Jupyter magics it provides
    %load_ext nyptune
    
    # Give your cache a name. Usually this should be similar to the notebook name.
    %nyptune_name readme
    
    # Optionally, if you want to encrypt your entries for a private cache
    %nyptune_secret mypassword

_

    %%cache cifar10path
    # A cache cell magic (starts with double-percent) will instruct the system that this cell's 
    # purpose is to generate the variable 'cifar10path'.  Since we already have cached this file,
    # it won't be downloaded again!
    
    import urllib.request, pathlib
    
    urllib.request.retrieve("http://files.fast.ai/data/cifar10.tgz", "cifar10.tgz")
    
    cifar10path = pathlib.Path("cifar10.tgz")
_

    # You can also use line magic to cache the results of lines.
    %cache calc = sum([x*x for x in range(100000000)])

    import torch.nn
    %cache model = torch.nn.Sequential(
      torch.nn.Linear(100, 100),
      torch.nn.ReLU(),
      torch.nn.Linear(100, 100),
    )

### Sharing the cache data

The simplest way to seed the data is to run `:$ nyptune start` and leave it running for as long as you want to share your data.  Data remains on the network for as long as at least one server is seeding it.

For long-term sharing, its helpful to "pin" the data into the ipfs network by uploading it to a server that is online all of the time.  Check out https://github.com/fizx/pincer for a free pinning service that has a one-click install into AWS, or https://www.eternum.io/ for someone's commercial service.

### Loading a Nyptune notebook

To get the most out of Nyptune's caching, when you download a notebook, you should download the relevant cache entries by running `:$ nyptune add somenotebook.ipynb; nyptune pull`

This will let you pick up where the author left off.  Often a model that took days to train will only be a few hundred megabytes, which will be much faster to download.

If you want to re-create the Python environment the notebook was saved in, you can run `:$ nyptune recreate ENV_NAME somenotebook.ipynb`.  This will create a conda environment called ENV_NAME, and install the same libraries that the notebook was using.  Run `:$ source activate ENV_NAME` as usual, and start jupyter!

### Usage

[Please read the README notebook for example usage inside Jupyter](README.ipynb)

