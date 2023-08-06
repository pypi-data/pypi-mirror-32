__author__ = 'kasi'
#coding=utf-8

import tempfile
import os
import re
import xml.etree.cElementTree as ET

from Core.Utils.adb_interface import AdbInterface
from Core.Info.system import SystemInfo
from Core.Action.app import LocalAction
from Core.Action.system import SystemAction
from uiautomator import device as d


shell=AdbInterface()
si=SystemInfo()
l=LocalAction()
sa=SystemAction()
PATH = lambda p: os.path.abspath(p)

class Element(object):
    """
    通过元素定位
    """
    def __init__(self):
        """
        初始化，获取系统临时文件存储目录，定义匹配数字模式
        """
        self.tempFile = tempfile.gettempdir()
        self.pattern = re.compile(r"\d+")

    """
    def __uiDump(self):

        获取当前Activity的控件树

        if int(si.getSdkVersion())>18:
            shell.SendShellCommand("uiautomator dump /data/local/tmp/uidump.xml")
        else:
            shell.SendShellCommand("uiautomator --compressed dump /data/local/tmp/uidump.xml")


    def __uidump(self):
        self.__uiDump()
        print 1
        l.pullFile("/data/local/tmp/uidump.xml",self.tempFile)
        print 2
        sa.delFile("/data/local/tmp","uidump.xml")
    """
    def __uidump(self):
        d.dump(PATH(self.tempFile + "/uidump.xml"))

    def __element(self, attrib, name):
        """
        同属性单个元素，返回单个坐标元组，(x, y)
        :args:
        - attrib - node节点中某个属性
        - name - node节点中某个属性对应的值
        """
        Xpoint = None
        Ypoint = None

        self.__uidump()
        tree = ET.ElementTree(file=PATH(self.tempFile + "/uidump.xml"))
        treeIter = tree.iter(tag="node")
        for elem in treeIter:
            if elem.attrib[attrib] == name:
                #获取元素所占区域坐标[x, y][x, y]
                bounds = elem.attrib["bounds"]

                #通过正则获取坐标列表
                coord = self.pattern.findall(bounds)

                #求取元素区域中心点坐标
                Xpoint = (int(coord[2]) - int(coord[0])) / 2.0 + int(coord[0])
                Ypoint = (int(coord[3]) - int(coord[1])) / 2.0 + int(coord[1])
                break

        if Xpoint is None or Ypoint is None:
            raise Exception("Not found this element(%s) in current activity"%name)

        return (Xpoint, Ypoint)

    def __elements(self, attrib, name):
        """
        同属性多个元素，返回坐标元组列表，[(x1, y1), (x2, y2)]
        """
        pointList = []
        self.__uidump()
        tree = ET.ElementTree(file=PATH(self.tempFile + "/uidump.xml"))
        treeIter = tree.iter(tag="node")
        for elem in treeIter:
            if elem.attrib[attrib] == name:
                bounds = elem.attrib["bounds"]
                coord = self.pattern.findall(bounds)
                Xpoint = (int(coord[2]) - int(coord[0])) / 2.0 + int(coord[0])
                Ypoint = (int(coord[3]) - int(coord[1])) / 2.0 + int(coord[1])

                #将匹配的元素区域的中心点添加进pointList中
                pointList.append((Xpoint, Ypoint))

        return pointList

    def __bound(self, attrib, name):
        """
        同属性单个元素，返回单个坐标区域元组,(x1, y1, x2, y2)
        """
        coord = []

        self.__uidump()
        tree = ET.ElementTree(file=PATH(self.tempFile + "/uidump.xml"))
        treeIter = tree.iter(tag="node")
        for elem in treeIter:
            if elem.attrib[attrib] == name:
                bounds = elem.attrib["bounds"]
                coord = self.pattern.findall(bounds)

        if not coord:
            raise Exception("Not found this element(%s) in current activity"%name)

        return (int(coord[0]), int(coord[1]), int(coord[2]), int(coord[3]))

    def __bounds(self, attrib, name):
        """
        同属性多个元素，返回坐标区域列表，[(x1, y1, x2, y2), (x3, y3, x4, y4)]
        """

        pointList = []
        self.__uidump()
        tree = ET.ElementTree(file=PATH(self.tempFile + "/uidump.xml"))
        treeIter = tree.iter(tag="node")
        for elem in treeIter:
            if elem.attrib[attrib] == name:
                bounds = elem.attrib["bounds"]
                coord = self.pattern.findall(bounds)
                pointList.append((int(coord[0]), int(coord[1]), int(coord[2]), int(coord[3])))

        return pointList

    def __checked(self, attrib, name):
        """
        返回布尔值列表
        """
        boolList = []
        self.__uidump()
        tree = ET.ElementTree(file=PATH(self.tempFile + "/uidump.xml"))
        treeIter = tree.iter(tag="node")
        for elem in treeIter:
            if elem.attrib[attrib] == name:
                checked = elem.attrib["checked"]
                if checked == "true":
                    boolList.append(True)
                else:
                    boolList.append(False)

        return boolList

    def __search(self,attrib,name):
        """
        返回布尔值
        """
        treelenth=self.getElementsNo()
        count=0
        tree = ET.ElementTree(file=PATH(self.tempFile + "/uidump.xml"))
        treeIter = tree.iter(tag="node")
        for elem in treeIter:
            count+=1
            if elem.attrib[attrib]==name:
                return True
            if count==treelenth:
                return False

    def getElementsNo(self):
        """
        获取节点数
        """
        self.__uidump()
        tree = ET.ElementTree(file=PATH(self.tempFile + "/uidump.xml"))
        treeIter = tree.iter(tag="node")
        count=0
        for x in treeIter:
            count+=1
        return count

    def findElementByContent(self, contentName):
        """
        通过元素的content-desc定位单个contentName的元素
        """
        return self.__element("content-desc",contentName)

    def findElementsByContent(self, contentName):
        """
        通过元素的content-desc定位多个相同contentName的元素
        """
        return self.__elements("content-desc",contentName)

    def findElementByName(self, name):
        """
        通过元素名称定位单个元素
        usage: findElementByName(u"设置")
        """
        return self.__element("text", name)

    def findElementsByName(self, name):
        """
        通过元素名称定位多个相同text的元素
        """
        return self.__elements("text", name)

    def findElementByClass(self, className):
        """
        通过元素类名定位单个元素
        usage: findElementByClass("android.widget.TextView")
        """
        return self.__element("class", className)

    def findElementsByClass(self, className):
        """
        通过元素类名定位多个相同class的元素
        """
        return self.__elements("class", className)

    def findElementById(self, id):
        """
        通过元素的resource-id定位单个元素
        usage: findElementsById("com.android.deskclock:id/imageview")
        """
        return self.__element("resource-id",id)

    def findElementsById(self, id):
        """
        通过元素的resource-id定位多个相同id的元素
        """
        return self.__elements("resource-id",id)


    def getElementBoundByContent(self, contentName):
        """
        通过元素描述信息获取单个元素的区域
        """
        return self.__bound("content-desc", contentName)

    def getElementBoundsByContent(self, contentName):
        """
        通过元素描述信息获取多个相同content元素的区域
        """
        return self.__bounds("content-desc", contentName)

    def getElementBoundByName(self, name):
        """
        通过元素名称获取单个元素的区域
        """
        return self.__bound("text", name)

    def getElementBoundsByName(self, name):
        """
        通过元素名称获取多个相同text元素的区域
        """
        return self.__bounds("text", name)

    def getElementBoundByClass(self, className):
        """
        通过元素类名获取单个元素的区域
        """
        return self.__bound("class", className)

    def getElementBoundsByClass(self, className):
        """
        通过元素类名获取多个相同class元素的区域
        """
        return self.__bounds("class", className)

    def getElementBoundById(self, id):
        """
        通过元素id获取单个元素的区域
        """
        return self.__bound("resource-id", id)

    def getElementBoundsById(self, id):
        """
        通过元素id获取多个相同resource-id元素的区域
        """
        return self.__bounds("resource-id", id)

    def isElementsCheckedByName(self, name):
        """
        通过元素名称判断checked的布尔值，返回布尔值列表
        """
        return self.__checked("text", name)

    def isElementsCheckedById(self, id):
        """
        通过元素id判断checked的布尔值，返回布尔值列表
        """
        return self.__checked("resource-id", id)

    def isElementsCheckedByClass(self, className):
        """
        通过元素类名判断checked的布尔值，返回布尔值列表
        """
        return self.__checked("class", className)

    def isElementsCheckedByContent(self,contentName):
        """
        通过元素描述信息判断checked的布尔值，返回布尔值列表
        """
        return self.__checked("content-desc", contentName)

    def searchForById(self,id):
        """
        通过id判断页面是否存在，返回布尔值
        """
        return self.__search("resource-id",id)

    def searchForByName(self,text):
        """
        通过文本信息判断页面是否存在，返回布尔值
        """
        return self.__search("text",text)

    def searchForByClass(self,className):
        """
        通过类名判断页面是否存在，返回布尔值
        """
        return self.__search("class",className)

    def searchForByContent(self,contentName):
        """
        通过描述信息判断页面是否存在，返回布尔值
        """
        return self.__search("content-desc",contentName)