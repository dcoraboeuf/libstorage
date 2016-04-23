package http

// NewRequestObjFunc is a function that creates a new instance of the type to
// which the request body is serialized.
type NewRequestObjFunc func() interface{}

// VolumeCreateRequest is the JSON body for creating a new volume.
type VolumeCreateRequest struct {
	Name             string                 `json:"name"`
	AvailabilityZone *string                `json:"availabilityZone"`
	IOPS             *int64                 `json:"iops"`
	Size             *int64                 `json:"size"`
	Type             *string                `json:"type"`
	Opts             map[string]interface{} `json:"opts"`
}

// VolumeCopyRequest is the JSON body for copying a volume.
type VolumeCopyRequest struct {
	VolumeName string                 `json:"volumeName"`
	Opts       map[string]interface{} `json:"opts"`
}

// VolumeSnapshotRequest is the JSON body for snapshotting a volume.
type VolumeSnapshotRequest struct {
	SnapshotName string                 `json:"snapshotName"`
	Opts         map[string]interface{} `json:"opts"`
}

// VolumeAttachRequest is the JSON body for attaching a volume to an instance.
type VolumeAttachRequest struct {
	NextDeviceName string                 `json:"nextDeviceName"`
	Opts           map[string]interface{} `json:"opts"`
}

// VolumeDetachRequest is the JSON body for detaching a volume from an instance.
type VolumeDetachRequest struct {
	Opts map[string]interface{} `json:"opts"`
}

// SnapshotCopyRequest is the JSON body for copying a snapshot.
type SnapshotCopyRequest struct {
	SnapshotName  string                 `json:"snapshotName"`
	DestinationID string                 `json:"destinationID"`
	Opts          map[string]interface{} `json:"opts"`
}

// SnapshotRemoveRequest is the JSON body for removing a snapshot.
type SnapshotRemoveRequest struct {
	Opts map[string]interface{} `json:"opts"`
}