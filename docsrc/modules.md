# Modules and sequences

Here's Modules' responsibilities and relations.

## Basic event loop

```uml
@startuml

hide footbox

participant "Main loop\n(Drive)" as main
participant "Event\nQueue"       as eq
participant "Event\nSubscribers" as es

loop while game scene
  main -> eq : fetch events
  main <- eq : events
  loop len(events)
    main -> es : fire events
    es -> es : do something
    note right of es : each events may\npush new event to\nevent queue
    eq <- es : (may push a new events)
  end
end

@enduml
```

* It is necessary to care the amount of events not to slow down game loop.
  * All events fired within a Drive function call,  
  should be completed within 1/60 sec if the game want to keep 60 fps.

## Strategy of balloon spawning

* Basically ballon spawns in random. Following factors are in random.
  * Spawn interval.
  * Balloon appears from bottom line but X position.
  * Balloon's moving speed.
  * Balloon's size.
  * Balloon's color.
  * A frog is in balloon or not.
* But obey following rules.
  * Don't spawn multiple balloons in same time.

```uml
@startuml

hide footbox

participant "Main loop\n(Drive)" as m
participant "Event\nQueue"       as eq
participant Spawner              as s
participant Factory              as f

create s
m -> s : new Spawner(frame counter)
s -> s : event subscribe
m -> s : run
eq <- s : push(JudgeSpawn)
create f
m -> f : new Factory
f -> f : event subscribe
loop while game scene
  m -> eq : fetch events
  m <- eq : events
  loop len(events)
  == event is JudgeSpawn ==
    m -> s : event publish
    m -> f : event publish
    note right of f : ignore
    alt
      note right of s : push event\nin random
      s -> eq : push(Spawn)
    end
    s -> eq : push(JudgeSpawn)
  == event is Spawn ==
    m -> s : event publish
    note right of s : ignore
    m -> f : event publish
    f -> f : spawn()
    note right of f : spawn a new balloon
  end
end

@enduml
```
