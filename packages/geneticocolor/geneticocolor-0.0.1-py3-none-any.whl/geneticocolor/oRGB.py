import math
import numpy as np
from random import random
"""
Color conversion from RGB/sRGB to oRGB and vice versa.
Random RGB/oRGB generation
"""

class Converter:
    """Convert colors from RGB/sRGB to oRGB and vice versa""" 
    _matrix_to_LCC = 0.0
    _matrix_to_sRGB = 0.0
    def __init__(self):
        self._matrix_to_LCC = np.matrix([[0.2990, 0.5870, 0.1140],[0.5, 0.5, -1.0],[0.866, -0.866, 0]])
        self._matrix_to_sRGB = np.matrix([[1.0, 0.1140, 0.7436],[1.0, 0.1140, -0.4111],[1.0, -0.886, 0.1663]])

    def sRGB_to_RGB(self,sR,sG,sB):
        sR = abs(sR)
        sG = abs(sG)
        sB = abs(sB)
        return [math.pow(sR,2.2),math.pow(sG,2.2),math.pow(sB,2.2)]

    def RGB_to_sRGB(self,R,G,B):
        return [math.pow(R,1.0/2.2),math.pow(G,1.0/2.2),math.pow(B,1.0/2.2)]

    def RGB_to_oRGB(self,R,G,B):
        sR,sG,sB = self.RGB_to_sRGB(R,G,B)
        return self.sRGB_to_oRGB(sR,sG,sB)

    def oRGB_to_RGB(self,L,CYB,CRG):
        sR,sG,sB = self.oRGB_to_sRGB(L,CYB,CRG)
        return self.sRGB_to_RGB(sR,sG,sB)

    def sRGB_to_oRGB(self,sR,sG,sB):
        L,C1,C2 = self._sRGB_to_LCC(sR,sG,sB)
        return self._LCC_to_LCYBCRG(L,C1,C2)

    def oRGB_to_sRGB(self,L,CYB,CRG):
        L,C1,C2 = self._LCYBCRG_to_LCC(L,CYB, CRG)
        return self._LCC_to_sRGB(L,C1,C2)

    def _sRGB_to_LCC(self,R,G,B):
        RGB = np.matrix([[R],[G],[B]])
        LCC = self._matrix_to_LCC * RGB;
        LCC = np.array(LCC.reshape(3,))[0]
        return [LCC[0], LCC[1], LCC[2]]

    def _LCC_to_sRGB(self,L,C1,C2):
        LCC = np.matrix([[L],[C1],[C2]])
        sRGB = self._matrix_to_sRGB * LCC;
        sRGB = np.array(sRGB.reshape(3,))[0]
        return [sRGB[0], sRGB[1], sRGB[2]]

    def _LCC_to_LCYBCRG(self,L,C1,C2):
        a = self._getAngleLCYBCRG(C1, C2);
        R = self._rotationMatrix(a)
        CC = R * np.matrix([[C1],[C2]])
        CC = np.array(CC.reshape(2,))[0]
        return [L, CC[0], CC[1]]

    def _LCYBCRG_to_LCC(self,L,CYB,CRG):
        a = self._getAngleLCC(CYB,CRG)
        R = self._rotationMatrix(a)
        CC = R * np.matrix([[CYB],[CRG]])
        CC = np.array(CC.reshape(2,))[0]
        return [L, CC[0], CC[1]]

    def _getAngleLCYBCRG(self,C1,C2):
        t = math.fabs(math.atan2(C2,C1))
        ot = self._ot(t)
        a = ot - t
        if C2 < 0:
            a = - a
        return a

    def _getAngleLCC(self,CYB,CRG):
        ot = math.fabs(math.atan2(CRG,CYB))
        t = self._t(ot)
        a = ot - t
        if CRG > 0:
            a = - a
        return a

    def _ot(self,t):
        if t == 0:
            ot = 0
        elif t < math.pi / 3:
            ot = (3.0/2.0) * t
        else:
            ot = (math.pi / 2) + (3.0/4.0) * (t - math.pi/3)
        return ot

    def _t(self,ot):
        if ot == 0:
            t = 0
        elif ot < math.pi / 2:
            t = (2.0/3.0) * ot
        else:
            t = (math.pi / 3) + (4.0/3.0) * (ot - math.pi/2)
        return t

    def _rotationMatrix(self, a):
        return np.matrix([[math.cos(a), -math.sin(a)],
            [math.sin(a), math.cos(a)]])

converter = Converter()

def sRGB_to_RGB(sR,sG,sB):
    return converter.sRGB_to_RGB(sR,sG,sB)

def RGB_to_sRGB(R,G,B):
    converter.RGB_to_sRGB(R,G,B)

def RGB_to_oRGB(R,G,B):
    return converter.RGB_to_oRGB(R,G,B)

def oRGB_to_RGB(L,CYB,CRG):
    return converter.oRGB_to_RGB(L,CYB,CRG)

def sRGB_to_oRGB(sR,sG,sB):
    return converter.sRGB_to_oRGB(sR,sG,sB)

def oRGB_to_sRGB(L,CYB,CRG):
    return converter.oRGB_to_sRGB(L,CYB,CRG)

def randcolor_RGB():
    R = random()
    G = random()
    B = random()
    return[R,G,B]

def randcolor_oRGB():
    R, G, B = randcolor_RGB()
    oRGB = RGB_to_oRGB(R,G,B)
    return oRGB

