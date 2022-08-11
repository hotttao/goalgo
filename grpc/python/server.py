import grpc
from concurrent import futures

from helloworld.helloword_pb2 import HelloReply
from helloworld.helloword_pb2_grpc import GreeterServicer
from helloworld.helloword_pb2_grpc import add_GreeterServicer_to_server


class Greeter(GreeterServicer):
    """The greeting service definition.
    """

    def SayHello(self, request, context):
        """Sends multiple greetings
        """
        return HelloReply(message=f'hello {request.name}')

    def SayHelloReplyStream(self, request, context):
        """Missing associated documentation comment in .proto file."""
        for i in range(3):
            yield HelloReply(message=f'hello {request.name} - {i}')

    def SayHelloRequestStream(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        print('start SayHelloRequestStream')
        for request in request_iterator:
            print(f"get request: {request.name} - {request.num_greetings}")
        return HelloReply(message=f'SayHelloRequestStream end')

    def SayHelloStream(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        for request in request_iterator:
            yield HelloReply(message=f'hello {request.name}')


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_GreeterServicer_to_server(
        Greeter(), server)
    server.add_insecure_port('[::]:50052')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()