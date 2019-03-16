// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/compiler/plugin.proto

package plugin_proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	descriptor "github.com/golang/protobuf/v2/types/descriptor"
	reflect "reflect"
)

// The version number of protocol compiler.
type Version struct {
	Major *int32 `protobuf:"varint,1,opt,name=major" json:"major,omitempty"`
	Minor *int32 `protobuf:"varint,2,opt,name=minor" json:"minor,omitempty"`
	Patch *int32 `protobuf:"varint,3,opt,name=patch" json:"patch,omitempty"`
	// A suffix for alpha, beta or rc release, e.g., "alpha-1", "rc2". It should
	// be empty for mainline stable releases.
	Suffix               *string  `protobuf:"bytes,4,opt,name=suffix" json:"suffix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_compiler_plugin_proto_messageTypes[0].MessageOf(m)
}
func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}

// Deprecated: Use Version.ProtoReflect.Type instead.
func (*Version) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_compiler_plugin_proto_rawdesc_gzipped, []int{0}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetMajor() int32 {
	if m != nil && m.Major != nil {
		return *m.Major
	}
	return 0
}

func (m *Version) GetMinor() int32 {
	if m != nil && m.Minor != nil {
		return *m.Minor
	}
	return 0
}

func (m *Version) GetPatch() int32 {
	if m != nil && m.Patch != nil {
		return *m.Patch
	}
	return 0
}

func (m *Version) GetSuffix() string {
	if m != nil && m.Suffix != nil {
		return *m.Suffix
	}
	return ""
}

// An encoded CodeGeneratorRequest is written to the plugin's stdin.
type CodeGeneratorRequest struct {
	// The .proto files that were explicitly listed on the command-line.  The
	// code generator should generate code only for these files.  Each file's
	// descriptor will be included in proto_file, below.
	FileToGenerate []string `protobuf:"bytes,1,rep,name=file_to_generate,json=fileToGenerate" json:"file_to_generate,omitempty"`
	// The generator parameter passed on the command-line.
	Parameter *string `protobuf:"bytes,2,opt,name=parameter" json:"parameter,omitempty"`
	// FileDescriptorProtos for all files in files_to_generate and everything
	// they import.  The files will appear in topological order, so each file
	// appears before any file that imports it.
	//
	// protoc guarantees that all proto_files will be written after
	// the fields above, even though this is not technically guaranteed by the
	// protobuf wire format.  This theoretically could allow a plugin to stream
	// in the FileDescriptorProtos and handle them one by one rather than read
	// the entire set into memory at once.  However, as of this writing, this
	// is not similarly optimized on protoc's end -- it will store all fields in
	// memory at once before sending them to the plugin.
	//
	// Type names of fields and extensions in the FileDescriptorProto are always
	// fully qualified.
	ProtoFile []*descriptor.FileDescriptorProto `protobuf:"bytes,15,rep,name=proto_file,json=protoFile" json:"proto_file,omitempty"`
	// The version number of protocol compiler.
	CompilerVersion      *Version `protobuf:"bytes,3,opt,name=compiler_version,json=compilerVersion" json:"compiler_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeGeneratorRequest) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_compiler_plugin_proto_messageTypes[1].MessageOf(m)
}
func (m *CodeGeneratorRequest) Reset()         { *m = CodeGeneratorRequest{} }
func (m *CodeGeneratorRequest) String() string { return proto.CompactTextString(m) }
func (*CodeGeneratorRequest) ProtoMessage()    {}

// Deprecated: Use CodeGeneratorRequest.ProtoReflect.Type instead.
func (*CodeGeneratorRequest) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_compiler_plugin_proto_rawdesc_gzipped, []int{1}
}

