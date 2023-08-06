"""
Org-mode compatible calendar implementation - handles a number of events in
time.
"""
import datetime as dt


from orgassist import log
from orgassist.calendar import DateType
from orgassist import helpers

class Calendar:
    """
    Manages multiple events, generates agenda
    """
    def __init__(self, agenda_path=None, agenda_content=None):
        "Initialize calendar"

        # Events sorted by sort_date
        self.events = []

        # Path to agenda
        self.agenda_path = agenda_path
        self.agenda_content = agenda_content
        assert self.agenda_content != self.agenda_path

    def add_events(self, events, internal_tag=None):
        "Add new events to the calendar"
        for event in events:
            event.meta['calendar_tag'] = internal_tag
        self.events += events
        self.events.sort()

    def del_events(self, internal_tag=None):
        "Delete events by internal tag"
        if internal_tag is None:
            self.events = []

        else:
            self.events = [
                event
                for event in self.events
                if event.meta['calendar_tag'] != internal_tag
            ]

    def get_unfinished(self, horizon,
                       list_unfinished_appointments,
                       relative_to=None):
        """
        Get a list of past unfinished events

        Args:
          horizon (datetime): The oldest date to consider.
          relative_to (datetime): The relative "now" time.
          list_unfinished_appointments (bool): Return all open or just scheduled.
        """
        if relative_to is None:
            relative_to = dt.datetime.now()

        print("GET UNFINISHED")
        unfinished = []
        for event in self.events:
            print("  ", event)
            date = event.relevant_date.sort_date
            if date < horizon:
                print("  BEFORE HORIZON")
                continue
            if event.state is None:
                print("  NO STATE, SO NOT AN OPEN STATE")
                continue
            if not event.state.is_open:
                print("  NOT OPEN STATE")
                continue
            if list_unfinished_appointments is False:
                print("  CHECK IF APP", event.date_types)
                if (DateType.SCHEDULED not in event.date_types and
                    DateType.DEADLINE not in event.date_types):
                    print("  NOT SCHEDULED")
                    continue
            if date > relative_to:
                # We are in future - not unfinished.
                break

            print("  GOT YOU")
            unfinished.append(event)
        return unfinished

    def get_appointments(self, since, horizon):
        "Get a list of scheduled and planned events"
        appointments = []

        print("GET APPOINTMENTS")
        appointments = []
        for event in self.events:
            # Include only appointments
            if not event.relevant_date.appointment:
                continue

            date = event.relevant_date.sort_date
            print("  ", event)
            if date < since:
                print("    IN PAST")
                continue
            if date > horizon:
                print("    OVER HORIZON")
                break

            print("  GOT YOU")
            appointments.append(event)
        return appointments

    def get_scheduled(self, horizon, relative_to=None):
        "Get tasks scheduled or deadlining in given period"

        print("GET SCHEDULED")
        scheduled = []
        for event in self.events:
            date = event.relevant_date.sort_date
            print("  ", event)
            if date < relative_to:
                print("    IN PAST")
                continue
            if date > horizon:
                print("    OVER HORIZON")
                break

            # State doesn't matter as long as the date is accurate
            if not event.relevant_date.appointment:
                continue

            print("  GOT YOU")
            scheduled.append(event)

        return scheduled

    def get_agenda(self, horizon_incoming, horizon_unfinished,
                   list_unfinished_appointments, relative_to=None):
        "Generate agenda in a text format"
        if relative_to is None:
            relative_to = dt.datetime.now()

        log.info("Getting agenda from %r to %r",
                 horizon_unfinished, horizon_incoming)

        # Open and read when needed so the file can be updated
        # without restarting bot.
        template = helpers.get_template(self.agenda_path, self.agenda_content)

        since = relative_to.replace(hour=0, minute=0)
        if (relative_to - since).total_seconds() < 4*60*60:
            # Include more past
            since -= dt.timedelta(hours=4)

        ctx = {
            'unfinished': self.get_unfinished(horizon_unfinished,
                                              list_unfinished_appointments,
                                              relative_to),
            # Incoming - from the whole day
            'appointments': self.get_appointments(since=since,
                                                  horizon=horizon_incoming),
            'now': relative_to,
        }

        return template.render(ctx)

    def __repr__(self):
        txt = "<Calendar events=%d>"
        return txt % len(self.events)
