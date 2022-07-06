/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: datamovement.proto */

/* Do not generate deprecated warnings for self */
#ifndef PROTOBUF_C__NO_DEPRECATED
#define PROTOBUF_C__NO_DEPRECATED
#endif

#include "datamovement.pb-c.h"
void   datamovement__data_movement_create_request__init
                     (Datamovement__DataMovementCreateRequest         *message)
{
  static const Datamovement__DataMovementCreateRequest init_value = DATAMOVEMENT__DATA_MOVEMENT_CREATE_REQUEST__INIT;
  *message = init_value;
}
size_t datamovement__data_movement_create_request__get_packed_size
                     (const Datamovement__DataMovementCreateRequest *message)
{
  assert(message->base.descriptor == &datamovement__data_movement_create_request__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t datamovement__data_movement_create_request__pack
                     (const Datamovement__DataMovementCreateRequest *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &datamovement__data_movement_create_request__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t datamovement__data_movement_create_request__pack_to_buffer
                     (const Datamovement__DataMovementCreateRequest *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &datamovement__data_movement_create_request__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Datamovement__DataMovementCreateRequest *
       datamovement__data_movement_create_request__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Datamovement__DataMovementCreateRequest *)
     protobuf_c_message_unpack (&datamovement__data_movement_create_request__descriptor,
                                allocator, len, data);
}
void   datamovement__data_movement_create_request__free_unpacked
                     (Datamovement__DataMovementCreateRequest *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &datamovement__data_movement_create_request__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   datamovement__data_movement_create_response__init
                     (Datamovement__DataMovementCreateResponse         *message)
{
  static const Datamovement__DataMovementCreateResponse init_value = DATAMOVEMENT__DATA_MOVEMENT_CREATE_RESPONSE__INIT;
  *message = init_value;
}
size_t datamovement__data_movement_create_response__get_packed_size
                     (const Datamovement__DataMovementCreateResponse *message)
{
  assert(message->base.descriptor == &datamovement__data_movement_create_response__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t datamovement__data_movement_create_response__pack
                     (const Datamovement__DataMovementCreateResponse *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &datamovement__data_movement_create_response__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t datamovement__data_movement_create_response__pack_to_buffer
                     (const Datamovement__DataMovementCreateResponse *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &datamovement__data_movement_create_response__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Datamovement__DataMovementCreateResponse *
       datamovement__data_movement_create_response__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Datamovement__DataMovementCreateResponse *)
     protobuf_c_message_unpack (&datamovement__data_movement_create_response__descriptor,
                                allocator, len, data);
}
void   datamovement__data_movement_create_response__free_unpacked
                     (Datamovement__DataMovementCreateResponse *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &datamovement__data_movement_create_response__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   datamovement__data_movement_status_request__init
                     (Datamovement__DataMovementStatusRequest         *message)
{
  static const Datamovement__DataMovementStatusRequest init_value = DATAMOVEMENT__DATA_MOVEMENT_STATUS_REQUEST__INIT;
  *message = init_value;
}
size_t datamovement__data_movement_status_request__get_packed_size
                     (const Datamovement__DataMovementStatusRequest *message)
{
  assert(message->base.descriptor == &datamovement__data_movement_status_request__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t datamovement__data_movement_status_request__pack
                     (const Datamovement__DataMovementStatusRequest *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &datamovement__data_movement_status_request__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t datamovement__data_movement_status_request__pack_to_buffer
                     (const Datamovement__DataMovementStatusRequest *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &datamovement__data_movement_status_request__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Datamovement__DataMovementStatusRequest *
       datamovement__data_movement_status_request__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Datamovement__DataMovementStatusRequest *)
     protobuf_c_message_unpack (&datamovement__data_movement_status_request__descriptor,
                                allocator, len, data);
}
void   datamovement__data_movement_status_request__free_unpacked
                     (Datamovement__DataMovementStatusRequest *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &datamovement__data_movement_status_request__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   datamovement__data_movement_status_response__init
                     (Datamovement__DataMovementStatusResponse         *message)
{
  static const Datamovement__DataMovementStatusResponse init_value = DATAMOVEMENT__DATA_MOVEMENT_STATUS_RESPONSE__INIT;
  *message = init_value;
}
size_t datamovement__data_movement_status_response__get_packed_size
                     (const Datamovement__DataMovementStatusResponse *message)
{
  assert(message->base.descriptor == &datamovement__data_movement_status_response__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t datamovement__data_movement_status_response__pack
                     (const Datamovement__DataMovementStatusResponse *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &datamovement__data_movement_status_response__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t datamovement__data_movement_status_response__pack_to_buffer
                     (const Datamovement__DataMovementStatusResponse *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &datamovement__data_movement_status_response__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Datamovement__DataMovementStatusResponse *
       datamovement__data_movement_status_response__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Datamovement__DataMovementStatusResponse *)
     protobuf_c_message_unpack (&datamovement__data_movement_status_response__descriptor,
                                allocator, len, data);
}
void   datamovement__data_movement_status_response__free_unpacked
                     (Datamovement__DataMovementStatusResponse *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &datamovement__data_movement_status_response__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   datamovement__data_movement_delete_request__init
                     (Datamovement__DataMovementDeleteRequest         *message)
{
  static const Datamovement__DataMovementDeleteRequest init_value = DATAMOVEMENT__DATA_MOVEMENT_DELETE_REQUEST__INIT;
  *message = init_value;
}
size_t datamovement__data_movement_delete_request__get_packed_size
                     (const Datamovement__DataMovementDeleteRequest *message)
{
  assert(message->base.descriptor == &datamovement__data_movement_delete_request__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t datamovement__data_movement_delete_request__pack
                     (const Datamovement__DataMovementDeleteRequest *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &datamovement__data_movement_delete_request__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t datamovement__data_movement_delete_request__pack_to_buffer
                     (const Datamovement__DataMovementDeleteRequest *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &datamovement__data_movement_delete_request__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Datamovement__DataMovementDeleteRequest *
       datamovement__data_movement_delete_request__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Datamovement__DataMovementDeleteRequest *)
     protobuf_c_message_unpack (&datamovement__data_movement_delete_request__descriptor,
                                allocator, len, data);
}
void   datamovement__data_movement_delete_request__free_unpacked
                     (Datamovement__DataMovementDeleteRequest *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &datamovement__data_movement_delete_request__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   datamovement__data_movement_delete_response__init
                     (Datamovement__DataMovementDeleteResponse         *message)
{
  static const Datamovement__DataMovementDeleteResponse init_value = DATAMOVEMENT__DATA_MOVEMENT_DELETE_RESPONSE__INIT;
  *message = init_value;
}
size_t datamovement__data_movement_delete_response__get_packed_size
                     (const Datamovement__DataMovementDeleteResponse *message)
{
  assert(message->base.descriptor == &datamovement__data_movement_delete_response__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t datamovement__data_movement_delete_response__pack
                     (const Datamovement__DataMovementDeleteResponse *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &datamovement__data_movement_delete_response__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t datamovement__data_movement_delete_response__pack_to_buffer
                     (const Datamovement__DataMovementDeleteResponse *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &datamovement__data_movement_delete_response__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Datamovement__DataMovementDeleteResponse *
       datamovement__data_movement_delete_response__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Datamovement__DataMovementDeleteResponse *)
     protobuf_c_message_unpack (&datamovement__data_movement_delete_response__descriptor,
                                allocator, len, data);
}
void   datamovement__data_movement_delete_response__free_unpacked
                     (Datamovement__DataMovementDeleteResponse *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &datamovement__data_movement_delete_response__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
static const ProtobufCFieldDescriptor datamovement__data_movement_create_request__field_descriptors[5] =
{
  {
    "workflow",
    3,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__DataMovementCreateRequest, workflow),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "namespace",
    4,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__DataMovementCreateRequest, namespace_),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "source",
    5,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__DataMovementCreateRequest, source),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "destination",
    6,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__DataMovementCreateRequest, destination),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "dryrun",
    7,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_BOOL,
    0,   /* quantifier_offset */
    offsetof(Datamovement__DataMovementCreateRequest, dryrun),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned datamovement__data_movement_create_request__field_indices_by_name[] = {
  3,   /* field[3] = destination */
  4,   /* field[4] = dryrun */
  1,   /* field[1] = namespace */
  2,   /* field[2] = source */
  0,   /* field[0] = workflow */
};
static const ProtobufCIntRange datamovement__data_movement_create_request__number_ranges[1 + 1] =
{
  { 3, 0 },
  { 0, 5 }
};
const ProtobufCMessageDescriptor datamovement__data_movement_create_request__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "datamovement.DataMovementCreateRequest",
  "DataMovementCreateRequest",
  "Datamovement__DataMovementCreateRequest",
  "datamovement",
  sizeof(Datamovement__DataMovementCreateRequest),
  5,
  datamovement__data_movement_create_request__field_descriptors,
  datamovement__data_movement_create_request__field_indices_by_name,
  1,  datamovement__data_movement_create_request__number_ranges,
  (ProtobufCMessageInit) datamovement__data_movement_create_request__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCEnumValue datamovement__data_movement_create_response__status__enum_values_by_number[2] =
{
  { "CREATED", "DATAMOVEMENT__DATA_MOVEMENT_CREATE_RESPONSE__STATUS__CREATED", 0 },
  { "FAILED", "DATAMOVEMENT__DATA_MOVEMENT_CREATE_RESPONSE__STATUS__FAILED", 1 },
};
static const ProtobufCIntRange datamovement__data_movement_create_response__status__value_ranges[] = {
{0, 0},{0, 2}
};
static const ProtobufCEnumValueIndex datamovement__data_movement_create_response__status__enum_values_by_name[2] =
{
  { "CREATED", 0 },
  { "FAILED", 1 },
};
const ProtobufCEnumDescriptor datamovement__data_movement_create_response__status__descriptor =
{
  PROTOBUF_C__ENUM_DESCRIPTOR_MAGIC,
  "datamovement.DataMovementCreateResponse.Status",
  "Status",
  "Datamovement__DataMovementCreateResponse__Status",
  "datamovement",
  2,
  datamovement__data_movement_create_response__status__enum_values_by_number,
  2,
  datamovement__data_movement_create_response__status__enum_values_by_name,
  1,
  datamovement__data_movement_create_response__status__value_ranges,
  NULL,NULL,NULL,NULL   /* reserved[1234] */
};
static const ProtobufCFieldDescriptor datamovement__data_movement_create_response__field_descriptors[3] =
{
  {
    "uid",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__DataMovementCreateResponse, uid),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "status",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_ENUM,
    0,   /* quantifier_offset */
    offsetof(Datamovement__DataMovementCreateResponse, status),
    &datamovement__data_movement_create_response__status__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "message",
    3,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__DataMovementCreateResponse, message),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned datamovement__data_movement_create_response__field_indices_by_name[] = {
  2,   /* field[2] = message */
  1,   /* field[1] = status */
  0,   /* field[0] = uid */
};
static const ProtobufCIntRange datamovement__data_movement_create_response__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 3 }
};
const ProtobufCMessageDescriptor datamovement__data_movement_create_response__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "datamovement.DataMovementCreateResponse",
  "DataMovementCreateResponse",
  "Datamovement__DataMovementCreateResponse",
  "datamovement",
  sizeof(Datamovement__DataMovementCreateResponse),
  3,
  datamovement__data_movement_create_response__field_descriptors,
  datamovement__data_movement_create_response__field_indices_by_name,
  1,  datamovement__data_movement_create_response__number_ranges,
  (ProtobufCMessageInit) datamovement__data_movement_create_response__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor datamovement__data_movement_status_request__field_descriptors[2] =
{
  {
    "uid",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__DataMovementStatusRequest, uid),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "maxWaitTime",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_INT64,
    0,   /* quantifier_offset */
    offsetof(Datamovement__DataMovementStatusRequest, maxwaittime),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned datamovement__data_movement_status_request__field_indices_by_name[] = {
  1,   /* field[1] = maxWaitTime */
  0,   /* field[0] = uid */
};
static const ProtobufCIntRange datamovement__data_movement_status_request__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 2 }
};
const ProtobufCMessageDescriptor datamovement__data_movement_status_request__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "datamovement.DataMovementStatusRequest",
  "DataMovementStatusRequest",
  "Datamovement__DataMovementStatusRequest",
  "datamovement",
  sizeof(Datamovement__DataMovementStatusRequest),
  2,
  datamovement__data_movement_status_request__field_descriptors,
  datamovement__data_movement_status_request__field_indices_by_name,
  1,  datamovement__data_movement_status_request__number_ranges,
  (ProtobufCMessageInit) datamovement__data_movement_status_request__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCEnumValue datamovement__data_movement_status_response__state__enum_values_by_number[5] =
{
  { "PENDING", "DATAMOVEMENT__DATA_MOVEMENT_STATUS_RESPONSE__STATE__PENDING", 0 },
  { "STARTING", "DATAMOVEMENT__DATA_MOVEMENT_STATUS_RESPONSE__STATE__STARTING", 1 },
  { "RUNNING", "DATAMOVEMENT__DATA_MOVEMENT_STATUS_RESPONSE__STATE__RUNNING", 2 },
  { "COMPLETED", "DATAMOVEMENT__DATA_MOVEMENT_STATUS_RESPONSE__STATE__COMPLETED", 3 },
  { "UNKNOWN_STATE", "DATAMOVEMENT__DATA_MOVEMENT_STATUS_RESPONSE__STATE__UNKNOWN_STATE", 4 },
};
static const ProtobufCIntRange datamovement__data_movement_status_response__state__value_ranges[] = {
{0, 0},{0, 5}
};
static const ProtobufCEnumValueIndex datamovement__data_movement_status_response__state__enum_values_by_name[5] =
{
  { "COMPLETED", 3 },
  { "PENDING", 0 },
  { "RUNNING", 2 },
  { "STARTING", 1 },
  { "UNKNOWN_STATE", 4 },
};
const ProtobufCEnumDescriptor datamovement__data_movement_status_response__state__descriptor =
{
  PROTOBUF_C__ENUM_DESCRIPTOR_MAGIC,
  "datamovement.DataMovementStatusResponse.State",
  "State",
  "Datamovement__DataMovementStatusResponse__State",
  "datamovement",
  5,
  datamovement__data_movement_status_response__state__enum_values_by_number,
  5,
  datamovement__data_movement_status_response__state__enum_values_by_name,
  1,
  datamovement__data_movement_status_response__state__value_ranges,
  NULL,NULL,NULL,NULL   /* reserved[1234] */
};
static const ProtobufCEnumValue datamovement__data_movement_status_response__status__enum_values_by_number[5] =
{
  { "INVALID", "DATAMOVEMENT__DATA_MOVEMENT_STATUS_RESPONSE__STATUS__INVALID", 0 },
  { "NOT_FOUND", "DATAMOVEMENT__DATA_MOVEMENT_STATUS_RESPONSE__STATUS__NOT_FOUND", 1 },
  { "SUCCESS", "DATAMOVEMENT__DATA_MOVEMENT_STATUS_RESPONSE__STATUS__SUCCESS", 2 },
  { "FAILED", "DATAMOVEMENT__DATA_MOVEMENT_STATUS_RESPONSE__STATUS__FAILED", 3 },
  { "UNKNOWN_STATUS", "DATAMOVEMENT__DATA_MOVEMENT_STATUS_RESPONSE__STATUS__UNKNOWN_STATUS", 4 },
};
static const ProtobufCIntRange datamovement__data_movement_status_response__status__value_ranges[] = {
{0, 0},{0, 5}
};
static const ProtobufCEnumValueIndex datamovement__data_movement_status_response__status__enum_values_by_name[5] =
{
  { "FAILED", 3 },
  { "INVALID", 0 },
  { "NOT_FOUND", 1 },
  { "SUCCESS", 2 },
  { "UNKNOWN_STATUS", 4 },
};
const ProtobufCEnumDescriptor datamovement__data_movement_status_response__status__descriptor =
{
  PROTOBUF_C__ENUM_DESCRIPTOR_MAGIC,
  "datamovement.DataMovementStatusResponse.Status",
  "Status",
  "Datamovement__DataMovementStatusResponse__Status",
  "datamovement",
  5,
  datamovement__data_movement_status_response__status__enum_values_by_number,
  5,
  datamovement__data_movement_status_response__status__enum_values_by_name,
  1,
  datamovement__data_movement_status_response__status__value_ranges,
  NULL,NULL,NULL,NULL   /* reserved[1234] */
};
static const ProtobufCFieldDescriptor datamovement__data_movement_status_response__field_descriptors[3] =
{
  {
    "state",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_ENUM,
    0,   /* quantifier_offset */
    offsetof(Datamovement__DataMovementStatusResponse, state),
    &datamovement__data_movement_status_response__state__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "status",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_ENUM,
    0,   /* quantifier_offset */
    offsetof(Datamovement__DataMovementStatusResponse, status),
    &datamovement__data_movement_status_response__status__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "message",
    3,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__DataMovementStatusResponse, message),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned datamovement__data_movement_status_response__field_indices_by_name[] = {
  2,   /* field[2] = message */
  0,   /* field[0] = state */
  1,   /* field[1] = status */
};
static const ProtobufCIntRange datamovement__data_movement_status_response__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 3 }
};
const ProtobufCMessageDescriptor datamovement__data_movement_status_response__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "datamovement.DataMovementStatusResponse",
  "DataMovementStatusResponse",
  "Datamovement__DataMovementStatusResponse",
  "datamovement",
  sizeof(Datamovement__DataMovementStatusResponse),
  3,
  datamovement__data_movement_status_response__field_descriptors,
  datamovement__data_movement_status_response__field_indices_by_name,
  1,  datamovement__data_movement_status_response__number_ranges,
  (ProtobufCMessageInit) datamovement__data_movement_status_response__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor datamovement__data_movement_delete_request__field_descriptors[1] =
{
  {
    "uid",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__DataMovementDeleteRequest, uid),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned datamovement__data_movement_delete_request__field_indices_by_name[] = {
  0,   /* field[0] = uid */
};
static const ProtobufCIntRange datamovement__data_movement_delete_request__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 1 }
};
const ProtobufCMessageDescriptor datamovement__data_movement_delete_request__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "datamovement.DataMovementDeleteRequest",
  "DataMovementDeleteRequest",
  "Datamovement__DataMovementDeleteRequest",
  "datamovement",
  sizeof(Datamovement__DataMovementDeleteRequest),
  1,
  datamovement__data_movement_delete_request__field_descriptors,
  datamovement__data_movement_delete_request__field_indices_by_name,
  1,  datamovement__data_movement_delete_request__number_ranges,
  (ProtobufCMessageInit) datamovement__data_movement_delete_request__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCEnumValue datamovement__data_movement_delete_response__status__enum_values_by_number[5] =
{
  { "INVALID", "DATAMOVEMENT__DATA_MOVEMENT_DELETE_RESPONSE__STATUS__INVALID", 0 },
  { "NOT_FOUND", "DATAMOVEMENT__DATA_MOVEMENT_DELETE_RESPONSE__STATUS__NOT_FOUND", 1 },
  { "DELETED", "DATAMOVEMENT__DATA_MOVEMENT_DELETE_RESPONSE__STATUS__DELETED", 2 },
  { "ACTIVE", "DATAMOVEMENT__DATA_MOVEMENT_DELETE_RESPONSE__STATUS__ACTIVE", 3 },
  { "UNKNOWN", "DATAMOVEMENT__DATA_MOVEMENT_DELETE_RESPONSE__STATUS__UNKNOWN", 4 },
};
static const ProtobufCIntRange datamovement__data_movement_delete_response__status__value_ranges[] = {
{0, 0},{0, 5}
};
static const ProtobufCEnumValueIndex datamovement__data_movement_delete_response__status__enum_values_by_name[5] =
{
  { "ACTIVE", 3 },
  { "DELETED", 2 },
  { "INVALID", 0 },
  { "NOT_FOUND", 1 },
  { "UNKNOWN", 4 },
};
const ProtobufCEnumDescriptor datamovement__data_movement_delete_response__status__descriptor =
{
  PROTOBUF_C__ENUM_DESCRIPTOR_MAGIC,
  "datamovement.DataMovementDeleteResponse.Status",
  "Status",
  "Datamovement__DataMovementDeleteResponse__Status",
  "datamovement",
  5,
  datamovement__data_movement_delete_response__status__enum_values_by_number,
  5,
  datamovement__data_movement_delete_response__status__enum_values_by_name,
  1,
  datamovement__data_movement_delete_response__status__value_ranges,
  NULL,NULL,NULL,NULL   /* reserved[1234] */
};
static const ProtobufCFieldDescriptor datamovement__data_movement_delete_response__field_descriptors[1] =
{
  {
    "status",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_ENUM,
    0,   /* quantifier_offset */
    offsetof(Datamovement__DataMovementDeleteResponse, status),
    &datamovement__data_movement_delete_response__status__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned datamovement__data_movement_delete_response__field_indices_by_name[] = {
  0,   /* field[0] = status */
};
static const ProtobufCIntRange datamovement__data_movement_delete_response__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 1 }
};
const ProtobufCMessageDescriptor datamovement__data_movement_delete_response__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "datamovement.DataMovementDeleteResponse",
  "DataMovementDeleteResponse",
  "Datamovement__DataMovementDeleteResponse",
  "datamovement",
  sizeof(Datamovement__DataMovementDeleteResponse),
  1,
  datamovement__data_movement_delete_response__field_descriptors,
  datamovement__data_movement_delete_response__field_indices_by_name,
  1,  datamovement__data_movement_delete_response__number_ranges,
  (ProtobufCMessageInit) datamovement__data_movement_delete_response__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCMethodDescriptor datamovement__data_mover__method_descriptors[3] =
{
  { "Create", &datamovement__data_movement_create_request__descriptor, &datamovement__data_movement_create_response__descriptor },
  { "Status", &datamovement__data_movement_status_request__descriptor, &datamovement__data_movement_status_response__descriptor },
  { "Delete", &datamovement__data_movement_delete_request__descriptor, &datamovement__data_movement_delete_response__descriptor },
};
const unsigned datamovement__data_mover__method_indices_by_name[] = {
  0,        /* Create */
  2,        /* Delete */
  1         /* Status */
};
const ProtobufCServiceDescriptor datamovement__data_mover__descriptor =
{
  PROTOBUF_C__SERVICE_DESCRIPTOR_MAGIC,
  "datamovement.DataMover",
  "DataMover",
  "Datamovement__DataMover",
  "datamovement",
  3,
  datamovement__data_mover__method_descriptors,
  datamovement__data_mover__method_indices_by_name
};
void datamovement__data_mover__create(ProtobufCService *service,
                                      const Datamovement__DataMovementCreateRequest *input,
                                      Datamovement__DataMovementCreateResponse_Closure closure,
                                      void *closure_data)
{
  assert(service->descriptor == &datamovement__data_mover__descriptor);
  service->invoke(service, 0, (const ProtobufCMessage *) input, (ProtobufCClosure) closure, closure_data);
}
void datamovement__data_mover__status(ProtobufCService *service,
                                      const Datamovement__DataMovementStatusRequest *input,
                                      Datamovement__DataMovementStatusResponse_Closure closure,
                                      void *closure_data)
{
  assert(service->descriptor == &datamovement__data_mover__descriptor);
  service->invoke(service, 1, (const ProtobufCMessage *) input, (ProtobufCClosure) closure, closure_data);
}
void datamovement__data_mover__delete(ProtobufCService *service,
                                      const Datamovement__DataMovementDeleteRequest *input,
                                      Datamovement__DataMovementDeleteResponse_Closure closure,
                                      void *closure_data)
{
  assert(service->descriptor == &datamovement__data_mover__descriptor);
  service->invoke(service, 2, (const ProtobufCMessage *) input, (ProtobufCClosure) closure, closure_data);
}
void datamovement__data_mover__init (Datamovement__DataMover_Service *service,
                                     Datamovement__DataMover_ServiceDestroy destroy)
{
  protobuf_c_service_generated_init (&service->base,
                                     &datamovement__data_mover__descriptor,
                                     (ProtobufCServiceDestroy) destroy);
}
