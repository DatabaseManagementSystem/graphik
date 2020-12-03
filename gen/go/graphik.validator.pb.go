// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: graphik.proto

package apipb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Bytes) Validate() error {
	return nil
}

var _regex_Path_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_Path_Gid = regexp.MustCompile(`^.{1,225}$`)

func (this *Path) Validate() error {
	if !_regex_Path_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !_regex_Path_Gid.MatchString(this.Gid) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gid", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gid))
	}
	return nil
}

var _regex_PathConstructor_Gtype = regexp.MustCompile(`^.{1,225}$`)

func (this *PathConstructor) Validate() error {
	if !_regex_PathConstructor_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	return nil
}
func (this *Metadata) Validate() error {
	if nil == this.CreatedAt {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", fmt.Errorf("message must exist"))
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if nil == this.UpdatedAt {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", fmt.Errorf("message must exist"))
	}
	if this.UpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	if this.CreatedBy != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedBy); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedBy", err)
		}
	}
	if nil == this.UpdatedBy {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", fmt.Errorf("message must exist"))
	}
	if this.UpdatedBy != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedBy); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", err)
		}
	}
	if !(this.Version > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Version", fmt.Errorf(`value '%v' must be greater than '0'`, this.Version))
	}
	return nil
}
func (this *Paths) Validate() error {
	for _, item := range this.Paths {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Paths", err)
			}
		}
	}
	return nil
}
func (this *Doc) Validate() error {
	if nil == this.Path {
		return github_com_mwitkow_go_proto_validators.FieldError("Path", fmt.Errorf("message must exist"))
	}
	if this.Path != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Path); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Path", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if nil == this.Metadata {
		return github_com_mwitkow_go_proto_validators.FieldError("Metadata", fmt.Errorf("message must exist"))
	}
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	return nil
}
func (this *DocConstructor) Validate() error {
	if nil == this.Path {
		return github_com_mwitkow_go_proto_validators.FieldError("Path", fmt.Errorf("message must exist"))
	}
	if this.Path != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Path); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Path", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *DocConstructors) Validate() error {
	for _, item := range this.Docs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Docs", err)
			}
		}
	}
	return nil
}
func (this *Docs) Validate() error {
	for _, item := range this.Docs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Docs", err)
			}
		}
	}
	return nil
}
func (this *DocTraversal) Validate() error {
	if this.Doc != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Doc); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Doc", err)
		}
	}
	if this.RelativePath != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RelativePath); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RelativePath", err)
		}
	}
	return nil
}
func (this *DocTraversals) Validate() error {
	for _, item := range this.Traversals {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Traversals", err)
			}
		}
	}
	return nil
}
func (this *DocDetail) Validate() error {
	if nil == this.Path {
		return github_com_mwitkow_go_proto_validators.FieldError("Path", fmt.Errorf("message must exist"))
	}
	if this.Path != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Path); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Path", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if this.ConnectionsFrom != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConnectionsFrom); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConnectionsFrom", err)
		}
	}
	if this.ConnectionsTo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConnectionsTo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConnectionsTo", err)
		}
	}
	if nil == this.Metadata {
		return github_com_mwitkow_go_proto_validators.FieldError("Metadata", fmt.Errorf("message must exist"))
	}
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	return nil
}
func (this *DocDetails) Validate() error {
	for _, item := range this.DocDetails {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DocDetails", err)
			}
		}
	}
	return nil
}
func (this *DocDetailFilter) Validate() error {
	if nil == this.Path {
		return github_com_mwitkow_go_proto_validators.FieldError("Path", fmt.Errorf("message must exist"))
	}
	if this.Path != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Path); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Path", err)
		}
	}
	if this.FromConnections != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.FromConnections); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("FromConnections", err)
		}
	}
	if this.ToConnections != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ToConnections); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ToConnections", err)
		}
	}
	return nil
}
func (this *Connection) Validate() error {
	if nil == this.Path {
		return github_com_mwitkow_go_proto_validators.FieldError("Path", fmt.Errorf("message must exist"))
	}
	if this.Path != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Path); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Path", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if nil == this.From {
		return github_com_mwitkow_go_proto_validators.FieldError("From", fmt.Errorf("message must exist"))
	}
	if this.From != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.From); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("From", err)
		}
	}
	if nil == this.To {
		return github_com_mwitkow_go_proto_validators.FieldError("To", fmt.Errorf("message must exist"))
	}
	if this.To != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.To); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("To", err)
		}
	}
	if nil == this.Metadata {
		return github_com_mwitkow_go_proto_validators.FieldError("Metadata", fmt.Errorf("message must exist"))
	}
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	return nil
}
func (this *ConnectionConstructor) Validate() error {
	if nil == this.Path {
		return github_com_mwitkow_go_proto_validators.FieldError("Path", fmt.Errorf("message must exist"))
	}
	if this.Path != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Path); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Path", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if nil == this.From {
		return github_com_mwitkow_go_proto_validators.FieldError("From", fmt.Errorf("message must exist"))
	}
	if this.From != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.From); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("From", err)
		}
	}
	if nil == this.To {
		return github_com_mwitkow_go_proto_validators.FieldError("To", fmt.Errorf("message must exist"))
	}
	if this.To != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.To); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("To", err)
		}
	}
	return nil
}
func (this *ConnectionConstructors) Validate() error {
	for _, item := range this.Connections {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Connections", err)
			}
		}
	}
	return nil
}
func (this *Connections) Validate() error {
	for _, item := range this.Connections {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Connections", err)
			}
		}
	}
	return nil
}
func (this *ConnectionDetail) Validate() error {
	if nil == this.Path {
		return github_com_mwitkow_go_proto_validators.FieldError("Path", fmt.Errorf("message must exist"))
	}
	if this.Path != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Path); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Path", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if nil == this.From {
		return github_com_mwitkow_go_proto_validators.FieldError("From", fmt.Errorf("message must exist"))
	}
	if this.From != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.From); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("From", err)
		}
	}
	if nil == this.To {
		return github_com_mwitkow_go_proto_validators.FieldError("To", fmt.Errorf("message must exist"))
	}
	if this.To != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.To); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("To", err)
		}
	}
	if nil == this.Metadata {
		return github_com_mwitkow_go_proto_validators.FieldError("Metadata", fmt.Errorf("message must exist"))
	}
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	return nil
}
func (this *ConnectionDetails) Validate() error {
	for _, item := range this.Connections {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Connections", err)
			}
		}
	}
	return nil
}