func (m *CodeGeneratorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeGeneratorRequest.Unmarshal(m, b)
}
func (m *CodeGeneratorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeGeneratorRequest.Marshal(b, m, deterministic)
}
func (m *CodeGeneratorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeGeneratorRequest.Merge(m, src)
}
func (m *CodeGeneratorRequest) XXX_Size() int {
	return xxx_messageInfo_CodeGeneratorRequest.Size(m)
}
func (m *CodeGeneratorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeGeneratorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CodeGeneratorRequest proto.InternalMessageInfo

func (m *CodeGeneratorRequest) GetFileToGenerate() []string {
	if m != nil {
		return m.FileToGenerate
	}
	return nil
}

func (m *CodeGeneratorRequest) GetParameter() string {
	if m != nil && m.Parameter != nil {
		return *m.Parameter
	}
	return ""
}

func (m *CodeGeneratorRequest) GetProtoFile() []*descriptor.FileDescriptorProto {
	if m != nil {
		return m.ProtoFile
	}
	return nil
}

func (m *CodeGeneratorRequest) GetCompilerVersion() *Version {
	if m != nil {
		return m.CompilerVersion
	}
	return nil
}

// The plugin writes an encoded CodeGeneratorResponse to stdout.
type CodeGeneratorResponse struct {
	// Error message.  If non-empty, code generation failed.  The plugin process
	// should exit with status code zero even if it reports an error in this way.
	//
	// This should be used to indicate errors in .proto files which prevent the
	// code generator from generating correct code.  Errors which indicate a
	// problem in protoc itself -- such as the input CodeGeneratorRequest being
	// unparseable -- should be reported by writing a message to stderr and
	// exiting with a non-zero status code.
	Error                *string                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	File                 []*CodeGeneratorResponse_File `protobuf:"bytes,15,rep,name=file" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CodeGeneratorResponse) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_compiler_plugin_proto_messageTypes[2].MessageOf(m)
}
func (m *CodeGeneratorResponse) Reset()         { *m = CodeGeneratorResponse{} }
func (m *CodeGeneratorResponse) String() string { return proto.CompactTextString(m) }
func (*CodeGeneratorResponse) ProtoMessage()    {}

// Deprecated: Use CodeGeneratorResponse.ProtoReflect.Type instead.
func (*CodeGeneratorResponse) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_compiler_plugin_proto_rawdesc_gzipped, []int{2}
}

func (m *CodeGeneratorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeGeneratorResponse.Unmarshal(m, b)
}
func (m *CodeGeneratorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeGeneratorResponse.Marshal(b, m, deterministic)
}
func (m *CodeGeneratorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeGeneratorResponse.Merge(m, src)
}
func (m *CodeGeneratorResponse) XXX_Size() int {
	return xxx_messageInfo_CodeGeneratorResponse.Size(m)
}
func (m *CodeGeneratorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeGeneratorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CodeGeneratorResponse proto.InternalMessageInfo

func (m *CodeGeneratorResponse) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *CodeGeneratorResponse) GetFile() []*CodeGeneratorResponse_File {
	if m != nil {
		return m.File
	}
	return nil
}

// Represents a single generated file.
type CodeGeneratorResponse_File struct {
	// The file name, relative to the output directory.  The name must not
	// contain "." or ".." components and must be relative, not be absolute (so,
	// the file cannot lie outside the output directory).  "/" must be used as
	// the path separator, not "\".
	//
	// If the name is omitted, the content will be appended to the previous
	// file.  This allows the generator to break large files into small chunks,
	// and allows the generated text to be streamed back to protoc so that large
	// files need not reside completely in memory at one time.  Note that as of
	// this writing protoc does not optimize for this -- it will read the entire
	// CodeGeneratorResponse before writing files to disk.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// If non-empty, indicates that the named file should already exist, and the
	// content here is to be inserted into that file at a defined insertion
	// point.  This feature allows a code generator to extend the output
	// produced by another code generator.  The original generator may provide
	// insertion points by placing special annotations in the file that look
	// like:
	//   @@protoc_insertion_point(NAME)
	// The annotation can have arbitrary text before and after it on the line,
	// which allows it to be placed in a comment.  NAME should be replaced with
	// an identifier naming the point -- this is what other generators will use
	// as the insertion_point.  Code inserted at this point will be placed
	// immediately above the line containing the insertion point (thus multiple
	// insertions to the same point will come out in the order they were added).
	// The double-@ is intended to make it unlikely that the generated code
	// could contain things that look like insertion points by accident.
	//
	// For example, the C++ code generator places the following line in the
	// .pb.h files that it generates:
	//   // @@protoc_insertion_point(namespace_scope)
	// This line appears within the scope of the file's package namespace, but
	// outside of any particular class.  Another plugin can then specify the
	// insertion_point "namespace_scope" to generate additional classes or
	// other declarations that should be placed in this scope.
	//
	// Note that if the line containing the insertion point begins with
	// whitespace, the same whitespace will be added to every line of the
	// inserted text.  This is useful for languages like Python, where
	// indentation matters.  In these languages, the insertion point comment
	// should be indented the same amount as any inserted code will need to be
	// in order to work correctly in that context.
	//
	// The code generator that generates the initial file and the one which
	// inserts into it must both run as part of a single invocation of protoc.
	// Code generators are executed in the order in which they appear on the
	// command line.
	//
	// If |insertion_point| is present, |name| must also be present.
	InsertionPoint *string `protobuf:"bytes,2,opt,name=insertion_point,json=insertionPoint" json:"insertion_point,omitempty"`
	// The file contents.
	Content              *string  `protobuf:"bytes,15,opt,name=content" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeGeneratorResponse_File) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_compiler_plugin_proto_messageTypes[3].MessageOf(m)
}
func (m *CodeGeneratorResponse_File) Reset()         { *m = CodeGeneratorResponse_File{} }
func (m *CodeGeneratorResponse_File) String() string { return proto.CompactTextString(m) }
func (*CodeGeneratorResponse_File) ProtoMessage()    {}

// Deprecated: Use CodeGeneratorResponse_File.ProtoReflect.Type instead.
func (*CodeGeneratorResponse_File) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_compiler_plugin_proto_rawdesc_gzipped, []int{2, 0}
}

