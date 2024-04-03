import grpc
from concurrent.futures import ThreadPoolExecutor

def serve():
    server = grpc.server(ThreadPoolExecutor(max_workers=10))
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    print("Welcome to the world's most over-complicated calculator!")
    serve()
