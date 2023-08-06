Python library for PHAT BEAT, a stereo DAC, AMP and VU meter with 5 input buttons for the Raspberry Pi.

Learn more: https://shop.pimoroni.com/products/phat-beat
0.1.1
-----

* Bugfix: Added sleep on clock to prevent unstable output on Pi 3B+ under load

0.1.0
-----

* Bugfix: Deferred initialisation to prevent import side-effects
* Bugfix: Implemented EOF for newer APA102 part

0.0.2
-----

* Bugfix: Corrected (reversed) order of pixels in second bar
* New: Optional threading for button handlers
* New: Added hold() to attach a handler to a button hold event

0.0.1
-----

* Initial release



