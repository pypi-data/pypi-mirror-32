===================================================
Google Earth Engine Batch Asset Manager with Addons
===================================================
Google Earth Engine Batch Asset Manager with Addons is an extension of the one developed by Lukasz `Lukasz's asset manager tool <https://github.com/tracek/gee_asset_manager>`_ and additional tools were added to include functionality that was missing from the EarthEngine CLI. The ambition is apart from helping user with batch actions on assets along with interacting and extending capabilities of existing `Google Earth Engine CLI <https://pypi.org/project/earthengine-api/#description>`_ . It is developed case by case basis to include more features in the future as it becomes available or as need arises.


Current Tools:
~~~~~~~~~~~~~~

* Print Earth Engine home folder quota
* Create folder or collection
* Batch uploading data
* Getting asset list and size
* Print report of all images and image collections
* Batch move and copy assets
* Batch remove collections
* Batch change properties of images or assets in a folder
* Set collection properties

Changelog
=========

v0.2.5
``````
* Handles bandnames during upload thanks to Lukasz for original upload code
* Removed manifest option for Planetscope, that can be handled by seperate tool `(ppipe) <https://github.com/samapriya/Planet-GEE-Pipeline-CLI>`_

v0.2.4
``````
* Major improvements to ingestion using manifest ingest in Google Earth Engine
* Contains manifest for all commonly used Planet Data item and asset combinations
* Added additional tool to Earth Engine Enhancement including quota check before upload to GEE

v0.2.3
``````
* Removing the initialization loop error

v0.2.2
``````
* Added improvement to earthengine authorization

v0.2.1
``````
* Added capability to handle PlanetScope 4Band Surface Reflectance Metadata Type
* General Improvements

v0.2.0
``````
* Tool improvements and enhancements

v0.1.9
``````
* New tool EE_Report was added

v0.1.8
``````
* Fixed issues with install
* Dependencies now part of setup.py
* Updated Parser and general improvements
