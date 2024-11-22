import grpc
import event_service_pb2
import event_service_pb2_grpc
from typing import List
from lib import SportEvent

from config import GRPC_SERVER_ADDRESS


# gRPC Client Setup
def get_grpc_stub():
    # Create a gRPC channel
    channel = grpc.insecure_channel(GRPC_SERVER_ADDRESS)
    # Create a stub (client)
    stub = event_service_pb2_grpc.EventServiceStub(channel)
    return stub


def send_events_via_grpc(events: List[SportEvent]):
    stub = get_grpc_stub()

    def generate_load_requests():
        for event in events:
            load_request = event_service_pb2.LoadRequest(
                info=event_service_pb2.EventInfo(
                    ekpId=event.id,
                    sportType=event.sport_type,
                    sportSubtype=event.sport_subtype,
                    name=event.name,
                    description=event.description,
                    dates=event_service_pb2.DateRange(
                        date_from=event.dates.from_, date_to=event.dates.to
                    ),
                    location=event.location,
                    participants=event.participants,
                    participantRequirements=[
                        event_service_pb2.ParticipantRequirements(
                            gender=info["gender"],
                            min_age=info["age"][0],
                            max_age=info["age"][1],
                        )
                        for info in event.gender_age_info
                    ],
                )
            )
            yield load_request

    try:
        print("Отправка событий на gRPC сервер...")
        response = stub.Load(generate_load_requests())
        print(f"Сохранено событий на сервере: {response.saved}")
    except grpc.RpcError as e:
        print(f"gRPC ошибка: {e.code()} - {e.details()}")
