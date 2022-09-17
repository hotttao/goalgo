import grpc
from helloworld.helloword_pb2_grpc import GreeterStub
from helloworld.helloword_pb2 import HelloRequest

def say_hello(stub: GreeterStub):
    request = HelloRequest(name='tao', num_greetings="2")
    response = stub.SayHello(request, wait_for_ready=True)
    print(f"{response.message}")
    
def say_hello_request_stream(stub: GreeterStub):
    requests = iter([
        HelloRequest(name='tao', num_greetings="2"),
        HelloRequest(name='song', num_greetings="2")
    ])
    response = stub.SayHelloRequestStream(requests)
    print(f"{response.message}")

def say_hello_reply_stream(stub: GreeterStub):
    request = HelloRequest(name='tao', num_greetings="2")
    for response in stub.SayHelloReplyStream(request):
        print(f"{response.message}")

def say_hello_stream(stub: GreeterStub):
    requests = iter([
        HelloRequest(name='tao', num_greetings="2"),
        HelloRequest(name='song', num_greetings="2")
    ])
    response_iter = stub.SayHelloStream(requests)
    for response in response_iter:
        print(f"{response.message}")


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('localhost:50052') as channel:
        stub = GreeterStub(channel)
        # stub.initial_rpc_call()
        print("-------------- say_hello --------------")
        say_hello(stub)
        print("-------------- say_hello_request_stream --------------")
        say_hello_request_stream(stub)
        print("-------------- say_hello_reply_stream --------------")
        say_hello_reply_stream(stub)
        print("-------------- say_hello_stream --------------")
        say_hello_stream(stub)


if __name__ == '__main__':
    run()