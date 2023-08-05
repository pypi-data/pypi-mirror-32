flamepy_sg Module Repository
----------



-
2018/05/20

Can now be installed using pip:
pip install flamepy_sg

force version update with: 
pip install --upgrade --force-reinstall flamepy_sg

Once the package is intalled, source the appropriate bash_profile_... file (found in /for_install) and start Flame from the same terminal.
But first, once the file located, modify the path to be sourced according to your system and where you installed the package.
The target folder/module is 'flame_start' which should be right next to the 'flamepy_sg' folder/module.


-
2018/05/18

How to install and add to the env.

The 'flame_start' folder should be copied or moved  next to the 'flamepy_sg' module.
Then source the OS appropriate env lines from the '...bash_profile' file. You'll need to edit this file for your specific path.
The filepath to add is to the 'flame_start', and the 'flamepy_sg' folder must be next to it, inside the same folder.

updates:
- Removed the 'lib' folder and put its content into the 'script' folder
- Started to look at how to make this an installable module. (added the begining of a setup.py ...)

-
2018/05/15

About this new Flame python rep:
The FlamePy_sg folder can be placed anywhere on your computer (i.e. your working folder or /opt/... or  /var/temp/ ...)
You need the DL_PYTHON_PATH to be added to your environment, by either:
- Sourcing the text file named Paths_to_source.txt (in /lib) in the terminal/shell from which you will start Flame.
- Add the path to your .bash_profile (on mac) and then source it in the terminal from which you will start Flame.


-
2018/05/14

Stripped down from previous version (Flame python dev).
This is also an attempt to work in a more common/pythonic way, mostly regarding the structure (folders, modules, packages...) and how things are imported/linked ...
Work in progress, first step.





