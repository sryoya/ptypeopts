package stropt

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func GetAvailableChars(pb proto.Message) []string {
	m := pb.ProtoReflect()
	var res []string
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		opts := fd.Options().(*descriptorpb.FieldOptions)
		ext, ok := proto.GetExtension(opts, E_MaxLength).(string)
		if !ok {
			return true
		}
		res = append(res, ext)
		return true
	})
	return res
}
