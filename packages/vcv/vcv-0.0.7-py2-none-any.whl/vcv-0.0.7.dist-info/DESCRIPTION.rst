Info
====
`vcv 2018-06-15`

`Author: Zhao Mingming <471106585@qq.com>`

`Copyright: This module has been placed in the public domain.`

`version:0.0.7`

Classes:
- `hand_detector`: detect the hand in the image 

Functions:

- `test()`: test function  
- `hand_detector()`:  a class
- `hand_detector.detect_hand(image)`: return the hand_number,hand position,and the confidense

How To Use This Module
======================
.. image:: funny.gif
   :height: 100px
   :width: 100px
   :alt: funny cat picture
   :align: center

1. example code:


.. code:: python
from vcv import hand_detector as hd
import cv2

hd1=hd.hand_detector()

hd1.test()
hd3=hand_detector()
imf=os.path.join(self.site_package,'test.jpg')
print imf
image=cv2.imread(imf)
print(hd3.detect_hand(image))


Refresh
========



