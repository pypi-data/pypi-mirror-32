# Geosoft GX for Python Repository

This is the repository for Geosoft GX Developer support for Python development. Refer to the documentation for more information.

[GX Developer documentation](https://geosoftgxdev.atlassian.net/wiki/display/GD/Python+in+GX+Developer)

[Python Tutorial for Geosoft GX Developer](https://geosoftgxdev.atlassian.net/wiki/spaces/GXD93/pages/103153671/Python+Tutorial+for+Geosoft+GX+Developer)

[Python Packages](https://github.com/GeosoftInc/gxpy/wiki)

Also see the [Geosoft Inc. organization on Github](https://github.com/GeosoftInc) for the other programming language specific repos.

Quick Start
-----------

### Configuration ###

SeeÂ [Python Configuration Menu](https://github.com/GeosoftInc/gxpy/wiki/Python-menu-for-Geosoft-Desktop) to install a Python menu that simplifies Python configuration for an Oasis montaj installation.

To update an existing Python installation, load the Python menu from your User Menus and select Python > Configure Python... > update geosoft package.

If you encounter problems due to a non-standard installation you can also update Python manually (see below).Â  

### Manual Configuration ###

Uninstall Geosoft from Python, then install version 9.3 as follows (you must have the Geosoft Desktop 9.3 platform installed).

```
pip uninstall geosoft
pip install geosoft
```

Or, alternately:

```
pip install geosoft --upgrade
```

### Version Compatibility ###
The base GX API, which is exposed to Python by the ___geosoft.gxapi___ module, is consistent across versions. This means that earlier versions of ___geosoft.pxpy___ will work with Geosoft Desktop 9.3. While we recommend that older scripts be updated to conform to the 9.3 API, should you need support for multiple versions of ___geosoft.gxpy___ you can create separate Anaconda Python environments for each version. For example, you might create an environment ___'py35_gx91'___ for Python 3.5 and the GX API version 9.1, ___'py36_gx92'___ for Python 3.6 and GX Developer 9.2 and 'py36_gx93' for GX Developer 9.3. If you do not depend on earlier versions of the GX Developer Python API it is best to use only the most recently released API.

Vesion 9.3 supports both Python 3.5 and 3.6.Â  If you need Python 3.4 support, install geosoft version 9.2.1, which will work with both Geosoft Desktop versions 9.2 and 9.3.

License
-------

Any source code found here are released under the [BSD 2-clause license](https://github.com/GeosoftInc/gxpy/blob/master/LICENSE). Core functionality exposed by the GX API may have additional license implications. For more information consult the [License page in the GX Developer Wiki](https://geosoftgxdev.atlassian.net/wiki/spaces/GD/pages/2359406/License)


