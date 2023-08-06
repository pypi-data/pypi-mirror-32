import os
import numbers

def isNumber(x):
	return isinstance(x, numbers.Number)

def TellDuration(seq):
	duration = 0
	for item in seq:
		if type(item)== tuple:
			_item = item[0]	
			if type(_item) == str: # singing
				tupleSize=len(item)
				j=0
				while j<tupleSize:
					j+=1  # by-pass lyric
					_item=item[j]
					if type(_item) == tuple: # singing note
						while j<tupleSize:
							_item = item[j]
							if type(_item) != tuple:
								break
							numCtrlPnt= len(_item)//2
							for k in range(numCtrlPnt):
								duration+=_item[k*2+1]
							j+=1
					elif isNumber(_item): # singing rap
						duration += item[j]
						j+=3
			elif isNumber(_item): # note
				duration += item[1]
	return duration

from .Catalog import Catalog
from .Catalog import PrintCatalog

from .Instrument import Instrument
from .Percussion import Percussion
from .Singer import Singer
from .Document import Document

from .Meteor import Document as MeteorDocument

from .UtauDraft import GetVoiceBank as GetVoiceBankUTAU
from .UtauDraft import UtauDraft
from .CVVCChineseConverter import CVVCChineseConverter
from .XiaYYConverter import XiaYYConverter
from .JPVCVConverter import JPVCVConverter
from .TsuroVCVConverter import TsuroVCVConverter
from .TTEnglishConverter import TTEnglishConverter
from .VCCVEnglishConverter import VCCVEnglishConverter

UTAU_VB_ROOT='UTAUVoice'
UTAU_VB_SUFFIX='_UTAU'
if os.path.isdir(UTAU_VB_ROOT):
	for item in os.listdir(UTAU_VB_ROOT):
		if os.path.isdir(UTAU_VB_ROOT+'/'+item):
			definition="""
def """+item+UTAU_VB_SUFFIX+"""(useCuda=True):
	return UtauDraft('"""+UTAU_VB_ROOT+"""/"""+item+"""',useCuda)
"""
			exec(definition)
			Catalog['Singers'] += [item+UTAU_VB_SUFFIX+' - UtauDraft']

from .SF2Instrument import ListPresets as ListPresetsSF2
from .SF2Instrument import SF2Instrument

SF2_ROOT='SF2'
if os.path.isdir(SF2_ROOT):
	for item in os.listdir(SF2_ROOT):
		sf2_path = SF2_ROOT+'/'+item
		if os.path.isfile(sf2_path) and item.endswith(".sf2"):
			name = item[0:len(item)-4]
			definition="""
def """+name+"""(preset_index):
	return SF2Instrument('"""+sf2_path+"""', preset_index)

def """+name+"""_List():
	ListPresetsSF2('"""+sf2_path+"""')
"""
			exec(definition)
			Catalog['Instruments'] += [name+' - SF2Instrument']

from .SimpleInstruments import PureSin
from .SimpleInstruments import Square
from .SimpleInstruments import Triangle
from .SimpleInstruments import Sawtooth
from .SimpleInstruments import NaivePiano
from .SimpleInstruments import BottleBlow

from .InstrumentSampler import InstrumentSampler_Single
from .InstrumentSampler import InstrumentSampler_Multi

INSTR_SAMPLE_ROOT='InstrumentSamples'
if os.path.isdir(INSTR_SAMPLE_ROOT):
	for item in os.listdir(INSTR_SAMPLE_ROOT):
		inst_path = INSTR_SAMPLE_ROOT+'/'+item
		if os.path.isfile(inst_path) and item.endswith(".wav"):
			name = item[0:len(item)-4]
			definition="""
def """+name+"""():
	return InstrumentSampler_Single('"""+inst_path+"""')
"""
			exec(definition)
			Catalog['Instruments'] += [name+' - InstrumentSampler_Single']
		elif os.path.isdir(inst_path):
			definition="""
def """+item+"""():
	return InstrumentSampler_Multi('"""+inst_path+"""')
"""
			exec(definition)
			Catalog['Instruments'] += [name+' - InstrumentSampler_Multi']

from .PercussionSampler import PercussionSampler

PERC_SAMPLE_ROOT='PercussionSamples'
if os.path.isdir(PERC_SAMPLE_ROOT):
	for item in os.listdir(PERC_SAMPLE_ROOT):
		file_path = PERC_SAMPLE_ROOT+'/'+item
		if os.path.isfile(file_path) and item.endswith(".wav"):
			name = item[0:len(item)-4]
			definition="""
def """+name+"""():
	return PercussionSampler('"""+file_path+"""')
"""
			exec(definition)
			Catalog['Percussions'] += [name+' - PercussionSampler']

