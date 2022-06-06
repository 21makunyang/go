include "base.thrift"
include "self_ref.thrift"
namespace go kitex.test.server

enum FOO {
    A = 1;
}

struct ExampleReq {
    1: required string Msg,
    2: required FOO Foo,
    255: base.Base Base,
}
struct ExampleResp {
    1: required string Msg,
    2: required string required_field
    3: string extra_field
    255: base.BaseResp BaseResp,
}
exception Exception {
    1: i32 code
    2: string msg
}

struct A {
    1: A self
    2: self_ref.A a
}

service ExampleService {
    ExampleResp ExampleMethod(1: ExampleReq req)throws(1: Exception err),
    A Foo(1: A req)
    string Ping(1: string msg)
    oneway void Oneway(1: string msg)
    void Void(1: string msg)
}