var _regex_ConnectionFilter_Gtype = regexp.MustCompile(`^.{1,225}$`)

func (this *ConnectionFilter) Validate() error {
	if nil == this.DocPath {
		return github_com_mwitkow_go_proto_validators.FieldError("DocPath", fmt.Errorf("message must exist"))
	}
	if this.DocPath != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DocPath); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DocPath", err)
		}
	}
	if !_regex_ConnectionFilter_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	return nil
}

var _regex_Filter_Gtype = regexp.MustCompile(`^.{1,225}$`)

func (this *Filter) Validate() error {
	if !_regex_Filter_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	return nil
}

var _regex_DepthFilter_ConnectionExpression = regexp.MustCompile(`^.{1,225}$`)

func (this *DepthFilter) Validate() error {
	if nil == this.Root {
		return github_com_mwitkow_go_proto_validators.FieldError("Root", fmt.Errorf("message must exist"))
	}
	if this.Root != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Root); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Root", err)
		}
	}
	if !_regex_DepthFilter_ConnectionExpression.MatchString(this.ConnectionExpression) {
		return github_com_mwitkow_go_proto_validators.FieldError("ConnectionExpression", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.ConnectionExpression))
	}
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	return nil
}

var _regex_IndexConstructor_Name = regexp.MustCompile(`^.{1,225}$`)
var _regex_IndexConstructor_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_IndexConstructor_Expression = regexp.MustCompile(`^.{1,225}$`)

func (this *IndexConstructor) Validate() error {
	if !_regex_IndexConstructor_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	if !_regex_IndexConstructor_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !_regex_IndexConstructor_Expression.MatchString(this.Expression) {
		return github_com_mwitkow_go_proto_validators.FieldError("Expression", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Expression))
	}
	return nil
}

var _regex_Authorizer_Name = regexp.MustCompile(`^.{1,225}$`)
var _regex_Authorizer_Expression = regexp.MustCompile(`^.{1,225}$`)

func (this *Authorizer) Validate() error {
	if !_regex_Authorizer_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	if !_regex_Authorizer_Expression.MatchString(this.Expression) {
		return github_com_mwitkow_go_proto_validators.FieldError("Expression", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Expression))
	}
	return nil
}
func (this *Authorizers) Validate() error {
	for _, item := range this.Authorizers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Authorizers", err)
			}
		}
	}
	return nil
}

var _regex_TypeValidator_Name = regexp.MustCompile(`^.{1,225}$`)
var _regex_TypeValidator_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_TypeValidator_Expression = regexp.MustCompile(`^.{1,225}$`)

func (this *TypeValidator) Validate() error {
	if !_regex_TypeValidator_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	if !_regex_TypeValidator_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !_regex_TypeValidator_Expression.MatchString(this.Expression) {
		return github_com_mwitkow_go_proto_validators.FieldError("Expression", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Expression))
	}
	return nil
}
func (this *TypeValidators) Validate() error {
	for _, item := range this.Validators {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Validators", err)
			}
		}
	}
	return nil
}

var _regex_Index_Name = regexp.MustCompile(`^.{1,225}$`)
var _regex_Index_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_Index_Expression = regexp.MustCompile(`^.{1,225}$`)

