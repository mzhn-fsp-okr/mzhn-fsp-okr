import os

GRPC_SERVER_ADDRESS = os.getenv("GRPC_SERVER_ADDRESS")

LOG_LEVEL = os.getenv("LOG_LEVEL", "ERROR").upper()
