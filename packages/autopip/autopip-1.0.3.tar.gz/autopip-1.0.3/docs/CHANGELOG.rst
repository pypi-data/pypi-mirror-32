Version 1.0.3
================================================================================

* Uninstall autopip last when doing a group

Version 1.0.2
--------------------------------------------------------------------------------

* Update readme

Version 1.0.1
--------------------------------------------------------------------------------

* Update readme

Version 1.0.0
--------------------------------------------------------------------------------

* Set status to prod/stable
* Support update frequency from autopip entry group
* Save/show update frequency
* Add update frequency info
* Terminate autopip if running for longer than an hour
* Add --update option to specify how often to update an app

Version 0.3.4
================================================================================

* Set keywords

Version 0.3.3
--------------------------------------------------------------------------------

* Fix link

Version 0.3.2
--------------------------------------------------------------------------------

* Add info about autopip entry points
* Support autopip entry points to install other apps

Version 0.3.1
--------------------------------------------------------------------------------

* Prevent autopip from being uninstalled when there are other apps

Version 0.3.0
--------------------------------------------------------------------------------

* Deactivate virtualenv after getting distribution

Version 0.2.9
================================================================================

* Skip script info in non-tty

Version 0.2.8
--------------------------------------------------------------------------------

* Soft fail for auto-update via cron

Version 0.2.7
--------------------------------------------------------------------------------

* Fall back to installed-files.txt if RECORD is not found

Version 0.2.6
--------------------------------------------------------------------------------

* Get scripts via entry point or installed file record

Version 0.2.5
--------------------------------------------------------------------------------

* Add optional name filter for list command
* Fix duplicate crontab entries and provide more info when already installed
* Update readme

Version 0.2.4
--------------------------------------------------------------------------------

* Use different system vs local install paths based on permission

Version 0.2.3
--------------------------------------------------------------------------------

* Override links to /opt/apps as our apps used to be there

Version 0.2.2
--------------------------------------------------------------------------------

* Check system base for permissions

Version 0.2.1
--------------------------------------------------------------------------------

* Check log parents for system permission

Version 0.2.0
--------------------------------------------------------------------------------

* Better words for sudo use and alternative to use virtual env

Version 0.1.2
================================================================================

* Switch to use /usr/local for system installs
  
  And also add note about using sudo and security

Version 0.1.1
--------------------------------------------------------------------------------

* Sort pkg versions from PyPI index
* Update readme

Version 0.1.0
--------------------------------------------------------------------------------

* Add note to use sudo to see apps installs in /usr/local/bin
* Prepend /usr/local/bin to PATH in crontab as brew installs python3 there

Version 0.0.9
================================================================================

* Move install comment to below the sudo command

Version 0.0.8
--------------------------------------------------------------------------------

* Redirect stderr for crontab calls
* Update readme

Version 0.0.7
--------------------------------------------------------------------------------

* Add notice to use sudo on first user install

Version 0.0.6
--------------------------------------------------------------------------------

* Add example using app and installing autopip itself

Version 0.0.5
--------------------------------------------------------------------------------

* Bump version
* Always override links for autopip

Version 0.0.4
--------------------------------------------------------------------------------

* Update readme
* Add link to pip conf
* Add note on doing user install

Version 0.0.3
--------------------------------------------------------------------------------

* Update description

Version 0.0.2
--------------------------------------------------------------------------------

* Add README and set status to Beta
* Add more tests
* Add tests
* Switch to use logging to show timestamp
* Support version requirements to pin version
* Add cron job when installing
* Failure of one install should not impact the rest
* Add app alias and implement uninstall
* Implement list packages

Version 0.0.1
--------------------------------------------------------------------------------

* Add package manager and crontab
* Initial commit

Version 0.0.1
--------------------------------------------------------------------------------

* Setup project and add crontab support
* Initial commit

Version 0.0.1
--------------------------------------------------------------------------------

* Setup project
* Initial commit
