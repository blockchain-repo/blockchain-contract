# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: contract.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='contract.proto',
  package='protos',
  syntax='proto3',
  serialized_pb=_b('\n\x0e\x63ontract.proto\x12\x06protos\"R\n\x11\x43ontractSignature\x12\x13\n\x0bOwnerPubkey\x18\x01 \x01(\t\x12\x11\n\tSignature\x18\x02 \x01(\t\x12\x15\n\rSignTimestamp\x18\x03 \x01(\t\"\x84\x01\n\rContractAsset\x12\x0f\n\x07\x41ssetId\x18\x01 \x01(\t\x12\x0c\n\x04Name\x18\x02 \x01(\t\x12\x0f\n\x07\x43\x61ption\x18\x03 \x01(\t\x12\x13\n\x0b\x44\x65scription\x18\x04 \x01(\t\x12\x0c\n\x04Unit\x18\x05 \x01(\t\x12\x0e\n\x06\x41mount\x18\x06 \x01(\x02\x12\x10\n\x08MetaData\x18\x07 \x01(\x0c\"2\n\x10\x45xpressionResult\x12\x10\n\x08Messsage\x18\x01 \x01(\t\x12\x0c\n\x04\x43ode\x18\x02 \x01(\t\"\xb9\x01\n\x14\x43omponentsExpression\x12\r\n\x05\x43name\x18\x01 \x01(\t\x12\r\n\x05\x43type\x18\x02 \x01(\t\x12\x0f\n\x07\x43\x61ption\x18\x03 \x01(\t\x12\x13\n\x0b\x44\x65scription\x18\x04 \x01(\t\x12\x15\n\rExpressionStr\x18\x05 \x01(\t\x12\x32\n\x10\x45xpressionResult\x18\x06 \x01(\x0b\x32\x18.protos.ExpressionResult\x12\x12\n\nLogicValue\x18\x07 \x01(\t\"\xba\x02\n\rComponentData\x12\r\n\x05\x43name\x18\x01 \x01(\t\x12\r\n\x05\x43type\x18\x02 \x01(\t\x12\x0f\n\x07\x43\x61ption\x18\x03 \x01(\t\x12\x13\n\x0b\x44\x65scription\x18\x04 \x01(\t\x12\x12\n\nModifyDate\x18\x05 \x01(\t\x12\x14\n\x0cHardConvType\x18\x06 \x01(\t\x12%\n\x06Parent\x18\x08 \x01(\x0b\x32\x15.protos.ComponentData\x12\x11\n\tMandatory\x18\t \x01(\x08\x12\x0c\n\x04Unit\x18\x0b \x01(\t\x12\x33\n\x07Options\x18\r \x03(\x0b\x32\".protos.ComponentData.OptionsEntry\x12\x0e\n\x06\x46ormat\x18\x0f \x01(\t\x1a.\n\x0cOptionsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\x05:\x02\x38\x01\"\xd6\x04\n\x11\x43ontractComponent\x12\r\n\x05\x43name\x18\x01 \x01(\t\x12\r\n\x05\x43type\x18\x02 \x01(\t\x12\x0f\n\x07\x43\x61ption\x18\x03 \x01(\t\x12\x13\n\x0b\x44\x65scription\x18\x04 \x01(\t\x12\r\n\x05State\x18\x05 \x01(\t\x12\x32\n\x0cPreCondition\x18\x06 \x03(\x0b\x32\x1c.protos.ComponentsExpression\x12\x37\n\x11\x43ompleteCondition\x18\x07 \x03(\x0b\x32\x1c.protos.ComponentsExpression\x12\x36\n\x10\x44isgardCondition\x18\x08 \x03(\x0b\x32\x1c.protos.ComponentsExpression\x12\x11\n\tNextTasks\x18\t \x03(\t\x12\'\n\x08\x44\x61taList\x18\n \x03(\x0b\x32\x15.protos.ComponentData\x12\x43\n\x1d\x44\x61taValueSetterExpressionList\x18\x0b \x03(\x0b\x32\x1c.protos.ComponentsExpression\x12\x30\n\rCandidateList\x18\x0c \x01(\x0b\x32\x19.protos.ContractComponent\x12\x31\n\x0e\x44\x65\x63isionResult\x18\r \x01(\x0b\x32\x19.protos.ContractComponent\x12\x10\n\x08TaskList\x18\x0e \x03(\t\x12\x18\n\x10SupportArguments\x18\x0f \x03(\t\x12\x18\n\x10\x41gainstArguments\x18\x10 \x03(\t\x12\x0f\n\x07Support\x18\x11 \x01(\x05\x12\x0c\n\x04Text\x18\x12 \x03(\t\"\xfc\x02\n\x0c\x43ontractBody\x12\x12\n\nContractId\x18\x01 \x01(\t\x12\r\n\x05\x43name\x18\x02 \x01(\t\x12\r\n\x05\x43type\x18\x03 \x01(\t\x12\x0f\n\x07\x43\x61ption\x18\x04 \x01(\t\x12\x13\n\x0b\x44\x65scription\x18\x05 \x01(\t\x12\x15\n\rContractState\x18\x06 \x01(\t\x12\x0f\n\x07\x43reator\x18\x07 \x01(\t\x12\x13\n\x0b\x43reatorTime\x18\x08 \x01(\t\x12\x11\n\tStartTime\x18\t \x01(\t\x12\x0f\n\x07\x45ndTime\x18\n \x01(\t\x12\x16\n\x0e\x43ontractOwners\x18\x0b \x03(\t\x12-\n\x0e\x43ontractAssets\x18\x0c \x03(\x0b\x32\x15.protos.ContractAsset\x12\x35\n\x12\x43ontractSignatures\x18\r \x03(\x0b\x32\x19.protos.ContractSignature\x12\x35\n\x12\x43ontractComponents\x18\x0e \x03(\x0b\x32\x19.protos.ContractComponent\"3\n\x0c\x43ontractHead\x12\x12\n\nMainPubkey\x18\x01 \x01(\t\x12\x0f\n\x07Version\x18\x02 \x01(\x05\"n\n\x08\x43ontract\x12\n\n\x02id\x18\x01 \x01(\t\x12*\n\x0c\x43ontractHead\x18\x02 \x01(\x0b\x32\x14.protos.ContractHead\x12*\n\x0c\x43ontractBody\x18\x03 \x01(\x0b\x32\x14.protos.ContractBodyB%\n\x14\x63om.uniledger.protosB\rProtoContractb\x06proto3')
)
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_CONTRACTSIGNATURE = _descriptor.Descriptor(
  name='ContractSignature',
  full_name='protos.ContractSignature',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='OwnerPubkey', full_name='protos.ContractSignature.OwnerPubkey', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Signature', full_name='protos.ContractSignature.Signature', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='SignTimestamp', full_name='protos.ContractSignature.SignTimestamp', index=2,
      number=3, type=9, cpp_type=9, label=1,
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
  serialized_start=26,
  serialized_end=108,
)


