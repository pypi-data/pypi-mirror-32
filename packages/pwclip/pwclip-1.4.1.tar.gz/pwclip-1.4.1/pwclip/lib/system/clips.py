#!/usr/bin/env python3
# -*- encoding: utf-8 -*-
#
# This file is free software by d0n <d0n@janeiskla.de>
#
# You can redistribute it and/or modify it under the terms of the GNU -
# Lesser General Public License as published by the Free Software Foundation
#
# This is distributed in the hope that it will be useful somehow.
#
# !WITHOUT ANY WARRANTY!
#
# Without even the implied warranty of MERCHANTABILITY or FITNESS FOR
# A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more
# details.
#
"""
clips - clipboard for various systems
"""
from os import name as osname

from platform import system

from time import sleep, time

from subprocess import Popen, PIPE, DEVNULL

def clips():
	"""return `copy`, `paste` as system independent functions"""
	def winclips():
		"""windows clipboards - the ugliest thing i've ever seen"""
		from ctypes import \
            windll, memmove, c_size_t, \
            sizeof, c_wchar_p, get_errno, c_wchar
		from ctypes.wintypes import \
            INT, HWND, DWORD, LPCSTR, HGLOBAL, \
            LPVOID, HINSTANCE, HMENU, BOOL, UINT, HANDLE
		from contextlib import contextmanager
		GMEM_MOVEABLE = 0x0002
		CF_UNICODETEXT = 13
		class CheckedCall(object):
			"""fancy windows callchecker"""
			def __init__(self, f):
				self.f = f
				super(CheckedCall, self).__setattr__("f", f)
			def __call__(self, *args):
				ret = self.f(*args)
				if not ret and get_errno():
					raise RuntimeError("Error calling %s"%self.f.__name__)
				return ret
			def __setattr__(self, key, value):
				setattr(self.f, key, value)
		mkwin = CheckedCall(windll.user32.CreateWindowExA)
		setattr('mkwin', 'argtypes', [
            DWORD, LPCSTR, LPCSTR, DWORD, INT, INT,
            INT, INT, HWND, HMENU, HINSTANCE, LPVOID])
		setattr('mkwin', 'restype', HWND)
		delwin = CheckedCall(windll.user32.DestroyWindow)
		setattr('delwin', 'argtypes', [HWND])
		setattr('delwin', 'restype', BOOL)
		clip = windll.user32.OpenClipboard
		setattr('clip', 'argtypes', [HWND])
		setattr('clip', 'restype', BOOL)
		clsclip = CheckedCall(windll.user32.CloseClipboard)
		setattr('clsclip', 'argtypes', [])
		setattr('clsclip', 'restype', BOOL)
		delclip = CheckedCall(windll.user32.EmptyClipboard)
		setattr('delclip', 'argtypes', [])
		setattr('delclip', 'restype', BOOL)
		getclip = CheckedCall(windll.user32.GetClipboardData)
		setattr('getclip', 'argtypes', [UINT])
		setattr('getclip', 'restype', HANDLE)
		setclip = CheckedCall(windll.user32.SetClipboardData)
		setattr('setclip', 'argtypes', [UINT, HANDLE])
		setattr('setclip', 'restype', HANDLE)
		allock = CheckedCall(windll.kernel32.GlobalAlloc)
		setattr('allock', 'argtypes', [UINT, c_size_t])
		setattr('setclip', 'restype', HGLOBAL)
		dolock = CheckedCall(windll.kernel32.GlobalLock)
		setattr('dolock', 'argtypes', [HGLOBAL])
		setattr('doclip', 'restype', LPVOID)
		unlock = CheckedCall(windll.kernel32.GlobalUnlock)
		setattr('unlock', 'argtypes', [HGLOBAL])
		setattr('unclip', 'restype', BOOL)
		@contextmanager
		def window():
			"""context that provides a valid window (hwnd)"""
			hwnd = mkwin(0, b"STATIC", None, 0, 0, 0, 0, 0,
									   None, None, None, None)
			try:
				yield hwnd
			finally:
				delwin(hwnd)
		@contextmanager
		def clipboard(hwnd):
			"""windows clipboard context manager"""
			t = time() + 0.5
			success = False
			while time() < t:
				success = clip(hwnd)
				if success:
					break
				sleep(0.01)
			if not success:
				raise Exception("could not open clipboard")
			try:
				yield
			finally:
				clsclip()
		def _copy(text, mode=None):
			"""windows copy function"""
			if mode == 'b':
				return
			with window() as hwnd:
				with clipboard(hwnd):
					delclip()
					if text:
						count = len(text) + 1
						handle = allock(GMEM_MOVEABLE, count*sizeof(c_wchar))
						locked_handle = dolock(handle)
						memmove(
						    c_wchar_p(locked_handle),
                            c_wchar_p(text), count*sizeof(c_wchar))
						unlock(handle)
						setclip(CF_UNICODETEXT, handle)
		def _paste(_=None):
			"""windows paste function"""
			with clipboard(None):
				handle = getclip(CF_UNICODETEXT)
				if not handle:
					return "", ""
			return c_wchar_p(handle).value, c_wchar_p(handle).value
		return _copy, _paste

	def osxclips():
		""""OSX clipboards"""
		def _copy(text, mode=None):
			"""osx copy function"""
			text = text if text else ''
			if mode != 'b':
				with Popen(['pbcopy'], stdin=PIPE, close_fds=True) as prc:
					prc.communicate(input=str(text).encode('utf-8'))
			return False
		def _paste(_=None):
			"""osx paste function"""
			with Popen(['pbpaste'], stdout=PIPE, close_fds=True) as prc:
				out, _ = prc.communicate()
				return out.decode('utf-8'), None
			return False
		return _copy, _paste

	def linclips():
		"""linux clipboards"""
		def _copy(text, mode='p'):
			"""linux copy function"""
			xsel = ['xsel', '-l', '/dev/null', '-i']
			text = text if text else ''
			for m in mode:
				with Popen(
                      ['xsel', '-%s'%m], stdin=PIPE, stderr=DEVNULL) as prc:
					prc.communicate(input=str(text).encode('utf-8'))
					prc.terminate()
		def _paste(mode='p'):
			"""linux paste function"""
			if mode == 'p':
				out, _ = Popen(
                   ['xsel', '-l', '/dev/null', '-o', '-p'],
                   stdout=PIPE, stderr=DEVNULL).communicate()
				ret = out.decode()
			elif mode == 'b':
				out, _ = Popen(
                   ['xsel', '-l', '/dev/null', '-o', '-b'],
                   stdout=PIPE, stderr=DEVNULL).communicate()
				ret = out.decode()
			else:
				pout, _ = Popen(
                    ['xsel', '-l', '/dev/null', '-o', '-p'],
                    stdout=PIPE, stderr=DEVNULL).communicate()
				bout, _ = Popen(
                    ['xsel', '-l', '/dev/null', '-o', '-b'],
                    stdout=PIPE, stderr=DEVNULL).communicate()
				ret = pout.decode(), bout.decode()
			return ret
		return _copy, _paste
	# decide which copy, paste functions to return [windows|mac|linux] mainly
	if osname == 'nt' or system() == 'Windows':
		return winclips()
	elif osname == 'mac' or system() == 'Darwin':
		return osxclips()
	return linclips()

copy, paste = clips()
