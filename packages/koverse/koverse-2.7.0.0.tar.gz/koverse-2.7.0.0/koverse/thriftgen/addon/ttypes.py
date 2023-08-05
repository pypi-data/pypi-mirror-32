#
# Autogenerated by Thrift Compiler (0.9.3)
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#
#  options string: py
#

from thrift.Thrift import TType, TMessageType, TException, TApplicationException
import koverse.thriftgen.ttypes
import koverse.thriftgen.security.ttypes


from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol, TProtocol
try:
  from thrift.protocol import fastbinary
except:
  fastbinary = None



class TAddOn:
  """
  Attributes:
   - id
   - type
   - fileName
   - disabled
   - errorsDuringInspection
   - installTimestamp
   - disabledTimestamp
   - crc32
   - responsibleUserId
   - transformTypeIds
   - sourceTypeDescriptionIds
   - sinkTypeDescriptionIds
   - applicationDescriptionIds
   - securityLabelParserDescriptionIds
   - displayName
   - importTransformDescriptionIds
   - exportTransformDescriptionIds
   - exportFileFormatDescriptionIds
  """

  thrift_spec = (
    None, # 0
    (1, TType.I64, 'id', None, None, ), # 1
    (2, TType.STRING, 'type', None, None, ), # 2
    (3, TType.STRING, 'fileName', None, None, ), # 3
    (4, TType.BOOL, 'disabled', None, None, ), # 4
    (5, TType.STRING, 'errorsDuringInspection', None, None, ), # 5
    (6, TType.I64, 'installTimestamp', None, None, ), # 6
    (7, TType.I64, 'disabledTimestamp', None, None, ), # 7
    (8, TType.I64, 'crc32', None, None, ), # 8
    (9, TType.I64, 'responsibleUserId', None, None, ), # 9
    (10, TType.LIST, 'transformTypeIds', (TType.I64,None), None, ), # 10
    (11, TType.LIST, 'sourceTypeDescriptionIds', (TType.I64,None), None, ), # 11
    (12, TType.LIST, 'sinkTypeDescriptionIds', (TType.I64,None), None, ), # 12
    (13, TType.LIST, 'applicationDescriptionIds', (TType.I64,None), None, ), # 13
    (14, TType.LIST, 'securityLabelParserDescriptionIds', (TType.I64,None), None, ), # 14
    (15, TType.STRING, 'displayName', None, None, ), # 15
    (16, TType.LIST, 'importTransformDescriptionIds', (TType.I64,None), None, ), # 16
    (17, TType.LIST, 'exportTransformDescriptionIds', (TType.I64,None), None, ), # 17
    (18, TType.LIST, 'exportFileFormatDescriptionIds', (TType.I64,None), None, ), # 18
  )

  def __init__(self, id=None, type=None, fileName=None, disabled=None, errorsDuringInspection=None, installTimestamp=None, disabledTimestamp=None, crc32=None, responsibleUserId=None, transformTypeIds=None, sourceTypeDescriptionIds=None, sinkTypeDescriptionIds=None, applicationDescriptionIds=None, securityLabelParserDescriptionIds=None, displayName=None, importTransformDescriptionIds=None, exportTransformDescriptionIds=None, exportFileFormatDescriptionIds=None,):
    self.id = id
    self.type = type
    self.fileName = fileName
    self.disabled = disabled
    self.errorsDuringInspection = errorsDuringInspection
    self.installTimestamp = installTimestamp
    self.disabledTimestamp = disabledTimestamp
    self.crc32 = crc32
    self.responsibleUserId = responsibleUserId
    self.transformTypeIds = transformTypeIds
    self.sourceTypeDescriptionIds = sourceTypeDescriptionIds
    self.sinkTypeDescriptionIds = sinkTypeDescriptionIds
    self.applicationDescriptionIds = applicationDescriptionIds
    self.securityLabelParserDescriptionIds = securityLabelParserDescriptionIds
    self.displayName = displayName
    self.importTransformDescriptionIds = importTransformDescriptionIds
    self.exportTransformDescriptionIds = exportTransformDescriptionIds
    self.exportFileFormatDescriptionIds = exportFileFormatDescriptionIds

  def read(self, iprot):
    if iprot.__class__ == TBinaryProtocol.TBinaryProtocolAccelerated and isinstance(iprot.trans, TTransport.CReadableTransport) and self.thrift_spec is not None and fastbinary is not None:
      fastbinary.decode_binary(self, iprot.trans, (self.__class__, self.thrift_spec))
      return
    iprot.readStructBegin()
    while True:
      (fname, ftype, fid) = iprot.readFieldBegin()
      if ftype == TType.STOP:
        break
      if fid == 1:
        if ftype == TType.I64:
          self.id = iprot.readI64()
        else:
          iprot.skip(ftype)
      elif fid == 2:
        if ftype == TType.STRING:
          self.type = iprot.readString()
        else:
          iprot.skip(ftype)
      elif fid == 3:
        if ftype == TType.STRING:
          self.fileName = iprot.readString()
        else:
          iprot.skip(ftype)
      elif fid == 4:
        if ftype == TType.BOOL:
          self.disabled = iprot.readBool()
        else:
          iprot.skip(ftype)
      elif fid == 5:
        if ftype == TType.STRING:
          self.errorsDuringInspection = iprot.readString()
        else:
          iprot.skip(ftype)
      elif fid == 6:
        if ftype == TType.I64:
          self.installTimestamp = iprot.readI64()
        else:
          iprot.skip(ftype)
      elif fid == 7:
        if ftype == TType.I64:
          self.disabledTimestamp = iprot.readI64()
        else:
          iprot.skip(ftype)
      elif fid == 8:
        if ftype == TType.I64:
          self.crc32 = iprot.readI64()
        else:
          iprot.skip(ftype)
      elif fid == 9:
        if ftype == TType.I64:
          self.responsibleUserId = iprot.readI64()
        else:
          iprot.skip(ftype)
      elif fid == 10:
        if ftype == TType.LIST:
          self.transformTypeIds = []
          (_etype3, _size0) = iprot.readListBegin()
          for _i4 in xrange(_size0):
            _elem5 = iprot.readI64()
            self.transformTypeIds.append(_elem5)
          iprot.readListEnd()
        else:
          iprot.skip(ftype)
      elif fid == 11:
        if ftype == TType.LIST:
          self.sourceTypeDescriptionIds = []
          (_etype9, _size6) = iprot.readListBegin()
          for _i10 in xrange(_size6):
            _elem11 = iprot.readI64()
            self.sourceTypeDescriptionIds.append(_elem11)
          iprot.readListEnd()
        else:
          iprot.skip(ftype)
      elif fid == 12:
        if ftype == TType.LIST:
          self.sinkTypeDescriptionIds = []
          (_etype15, _size12) = iprot.readListBegin()
          for _i16 in xrange(_size12):
            _elem17 = iprot.readI64()
            self.sinkTypeDescriptionIds.append(_elem17)
          iprot.readListEnd()
        else:
          iprot.skip(ftype)
      elif fid == 13:
        if ftype == TType.LIST:
          self.applicationDescriptionIds = []
          (_etype21, _size18) = iprot.readListBegin()
          for _i22 in xrange(_size18):
            _elem23 = iprot.readI64()
            self.applicationDescriptionIds.append(_elem23)
          iprot.readListEnd()
        else:
          iprot.skip(ftype)
      elif fid == 14:
        if ftype == TType.LIST:
          self.securityLabelParserDescriptionIds = []
          (_etype27, _size24) = iprot.readListBegin()
          for _i28 in xrange(_size24):
            _elem29 = iprot.readI64()
            self.securityLabelParserDescriptionIds.append(_elem29)
          iprot.readListEnd()
        else:
          iprot.skip(ftype)
      elif fid == 15:
        if ftype == TType.STRING:
          self.displayName = iprot.readString()
        else:
          iprot.skip(ftype)
      elif fid == 16:
        if ftype == TType.LIST:
          self.importTransformDescriptionIds = []
          (_etype33, _size30) = iprot.readListBegin()
          for _i34 in xrange(_size30):
            _elem35 = iprot.readI64()
            self.importTransformDescriptionIds.append(_elem35)
          iprot.readListEnd()
        else:
          iprot.skip(ftype)
      elif fid == 17:
        if ftype == TType.LIST:
          self.exportTransformDescriptionIds = []
          (_etype39, _size36) = iprot.readListBegin()
          for _i40 in xrange(_size36):
            _elem41 = iprot.readI64()
            self.exportTransformDescriptionIds.append(_elem41)
          iprot.readListEnd()
        else:
          iprot.skip(ftype)
      elif fid == 18:
        if ftype == TType.LIST:
          self.exportFileFormatDescriptionIds = []
          (_etype45, _size42) = iprot.readListBegin()
          for _i46 in xrange(_size42):
            _elem47 = iprot.readI64()
            self.exportFileFormatDescriptionIds.append(_elem47)
          iprot.readListEnd()
        else:
          iprot.skip(ftype)
      else:
        iprot.skip(ftype)
      iprot.readFieldEnd()
    iprot.readStructEnd()

  def write(self, oprot):
    if oprot.__class__ == TBinaryProtocol.TBinaryProtocolAccelerated and self.thrift_spec is not None and fastbinary is not None:
      oprot.trans.write(fastbinary.encode_binary(self, (self.__class__, self.thrift_spec)))
      return
    oprot.writeStructBegin('TAddOn')
    if self.id is not None:
      oprot.writeFieldBegin('id', TType.I64, 1)
      oprot.writeI64(self.id)
      oprot.writeFieldEnd()
    if self.type is not None:
      oprot.writeFieldBegin('type', TType.STRING, 2)
      oprot.writeString(self.type)
      oprot.writeFieldEnd()
    if self.fileName is not None:
      oprot.writeFieldBegin('fileName', TType.STRING, 3)
      oprot.writeString(self.fileName)
      oprot.writeFieldEnd()
    if self.disabled is not None:
      oprot.writeFieldBegin('disabled', TType.BOOL, 4)
      oprot.writeBool(self.disabled)
      oprot.writeFieldEnd()
    if self.errorsDuringInspection is not None:
      oprot.writeFieldBegin('errorsDuringInspection', TType.STRING, 5)
      oprot.writeString(self.errorsDuringInspection)
      oprot.writeFieldEnd()
    if self.installTimestamp is not None:
      oprot.writeFieldBegin('installTimestamp', TType.I64, 6)
      oprot.writeI64(self.installTimestamp)
      oprot.writeFieldEnd()
    if self.disabledTimestamp is not None:
      oprot.writeFieldBegin('disabledTimestamp', TType.I64, 7)
      oprot.writeI64(self.disabledTimestamp)
      oprot.writeFieldEnd()
    if self.crc32 is not None:
      oprot.writeFieldBegin('crc32', TType.I64, 8)
      oprot.writeI64(self.crc32)
      oprot.writeFieldEnd()
    if self.responsibleUserId is not None:
      oprot.writeFieldBegin('responsibleUserId', TType.I64, 9)
      oprot.writeI64(self.responsibleUserId)
      oprot.writeFieldEnd()
    if self.transformTypeIds is not None:
      oprot.writeFieldBegin('transformTypeIds', TType.LIST, 10)
      oprot.writeListBegin(TType.I64, len(self.transformTypeIds))
      for iter48 in self.transformTypeIds:
        oprot.writeI64(iter48)
      oprot.writeListEnd()
      oprot.writeFieldEnd()
    if self.sourceTypeDescriptionIds is not None:
      oprot.writeFieldBegin('sourceTypeDescriptionIds', TType.LIST, 11)
      oprot.writeListBegin(TType.I64, len(self.sourceTypeDescriptionIds))
      for iter49 in self.sourceTypeDescriptionIds:
        oprot.writeI64(iter49)
      oprot.writeListEnd()
      oprot.writeFieldEnd()
    if self.sinkTypeDescriptionIds is not None:
      oprot.writeFieldBegin('sinkTypeDescriptionIds', TType.LIST, 12)
      oprot.writeListBegin(TType.I64, len(self.sinkTypeDescriptionIds))
      for iter50 in self.sinkTypeDescriptionIds:
        oprot.writeI64(iter50)
      oprot.writeListEnd()
      oprot.writeFieldEnd()
    if self.applicationDescriptionIds is not None:
      oprot.writeFieldBegin('applicationDescriptionIds', TType.LIST, 13)
      oprot.writeListBegin(TType.I64, len(self.applicationDescriptionIds))
      for iter51 in self.applicationDescriptionIds:
        oprot.writeI64(iter51)
      oprot.writeListEnd()
      oprot.writeFieldEnd()
    if self.securityLabelParserDescriptionIds is not None:
      oprot.writeFieldBegin('securityLabelParserDescriptionIds', TType.LIST, 14)
      oprot.writeListBegin(TType.I64, len(self.securityLabelParserDescriptionIds))
      for iter52 in self.securityLabelParserDescriptionIds:
        oprot.writeI64(iter52)
      oprot.writeListEnd()
      oprot.writeFieldEnd()
    if self.displayName is not None:
      oprot.writeFieldBegin('displayName', TType.STRING, 15)
      oprot.writeString(self.displayName)
      oprot.writeFieldEnd()
    if self.importTransformDescriptionIds is not None:
      oprot.writeFieldBegin('importTransformDescriptionIds', TType.LIST, 16)
      oprot.writeListBegin(TType.I64, len(self.importTransformDescriptionIds))
      for iter53 in self.importTransformDescriptionIds:
        oprot.writeI64(iter53)
      oprot.writeListEnd()
      oprot.writeFieldEnd()
    if self.exportTransformDescriptionIds is not None:
      oprot.writeFieldBegin('exportTransformDescriptionIds', TType.LIST, 17)
      oprot.writeListBegin(TType.I64, len(self.exportTransformDescriptionIds))
      for iter54 in self.exportTransformDescriptionIds:
        oprot.writeI64(iter54)
      oprot.writeListEnd()
      oprot.writeFieldEnd()
    if self.exportFileFormatDescriptionIds is not None:
      oprot.writeFieldBegin('exportFileFormatDescriptionIds', TType.LIST, 18)
      oprot.writeListBegin(TType.I64, len(self.exportFileFormatDescriptionIds))
      for iter55 in self.exportFileFormatDescriptionIds:
        oprot.writeI64(iter55)
      oprot.writeListEnd()
      oprot.writeFieldEnd()
    oprot.writeFieldStop()
    oprot.writeStructEnd()

  def validate(self):
    return


  def __hash__(self):
    value = 17
    value = (value * 31) ^ hash(self.id)
    value = (value * 31) ^ hash(self.type)
    value = (value * 31) ^ hash(self.fileName)
    value = (value * 31) ^ hash(self.disabled)
    value = (value * 31) ^ hash(self.errorsDuringInspection)
    value = (value * 31) ^ hash(self.installTimestamp)
    value = (value * 31) ^ hash(self.disabledTimestamp)
    value = (value * 31) ^ hash(self.crc32)
    value = (value * 31) ^ hash(self.responsibleUserId)
    value = (value * 31) ^ hash(self.transformTypeIds)
    value = (value * 31) ^ hash(self.sourceTypeDescriptionIds)
    value = (value * 31) ^ hash(self.sinkTypeDescriptionIds)
    value = (value * 31) ^ hash(self.applicationDescriptionIds)
    value = (value * 31) ^ hash(self.securityLabelParserDescriptionIds)
    value = (value * 31) ^ hash(self.displayName)
    value = (value * 31) ^ hash(self.importTransformDescriptionIds)
    value = (value * 31) ^ hash(self.exportTransformDescriptionIds)
    value = (value * 31) ^ hash(self.exportFileFormatDescriptionIds)
    return value

  def __repr__(self):
    L = ['%s=%r' % (key, value)
      for key, value in self.__dict__.iteritems()]
    return '%s(%s)' % (self.__class__.__name__, ', '.join(L))

  def __eq__(self, other):
    return isinstance(other, self.__class__) and self.__dict__ == other.__dict__

  def __ne__(self, other):
    return not (self == other)
