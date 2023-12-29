# ScanIt
This is a public go project that will walk the specifid directory for all music, movie, and image type files and move them the the designated directory for quick easy cleanup. Please note that this does not take into consideration other program files therefore if you run this at root you may move unwanted files.


### Variables:

|   Variable     |  Definition | Required |
|------        |---            |---|
|root-dir    | The entry point of the scanner. If the root dir is selected then the entire machine will be scanned and moved... DO NOT PUT DOWN ROOT.                |No|
|image-dir    | Directory where you want to images to be moved to.           |No|
|movie-dir      | Directory where you want to movies to be moved to.          |No|
|music-dir      | Directory where you want to music to be moved to.         |No|


### Invocation
#### Linux
ScanIt -root-dir ./ScanIt/test -image-dir ./ScanIt/test/image -movie-dir ./ScanIt/test/movie -music-dir ./ScanIt/test/music
#### Windows
ScanIt.exe --root-dir .\test --image-dir ..\ScanIt\test\image --movie-dir ..\ScanIt\test\movie --music-dir ..\ScanIt\test\music