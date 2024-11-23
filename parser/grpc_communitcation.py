import grpc
import event_service_pb2
import event_service_pb2_grpc
from typing import List
from lib import SportEvent
from config import EVENT_SERVICE_ADDRESS
from logger import logging


def get_grpc_stub():
    """
    Создает и возвращает gRPC-клиентский объект (stub) для взаимодействия с сервером.
    """
    channel = grpc.insecure_channel(EVENT_SERVICE_ADDRESS)
    stub = event_service_pb2_grpc.EventServiceStub(channel)
    return stub


def send_events_via_grpc(events: List[SportEvent]):
    """
    Отправляет список спортивных событий на gRPC-сервер.

    :param events: Список объектов SportEvent для отправки.
    """
    stub = get_grpc_stub()

    def generate_load_requests():
        """
        Генератор для создания и отправки LoadRequest сообщений на сервер.
        """
        for event in events:
            yield event_service_pb2.LoadRequest(
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

    try:
        logging.info("Начало отправки событий на gRPC сервер...")
        response = stub.Load(generate_load_requests())
        logging.info(f"Сохранено событий на сервере: {response.saved}")
    except grpc.RpcError as e:
        # Обработка исключений с выводом кода ошибки и подробностей
        logging.error(f"gRPC ошибка: {e.code()} - {e.details()}")
