# $Id: ai_eye.py 2018-05-25 $
# Author: Zhao Mingming <471106585@qq.com>
# Copyright: This module has been placed in the public domain.

__docformat__ = 'restructuredtext'
"""
ai_eye
version:0.0.3
ai_eye can tell u the eye's open degree from the landmarks

Functions:

- `has_closed_eye`: eye's open degree


How To Use This Module
======================


1. when u use pip install ldm==0.0.2

__docformat__ = 'restructuredtext'
from ldm import landmarks
from ai_eye import has_closed_eye
from skimage import io

imagepath="closed_eye/10.jfif"

img=io.imread(imagepath)
ldl,helptxt=landmarks(img)
print helptxt
for ld in ldl:
&nbsp;&nbsp;&nbsp;&nbsp;print has_closed_eye(ld)


2. when u use  pip install ldm==0.0.4




import  ldm
from ai_eye import has_closed_eye
from skimage import io

imagepath="closed_eye/10.jfif"
ldmer=ldm.LDM()
img=io.imread(imagepath)
ldl,facel,helptxt=ldmer.landmarks(img)
print helptxt
for ld in ldl:
&nbsp;&nbsp;&nbsp;&nbsp;print has_closed_eye(ld)
&nbsp;&nbsp;&nbsp;&nbsp;print 10*'-'
&nbsp;&nbsp;&nbsp;&nbsp;print 'nose:'
&nbsp;&nbsp;&nbsp;&nbsp;print ld['nose']
for face in facel:
&nbsp;&nbsp;&nbsp;&nbsp;print 10*'-'
&nbsp;&nbsp;&nbsp;&nbsp;print 'face:'
&nbsp;&nbsp;&nbsp;&nbsp;print face
"""