func (this *Index) Validate() error {
	if !_regex_Index_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	if !_regex_Index_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !_regex_Index_Expression.MatchString(this.Expression) {
		return github_com_mwitkow_go_proto_validators.FieldError("Expression", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Expression))
	}
	return nil
}
func (this *Indexes) Validate() error {
	for _, item := range this.Indexes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Indexes", err)
			}
		}
	}
	return nil
}
func (this *MeFilter) Validate() error {
	if this.ConnectionsFrom != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConnectionsFrom); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConnectionsFrom", err)
		}
	}
	if this.ConnectionsTo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConnectionsTo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConnectionsTo", err)
		}
	}
	return nil
}

var _regex_ChannelFilter_Channel = regexp.MustCompile(`^.{1,225}$`)

func (this *ChannelFilter) Validate() error {
	if !_regex_ChannelFilter_Channel.MatchString(this.Channel) {
		return github_com_mwitkow_go_proto_validators.FieldError("Channel", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Channel))
	}
	return nil
}
func (this *SubGraphFilter) Validate() error {
	if nil == this.DocFilter {
		return github_com_mwitkow_go_proto_validators.FieldError("DocFilter", fmt.Errorf("message must exist"))
	}
	if this.DocFilter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DocFilter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DocFilter", err)
		}
	}
	if nil == this.ConnectionFilter {
		return github_com_mwitkow_go_proto_validators.FieldError("ConnectionFilter", fmt.Errorf("message must exist"))
	}
	if this.ConnectionFilter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConnectionFilter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConnectionFilter", err)
		}
	}
	return nil
}
func (this *Graph) Validate() error {
	if this.Docs != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Docs); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Docs", err)
		}
	}
	if this.Connections != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Connections); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Connections", err)
		}
	}
	return nil
}
func (this *Flags) Validate() error {
	return nil
}
func (this *Edit) Validate() error {
	if nil == this.Path {
		return github_com_mwitkow_go_proto_validators.FieldError("Path", fmt.Errorf("message must exist"))
	}
	if this.Path != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Path); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Path", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *EditFilter) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *Pong) Validate() error {
	return nil
}

var _regex_OutboundMessage_Channel = regexp.MustCompile(`^.{1,225}$`)

func (this *OutboundMessage) Validate() error {
	if !_regex_OutboundMessage_Channel.MatchString(this.Channel) {
		return github_com_mwitkow_go_proto_validators.FieldError("Channel", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Channel))
	}
	if nil == this.Data {
		return github_com_mwitkow_go_proto_validators.FieldError("Data", fmt.Errorf("message must exist"))
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}

var _regex_Message_Channel = regexp.MustCompile(`^.{1,225}$`)

func (this *Message) Validate() error {
	if !_regex_Message_Channel.MatchString(this.Channel) {
		return github_com_mwitkow_go_proto_validators.FieldError("Channel", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Channel))
	}
	if nil == this.Data {
		return github_com_mwitkow_go_proto_validators.FieldError("Data", fmt.Errorf("message must exist"))
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if nil == this.Sender {
		return github_com_mwitkow_go_proto_validators.FieldError("Sender", fmt.Errorf("message must exist"))
	}
	if this.Sender != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Sender); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Sender", err)
		}
	}
	if nil == this.Timestamp {
		return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", fmt.Errorf("message must exist"))
	}
	if this.Timestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", err)
		}
	}
	return nil
}
func (this *Schema) Validate() error {
	if this.Authorizers != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Authorizers); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Authorizers", err)
		}
	}
	if this.Validators != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Validators); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Validators", err)
		}
	}
	if this.Indexes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Indexes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Indexes", err)
		}
	}
	return nil
}

var _regex_Change_Method = regexp.MustCompile(`^.{1,225}$`)

func (this *Change) Validate() error {
	if !_regex_Change_Method.MatchString(this.Method) {
		return github_com_mwitkow_go_proto_validators.FieldError("Method", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Method))
	}
	if nil == this.Identity {
		return github_com_mwitkow_go_proto_validators.FieldError("Identity", fmt.Errorf("message must exist"))
	}
	if this.Identity != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Identity); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Identity", err)
		}
	}
	if nil == this.Timestamp {
		return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", fmt.Errorf("message must exist"))
	}
	if this.Timestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", err)
		}
	}
	if this.PathsAffected != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PathsAffected); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PathsAffected", err)
		}
	}
	return nil
}
func (this *ExpressionFilter) Validate() error {
	return nil
}

var _regex_Request_Method = regexp.MustCompile(`^.{1,225}$`)

func (this *Request) Validate() error {
	if !_regex_Request_Method.MatchString(this.Method) {
		return github_com_mwitkow_go_proto_validators.FieldError("Method", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Method))
	}
	if nil == this.Identity {
		return github_com_mwitkow_go_proto_validators.FieldError("Identity", fmt.Errorf("message must exist"))
	}
	if this.Identity != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Identity); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Identity", err)
		}
	}
	if nil == this.Timestamp {
		return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", fmt.Errorf("message must exist"))
	}
	if this.Timestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", err)
		}
	}
	if this.Request != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Request); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Request", err)
		}
	}
	return nil
}
