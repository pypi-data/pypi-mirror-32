## RoboBrakeman

Robot Framework Library for Ruby's Brakeman SAST Tool

**Supports Python 2.7.x for now**

### Install Instructions
* You need docker to run this program
* Pull the brakeman docker image: `docker pull abhaybhargav/brakeman`
* Install the RoboBrakeman Library with `pip install RoboBrakeman`
* Create a `.robot` file that includes the keywords used by RoboBrakeman Library


### Keywords

`run brakeman against source`

`| run brakeman against source  | source code path  | results path`

* source code path: where your ruby source code is located currently
* results path: where your results will be stored. An `.html` file and `.json` are generated as outputs