_CONTRACTASSET = _descriptor.Descriptor(
  name='ContractAsset',
  full_name='protos.ContractAsset',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='AssetId', full_name='protos.ContractAsset.AssetId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Name', full_name='protos.ContractAsset.Name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Caption', full_name='protos.ContractAsset.Caption', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Description', full_name='protos.ContractAsset.Description', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Unit', full_name='protos.ContractAsset.Unit', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Amount', full_name='protos.ContractAsset.Amount', index=5,
      number=6, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='MetaData', full_name='protos.ContractAsset.MetaData', index=6,
      number=7, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
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
  serialized_start=111,
  serialized_end=243,
)


_EXPRESSIONRESULT = _descriptor.Descriptor(
  name='ExpressionResult',
  full_name='protos.ExpressionResult',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Messsage', full_name='protos.ExpressionResult.Messsage', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Code', full_name='protos.ExpressionResult.Code', index=1,
      number=2, type=9, cpp_type=9, label=1,
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
  serialized_start=245,
  serialized_end=295,
)


_COMPONENTSEXPRESSION = _descriptor.Descriptor(
  name='ComponentsExpression',
  full_name='protos.ComponentsExpression',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Cname', full_name='protos.ComponentsExpression.Cname', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Ctype', full_name='protos.ComponentsExpression.Ctype', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Caption', full_name='protos.ComponentsExpression.Caption', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Description', full_name='protos.ComponentsExpression.Description', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ExpressionStr', full_name='protos.ComponentsExpression.ExpressionStr', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ExpressionResult', full_name='protos.ComponentsExpression.ExpressionResult', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='LogicValue', full_name='protos.ComponentsExpression.LogicValue', index=6,
      number=7, type=9, cpp_type=9, label=1,
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
  serialized_start=298,
  serialized_end=483,
)


