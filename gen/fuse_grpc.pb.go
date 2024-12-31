// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: fuse.proto

package fuse

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Sds_ExecuteSql_FullMethodName           = "/sds.sds/ExecuteSql"
	Sds_CreateSession_FullMethodName        = "/sds.sds/CreateSession"
	Sds_LoadCSV_FullMethodName              = "/sds.sds/LoadCSV"
	Sds_LoadParquet_FullMethodName          = "/sds.sds/LoadParquet"
	Sds_LoadJSON_FullMethodName             = "/sds.sds/LoadJSON"
	Sds_LoadDeltaTable_FullMethodName       = "/sds.sds/LoadDeltaTable"
	Sds_CloseSession_FullMethodName         = "/sds.sds/CloseSession"
	Sds_SaveDataFrameAsTable_FullMethodName = "/sds.sds/SaveDataFrameAsTable"
	Sds_PrettyPrintDataframe_FullMethodName = "/sds.sds/PrettyPrintDataframe"
	Sds_LimitDataFrame_FullMethodName       = "/sds.sds/LimitDataFrame"
	Sds_SortDataFrame_FullMethodName        = "/sds.sds/SortDataFrame"
	Sds_FilterDataFrame_FullMethodName      = "/sds.sds/FilterDataFrame"
	Sds_Aggregate_FullMethodName            = "/sds.sds/Aggregate"
	Sds_WithColumn_FullMethodName           = "/sds.sds/WithColumn"
	Sds_Select_FullMethodName               = "/sds.sds/Select"
	Sds_Collect_FullMethodName              = "/sds.sds/Collect"
	Sds_Join_FullMethodName                 = "/sds.sds/Join"
	Sds_Distinct_FullMethodName             = "/sds.sds/Distinct"
	Sds_Union_FullMethodName                = "/sds.sds/Union"
	Sds_Drop_FullMethodName                 = "/sds.sds/Drop"
	Sds_ExportCSV_FullMethodName            = "/sds.sds/ExportCSV"
)

// SdsClient is the client API for Sds service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SdsClient interface {
	// execute SQL against a session and return a new DataFrame representing
	// the result
	ExecuteSql(ctx context.Context, in *ExecuteSqlRequest, opts ...grpc.CallOption) (*DataFrameUID, error)
	// create a new session with no DataFrames therein, then return its UID
	CreateSession(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SessionUID, error)
	// Load a CSV into a DataFrame, then return its UID
	LoadCSV(ctx context.Context, in *LoadFileRequest, opts ...grpc.CallOption) (*DataFrameUID, error)
	// Load a Parquet file into a DataFrame, then return its UID
	LoadParquet(ctx context.Context, in *LoadFileRequest, opts ...grpc.CallOption) (*DataFrameUID, error)
	// Load a JSON file into a DataFrame, then return its UID
	LoadJSON(ctx context.Context, in *LoadFileRequest, opts ...grpc.CallOption) (*DataFrameUID, error)
	// Load a Delta Table into a DataFrame, then return its UID
	LoadDeltaTable(ctx context.Context, in *LoadFileRequest, opts ...grpc.CallOption) (*DataFrameUID, error)
	// Close a session and free all associated dataframes.
	//
	// This is a highly destructive method, because all dataframes
	// created directly or indirectly as part of this session
	// will be instantly deleted.
	CloseSession(ctx context.Context, in *SessionUID, opts ...grpc.CallOption) (*Empty, error)
	// save a dataframe as a table, so that you can execute SQL queries
	// against it
	SaveDataFrameAsTable(ctx context.Context, in *SaveDataFrameAsTableRequest, opts ...grpc.CallOption) (*Empty, error)
	// pretty-print a given dataframe.
	//
	// TODO: throw an error if the dataframe is too big to pretty-print
	PrettyPrintDataframe(ctx context.Context, in *DataFrameUID, opts ...grpc.CallOption) (*PrettyPrintDataframeResponse, error)
	// filter an existing DataFrame, and return a new DataFrame
	LimitDataFrame(ctx context.Context, in *LimitDataFrameRequest, opts ...grpc.CallOption) (*DataFrameUID, error)
	// sort a dataframe by 1 or more column(s)
	SortDataFrame(ctx context.Context, in *SortDataFrameRequest, opts ...grpc.CallOption) (*DataFrameUID, error)
	// filter a dataframe according to 1 or more filter conditions
	FilterDataFrame(ctx context.Context, in *FilterDataFrameRequest, opts ...grpc.CallOption) (*DataFrameUID, error)
	// Group rows in a DataFrame, then compute an aggregate across all the
	// rows in each group
	Aggregate(ctx context.Context, in *AggregateRequest, opts ...grpc.CallOption) (*DataFrameUID, error)
	// add a new column -- optionally by doing some calculation -- to a
	// given dataframe
	WithColumn(ctx context.Context, in *WithColumnRequest, opts ...grpc.CallOption) (*DataFrameUID, error)
	// project a DataFrame onto a new one, optionally by calculating new
	// values
	Select(ctx context.Context, in *SelectRequest, opts ...grpc.CallOption) (*DataFrameUID, error)
	// eagerly evaluate the dataframe, then return its contents
	Collect(ctx context.Context, in *DataFrameUID, opts ...grpc.CallOption) (*DataFrameContents, error)
	// Join 2 dataframes together into one
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*DataFrameUID, error)
	// Return a new DataFrame whose contents are the same as the given
	// DataFrame, except with duplicate rows removed
	Distinct(ctx context.Context, in *DataFrameUID, opts ...grpc.CallOption) (*DataFrameUID, error)
	// Combine two DataFrames together to calculate the union of the two.
	//
	// Both DataFrames must have the same schema. If they do not, return
	// an error. If they do, preserve duplicate rows in the final result.
	Union(ctx context.Context, in *UnionRequest, opts ...grpc.CallOption) (*DataFrameUID, error)
	// return a new DataFrame with the given columns missing.
	//
	// if you pass a column that does not exist, this entire operation
	// will be a no-op and you'll get the same UUID back
	Drop(ctx context.Context, in *DropRequest, opts ...grpc.CallOption) (*DataFrameUID, error)
	// export a CSV file and return it in the response
	ExportCSV(ctx context.Context, in *DataFrameUID, opts ...grpc.CallOption) (*CSVOutput, error)
}

