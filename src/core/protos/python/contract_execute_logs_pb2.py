# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: contract_execute_logs.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import contract_pb2 as contract__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='contract_execute_logs.proto',
  package='protos',
  syntax='proto3',
  serialized_pb=_b('\n\x1b\x63ontract_execute_logs.proto\x12\x06protos\x1a\x0e\x63ontract.proto\"\xc7\x02\n\x12\x43ontractExecuteLog\x12\x16\n\x0e\x43ontractHashId\x18\x01 \x01(\t\x12\x0e\n\x06TaskId\x18\x02 \x01(\t\x12\x11\n\tTimestamp\x18\x03 \x01(\t\x12\x0f\n\x07\x43\x61ption\x18\x04 \x01(\t\x12\r\n\x05\x43name\x18\x05 \x01(\t\x12\x12\n\nContractId\x18\x06 \x01(\t\x12\x16\n\x0e\x43ontractOwners\x18\x07 \x03(\t\x12\x35\n\x12\x43ontractSignatures\x18\x08 \x03(\x0b\x32\x19.protos.ContractSignature\x12\x15\n\rContractState\x18\t \x01(\t\x12\x12\n\nCreateTime\x18\n \x01(\t\x12\x0f\n\x07\x43reator\x18\x0b \x01(\t\x12\x13\n\x0b\x44\x65scription\x18\x0c \x01(\t\x12\x11\n\tStartTime\x18\r \x01(\t\x12\x0f\n\x07\x45ndTime\x18\x0e \x01(\t\"G\n\x13\x43ontractExecuteLogs\x12\x30\n\x0c\x43ontractLogs\x18\x01 \x03(\x0b\x32\x1a.protos.ContractExecuteLogB)\n\x14\x63om.uniledger.protosB\x11ProtoContractListb\x06proto3')
  ,
  dependencies=[contract__pb2.DESCRIPTOR,])




_CONTRACTEXECUTELOG = _descriptor.Descriptor(
  name='ContractExecuteLog',
  full_name='protos.ContractExecuteLog',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ContractHashId', full_name='protos.ContractExecuteLog.ContractHashId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='TaskId', full_name='protos.ContractExecuteLog.TaskId', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Timestamp', full_name='protos.ContractExecuteLog.Timestamp', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Caption', full_name='protos.ContractExecuteLog.Caption', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Cname', full_name='protos.ContractExecuteLog.Cname', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ContractId', full_name='protos.ContractExecuteLog.ContractId', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ContractOwners', full_name='protos.ContractExecuteLog.ContractOwners', index=6,
      number=7, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ContractSignatures', full_name='protos.ContractExecuteLog.ContractSignatures', index=7,
      number=8, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ContractState', full_name='protos.ContractExecuteLog.ContractState', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='CreateTime', full_name='protos.ContractExecuteLog.CreateTime', index=9,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Creator', full_name='protos.ContractExecuteLog.Creator', index=10,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Description', full_name='protos.ContractExecuteLog.Description', index=11,
      number=12, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='StartTime', full_name='protos.ContractExecuteLog.StartTime', index=12,
      number=13, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='EndTime', full_name='protos.ContractExecuteLog.EndTime', index=13,
      number=14, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=56,
  serialized_end=383,
)


_CONTRACTEXECUTELOGS = _descriptor.Descriptor(
  name='ContractExecuteLogs',
  full_name='protos.ContractExecuteLogs',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ContractLogs', full_name='protos.ContractExecuteLogs.ContractLogs', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=385,
  serialized_end=456,
)

_CONTRACTEXECUTELOG.fields_by_name['ContractSignatures'].message_type = contract__pb2._CONTRACTSIGNATURE
_CONTRACTEXECUTELOGS.fields_by_name['ContractLogs'].message_type = _CONTRACTEXECUTELOG
DESCRIPTOR.message_types_by_name['ContractExecuteLog'] = _CONTRACTEXECUTELOG
DESCRIPTOR.message_types_by_name['ContractExecuteLogs'] = _CONTRACTEXECUTELOGS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ContractExecuteLog = _reflection.GeneratedProtocolMessageType('ContractExecuteLog', (_message.Message,), dict(
  DESCRIPTOR = _CONTRACTEXECUTELOG,
  __module__ = 'contract_execute_logs_pb2'
  # @@protoc_insertion_point(class_scope:protos.ContractExecuteLog)
  ))
_sym_db.RegisterMessage(ContractExecuteLog)

ContractExecuteLogs = _reflection.GeneratedProtocolMessageType('ContractExecuteLogs', (_message.Message,), dict(
  DESCRIPTOR = _CONTRACTEXECUTELOGS,
  __module__ = 'contract_execute_logs_pb2'
  # @@protoc_insertion_point(class_scope:protos.ContractExecuteLogs)
  ))
_sym_db.RegisterMessage(ContractExecuteLogs)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('\n\024com.uniledger.protosB\021ProtoContractList'))
# @@protoc_insertion_point(module_scope)
