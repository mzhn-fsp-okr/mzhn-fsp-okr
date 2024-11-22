from dataclasses import dataclass, field, asdict


@dataclass
class DatesRange:
    from_: str = ""
    to: str = ""

    def __str__(self) -> str:
        return f"{self.from_} - {self.to}"


@dataclass
class SportEvent:
    page_number: int
    event_order: int

    id: str
    sport_type: str
    sport_subtype: str

    name: str = ""
    description: str = ""
    dates: DatesRange = field(default_factory=DatesRange)
    location: str = ""
    participants: int = 0

    gender_age_info: dict = field(default_factory=dict)

    def __str__(self) -> str:
        details = [
            f"id: {self.id}",
            f"sport: {self.sport_type}",
            f"subtype: {self.sport_subtype}",
            f"name: {self.name}" if self.name else "",
            f"dates: {self.dates}" if self.dates.from_ or self.dates.to else "",
            f"description: {self.description}" if self.description else "",
            f"location: {self.location}" if self.location else "",
            f"participants: {self.participants}" if self.participants > 0 else "",
            f"page: {self.page_number}",
            f"order: {self.event_order}",
        ]
        return "\n".join(filter(bool, details))


def sport_event_to_dict(event: SportEvent) -> dict:
    """
    Преобразует объект SportEvent в словарь.
    """
    event_dict = asdict(event)

    if "dates" in event_dict:
        dates = event_dict["dates"]
        event_dict["dates"] = {"from": dates["from_"], "to": dates["to"]}

    return event_dict
