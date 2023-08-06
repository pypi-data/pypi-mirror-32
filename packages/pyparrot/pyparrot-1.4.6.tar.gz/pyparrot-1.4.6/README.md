# pyparrot
Python interface for Parrot Drones

pyparrot is designed to program Parrot Mambo and Parrot Bebop 2 drones using python.  This interface was developed to teach kids (K-12) STEM concepts (programming, math, and more) by having them program a drone to fly autonomously.  Anyone can use it who is interested in autonomous drone programming!   

# Installation, Quick-start, Documenation, FAQs
The GitHub [wiki page for pyparrot](https://github.com/amymcgovern/pyparrot/wiki) has extensive documentation on installing and using pyparrot.  If you are looking for documenation, please read the wiki!  

# Major updates and releases:
* 5/28/2018: Version 1.4.5 Fixed imports for new pypi structure and added xml files to pypi.
* 5/25/2018: Version 1.4.3. Uploaded to pypi so pyparrot can now be installed directory from pip.  Updated documentation for new vision.
* 5/23/2018: Updated function (contributed) to download pictures from Mambo's downward facing camera. 
* 3/25/2018: Added DroneVisionGUI which is a version of the vision that shows the video stream (for Bebop or Mambo) in real time.
* 2/22/2018: Version 1.3.2.  Updated DroneVision to make the vision processing faster.  Interface changed to only have the user call open_vision and close_vision (and not start_video_buffering)
* 2/10/2018: Version 1.3.1. Updated DroneVision to work on Windows.
* 2/8/2018: Version 1.3. Vision is working for both the Mambo and Bebop in a general interface called DroneVision.  Major documenation updates as well.
* 2/6/2018: Updated Mambo to add speed settings for tilt & vertical.  Needed for class.
* 2/4/2018: Unofficial updates to add ffmpeg support to the vision (will make an official release with examples soon)
* 12/09/2017: Version 1.2.  Mambo now gives estimated orientation using quaternions.  Bebop now streams vision, which is accessible via VLC or other video clients.  Coming soon: opencv hooks into the vision.  
* 12/02/2017: Version 1.1.  Fixed sensors with multiple values for Mambo and Bebop.
* 11/26/2017: Initial release, version 1.0.  Working wifi and BLE for Mambo, initial flight for Bebop.

# Programming and using your drones responsibly

It is your job to program and use your drones responsibly!  We are not responsible for any losses or damages of your drones or injuries.  Please fly safely and obey all laws.

