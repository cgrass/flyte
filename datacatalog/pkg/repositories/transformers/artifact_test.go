package transformers

import (
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"

	"github.com/flyteorg/flyte/datacatalog/pkg/repositories/models"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/datacatalog"
)

func getTestArtifactData() []*datacatalog.ArtifactData {
	testInteger := &core.Literal{
		Value: &core.Literal_Scalar{
			Scalar: &core.Scalar{
				Value: &core.Scalar_Primitive{
					Primitive: &core.Primitive{Value: &core.Primitive_Integer{Integer: 1}},
				},
			},
		},
	}
	return []*datacatalog.ArtifactData{
		{Name: "data1", Value: testInteger},
		{Name: "data2", Value: testInteger},
	}
}

func getTestPartitions() []models.Partition {
	return []models.Partition{
		{DatasetUUID: "test-uuid", Key: "key1", Value: "value1"},
		{DatasetUUID: "test-uuid", Key: "key2", Value: "value2"},
	}
}

func getTestTags() []models.Tag {
	return []models.Tag{
		{TagKey: models.TagKey{TagName: "test"}},
	}
}

func getDatasetModel() models.Dataset {
	return models.Dataset{
		DatasetKey: models.DatasetKey{
			Project: datasetID.Project,
			Domain:  datasetID.Domain,
			Name:    datasetID.Name,
			Version: datasetID.Version,
			UUID:    datasetID.UUID,
		},
	}
}

func TestCreateArtifactModel(t *testing.T) {

	createArtifactRequest := &datacatalog.CreateArtifactRequest{
		Artifact: &datacatalog.Artifact{
			Id:       "artifactID-1",
			Dataset:  datasetID,
			Data:     getTestArtifactData(),
			Metadata: metadata,
			Partitions: []*datacatalog.Partition{
				{Key: "key1", Value: "value1"},
				{Key: "key2", Value: "value2"},
			},
		},
	}

	testArtifactData := []models.ArtifactData{
		{Name: "data1", Location: "s3://test1"},
		{Name: "data3", Location: "s3://test2"},
	}

	artifactModel, err := CreateArtifactModel(createArtifactRequest, testArtifactData, getDatasetModel())
	assert.NoError(t, err)
	assert.Equal(t, artifactModel.ArtifactID, createArtifactRequest.Artifact.Id)
	assert.Equal(t, artifactModel.ArtifactKey.DatasetProject, datasetID.Project)
	assert.Equal(t, artifactModel.ArtifactKey.DatasetDomain, datasetID.Domain)
	assert.Equal(t, artifactModel.ArtifactKey.DatasetName, datasetID.Name)
	assert.Equal(t, artifactModel.ArtifactKey.DatasetVersion, datasetID.Version)
	assert.EqualValues(t, testArtifactData, artifactModel.ArtifactData)
	assert.EqualValues(t, getTestPartitions(), artifactModel.Partitions)
}

func TestCreateArtifactModelNoMetdata(t *testing.T) {
	createArtifactRequest := &datacatalog.CreateArtifactRequest{
		Artifact: &datacatalog.Artifact{
			Id:      "artifactID-1",
			Dataset: datasetID,
			Data:    getTestArtifactData(),
		},
	}

	testArtifactData := []models.ArtifactData{
		{Name: "data1", Location: "s3://test1"},
		{Name: "data3", Location: "s3://test2"},
	}
	artifactModel, err := CreateArtifactModel(createArtifactRequest, testArtifactData, getDatasetModel())
	assert.NoError(t, err)
	assert.Equal(t, []byte{}, artifactModel.SerializedMetadata)
	assert.Len(t, artifactModel.Partitions, 0)
}

func TestFromArtifactModel(t *testing.T) {
	createdAt := time.Now()

	artifactModel := models.Artifact{
		ArtifactKey: models.ArtifactKey{
			DatasetProject: "project1",
			DatasetDomain:  "domain1",
			DatasetName:    "name1",
			DatasetVersion: "version1",
			ArtifactID:     "id1",
		},
		SerializedMetadata: []byte{},
		Partitions:         getTestPartitions(),
		Tags:               getTestTags(),
		BaseModel: models.BaseModel{
			CreatedAt: createdAt,
		},
	}

	actual, err := FromArtifactModel(artifactModel)
	assert.NoError(t, err)
	assert.Equal(t, artifactModel.ArtifactID, actual.Id)
	assert.Equal(t, artifactModel.DatasetProject, actual.Dataset.Project)
	assert.Equal(t, artifactModel.DatasetDomain, actual.Dataset.Domain)
	assert.Equal(t, artifactModel.DatasetName, actual.Dataset.Name)
	assert.Equal(t, artifactModel.DatasetVersion, actual.Dataset.Version)

	assert.Len(t, actual.Partitions, 2)
	assert.EqualValues(t, artifactModel.Partitions[0].Key, actual.Partitions[0].Key)
	assert.EqualValues(t, artifactModel.Partitions[0].Value, actual.Partitions[0].Value)
	assert.EqualValues(t, artifactModel.Partitions[1].Value, actual.Partitions[1].Value)
	assert.EqualValues(t, artifactModel.Partitions[1].Value, actual.Partitions[1].Value)

	assert.Len(t, actual.Tags, 1)
	assert.EqualValues(t, artifactModel.Tags[0].TagName, actual.Tags[0].Name)

	timestampProto, err := ptypes.TimestampProto(createdAt)
	assert.NoError(t, err)
	assert.Equal(t, actual.CreatedAt, timestampProto)
}

func TestToArtifactKey(t *testing.T) {
	artifactKey := ToArtifactKey(datasetID, "artifactID-1")
	assert.Equal(t, datasetID.Project, artifactKey.DatasetProject)
	assert.Equal(t, datasetID.Domain, artifactKey.DatasetDomain)
	assert.Equal(t, datasetID.Name, artifactKey.DatasetName)
	assert.Equal(t, datasetID.Version, artifactKey.DatasetVersion)
	assert.Equal(t, artifactKey.ArtifactID, "artifactID-1")
}

func TestToArtifactKeyNoDataset(t *testing.T) {
	artifactKey := ToArtifactKey(nil, "artifactID-1")
	assert.Equal(t, artifactKey.DatasetProject, "")
	assert.Equal(t, artifactKey.DatasetDomain, "")
	assert.Equal(t, artifactKey.DatasetName, "")
	assert.Equal(t, artifactKey.DatasetVersion, "")
	assert.Equal(t, artifactKey.ArtifactID, "artifactID-1")
}
