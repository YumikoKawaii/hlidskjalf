package handler

import (
	"YumikoKawaii/hlidskjalf/golang/common/schemaregistry"
	pb "YumikoKawaii/hlidskjalf/golang/protobuf/schemaproxy"
)

func convertToSchemaRequest(request *pb.RegisterSchemaRequest) *schemaregistry.SchemaRequest {

	references := make([]*schemaregistry.Reference, 0)
	for _, reference := range request.References {
		references = append(references, &schemaregistry.Reference{
			Name:    reference.Name,
			Subject: reference.Subject,
			Version: int(reference.Version),
		})
	}

	return &schemaregistry.SchemaRequest{
		Schema:     request.Schema,
		SchemaType: request.SchemaType,
		References: references,
	}
}

//func convertToProtoSubjectVersions(svs []schemaregistry.SubjectVersion) []*pb.SubjectVersionInfo {
//	infos := make([]*pb.SubjectVersionInfo, 0)
//	for _, sv := range svs {
//		infos = append(infos, &pb.SubjectVersionInfo{
//			Subject: sv.Subject,
//			Version: int32(sv.Version),
//		})
//	}
//	return infos
//}

func convertToSubjectVersion(info *pb.SubjectVersionInfo) *schemaregistry.SubjectVersion {
	return &schemaregistry.SubjectVersion{
		Subject: info.Subject,
		Version: int(info.Version),
	}
}