type sdsClient struct {
	cc grpc.ClientConnInterface
}

func NewSdsClient(cc grpc.ClientConnInterface) SdsClient {
	return &sdsClient{cc}
}

func (c *sdsClient) ExecuteSql(ctx context.Context, in *ExecuteSqlRequest, opts ...grpc.CallOption) (*DataFrameUID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataFrameUID)
	err := c.cc.Invoke(ctx, Sds_ExecuteSql_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) CreateSession(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SessionUID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SessionUID)
	err := c.cc.Invoke(ctx, Sds_CreateSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) LoadCSV(ctx context.Context, in *LoadFileRequest, opts ...grpc.CallOption) (*DataFrameUID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataFrameUID)
	err := c.cc.Invoke(ctx, Sds_LoadCSV_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) LoadParquet(ctx context.Context, in *LoadFileRequest, opts ...grpc.CallOption) (*DataFrameUID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataFrameUID)
	err := c.cc.Invoke(ctx, Sds_LoadParquet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) LoadJSON(ctx context.Context, in *LoadFileRequest, opts ...grpc.CallOption) (*DataFrameUID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataFrameUID)
	err := c.cc.Invoke(ctx, Sds_LoadJSON_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) LoadDeltaTable(ctx context.Context, in *LoadFileRequest, opts ...grpc.CallOption) (*DataFrameUID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataFrameUID)
	err := c.cc.Invoke(ctx, Sds_LoadDeltaTable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) CloseSession(ctx context.Context, in *SessionUID, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Sds_CloseSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) SaveDataFrameAsTable(ctx context.Context, in *SaveDataFrameAsTableRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Sds_SaveDataFrameAsTable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) PrettyPrintDataframe(ctx context.Context, in *DataFrameUID, opts ...grpc.CallOption) (*PrettyPrintDataframeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrettyPrintDataframeResponse)
	err := c.cc.Invoke(ctx, Sds_PrettyPrintDataframe_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) LimitDataFrame(ctx context.Context, in *LimitDataFrameRequest, opts ...grpc.CallOption) (*DataFrameUID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataFrameUID)
	err := c.cc.Invoke(ctx, Sds_LimitDataFrame_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) SortDataFrame(ctx context.Context, in *SortDataFrameRequest, opts ...grpc.CallOption) (*DataFrameUID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataFrameUID)
	err := c.cc.Invoke(ctx, Sds_SortDataFrame_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) FilterDataFrame(ctx context.Context, in *FilterDataFrameRequest, opts ...grpc.CallOption) (*DataFrameUID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataFrameUID)
	err := c.cc.Invoke(ctx, Sds_FilterDataFrame_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) Aggregate(ctx context.Context, in *AggregateRequest, opts ...grpc.CallOption) (*DataFrameUID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataFrameUID)
	err := c.cc.Invoke(ctx, Sds_Aggregate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) WithColumn(ctx context.Context, in *WithColumnRequest, opts ...grpc.CallOption) (*DataFrameUID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataFrameUID)
	err := c.cc.Invoke(ctx, Sds_WithColumn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) Select(ctx context.Context, in *SelectRequest, opts ...grpc.CallOption) (*DataFrameUID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataFrameUID)
	err := c.cc.Invoke(ctx, Sds_Select_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) Collect(ctx context.Context, in *DataFrameUID, opts ...grpc.CallOption) (*DataFrameContents, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataFrameContents)
	err := c.cc.Invoke(ctx, Sds_Collect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*DataFrameUID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataFrameUID)
	err := c.cc.Invoke(ctx, Sds_Join_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) Distinct(ctx context.Context, in *DataFrameUID, opts ...grpc.CallOption) (*DataFrameUID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataFrameUID)
	err := c.cc.Invoke(ctx, Sds_Distinct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) Union(ctx context.Context, in *UnionRequest, opts ...grpc.CallOption) (*DataFrameUID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataFrameUID)
	err := c.cc.Invoke(ctx, Sds_Union_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) Drop(ctx context.Context, in *DropRequest, opts ...grpc.CallOption) (*DataFrameUID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataFrameUID)
	err := c.cc.Invoke(ctx, Sds_Drop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdsClient) ExportCSV(ctx context.Context, in *DataFrameUID, opts ...grpc.CallOption) (*CSVOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CSVOutput)
	err := c.cc.Invoke(ctx, Sds_ExportCSV_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SdsServer is the server API for Sds service.
// All implementations must embed UnimplementedSdsServer
// for forward compatibility.
type SdsServer interface {
	// execute SQL against a session and return a new DataFrame representing
	// the result
	ExecuteSql(context.Context, *ExecuteSqlRequest) (*DataFrameUID, error)
	// create a new session with no DataFrames therein, then return its UID
	CreateSession(context.Context, *Empty) (*SessionUID, error)
	// Load a CSV into a DataFrame, then return its UID
	LoadCSV(context.Context, *LoadFileRequest) (*DataFrameUID, error)
	// Load a Parquet file into a DataFrame, then return its UID
	LoadParquet(context.Context, *LoadFileRequest) (*DataFrameUID, error)
	// Load a JSON file into a DataFrame, then return its UID
	LoadJSON(context.Context, *LoadFileRequest) (*DataFrameUID, error)
	// Load a Delta Table into a DataFrame, then return its UID
	LoadDeltaTable(context.Context, *LoadFileRequest) (*DataFrameUID, error)
	// Close a session and free all associated dataframes.
	//
	// This is a highly destructive method, because all dataframes
	// created directly or indirectly as part of this session
	// will be instantly deleted.
	CloseSession(context.Context, *SessionUID) (*Empty, error)
	// save a dataframe as a table, so that you can execute SQL queries
	// against it
	SaveDataFrameAsTable(context.Context, *SaveDataFrameAsTableRequest) (*Empty, error)
	// pretty-print a given dataframe.
	//
	// TODO: throw an error if the dataframe is too big to pretty-print
	PrettyPrintDataframe(context.Context, *DataFrameUID) (*PrettyPrintDataframeResponse, error)
	// filter an existing DataFrame, and return a new DataFrame
	LimitDataFrame(context.Context, *LimitDataFrameRequest) (*DataFrameUID, error)
	// sort a dataframe by 1 or more column(s)
	SortDataFrame(context.Context, *SortDataFrameRequest) (*DataFrameUID, error)
	// filter a dataframe according to 1 or more filter conditions
	FilterDataFrame(context.Context, *FilterDataFrameRequest) (*DataFrameUID, error)
	// Group rows in a DataFrame, then compute an aggregate across all the
	// rows in each group
	Aggregate(context.Context, *AggregateRequest) (*DataFrameUID, error)
	// add a new column -- optionally by doing some calculation -- to a
	// given dataframe
	WithColumn(context.Context, *WithColumnRequest) (*DataFrameUID, error)
	// project a DataFrame onto a new one, optionally by calculating new
	// values
	Select(context.Context, *SelectRequest) (*DataFrameUID, error)
	// eagerly evaluate the dataframe, then return its contents
	Collect(context.Context, *DataFrameUID) (*DataFrameContents, error)
	// Join 2 dataframes together into one
	Join(context.Context, *JoinRequest) (*DataFrameUID, error)
	// Return a new DataFrame whose contents are the same as the given
	// DataFrame, except with duplicate rows removed
	Distinct(context.Context, *DataFrameUID) (*DataFrameUID, error)
	// Combine two DataFrames together to calculate the union of the two.
	//
	// Both DataFrames must have the same schema. If they do not, return
	// an error. If they do, preserve duplicate rows in the final result.
	Union(context.Context, *UnionRequest) (*DataFrameUID, error)
	// return a new DataFrame with the given columns missing.
	//
	// if you pass a column that does not exist, this entire operation
	// will be a no-op and you'll get the same UUID back
	Drop(context.Context, *DropRequest) (*DataFrameUID, error)
	// export a CSV file and return it in the response
	ExportCSV(context.Context, *DataFrameUID) (*CSVOutput, error)
	mustEmbedUnimplementedSdsServer()
}

// UnimplementedSdsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSdsServer struct{}

func (UnimplementedSdsServer) ExecuteSql(context.Context, *ExecuteSqlRequest) (*DataFrameUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteSql not implemented")
}
func (UnimplementedSdsServer) CreateSession(context.Context, *Empty) (*SessionUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedSdsServer) LoadCSV(context.Context, *LoadFileRequest) (*DataFrameUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadCSV not implemented")
}
func (UnimplementedSdsServer) LoadParquet(context.Context, *LoadFileRequest) (*DataFrameUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadParquet not implemented")
}
func (UnimplementedSdsServer) LoadJSON(context.Context, *LoadFileRequest) (*DataFrameUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadJSON not implemented")
}
func (UnimplementedSdsServer) LoadDeltaTable(context.Context, *LoadFileRequest) (*DataFrameUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadDeltaTable not implemented")
}
func (UnimplementedSdsServer) CloseSession(context.Context, *SessionUID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseSession not implemented")
}
func (UnimplementedSdsServer) SaveDataFrameAsTable(context.Context, *SaveDataFrameAsTableRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveDataFrameAsTable not implemented")
}
func (UnimplementedSdsServer) PrettyPrintDataframe(context.Context, *DataFrameUID) (*PrettyPrintDataframeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrettyPrintDataframe not implemented")
}
func (UnimplementedSdsServer) LimitDataFrame(context.Context, *LimitDataFrameRequest) (*DataFrameUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LimitDataFrame not implemented")
}
func (UnimplementedSdsServer) SortDataFrame(context.Context, *SortDataFrameRequest) (*DataFrameUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SortDataFrame not implemented")
}
func (UnimplementedSdsServer) FilterDataFrame(context.Context, *FilterDataFrameRequest) (*DataFrameUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterDataFrame not implemented")
}
func (UnimplementedSdsServer) Aggregate(context.Context, *AggregateRequest) (*DataFrameUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Aggregate not implemented")
}
func (UnimplementedSdsServer) WithColumn(context.Context, *WithColumnRequest) (*DataFrameUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithColumn not implemented")
}
func (UnimplementedSdsServer) Select(context.Context, *SelectRequest) (*DataFrameUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Select not implemented")
}
func (UnimplementedSdsServer) Collect(context.Context, *DataFrameUID) (*DataFrameContents, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}
func (UnimplementedSdsServer) Join(context.Context, *JoinRequest) (*DataFrameUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedSdsServer) Distinct(context.Context, *DataFrameUID) (*DataFrameUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Distinct not implemented")
}
func (UnimplementedSdsServer) Union(context.Context, *UnionRequest) (*DataFrameUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Union not implemented")
}
func (UnimplementedSdsServer) Drop(context.Context, *DropRequest) (*DataFrameUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Drop not implemented")
}
func (UnimplementedSdsServer) ExportCSV(context.Context, *DataFrameUID) (*CSVOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportCSV not implemented")
}
func (UnimplementedSdsServer) mustEmbedUnimplementedSdsServer() {}
func (UnimplementedSdsServer) testEmbeddedByValue()             {}

// UnsafeSdsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SdsServer will
// result in compilation errors.
type UnsafeSdsServer interface {
	mustEmbedUnimplementedSdsServer()
}

func RegisterSdsServer(s grpc.ServiceRegistrar, srv SdsServer) {
	// If the following call pancis, it indicates UnimplementedSdsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Sds_ServiceDesc, srv)
}

func _Sds_ExecuteSql_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteSqlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).ExecuteSql(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_ExecuteSql_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).ExecuteSql(ctx, req.(*ExecuteSqlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_CreateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).CreateSession(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_LoadCSV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).LoadCSV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_LoadCSV_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).LoadCSV(ctx, req.(*LoadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_LoadParquet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).LoadParquet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_LoadParquet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).LoadParquet(ctx, req.(*LoadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_LoadJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).LoadJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_LoadJSON_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).LoadJSON(ctx, req.(*LoadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_LoadDeltaTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).LoadDeltaTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_LoadDeltaTable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).LoadDeltaTable(ctx, req.(*LoadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_CloseSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).CloseSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_CloseSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).CloseSession(ctx, req.(*SessionUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_SaveDataFrameAsTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveDataFrameAsTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).SaveDataFrameAsTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_SaveDataFrameAsTable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).SaveDataFrameAsTable(ctx, req.(*SaveDataFrameAsTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_PrettyPrintDataframe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataFrameUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).PrettyPrintDataframe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_PrettyPrintDataframe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).PrettyPrintDataframe(ctx, req.(*DataFrameUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_LimitDataFrame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LimitDataFrameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).LimitDataFrame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_LimitDataFrame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).LimitDataFrame(ctx, req.(*LimitDataFrameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_SortDataFrame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SortDataFrameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).SortDataFrame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_SortDataFrame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).SortDataFrame(ctx, req.(*SortDataFrameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_FilterDataFrame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterDataFrameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).FilterDataFrame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_FilterDataFrame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).FilterDataFrame(ctx, req.(*FilterDataFrameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_Aggregate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AggregateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).Aggregate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_Aggregate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).Aggregate(ctx, req.(*AggregateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_WithColumn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithColumnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).WithColumn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_WithColumn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).WithColumn(ctx, req.(*WithColumnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_Select_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).Select(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_Select_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).Select(ctx, req.(*SelectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataFrameUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_Collect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).Collect(ctx, req.(*DataFrameUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_Join_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_Distinct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataFrameUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).Distinct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_Distinct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).Distinct(ctx, req.(*DataFrameUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_Union_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).Union(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_Union_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).Union(ctx, req.(*UnionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_Drop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).Drop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_Drop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).Drop(ctx, req.(*DropRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sds_ExportCSV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataFrameUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdsServer).ExportCSV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sds_ExportCSV_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdsServer).ExportCSV(ctx, req.(*DataFrameUID))
	}
	return interceptor(ctx, in, info, handler)
}

// Sds_ServiceDesc is the grpc.ServiceDesc for Sds service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sds_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sds.sds",
	HandlerType: (*SdsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteSql",
			Handler:    _Sds_ExecuteSql_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _Sds_CreateSession_Handler,
		},
		{
			MethodName: "LoadCSV",
			Handler:    _Sds_LoadCSV_Handler,
		},
		{
			MethodName: "LoadParquet",
			Handler:    _Sds_LoadParquet_Handler,
		},
		{
			MethodName: "LoadJSON",
			Handler:    _Sds_LoadJSON_Handler,
		},
		{
			MethodName: "LoadDeltaTable",
			Handler:    _Sds_LoadDeltaTable_Handler,
		},
		{
			MethodName: "CloseSession",
			Handler:    _Sds_CloseSession_Handler,
		},
		{
			MethodName: "SaveDataFrameAsTable",
			Handler:    _Sds_SaveDataFrameAsTable_Handler,
		},
		{
			MethodName: "PrettyPrintDataframe",
			Handler:    _Sds_PrettyPrintDataframe_Handler,
		},
		{
			MethodName: "LimitDataFrame",
			Handler:    _Sds_LimitDataFrame_Handler,
		},
		{
			MethodName: "SortDataFrame",
			Handler:    _Sds_SortDataFrame_Handler,
		},
		{
			MethodName: "FilterDataFrame",
			Handler:    _Sds_FilterDataFrame_Handler,
		},
		{
			MethodName: "Aggregate",
			Handler:    _Sds_Aggregate_Handler,
		},
		{
			MethodName: "WithColumn",
			Handler:    _Sds_WithColumn_Handler,
		},
		{
			MethodName: "Select",
			Handler:    _Sds_Select_Handler,
		},
		{
			MethodName: "Collect",
			Handler:    _Sds_Collect_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _Sds_Join_Handler,
		},
		{
			MethodName: "Distinct",
			Handler:    _Sds_Distinct_Handler,
		},
		{
			MethodName: "Union",
			Handler:    _Sds_Union_Handler,
		},
		{
			MethodName: "Drop",
			Handler:    _Sds_Drop_Handler,
		},
		{
			MethodName: "ExportCSV",
			Handler:    _Sds_ExportCSV_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fuse.proto",
}