_COMPONENTDATA_OPTIONSENTRY = _descriptor.Descriptor(
  name='OptionsEntry',
  full_name='protos.ComponentData.OptionsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='protos.ComponentData.OptionsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='value', full_name='protos.ComponentData.OptionsEntry.value', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=754,
  serialized_end=800,
)

_COMPONENTDATA = _descriptor.Descriptor(
  name='ComponentData',
  full_name='protos.ComponentData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Cname', full_name='protos.ComponentData.Cname', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Ctype', full_name='protos.ComponentData.Ctype', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Caption', full_name='protos.ComponentData.Caption', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Description', full_name='protos.ComponentData.Description', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ModifyDate', full_name='protos.ComponentData.ModifyDate', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='HardConvType', full_name='protos.ComponentData.HardConvType', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Parent', full_name='protos.ComponentData.Parent', index=6,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Mandatory', full_name='protos.ComponentData.Mandatory', index=7,
      number=9, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Unit', full_name='protos.ComponentData.Unit', index=8,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Options', full_name='protos.ComponentData.Options', index=9,
      number=13, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Format', full_name='protos.ComponentData.Format', index=10,
      number=15, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[_COMPONENTDATA_OPTIONSENTRY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=486,
  serialized_end=800,
)


_CONTRACTCOMPONENT = _descriptor.Descriptor(
  name='ContractComponent',
  full_name='protos.ContractComponent',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Cname', full_name='protos.ContractComponent.Cname', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Ctype', full_name='protos.ContractComponent.Ctype', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Caption', full_name='protos.ContractComponent.Caption', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Description', full_name='protos.ContractComponent.Description', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='State', full_name='protos.ContractComponent.State', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='PreCondition', full_name='protos.ContractComponent.PreCondition', index=5,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='CompleteCondition', full_name='protos.ContractComponent.CompleteCondition', index=6,
      number=7, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='DisgardCondition', full_name='protos.ContractComponent.DisgardCondition', index=7,
      number=8, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='NextTasks', full_name='protos.ContractComponent.NextTasks', index=8,
      number=9, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='DataList', full_name='protos.ContractComponent.DataList', index=9,
      number=10, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='DataValueSetterExpressionList', full_name='protos.ContractComponent.DataValueSetterExpressionList', index=10,
      number=11, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='CandidateList', full_name='protos.ContractComponent.CandidateList', index=11,
      number=12, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='DecisionResult', full_name='protos.ContractComponent.DecisionResult', index=12,
      number=13, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='TaskList', full_name='protos.ContractComponent.TaskList', index=13,
      number=14, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='SupportArguments', full_name='protos.ContractComponent.SupportArguments', index=14,
      number=15, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='AgainstArguments', full_name='protos.ContractComponent.AgainstArguments', index=15,
      number=16, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Support', full_name='protos.ContractComponent.Support', index=16,
      number=17, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Text', full_name='protos.ContractComponent.Text', index=17,
      number=18, type=9, cpp_type=9, label=3,
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
  serialized_start=803,
  serialized_end=1401,
)


_CONTRACTBODY = _descriptor.Descriptor(
  name='ContractBody',
  full_name='protos.ContractBody',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ContractId', full_name='protos.ContractBody.ContractId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Cname', full_name='protos.ContractBody.Cname', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Ctype', full_name='protos.ContractBody.Ctype', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Caption', full_name='protos.ContractBody.Caption', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Description', full_name='protos.ContractBody.Description', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ContractState', full_name='protos.ContractBody.ContractState', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Creator', full_name='protos.ContractBody.Creator', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='CreatorTime', full_name='protos.ContractBody.CreatorTime', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='StartTime', full_name='protos.ContractBody.StartTime', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='EndTime', full_name='protos.ContractBody.EndTime', index=9,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ContractOwners', full_name='protos.ContractBody.ContractOwners', index=10,
      number=11, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ContractAssets', full_name='protos.ContractBody.ContractAssets', index=11,
      number=12, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ContractSignatures', full_name='protos.ContractBody.ContractSignatures', index=12,
      number=13, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ContractComponents', full_name='protos.ContractBody.ContractComponents', index=13,
      number=14, type=11, cpp_type=10, label=3,
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
  serialized_start=1404,
  serialized_end=1784,
)


_CONTRACTHEAD = _descriptor.Descriptor(
  name='ContractHead',
  full_name='protos.ContractHead',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='MainPubkey', full_name='protos.ContractHead.MainPubkey', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Version', full_name='protos.ContractHead.Version', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=1786,
  serialized_end=1837,
)


_CONTRACT = _descriptor.Descriptor(
  name='Contract',
  full_name='protos.Contract',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='protos.Contract.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ContractHead', full_name='protos.Contract.ContractHead', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ContractBody', full_name='protos.Contract.ContractBody', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=1839,
  serialized_end=1949,
)

_COMPONENTSEXPRESSION.fields_by_name['ExpressionResult'].message_type = _EXPRESSIONRESULT
_COMPONENTDATA_OPTIONSENTRY.containing_type = _COMPONENTDATA
_COMPONENTDATA.fields_by_name['Parent'].message_type = _COMPONENTDATA
_COMPONENTDATA.fields_by_name['Options'].message_type = _COMPONENTDATA_OPTIONSENTRY
_CONTRACTCOMPONENT.fields_by_name['PreCondition'].message_type = _COMPONENTSEXPRESSION
_CONTRACTCOMPONENT.fields_by_name['CompleteCondition'].message_type = _COMPONENTSEXPRESSION
_CONTRACTCOMPONENT.fields_by_name['DisgardCondition'].message_type = _COMPONENTSEXPRESSION
_CONTRACTCOMPONENT.fields_by_name['DataList'].message_type = _COMPONENTDATA
_CONTRACTCOMPONENT.fields_by_name['DataValueSetterExpressionList'].message_type = _COMPONENTSEXPRESSION
_CONTRACTCOMPONENT.fields_by_name['CandidateList'].message_type = _CONTRACTCOMPONENT
_CONTRACTCOMPONENT.fields_by_name['DecisionResult'].message_type = _CONTRACTCOMPONENT
_CONTRACTBODY.fields_by_name['ContractAssets'].message_type = _CONTRACTASSET
_CONTRACTBODY.fields_by_name['ContractSignatures'].message_type = _CONTRACTSIGNATURE
_CONTRACTBODY.fields_by_name['ContractComponents'].message_type = _CONTRACTCOMPONENT
_CONTRACT.fields_by_name['ContractHead'].message_type = _CONTRACTHEAD
_CONTRACT.fields_by_name['ContractBody'].message_type = _CONTRACTBODY
DESCRIPTOR.message_types_by_name['ContractSignature'] = _CONTRACTSIGNATURE
DESCRIPTOR.message_types_by_name['ContractAsset'] = _CONTRACTASSET
DESCRIPTOR.message_types_by_name['ExpressionResult'] = _EXPRESSIONRESULT
DESCRIPTOR.message_types_by_name['ComponentsExpression'] = _COMPONENTSEXPRESSION
DESCRIPTOR.message_types_by_name['ComponentData'] = _COMPONENTDATA
DESCRIPTOR.message_types_by_name['ContractComponent'] = _CONTRACTCOMPONENT
DESCRIPTOR.message_types_by_name['ContractBody'] = _CONTRACTBODY
DESCRIPTOR.message_types_by_name['ContractHead'] = _CONTRACTHEAD
DESCRIPTOR.message_types_by_name['Contract'] = _CONTRACT

ContractSignature = _reflection.GeneratedProtocolMessageType('ContractSignature', (_message.Message,), dict(
  DESCRIPTOR = _CONTRACTSIGNATURE,
  __module__ = 'contract_pb2'
  # @@protoc_insertion_point(class_scope:protos.ContractSignature)
  ))
_sym_db.RegisterMessage(ContractSignature)

ContractAsset = _reflection.GeneratedProtocolMessageType('ContractAsset', (_message.Message,), dict(
  DESCRIPTOR = _CONTRACTASSET,
  __module__ = 'contract_pb2'
  # @@protoc_insertion_point(class_scope:protos.ContractAsset)
  ))
_sym_db.RegisterMessage(ContractAsset)

ExpressionResult = _reflection.GeneratedProtocolMessageType('ExpressionResult', (_message.Message,), dict(
  DESCRIPTOR = _EXPRESSIONRESULT,
  __module__ = 'contract_pb2'
  # @@protoc_insertion_point(class_scope:protos.ExpressionResult)
  ))
_sym_db.RegisterMessage(ExpressionResult)

ComponentsExpression = _reflection.GeneratedProtocolMessageType('ComponentsExpression', (_message.Message,), dict(
  DESCRIPTOR = _COMPONENTSEXPRESSION,
  __module__ = 'contract_pb2'
  # @@protoc_insertion_point(class_scope:protos.ComponentsExpression)
  ))
_sym_db.RegisterMessage(ComponentsExpression)

ComponentData = _reflection.GeneratedProtocolMessageType('ComponentData', (_message.Message,), dict(

  OptionsEntry = _reflection.GeneratedProtocolMessageType('OptionsEntry', (_message.Message,), dict(
    DESCRIPTOR = _COMPONENTDATA_OPTIONSENTRY,
    __module__ = 'contract_pb2'
    # @@protoc_insertion_point(class_scope:protos.ComponentData.OptionsEntry)
    ))
  ,
  DESCRIPTOR = _COMPONENTDATA,
  __module__ = 'contract_pb2'
  # @@protoc_insertion_point(class_scope:protos.ComponentData)
  ))
_sym_db.RegisterMessage(ComponentData)
_sym_db.RegisterMessage(ComponentData.OptionsEntry)

ContractComponent = _reflection.GeneratedProtocolMessageType('ContractComponent', (_message.Message,), dict(
  DESCRIPTOR = _CONTRACTCOMPONENT,
  __module__ = 'contract_pb2'
  # @@protoc_insertion_point(class_scope:protos.ContractComponent)
  ))
_sym_db.RegisterMessage(ContractComponent)

ContractBody = _reflection.GeneratedProtocolMessageType('ContractBody', (_message.Message,), dict(
  DESCRIPTOR = _CONTRACTBODY,
  __module__ = 'contract_pb2'
  # @@protoc_insertion_point(class_scope:protos.ContractBody)
  ))
_sym_db.RegisterMessage(ContractBody)

ContractHead = _reflection.GeneratedProtocolMessageType('ContractHead', (_message.Message,), dict(
  DESCRIPTOR = _CONTRACTHEAD,
  __module__ = 'contract_pb2'
  # @@protoc_insertion_point(class_scope:protos.ContractHead)
  ))
_sym_db.RegisterMessage(ContractHead)

Contract = _reflection.GeneratedProtocolMessageType('Contract', (_message.Message,), dict(
  DESCRIPTOR = _CONTRACT,
  __module__ = 'contract_pb2'
  # @@protoc_insertion_point(class_scope:protos.Contract)
  ))
_sym_db.RegisterMessage(Contract)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('\n\024com.uniledger.protosB\rProtoContract'))
_COMPONENTDATA_OPTIONSENTRY.has_options = True
_COMPONENTDATA_OPTIONSENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
# @@protoc_insertion_point(module_scope)