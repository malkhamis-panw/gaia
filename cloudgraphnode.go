package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudGraphNode represents the model of a cloudgraphnode
type CloudGraphNode struct {
	// The native ID of the node.
	NativeID string `json:"nativeID" msgpack:"nativeID" bson:"-" mapstructure:"nativeID,omitempty"`

	// Details about the node if the query type requests full details.
	NodeData *CloudNode `json:"nodeData" msgpack:"nodeData" bson:"-" mapstructure:"nodeData,omitempty"`

	// The policies that were applied to this node for each destination.
	Policies map[string]*CloudGraphNodeAction `json:"policies" msgpack:"policies" bson:"-" mapstructure:"policies,omitempty"`

	// The list of route tables IDs that forwarding was based on for the internal path,
	// if routing was
	// performed.
	RouteTableIDs map[string]string `json:"routeTableIDs" msgpack:"routeTableIDs" bson:"-" mapstructure:"routeTableIDs,omitempty"`

	// The type of the node as a string.
	Type string `json:"type" msgpack:"type" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudGraphNode returns a new *CloudGraphNode
func NewCloudGraphNode() *CloudGraphNode {

	return &CloudGraphNode{
		ModelVersion:  1,
		NodeData:      NewCloudNode(),
		Policies:      map[string]*CloudGraphNodeAction{},
		RouteTableIDs: map[string]string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudGraphNode) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudGraphNode{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudGraphNode) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudGraphNode{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudGraphNode) BleveType() string {

	return "cloudgraphnode"
}

// DeepCopy returns a deep copy if the CloudGraphNode.
func (o *CloudGraphNode) DeepCopy() *CloudGraphNode {

	if o == nil {
		return nil
	}

	out := &CloudGraphNode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudGraphNode.
func (o *CloudGraphNode) DeepCopyInto(out *CloudGraphNode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudGraphNode: %s", err))
	}

	*out = *target.(*CloudGraphNode)
}

// Validate valides the current information stored into the structure.
func (o *CloudGraphNode) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if o.NodeData != nil {
		elemental.ResetDefaultForZeroValues(o.NodeData)
		if err := o.NodeData.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.Policies {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

type mongoAttributesCloudGraphNode struct {
}
