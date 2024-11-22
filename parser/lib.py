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
        return f"""id: {self.id}
sport: {self.sport_type}
subtype: {self.sport_subtype}
name: {self.name}
dates: {self.dates}
description: {self.description}
location: {self.location}
participants: {self.participants}
page: {self.page_number}
order: {self.event_order}"""
    

def sport_event_to_dict(event: SportEvent) -> dict:
    event_dict = asdict(event)
    
    if 'dates' in event_dict:
        dates = event_dict['dates']
        event_dict['dates'] = {
            'from': dates['from_'],
            'to': dates['to']
        }
    
    return event_dict