func (m *CodeGeneratorResponse_File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeGeneratorResponse_File.Unmarshal(m, b)
}
func (m *CodeGeneratorResponse_File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeGeneratorResponse_File.Marshal(b, m, deterministic)
}
func (m *CodeGeneratorResponse_File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeGeneratorResponse_File.Merge(m, src)
}
func (m *CodeGeneratorResponse_File) XXX_Size() int {
	return xxx_messageInfo_CodeGeneratorResponse_File.Size(m)
}
func (m *CodeGeneratorResponse_File) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeGeneratorResponse_File.DiscardUnknown(m)
}

var xxx_messageInfo_CodeGeneratorResponse_File proto.InternalMessageInfo

func (m *CodeGeneratorResponse_File) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CodeGeneratorResponse_File) GetInsertionPoint() string {
	if m != nil && m.InsertionPoint != nil {
		return *m.InsertionPoint
	}
	return ""
}

func (m *CodeGeneratorResponse_File) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func init() {
	proto.RegisterFile("google/protobuf/compiler/plugin.proto", xxx_File_google_protobuf_compiler_plugin_proto_rawdesc_gzipped)
	proto.RegisterType((*Version)(nil), "google.protobuf.compiler.Version")
	proto.RegisterType((*CodeGeneratorRequest)(nil), "google.protobuf.compiler.CodeGeneratorRequest")
	proto.RegisterType((*CodeGeneratorResponse)(nil), "google.protobuf.compiler.CodeGeneratorResponse")
	proto.RegisterType((*CodeGeneratorResponse_File)(nil), "google.protobuf.compiler.CodeGeneratorResponse.File")
}

var xxx_File_google_protobuf_compiler_plugin_proto_rawdesc = []byte{
	// 764 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x72, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d,
	0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x22, 0xf1, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x64,
	0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x69, 0x6c,
	0x65, 0x54, 0x6f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x4c,
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x65, 0x72, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xd6, 0x01, 0x0a,
	0x15, 0x43, 0x6f, 0x64, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x48, 0x0a, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x5d, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x65, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x42, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x32, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x3b,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var xxx_File_google_protobuf_compiler_plugin_proto_rawdesc_gzipped = protoimpl.X.CompressGZIP(xxx_File_google_protobuf_compiler_plugin_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_google_protobuf_compiler_plugin_proto protoreflect.FileDescriptor

var xxx_File_google_protobuf_compiler_plugin_proto_messageTypes = make([]protoimpl.MessageType, 4)
var xxx_File_google_protobuf_compiler_plugin_proto_goTypes = []interface{}{
	(*Version)(nil),                        // 0: google.protobuf.compiler.Version
	(*CodeGeneratorRequest)(nil),           // 1: google.protobuf.compiler.CodeGeneratorRequest
	(*CodeGeneratorResponse)(nil),          // 2: google.protobuf.compiler.CodeGeneratorResponse
	(*CodeGeneratorResponse_File)(nil),     // 3: google.protobuf.compiler.CodeGeneratorResponse.File
	(*descriptor.FileDescriptorProto)(nil), // 4: google.protobuf.FileDescriptorProto
}
var xxx_File_google_protobuf_compiler_plugin_proto_depIdxs = []int32{
	4, // google.protobuf.compiler.CodeGeneratorRequest.proto_file:type_name -> google.protobuf.FileDescriptorProto
	0, // google.protobuf.compiler.CodeGeneratorRequest.compiler_version:type_name -> google.protobuf.compiler.Version
	3, // google.protobuf.compiler.CodeGeneratorResponse.file:type_name -> google.protobuf.compiler.CodeGeneratorResponse.File
}

func init() { xxx_File_google_protobuf_compiler_plugin_proto_init() }
func xxx_File_google_protobuf_compiler_plugin_proto_init() {
	if File_google_protobuf_compiler_plugin_proto != nil {
		return
	}
	messageTypes := make([]protoreflect.MessageType, 4)
	File_google_protobuf_compiler_plugin_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_google_protobuf_compiler_plugin_proto_rawdesc,
		GoTypes:            xxx_File_google_protobuf_compiler_plugin_proto_goTypes,
		DependencyIndexes:  xxx_File_google_protobuf_compiler_plugin_proto_depIdxs,
		MessageOutputTypes: messageTypes,
	}.Init()
	messageGoTypes := xxx_File_google_protobuf_compiler_plugin_proto_goTypes[0:][:4]
	for i, mt := range messageTypes {
		xxx_File_google_protobuf_compiler_plugin_proto_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_File_google_protobuf_compiler_plugin_proto_messageTypes[i].PBType = mt
	}
	xxx_File_google_protobuf_compiler_plugin_proto_goTypes = nil
	xxx_File_google_protobuf_compiler_plugin_proto_depIdxs = nil
}
