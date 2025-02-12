# -*- coding: utf-8 -*-
import time
import os
from utils.config import Config

class Logger():
	def __init__(self):
		self.prefix = Config.LOG_PATH + "/" + time.strftime('%Y-%m-%d_%H-%M-%S',time.localtime(time.time()))
		self.prefixFul = self.prefix
		self.init = False
		#self.prefix = "logs/" + time.strftime('%Y-%m-%d',time.localtime(time.time()))
		self.logfile = None
		self.logpath = ""
		self.filename = ""
		self.collectionfile = None

	def __del__(self):
		if self.init:
			pass

	def setPath(self, path):
		self.prefixFul = self.prefix + "/" + path

	def logPath(self):
		return self.logpath

	def open(self, filepath, title = None):
		try:
			self.logpath = self.prefixFul + "/" + filepath
			logdir = self.prefixFul + "/" + os.path.dirname(filepath)
			if not os.path.exists(logdir):
				os.makedirs(logdir)

			if not self.init:
				self.init = True

			self.filename = os.path.basename(self.logpath)
			self.logfile = open(self.logpath, "w")  # 打开文件
			self.logtitle = title if title else os.path.splitext(filepath)[0]
		except Exception as e:
			print("logger open: ", e)
	#write
	def print(self, str):
		try:
			strtime = time.strftime('%Y-%m-%d %H:%M:%S',time.localtime(time.time()))

			strlist = str.split("\n")
			str = ""
			for line in strlist:
				str = str + strtime + ": " + line + "\n"
			
			print(str, end='')
			if self.logfile:
				self.logfile.write(str)
		except Exception as e:
			print("logger print: ", e)

	def error(self, str):
		try:
			strtime = time.strftime('%Y-%m-%d %H:%M:%S',time.localtime(time.time()))
			str = "[ ERROR ]  " + str

			strlist = str.split("\n")
			str = ""
			for line in strlist:
				str = str + strtime + ": " + line + "\n"

			print(str, end='')
			if self.logfile:
				self.logfile.write(str)
		except Exception as e:
			print("logger error: ", e)

	def info(self, str):
		try:
			strtime = time.strftime('%Y-%m-%d %H:%M:%S',time.localtime(time.time()))
			str = "[ INFO ]  " + str

			strlist = str.split("\n")
			str = ""
			for line in strlist:
				str = str + strtime + ": " + line + "\n"

			print(str, end='')
			if self.logfile:
				self.logfile.write(str)
		except Exception as e:
			print("logger info: ", e)

	def close(self, result = None, msg = None):
		try:
			if not result is None:
				if result == "pass":
					self.print("[ OK       ] ")
					self.append_record(self.logtitle, "pass", self.filename)
				elif result == "fail":
					self.print("[ Failed   ] ")
					self.append_record(self.logtitle, "fail", self.filename)
				else:
					self.print("[ Block    ] ")
					self.append_record(self.logtitle, "block", self.filename)
			if self.logfile:
				self.logfile.close()
				self.logfile = None
		except Exception as e:
			print("logger close: ", e)

	def append_record(self, name, status, logpath, retrytimes = 0):
		filename = "collection_log"
		try:
			newfile = False
			if not os.path.exists(os.path.dirname(self.logpath) + "/" + filename + ".csv"):
				newfile = True
			self.collectionfile = open(os.path.dirname(self.logpath) + "/" + filename + ".csv", "a+")  # 打开文件
			
			if newfile:
				self.collectionfile.write("NAME,STATUS,LOG PATH\n")
			self.collectionfile.write(name + "," + status + "," + logpath + "\n")
			self.collectionfile.close()
		except Exception as e:
			print("logger append_record:", e)
			#append_record(name, status, logpath, retrytimes = retrytimes + 1)

LoggerInstance = Logger()