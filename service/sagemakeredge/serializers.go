// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakeredge

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sagemakeredge/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpGetDeviceRegistration struct {
}

func (*awsRestjson1_serializeOpGetDeviceRegistration) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetDeviceRegistration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetDeviceRegistrationInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/GetDeviceRegistration")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentGetDeviceRegistrationInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetDeviceRegistrationInput(v *GetDeviceRegistrationInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentGetDeviceRegistrationInput(v *GetDeviceRegistrationInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DeviceFleetName != nil {
		ok := object.Key("DeviceFleetName")
		ok.String(*v.DeviceFleetName)
	}

	if v.DeviceName != nil {
		ok := object.Key("DeviceName")
		ok.String(*v.DeviceName)
	}

	return nil
}

type awsRestjson1_serializeOpSendHeartbeat struct {
}

func (*awsRestjson1_serializeOpSendHeartbeat) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpSendHeartbeat) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*SendHeartbeatInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/SendHeartbeat")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentSendHeartbeatInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsSendHeartbeatInput(v *SendHeartbeatInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentSendHeartbeatInput(v *SendHeartbeatInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AgentMetrics != nil {
		ok := object.Key("AgentMetrics")
		if err := awsRestjson1_serializeDocumentEdgeMetrics(v.AgentMetrics, ok); err != nil {
			return err
		}
	}

	if v.AgentVersion != nil {
		ok := object.Key("AgentVersion")
		ok.String(*v.AgentVersion)
	}

	if v.DeviceFleetName != nil {
		ok := object.Key("DeviceFleetName")
		ok.String(*v.DeviceFleetName)
	}

	if v.DeviceName != nil {
		ok := object.Key("DeviceName")
		ok.String(*v.DeviceName)
	}

	if v.Models != nil {
		ok := object.Key("Models")
		if err := awsRestjson1_serializeDocumentModels(v.Models, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentEdgeMetric(v *types.EdgeMetric, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Dimension != nil {
		ok := object.Key("Dimension")
		ok.String(*v.Dimension)
	}

	if v.MetricName != nil {
		ok := object.Key("MetricName")
		ok.String(*v.MetricName)
	}

	if v.Timestamp != nil {
		ok := object.Key("Timestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.Timestamp))
	}

	if v.Value != 0 {
		ok := object.Key("Value")
		ok.Double(v.Value)
	}

	return nil
}

func awsRestjson1_serializeDocumentEdgeMetrics(v []types.EdgeMetric, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentEdgeMetric(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentModel(v *types.Model, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.LatestInference != nil {
		ok := object.Key("LatestInference")
		ok.Double(smithytime.FormatEpochSeconds(*v.LatestInference))
	}

	if v.LatestSampleTime != nil {
		ok := object.Key("LatestSampleTime")
		ok.Double(smithytime.FormatEpochSeconds(*v.LatestSampleTime))
	}

	if v.ModelMetrics != nil {
		ok := object.Key("ModelMetrics")
		if err := awsRestjson1_serializeDocumentEdgeMetrics(v.ModelMetrics, ok); err != nil {
			return err
		}
	}

	if v.ModelName != nil {
		ok := object.Key("ModelName")
		ok.String(*v.ModelName)
	}

	if v.ModelVersion != nil {
		ok := object.Key("ModelVersion")
		ok.String(*v.ModelVersion)
	}

	return nil
}

func awsRestjson1_serializeDocumentModels(v []types.Model, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentModel(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}
